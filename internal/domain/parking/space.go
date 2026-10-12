package parking

type SpaceStatus string

const (
	SpaceAvailable SpaceStatus = "AVAILABLE"
	SpaceOccupied  SpaceStatus = "OCCUPIED"
)

type ParkingSpace struct {
	id     string
	lotID  string
	status SpaceStatus
}

func NewParkingSpace(id, lotID string) (ParkingSpace, error) {
	if id == "" || lotID == "" {
		return ParkingSpace{}, ErrInvalidID
	}

	return ParkingSpace{
		id:     id,
		lotID:  lotID,
		status: SpaceAvailable,
	}, nil
}

func (s ParkingSpace) ID() string {
	return s.id
}

func (s ParkingSpace) LotID() string {
	return s.lotID
}

func (s ParkingSpace) Status() SpaceStatus {
	return s.status
}

func (s *ParkingSpace) Occupy() error {
	if s.status == SpaceOccupied {
		return ErrSpaceOccupied
	}

	s.status = SpaceOccupied
	return nil
}

func (s *ParkingSpace) Release() error {
	if s.status == SpaceAvailable {
		return ErrSpaceAvailable
	}

	s.status = SpaceAvailable
	return nil
}
