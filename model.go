package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type linear_model struct {
	gradient    float64
	y_intercept float64
}

// simple generator to generate a model.
func generate(my_model linear_model, number_of_values float64, variance float64) chan float64 {

	rand.Seed(time.Now().Unix())

	c := make(chan float64)

	go func() {
		y := my_model.y_intercept

		for k := 0.00; k < number_of_values; k++ {
			y = (my_model.gradient * k) + my_model.y_intercept
			y += (rand.Float64() * variance)
			c <- y
		}

		close(c)
	}()

	return c
}

func export(model [100]float64, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file
	defer f.Close()

	// create new buffer
	buffer := bufio.NewWriter(f)

	for _, line := range model {
		_, err := buffer.WriteString(fmt.Sprintf("%f\n", line))
		if err != nil {
			log.Fatal(err)
		}
	}

	// flush buffered data to the file
	if err := buffer.Flush(); err != nil {
		log.Fatal(err)
	}
}
