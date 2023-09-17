package entities

import (
	"time"
)

type Booking struct {
	TableID        []int
	CustomerAmount int
	IsAvailable    bool
	CreatedAt      time.Time
}

type Booked struct {
	BookID int
}

var sequenceIDBooking int

// struct map[BookID]Booking
var book map[int]Booking

func SetBook(bo Booking) (bookingID int) {
	sequenceIDBooking++

	if book == nil {
		book = make(map[int]Booking)
	}

	book[sequenceIDBooking] = bo
	return sequenceIDBooking
}

func GetBook() map[int]Booking {
	return book
}

func UpdateBookingAvailable(bookingID int, isAvailable bool) (IsSuccess bool) {
	if booking, ok := book[bookingID]; ok {
		booking.IsAvailable = isAvailable
		book[bookingID] = booking
		return true
	}
	return false
}
