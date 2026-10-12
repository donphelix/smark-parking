package parking

import "testing"

func TestNewParkingSpaceStartsAvailable(t *testing.T) {
	space, err := NewParkingSpace("A1", "LOT-1")
	if err != nil {
		t.Fatalf("NewParkingSpace() error = %v", err)
	}

	if got := space.Status(); got != SpaceAvailable {
		t.Errorf("Status() = %q, want %q", got, SpaceAvailable)
	}
}
