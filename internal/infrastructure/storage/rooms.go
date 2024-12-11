package storage

import (
	"sync"
	"time"
)

type RoomsStorage interface {
}

type RoomAvailability struct {
	HotelID     int
	RoomID      int
	Date        time.Time
	IsAvailable bool
	//{"reddison", "lux", date(2024, 1, 1), 1},
	//{"reddison", "lux", date(2024, 1, 2), 1},
	//{"reddison", "lux", date(2024, 1, 3), 1},
	//{"reddison", "lux", date(2024, 1, 4), 1},
	//{"reddison", "lux", date(2024, 1, 5), 0},
}

type Room struct {
	ID      int
	HotelID int
	Name    string
}

type roomsStorage struct {
	mu           *sync.RWMutex
	Availability []RoomAvailability
	Rooms        []Room
}

func NewRoomsStorage() RoomsStorage {
	return &roomsStorage{mu: &sync.RWMutex{}, Availability: []RoomAvailability{}, Rooms: []Room{}}
}

func (s *roomsStorage) CreateRoom(room Room) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Availability = append(s.Availability, room)

	return nil
}

func (s *roomsStorage) GetRoomByID(id int) (*Room, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Availability = append(s.Availability, room)

	return nil
}

func (s *roomsStorage) CreateRoomAvailability(room RoomAvailability) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Availability = append(s.Availability, room)

	return nil
}
