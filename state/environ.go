// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/txn"

	"launchpad.net/juju-core/errors"
	"launchpad.net/juju-core/names"
)

// environGlobalKey is the key for the environment, its
// settings and constraints.
const environGlobalKey = "e"

// Environment represents the state of an environment.
type Environment struct {
	st  *State
	doc environmentDoc
	annotator
}

// environmentDoc represents the internal state of the environment in MongoDB.
type environmentDoc struct {
	UUID string `bson:"_id"`
	Name string
	Life Life
}

// Environment returns the environment entity.
func (st *State) Environment() (*Environment, error) {
	env := &Environment{st: st}
	if err := env.refresh(st.environments.Find(nil)); err != nil {
		return nil, err
	}
	env.annotator = annotator{
		globalKey: env.globalKey(),
		tag:       env.Tag(),
		st:        st,
	}
	return env, nil
}

// Tag returns a name identifying the environment.
// The returned name will be different from other Tag values returned
// by any other entities from the same state.
func (e *Environment) Tag() string {
	return names.EnvironTag(e.doc.UUID)
}

// UUID returns the universally unique identifier of the environment.
func (e *Environment) UUID() string {
	return e.doc.UUID
}

// Life returns whether the environment is Alive, Dying or Dead.
func (e *Environment) Life() Life {
	return e.doc.Life
}

// globalKey returns the global database key for the environment.
func (e *Environment) globalKey() string {
	return environGlobalKey
}

func (e *Environment) Refresh() error {
	return e.refresh(e.st.environments.FindId(e.UUID()))
}

func (e *Environment) refresh(query *mgo.Query) error {
	err := query.One(&e.doc)
	if err == mgo.ErrNotFound {
		return errors.NotFoundf("environment")
	}
	return err
}

// Destroy sets the environment's lifecycle to Dying, preventing
// addition of services or machines to state.
func (e *Environment) Destroy() error {
	if e.Life() != Alive {
		return nil
	}
	// Note: we add a cleanup for services, but not for machines;
	// machines are destroyed via the provider interface. The
	// exception to this rule is manual machines; the API prevents
	// destroy-environment from succeeding if any non-manager
	// manual machines exist.
	ops := []txn.Op{{
		C:      e.st.environments.Name,
		Id:     e.doc.UUID,
		Update: D{{"$set", D{{"life", Dying}}}},
		Assert: isAliveDoc,
	}, e.st.newCleanupOp("services", "")}
	err := e.st.runTransaction(ops)
	switch err {
	case nil:
		e.doc.Life = Dying
	case txn.ErrAborted:
		// If the transaction aborted, the environment is either
		// Dying or Dead; neither case is an error.
		err = e.Refresh()
	}
	return err
}

// createEnvironmentOp returns the operation needed to create
// an environment document with the given name and UUID.
func createEnvironmentOp(st *State, name, uuid string) txn.Op {
	doc := &environmentDoc{uuid, name, Alive}
	return txn.Op{
		C:      st.environments.Name,
		Id:     uuid,
		Assert: txn.DocMissing,
		Insert: doc,
	}
}

// assertAliveOp returns a read-only txn.Op that asserts
// the environment is alive.
func (e *Environment) assertAliveOp() txn.Op {
	return txn.Op{
		C:      e.st.environments.Name,
		Id:     e.UUID(),
		Assert: isAliveDoc,
	}
}
