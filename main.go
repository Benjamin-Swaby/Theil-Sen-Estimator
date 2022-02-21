package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	grad, _ := strconv.ParseFloat(os.Args[1], 64)
	yinter, _ := strconv.ParseFloat(os.Args[2], 64)
	variance, _ := strconv.ParseFloat(os.Args[3], 64)

	var my_model = linear_model{
		gradient:    grad,
		y_intercept: yinter,
	}

	values := generate(my_model, 100, variance)

	var arr_values [100]float64
	index := 0
	for i := range values {
		arr_values[index] = i
		index++
	}

	m := estimate_gradient(arr_values)
	fmt.Println("Estimated Gradient is: ")
	fmt.Println(m)

	c := estimate_intercept(m, arr_values)
	fmt.Println("Estimated Intercept is: ")
	fmt.Println(c)

	var estimated_model = linear_model{
		gradient:    m,
		y_intercept: c,
	}

	var arr_estimated [100]float64
	index = 0
	for i := range generate(estimated_model, 100, 0.0) {
		arr_estimated[index] = i
		index++
	}

	export(arr_values, fmt.Sprintf("y=%fx+%f", my_model.gradient, my_model.y_intercept))
	export(arr_estimated, fmt.Sprintf("y=%fx+%f", estimated_model.gradient, estimated_model.y_intercept))

}
