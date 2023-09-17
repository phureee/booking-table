package entities

import (
	"sort"
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

type showBook struct {
	BookingID      int       `json:"booking_id"`
	TableID        []int     `json:"table_id"`
	CustomerAmount int       `json:"customer_amount"`
	IsAvailable    bool      `json:"available"`
	CreatedAt      time.Time `json:"create_time"`
}

func ShowBook() []showBook {
	var showB []showBook
	for bID, b := range book {
		showB = append(showB, showBook{
			BookingID:      bID,
			TableID:        b.TableID,
			CustomerAmount: b.CustomerAmount,
			IsAvailable:    b.IsAvailable,
			CreatedAt:      b.CreatedAt,
		})
	}

	sort.Slice(showB, func(i, j int) bool {
		return showB[i].BookingID < showB[j].BookingID
	})

	return showB
}
