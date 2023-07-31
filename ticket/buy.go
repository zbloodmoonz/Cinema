package ticket

import "fmt"

func init() {
	fmt.Println("init: ticket")
}

func BuyTicket(movie string) {
	fmt.Printf("Tickets for %s\n purchaced successfuly", movie)
}
