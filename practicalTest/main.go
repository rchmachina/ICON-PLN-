package main

import (
	"log"
	"practicalTest/handler"


	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())


    api := e.Group("/api/V1")
    {
         
        api.GET("/summary-bookings",handler.GetBookings)
        api.GET("/consumption",handler.GetConsumptions)
    }

   
    log.Fatal(e.Start(":8080"))
}