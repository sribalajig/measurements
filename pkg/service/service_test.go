package service

import (
	"testing"

	"reflect"

	"github.com/hpi/measurement/pkg/model"
)

// Test whether merging two structs works correctly
func TestMerge(t *testing.T) {
	ms := NewMeasurementService(NewProviderStub(nil))

	existing := []model.Measurement{
		{
			ID:    1,
			Value: 3.0,
		},
		{
			ID:    3,
			Value: 3.0,
		},
	}

	updated := []model.Measurement{
		{
			ID:    1,
			Value: 2.0,
		},
		{
			ID:    2,
			Value: 3.0,
		},
	}

	result := ms.merge(existing, updated)

	if len(result) != 3 {
		t.Fatalf("Expected %d, got %d", 3, len(result))
	}

	expected := []model.Measurement{
		{
			ID:    1,
			Value: 2.0,
		},
		{
			ID:    3,
			Value: 3.0,
		},
		{
			ID:    2,
			Value: 3.0,
		},
	}

	eq := reflect.DeepEqual(expected, result)

	if !eq {
		t.Fatalf("Failed Expected %v, Got %v", expected, result)
	}
}

// Test if create calls the right methods in data provider
func TestCreate(t *testing.T) {
	ps := NewProviderStub(nil)
	ms := NewMeasurementService(ps)

	ms.Create(&model.UserMeasurements{})

	if ps.callsToCreate != 1 {
		t.Fatalf("Expected %d calls to create, got %d", 1, ps.callsToCreate)
	}

	if ps.callsToGet != 1 {
		t.Fatalf("Expected %d calls to get, got %d", 1, ps.callsToGet)
	}
}

// Test if update calls the right methods in data provider
func TestUpdate(t *testing.T) {
	ps := NewProviderStub(&model.UserMeasurements{ID: 1})
	ms := NewMeasurementService(ps)

	ms.Create(&model.UserMeasurements{})

	if ps.callsToCreate != 0 {
		t.Fatalf("Expected %d calls to create, got %d", 0, ps.callsToCreate)
	}

	if ps.callsToGet != 1 {
		t.Fatalf("Expected %d calls to get, got %d", 1, ps.callsToGet)
	}

	if ps.callsToUpdate != 1 {
		t.Fatalf("Expected %d calls to update, got %d", 1, ps.callsToUpdate)
	}
}

// Stub definition
type ProviderStub struct {
	callsToCreate int
	callsToUpdate int
	callsToGet    int
	existing      *model.UserMeasurements
}

func NewProviderStub(existing *model.UserMeasurements) *ProviderStub {
	return &ProviderStub{
		existing: existing,
	}
}

func (p *ProviderStub) Create(userMeasurements *model.UserMeasurements) error {
	p.callsToCreate++

	return nil
}

func (p *ProviderStub) GetByUserID(id int) *model.UserMeasurements {
	p.callsToGet++

	return p.existing
}

func (p *ProviderStub) Update(userMeasurements *model.UserMeasurements) error {
	p.callsToUpdate++

	return nil
}
