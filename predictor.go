package main

import "fmt"

func main() {
    var i int
	fmt.Scanf("Type mileage: %d", &i)
	price := predictPrice(i)
	fmt.Printf("Predicted price: %.2f\n", price)
}

func predictPrice(mileage int) float64 {
	theta0 := 0.0
	theta1 := 0.0
	return theta0 + theta1*float64(mileage)
}
