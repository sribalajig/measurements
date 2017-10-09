package service

import (
	"github.com/hpi/measurement/pkg/datastore"
	"github.com/hpi/measurement/pkg/model"
)

// MeasurementService ..
type MeasurementService struct {
	dataProvider datastore.Provider
}

// NewMeasurementService returns a pointer to MeasurementService
func NewMeasurementService(dp datastore.Provider) *MeasurementService {
	return &MeasurementService{
		dataProvider: dp,
	}
}

// Create - creates UserMeasurements
func (ms *MeasurementService) Create(userMeasurements *model.UserMeasurements) error {
	existing := ms.dataProvider.GetByUserID(userMeasurements.ID)

	if existing == nil {
		return ms.dataProvider.Create(userMeasurements)
	}

	updated := ms.merge(existing.Measurements, userMeasurements.Measurements)

	existing.Measurements = updated

	return ms.dataProvider.Update(existing)
}

// Get - get user measurements by ID
func (ms *MeasurementService) Get(id int, measurementIDs []int) *model.UserMeasurements {
	userMeasurements := ms.dataProvider.GetByUserID(id)

	if userMeasurements == nil {
		return nil
	}

	if measurementIDs == nil || len(measurementIDs) == 0 {
		return userMeasurements
	}

	return ms.filter(userMeasurements, measurementIDs)
}

// merge existing and updates user measurements
func (ms *MeasurementService) merge(
	existing []model.Measurement, updated []model.Measurement) []model.Measurement {
	for _, m := range updated {
		exists := false

		for index, e := range existing {
			if e.ID == m.ID {
				existing[index].Value = m.Value
				exists = true

				break
			}
		}

		if !exists {
			existing = append(existing, m)
		}
	}

	return existing
}

// filter out only requested measurement id's
func (ms *MeasurementService) filter(
	userMeasurements *model.UserMeasurements,
	measurementIDs []int) *model.UserMeasurements {

	var um model.UserMeasurements
	um.ID = userMeasurements.ID

	for _, givenID := range measurementIDs {
		for _, m := range userMeasurements.Measurements {
			if m.ID == givenID {
				um.Measurements = append(um.Measurements, m)
			}
		}
	}

	return &um
}
