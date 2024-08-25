package model

import "time"

type ConsumptionData struct {
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	MaxPrice  int       `json:"maxPrice"`
	ID        string    `json:"id"`
}

type ConsumptionBooking struct {
	Name       string `json:"name"`
	TotalItems int    `json:"totalItems"`
}

type Booking struct {
	ID                      string               `json:"id"`
	BookingDate             time.Time            `json:"bookingDate"`
	OfficeName              string               `json:"officeName"`
	ListConsumption         []ConsumptionBooking `json:"listConsumption"`
	TotalPriceConsumption   int                  `json:"totalPriceConsumption"`
	Participants            int                  `json:"participants"`
	RoomName                string               `json:"roomName"`
	ParticipationAttendance float64              `json:"participationAttendance"`
}

type BookingList struct {
	Month string    `json:"month"`
	Data  []Booking `json:"data"`
}

type MonthlyBookings map[string][]Booking
