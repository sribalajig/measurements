package datastore

import (
	"github.com/hpi/measurement/pkg/model"
	"gopkg.in/mgo.v2/bson"
)

// Provider defines the interface for a data provider
type Provider interface {
	Create(userMeasurements *model.UserMeasurements) error
	GetByUserID(id int) *model.UserMeasurements
	Update(userMeasurements *model.UserMeasurements) error
}

type mongo struct {
	sessionFactory *SessionFactory
}

// NewMongo return a pointer to the mongo implementation of Provider
func NewMongo(s *SessionFactory) Provider {
	return &mongo{sessionFactory: s}
}

// Create - create user measurements
func (m *mongo) Create(userMeasurements *model.UserMeasurements) error {
	session := m.sessionFactory.Get()
	defer session.Close()

	return session.DB("hpi").C("user_measurements").Insert(userMeasurements)
}

// GetByUserId - get measurements by user id
func (m *mongo) GetByUserID(id int) *model.UserMeasurements {
	session := m.sessionFactory.Get()
	defer session.Close()

	var userMeasurements model.UserMeasurements
	find := session.DB("hpi").C("user_measurements").Find(bson.M{"id": id}).Iter()
	defer find.Close()

	found := find.Next(&userMeasurements)
	if !found {
		return nil
	}

	return &userMeasurements
}

// Update - update the given measurements for a user
func (m *mongo) Update(userMeasurements *model.UserMeasurements) error {
	session := m.sessionFactory.Get()

	return session.DB("hpi").C("user_measurements").Update(bson.M{"id": userMeasurements.ID}, userMeasurements)
}
