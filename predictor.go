package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var i int
	fmt.Print("Type mileage: ")
	fmt.Scanf("%d", &i)
	price := predictPrice(i)
	fmt.Printf("Predicted price: %.2f\n", price)
}

func predictPrice(mileage int) float64 {
	theta0, theta1 := loadParameters()
	minMileage, maxMileage, minPrice, maxPrice := loadNormalizationParams()

	// Normalize the input
	normalizedMileage := (float64(mileage) - minMileage) / (maxMileage - minMileage)
	normalizedPrice := theta0 + theta1*normalizedMileage

	// Denormalize the output
	return normalizedPrice*(maxPrice-minPrice) + minPrice
}

func loadParameters() (float64, float64) {
	file, err := os.Open("model.txt")
	if err != nil {
		return 0.0, 0.0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	theta0, _ := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
	scanner.Scan()
	theta1, _ := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)

	return theta0, theta1
}

func loadNormalizationParams() (float64, float64, float64, float64) {
	file, err := os.Open("normalization.txt")
	if err != nil {
		return 0, 1, 0, 1 // defaults
	}
	defer file.Close()

	var minMileage, maxMileage, minPrice, maxPrice float64
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		minMileage, _ = strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
	}
	if scanner.Scan() {
		maxMileage, _ = strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
	}
	if scanner.Scan() {
		minPrice, _ = strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
	}
	if scanner.Scan() {
		maxPrice, _ = strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
	}
	return minMileage, maxMileage, minPrice, maxPrice
}
