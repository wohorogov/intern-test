package main

import (
	"os"
	"strconv"
	c "app/coordinates"
)

func main() {
	var args = os.Args
	if len(args) > 1 {
		ancientCoord, _ := strconv.Atoi(args[1])
		ulamPos := c.Sieve(ancientCoord)
		println(ulamPos)
		if ulamPos == -1 {
			println("Given coordinate is prime number!")
		} else {
			x, y := c.UlamToCartesian(ulamPos)
			println(x, y)
		}
	} else {
		println("Please put ancient coordinate as first parameter!")
	}
}
