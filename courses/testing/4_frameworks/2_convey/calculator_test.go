package frameworks_convey

import (
	"testing"

	frameworks "github.com/minitauros/go-concurrency-training/courses/testing/4_frameworks"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Calculator_Sum(t *testing.T) {
	// Write assertions.
}

func Test_Calculator_Multiply(t *testing.T) {
	// Write assertions.
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
