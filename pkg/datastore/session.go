package datastore

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// SessionFactory implements a method to return a pointer to mgo.Session
type SessionFactory struct {
	session *mgo.Session
}

// NewSessionFactory returns a SessionFactory
func NewSessionFactory(host string, port string) (*SessionFactory, error) {
	session, err := mgo.Dial(fmt.Sprintf("%s:%s", host, port))

	if err != nil {
		return nil, err
	}

	return &SessionFactory{session: session}, nil
}

// Get returns a pointer to mgo.Session
func (s SessionFactory) Get() *mgo.Session {
	return s.session.Copy()
}
