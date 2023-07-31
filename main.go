package main

import (
	"fmt"

	"github.com/zbloodmoonz/Cinema/movie"
	"github.com/zbloodmoonz/Cinema/ticket"
)

func init() {
	fmt.Println("init: main")
}

func main() {
	movieID := movie.FindName("marv01")
	fmt.Println(movieID)
	movie.Review(movieID, 8.4)
	ticket.BuyTicket(movieID)
}
