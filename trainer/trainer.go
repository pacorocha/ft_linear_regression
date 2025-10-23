package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	train()
}

func train() {
	data := loadData("data.csv")

	// Normalize data
	data = normalize(data)

	theta0 := 0.0
	theta1 := 0.0
	learningRate := 0.1
	iterations := 1000

	n := float64(len(data))

	for i := 0; i < iterations; i++ {
		sum0 := 0.0
		sum1 := 0.0

		for _, point := range data {
			prediction := theta0 + theta1*point.mileage
			sum0 += prediction - point.price
			sum1 += (prediction - point.price) * point.mileage
		}

		theta0 -= learningRate * sum0 / n
		theta1 -= learningRate * sum1 / n
	}

	saveParameters(theta0, theta1)
	fmt.Printf("Training complete.\nTheta0: %.6f\nTheta1: %.6f\n", theta0, theta1)
}

type DataPoint struct {
	mileage float64
	price   float64
}

func loadData(filename string) []DataPoint {
	file, _ := os.Open(filename)
	defer file.Close()

	var data []DataPoint
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		mileage, _ := strconv.ParseFloat(parts[0], 64)
		price, _ := strconv.ParseFloat(parts[1], 64)
		data = append(data, DataPoint{mileage, price})
	}
	return data
}

func normalize(data []DataPoint) []DataPoint {
	var minMileage, maxMileage, minPrice, maxPrice float64
	minMileage = math.MaxFloat64
	minPrice = math.MaxFloat64

	for _, point := range data {
		if point.mileage < minMileage {
			minMileage = point.mileage
		}
		if point.mileage > maxMileage {
			maxMileage = point.mileage
		}
		if point.price < minPrice {
			minPrice = point.price
		}
		if point.price > maxPrice {
			maxPrice = point.price
		}
	}

	for i := range data {
		data[i].mileage = (data[i].mileage - minMileage) / (maxMileage - minMileage)
		data[i].price = (data[i].price - minPrice) / (maxPrice - minPrice)
	}
	saveNormalizationParams(minMileage, maxMileage, minPrice, maxPrice)

	return data
}

func saveParameters(theta0, theta1 float64) {
	file, _ := os.Create("model.txt")
	defer file.Close()
	fmt.Fprintf(file, "%.6f\n%.6f\n", theta0, theta1)
}

func saveNormalizationParams(minMileage, maxMileage, minPrice, maxPrice float64) {
	file, _ := os.Create("normalization.txt")
	defer file.Close()
	fmt.Fprintf(file, "%.6f\n%.6f\n%.6f\n%.6f\n", minMileage, maxMileage, minPrice, maxPrice)
}
