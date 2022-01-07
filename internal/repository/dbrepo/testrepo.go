package dbrepo

import (
	"errors"
	"log"
	"time"

	"github.com/Schattenbrot/bookings/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database, fails if roomID == 3
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	if res.RoomID == 3 {
		return 0, errors.New("insertReservation failed")
	}
	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into the database, fails if roomID == 4
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	if r.RoomID == 4 {
		return errors.New("insertRoomRestriction failed")
	}
	return nil
}

// SearchAvailabilityByDatesByRoomID returns true if availability exists, and false if no availability exists
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	log.Println(roomID)
	if roomID == 0 {
		return false, errors.New("searchAvailabilityByDatesByRoomID failed")

	}
	return true, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms, if any, for given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

// GetRoomByID gets a room by its ID
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 || id == 0 {
		return room, errors.New("some error")
	}
	return room, nil
}

func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User
	return u, nil
}
func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}
func (m *testDBRepo) Authenticate(email, typedPassword string) (int, string, error) {
	return 1, "", nil
}
