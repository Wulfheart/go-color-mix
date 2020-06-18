package colormix

import (
	"fmt"
	"image/color"
	"math"
)

type Fraction struct {
	Color color.Color
	Factor float64
}

// Combine joins some colors
func Combine(fra []Fraction)(color.Color, error){
	// Check if total factor equals 1
	var totalFraction, rTotal, gTotal, bTotal float64
	for _, p := range fra {
		totalFraction += p.Factor
		r, g, b, a := p.Color.RGBA()
		if uint8(a) != 255 {
			return nil, fmt.Errorf("currently the package colormix does only support rgba colors with an opacity of 100%%")
		}
		rTotal += float64(uint8(r)) * p.Factor
		gTotal += float64(uint8(g)) * p.Factor
		bTotal += float64(uint8(b)) * p.Factor

	}
	if totalFraction <= 0.95 || totalFraction >= 1.05 {
		return nil, fmt.Errorf("expecting the sum of all factors to be exactly 1 ± 0.05, it is %f", totalFraction)
	}
	return color.RGBA{
		R: uint8(math.Round(rTotal)),
		G: uint8(math.Round(gTotal)),
		B: uint8(math.Round(bTotal)),
		A: uint8(255),
	}, nil
}