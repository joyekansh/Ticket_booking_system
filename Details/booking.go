package details

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func BookTheSeat(w http.ResponseWriter, r *http.Request) {
	eventID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid event id", http.StatusBadRequest)
		return
	}
	seatID := r.PathValue("seat")
	if seatID == "" {
		http.Error(w, "seat id required", http.StatusBadRequest)
		return
	}

	bookMu.Lock()
	defer bookMu.Unlock()

	// find the event (Events/Event/Seat come from event_service.go — same package)
	var target *Event
	for i := range Events {
		if Events[i].ID == eventID {
			target = &Events[i]
			break
		}
	}
	if target == nil {
		http.Error(w, "event not found", http.StatusNotFound)
		return
	}
	var seat *Seat
	for i := range target.Seats {
		if target.Seats[i].ID == seatID {
			seat = &target.Seats[i]
			break
		}
	}
	if seat == nil {
		http.Error(w, "seat not found", http.StatusBadRequest)
		return
	}
	if seat.Booked {
		http.Error(w, "seat already booked", http.StatusConflict)
		return
	}

	// hold the seat BEFORE charging — so a second request can't grab it mid-payment
	seat.Booked = true
	if err := PaymentGateway(); err != nil {
		seat.Booked = false // payment failed — release the hold
		http.Error(w, "payment failed", http.StatusPaymentRequired)
		return
	}

	id := fmt.Sprintf("bk-%d", nextBookingID)
	nextBookingID++

	b := &Booking{
		ID:        id,
		EventID:   eventID,
		SeatID:    seatID,
		Status:    "CONFIRMED",
		CreatedAt: time.Now(),
	}
	bookings[id] = b

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}

func PaymentGateway() {

}
