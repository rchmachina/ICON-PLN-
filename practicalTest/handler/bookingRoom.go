package handler

import (
	"encoding/json"
	"net/http"
	"practicalTest/model"
	"sort"

	"github.com/labstack/echo/v4"
)

func GetBookings(c echo.Context) error {

	urlConsumption := "https://6686cb5583c983911b03a7f3.mockapi.io/api/dummy-data/masterJenisKonsumsi"

	respCons, err := http.Get(urlConsumption)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer respCons.Body.Close()

	// Decode the JSON response into a slice of Consumption structs
	var consumptions []model.ConsumptionData
	if err := json.NewDecoder(respCons.Body).Decode(&consumptions); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	url := "https://66876cc30bc7155dc017a662.mockapi.io/api/dummy-data/bookingList"

	resp, err := http.Get(url)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer resp.Body.Close()

	var bookings []model.Booking
	if err := json.NewDecoder(resp.Body).Decode(&bookings); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Sort the bookings based on BookingDate
	sort.Slice(bookings, func(i, j int) bool {
		return bookings[i].BookingDate.Before(bookings[j].BookingDate)
	})

	// Organize the data by month
	monthlyBookings := make(model.MonthlyBookings)
	for _, booking := range bookings {
		month := booking.BookingDate.Format("January 2006")
		booking.ParticipationAttendance = (float64(booking.Participants) / 70) * 100
		var totalPrice = 0
		for i := range booking.ListConsumption {
			for _, consumptionData := range consumptions {
				if booking.ListConsumption[i].Name == consumptionData.Name {
					if booking.ListConsumption[i].Name == "Snack Siang" {
						booking.ListConsumption[i].TotalItems = 140
						totalPrice += consumptionData.MaxPrice * 140
						break
					}
					if booking.ListConsumption[i].Name == "Makan Siang" {
						booking.ListConsumption[i].TotalItems = 280
						totalPrice += consumptionData.MaxPrice * 280
						break
					}
					if booking.ListConsumption[i].Name == "Snack Sore" {
						booking.ListConsumption[i].TotalItems = 140
						totalPrice += consumptionData.MaxPrice * 140
						break
					}
				}
			}
		}

		booking.TotalPriceConsumption = totalPrice
		monthlyBookings[month] = append(monthlyBookings[month], booking)
	}

	var result []model.BookingList
	for month, datas := range monthlyBookings {
		result = append(result, model.BookingList{
			Month: month,
			Data:  datas,
		})
	}

	return c.JSON(http.StatusOK, result)
}

func GetConsumptions(c echo.Context) error {
	url := "https://6686cb5583c983911b03a7f3.mockapi.io/api/dummy-data/masterJenisKonsumsi"

	resp, err := http.Get(url)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer resp.Body.Close()

	// Decode the JSON response into a slice of Consumption structs
	var consumptions []model.ConsumptionData
	if err := json.NewDecoder(resp.Body).Decode(&consumptions); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Optionally: Process or modify the data if needed

	return c.JSON(http.StatusOK, consumptions)
}
