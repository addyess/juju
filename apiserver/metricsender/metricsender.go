// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

// Package metricsender contains functions for sending
// metrics from a controller to a remote metric collector.
package metricsender

import (
	"time"

	"github.com/juju/errors"
	"github.com/juju/loggo"
	wireformat "github.com/juju/romulus/wireformat/metrics"

	"github.com/juju/juju/state"
)

var logger = loggo.GetLogger("juju.apiserver.metricsender")

// MetricSender defines the interface used to send metrics
// to a collection service.
type MetricSender interface {
	Send([]*wireformat.MetricBatch) (*wireformat.Response, error)
}

var (
	defaultMaxBatchesPerSend              = 10
	defaultSender            MetricSender = &HttpSender{}
)

func handleResponse(mm *state.MetricsManager, st ModelBackend, response wireformat.Response) {
	for _, envResp := range response.EnvResponses {
		err := st.SetMetricBatchesSent(envResp.AcknowledgedBatches)
		if err != nil {
			logger.Errorf("failed to set sent on metrics %v", err)
		}
		for unitName, status := range envResp.UnitStatuses {
			unit, err := st.Unit(unitName)
			if err != nil {
				logger.Errorf("failed to retrieve unit %q: %v", unitName, err)
				continue
			}
			err = unit.SetMeterStatus(status.Status, status.Info)
			if err != nil {
				logger.Errorf("failed to set unit %q meter status to %v: %v", unitName, status, err)
			}
		}
	}
	if response.NewGracePeriod > 0 {
		err := mm.SetGracePeriod(response.NewGracePeriod)
		if err != nil {
			logger.Errorf("failed to set new grace period %v", err)
		}
	}
}

// SendMetrics will send any unsent metrics
// over the MetricSender interface in batches
// no larger than batchSize.
func SendMetrics(st ModelBackend, sender MetricSender, batchSize int, transmitVendorMetrics bool) error {
	metricsManager, err := st.MetricsManager()
	if err != nil {
		return errors.Trace(err)
	}
	sent := 0
	held := 0
	for {
		metrics, err := st.MetricsToSend(batchSize)
		if err != nil {
			return errors.Trace(err)
		}
		lenM := len(metrics)
		if lenM == 0 {
			if sent == 0 {
				logger.Infof("nothing to send")
			} else {
				logger.Infof("done sending")
			}
			break
		}

		var wireData []*wireformat.MetricBatch
		var heldBatches []string
		for _, m := range metrics {
			if !transmitVendorMetrics && len(m.Credentials()) == 0 {
				heldBatches = append(heldBatches, m.UUID())
			} else {
				wireData = append(wireData, ToWire(m))
			}
		}
		response, err := sender.Send(wireData)
		if err != nil {
			logger.Errorf("%+v", err)
			if incErr := metricsManager.IncrementConsecutiveErrors(); incErr != nil {
				logger.Errorf("failed to increment error count %v", incErr)
				return errors.Trace(errors.Wrap(err, incErr))
			}
			return errors.Trace(err)
		}
		if response != nil {
			// TODO (mattyw) We are currently ignoring errors during response handling.
			handleResponse(metricsManager, st, *response)
			// TODO(fwereade): 2016-03-17 lp:1558657
			if err := metricsManager.SetLastSuccessfulSend(time.Now()); err != nil {
				err = errors.Annotate(err, "failed to set successful send time")
				logger.Warningf("%v", err)
				return errors.Trace(err)
			}
		}

		// Mark held metric batches as sent so that they can be cleaned up later.
		if len(heldBatches) > 0 {
			err := st.SetMetricBatchesSent(heldBatches)
			if err != nil {
				return errors.Annotatef(err, "failed to mark metric batches as sent for %s", st.ModelTag())
			}
		}

		sent += len(wireData)
		held += len(heldBatches)
	}

	unsent, err := st.CountOfUnsentMetrics()
	if err != nil {
		return errors.Trace(err)
	}
	sentStored, err := st.CountOfSentMetrics()
	if err != nil {
		return errors.Trace(err)
	}
	logger.Infof("metrics collection summary for %s: sent:%d unsent:%d held:%d (%d sent metrics stored)", st.ModelTag(), sent, unsent, held, sentStored)

	return nil
}

// DefaultMaxBatchesPerSend returns the default number of batches per send.
func DefaultMaxBatchesPerSend() int {
	return defaultMaxBatchesPerSend
}

// DefaultMetricSender returns the default metric sender.
func DefaultMetricSender() MetricSender {
	return defaultSender
}

// ToWire converts the state.MetricBatch into a type
// that can be sent over the wire to the collector.
func ToWire(mb *state.MetricBatch) *wireformat.MetricBatch {
	metrics := make([]wireformat.Metric, len(mb.Metrics()))
	for i, m := range mb.Metrics() {
		metrics[i] = wireformat.Metric{
			Key:   m.Key,
			Value: m.Value,
			Time:  m.Time.UTC(),
		}
	}
	return &wireformat.MetricBatch{
		UUID:        mb.UUID(),
		ModelUUID:   mb.ModelUUID(),
		UnitName:    mb.Unit(),
		CharmUrl:    mb.CharmURL(),
		Created:     mb.Created().UTC(),
		Metrics:     metrics,
		Credentials: mb.Credentials(),
	}
}
