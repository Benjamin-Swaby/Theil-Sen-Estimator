package main

import "fmt"

func estimate_gradient(values [100]float64) float64 {

	total := 0.00

	for i := 0; i < len(values)-2; i++ {

		value := values[i]
		local_gradient := values[i+1] - value
		total += local_gradient

	}

	return (total / float64(len(values)-1))
}

func estimate_intercept(m float64, values [100]float64) float64 {
	median_index := int((len(values) / 2))
	fmt.Println(median_index)
	return values[median_index] - m*float64(median_index)
}
