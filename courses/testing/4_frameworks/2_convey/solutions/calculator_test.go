package solutions

import (
	"fmt"
	"testing"

	frameworks "github.com/minitauros/go-concurrency-training/courses/testing/4_frameworks"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Calculator_Sum(t *testing.T) {
	// Also with GoConvey we can use table driven tests.

	testCases := []struct {
		description string
		inputs      []int64
		expected    int64
	}{
		{
			description: "Works with numbers above zero",
			inputs:      []int64{1, 2, 5},
			expected:    8,
		},
		{
			description: "Works with numbers below zero",
			inputs:      []int64{-1, -2, -4},
			expected:    -7,
		},
	}

	Convey("Calculator.Sum()", t, func() {
		for i, tc := range testCases {
			Convey(fmt.Sprintf("%d: %s", i, tc.description), func() {
				calculator := &frameworks.Calculator{}
				res := calculator.Sum(tc.inputs...)
				So(res, ShouldEqual, tc.expected)
			})
		}
	})
}

func Test_Calculator_Multiply(t *testing.T) {
	testCases := []struct {
		description string
		start       int64
		multiplyBy  int64
		expected    int64
	}{
		{
			description: "Works with numbers above zero",
			start:       10,
			multiplyBy:  10,
			expected:    100,
		},
		{
			description: "Works with numbers below zero",
			start:       -10,
			multiplyBy:  8,
			expected:    -80,
		},
	}

	Convey("Calculator.Multiply()", t, func() {
		for i, tc := range testCases {
			Convey(fmt.Sprintf("%d: %s", i, tc.description), func() {
				calculator := &frameworks.Calculator{}
				res := calculator.Multiply(tc.start, tc.multiplyBy)
				So(res, ShouldEqual, tc.expected)
			})
		}
	})
}

func Test_Calculator_SpecialSub(t *testing.T) {
	Convey("Calculator.SpecialSub()", t, func() {
		calculator := &frameworks.Calculator{}
		var input, subtract int64

		Convey("If the input is above 0", func() {
			input = 10

			Convey("If the value to subtract is above 0", func() {
				subtract = 5

				Convey("Correctly subtracts", func() {
					So(calculator.SpecialSub(input, subtract), ShouldEqual, 5)
				})
			})

			Convey("If the value to subtract is below 0", func() {
				subtract = -5

				Convey("Correctly subtracts", func() {
					So(calculator.SpecialSub(input, subtract), ShouldEqual, 15)
				})
			})

			Convey("If the input is 100", func() {
				input = 100
				subtract = 5

				Convey("Correctly subtracts", func() {
					So(calculator.SpecialSub(input, subtract), ShouldEqual, 95)
				})
			})

			Convey("If the input is above 100", func() {
				input = 101
				subtract = 5

				Convey("Subtracts double the amount", func() {
					So(calculator.SpecialSub(input, subtract), ShouldEqual, 91)
				})
			})

			Convey("If the input is exactly 110", func() {
				input = 110
				subtract = 5

				Convey("Returns the answer to the ultimate question of life, the universe, and everything", func() {
					So(calculator.SpecialSub(input, subtract), ShouldEqual, 42)
				})

			})
		})
	})
}
