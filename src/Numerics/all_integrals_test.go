package Numerics

import (
	"fmt"
	"testing"
	"time"
)

func TestAdaptiveBoole(t *testing.T) {

	start := time.Now()

	for i := 0; i < 100000; i++ {

		AdaptiveBoole(-5, 0)

	}

	end := time.Now()

	fmt.Println("Adaptive Boole Method: ", (end.Sub(start)))
	fmt.Println("Area: ", AdaptiveBoole(-5, 0))

}

func TestAdaptiveSimpson(t *testing.T) {

	start := time.Now()

	for i := 0; i < 100000; i++ {

		AdaptiveSimpson(-5, 0)

	}

	end := time.Now()

	fmt.Println("Adaptive Simpson Method: ", (end.Sub(start)))
	fmt.Println("Area: ", AdaptiveSimpson(-5, 0))

}

func TestAdaptiveTrapezoid(t *testing.T) {

	start := time.Now()

	for i := 0; i < 100000; i++ {

		AdaptiveTrapezoid(-5, 0)

	}

	end := time.Now()

	fmt.Println("Adaptive Trapezoid Method: ", (end.Sub(start)))
	fmt.Println("Area: ", AdaptiveTrapezoid(-5, 0))

}
