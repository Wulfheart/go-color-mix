package colormix

import (
	"fmt"
	"image/color"
	"testing"
)

type testFraction struct {
	R      int
	G      int
	B      int
	Amount int
}

func (t testFraction) colorize() color.Color {
	return color.RGBA{
		R: uint8(t.R),
		G: uint8(t.G),
		B: uint8(t.B),
		A: uint8(255),
	}
}

func (t testFraction) fractionize(total int) Fraction {
	fmt.Printf("%v \n", t)
	return Fraction{
		Color:  t.colorize(),
		Factor: float64(t.Amount) / float64(total),
	}
}

func compare(col1 color.Color, col2 color.Color) error {
	fmt.Println(col1)
	if col1 != col2 {
		return fmt.Errorf("the two colors don't match \n%v %v", col1, col2)
	}
	return nil
}

func TestColorCombine(t *testing.T) {
	for _, tc := range []struct {
		Input  []testFraction
		Output testFraction
	}{
		{ Input: []testFraction{
			{255, 0, 0, 1},
			{255, 165, 0, 1},
			{255, 255, 0, 1},
		},
		Output: testFraction{255, 140, 0, 0}},
	} {
		fmt.Println("Input:", tc.Input)
		sum := 0
		for _, s := range tc.Input {
			sum += s.Amount
		}
		var fractions []Fraction
		for _, f := range tc.Input {
			fractions = append(fractions, f.fractionize(sum))
		}
		result := tc.Output.colorize()

		c, err := Combine(fractions)
		if err != nil {
			t.Error(err)
		}
		err = compare(c, result)
		if err != nil {
			t.Error(err)
		}

	}
}
