package parking

import "errors"

var (
	ErrInvalidID        = errors.New("invalid identifier")
	ErrInvalidVehicle   = errors.New("invalid vehicle registration")
	ErrSpaceOccupied    = errors.New("parking space is occupied")
	ErrSpaceAvailable   = errors.New("parking space is not occupied")
	ErrSessionClosed    = errors.New("parking session is already closed")
	ErrInvalidTimestamp = errors.New("invalid timestamp")
)
