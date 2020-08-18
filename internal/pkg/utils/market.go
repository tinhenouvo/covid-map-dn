package utils

import (
	"time"

	"github.com/google/uuid"
)

// Market front a few common objects so that they can be mock in tests
type Market interface {
	// equivalent to github.com/google/uuid.New().String()
	NewUUIDString() string

	// equivalent to time.Now()
	TimeNow() time.Time
}

// NewMarket returns a real Market instance
func NewMarket() Market {
	return &market{}
}

type market struct{}

func (m *market) NewUUIDString() string {
	return uuid.New().String()
}

func (m *market) TimeNow() time.Time {
	return time.Now()
}

// MockMarket mock Market interface.
type MockMarket struct {
	UUIDString string
	Now        time.Time
}

// NewUUIDString returns field UUIDString
func (m *MockMarket) NewUUIDString() string {
	return m.UUIDString
}

// TimeNow returns field Now
func (m *MockMarket) TimeNow() time.Time {
	return m.Now
}
