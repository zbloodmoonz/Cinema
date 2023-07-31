package movie

import (
	"fmt"
)

func init() {
	fmt.Println("init: movie")
}

func Review(name string, rating float64) {
	fmt.Printf("You reviewed %s and it's rating is %0.1f\n", name, rating)
}
