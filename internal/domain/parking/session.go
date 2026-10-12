package parking

import "time"

type ParkingSession struct {
	id        string
	lotID     string
	spaceID   string
	vehicle   Vehicle
	entryTime time.Time
	exitTime  *time.Time
}

func NewParkingSession(
	id, lotID, spaceID string,
	vehicle Vehicle,
	entryTime time.Time,
) (ParkingSession, error) {
	if id == "" || lotID == "" || spaceID == "" {
		return ParkingSession{}, ErrInvalidID
	}

	if entryTime.IsZero() {
		return ParkingSession{}, ErrInvalidTimestamp
	}

	return ParkingSession{
		id:        id,
		lotID:     lotID,
		spaceID:   spaceID,
		vehicle:   vehicle,
		entryTime: entryTime,
	}, nil
}

func (s ParkingSession) IsActive() bool {
	return s.exitTime == nil
}

func (s ParkingSession) EntryTime() time.Time {
	return s.entryTime
}

func (s ParkingSession) ExitTime() *time.Time {
	if s.exitTime == nil {
		return nil
	}

	exit := *s.exitTime
	return &exit
}

func (s *ParkingSession) Close(exitTime time.Time) error {
	if !s.IsActive() {
		return ErrSessionClosed
	}

	if exitTime.IsZero() || exitTime.Before(s.entryTime) {
		return ErrInvalidTimestamp
	}

	s.exitTime = &exitTime
	return nil
}
