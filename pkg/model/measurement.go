package model

// Measurement - datastructure to hold a measurement
type Measurement struct {
	ID    int
	Type  string
	Value float64
	Units string
}

// UserMeasurements - measurement for a user
type UserMeasurements struct {
	ID           int
	Measurements []Measurement
}
