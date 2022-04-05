package pkg

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/rand"
	"sort"
)

func generateFunc() func(x, y int) int {
	k := make([]float64, 0, 5)
	k1 := make([]float64, 0, 11)
	lineK := make([]float64, 0, 2)
	for i := 0; i < cap(k); i++ {
		k = append(k, rand.Float64()*10-5)
	}

	for i := 0; i < cap(k1); i++ {
		k1 = append(k1, rand.Float64()*5)
	}

	for i := 0; i < cap(lineK); i++ {
		lineK = append(lineK, rand.Float64()*10-5)
	}

	zoom := rand.NormFloat64()*50 + 75

	return func(x, y int) int {
		fX, fY := float64(x)/zoom, float64(y)/zoom
		return int(100 * math.Abs(lineK[0]*(k1[9]*fX+k1[10]*fY)+k[1]*math.Cos(k1[1]*fX+k1[2]*fY)+k[2]*math.Sin(k1[3]*fX+k1[4]*fY)+k[3]*math.Cos(k1[5]*fX+k1[6]*fY)+k[4]*math.Sin(k1[7]*fX+k1[8]*fY)))
	}
}

// colorWithLinearPosition is struct which contains colour and it's corresponding position on the [0,255] section.
type colorWithLinearPosition struct {
	color.RGBA
	pos int
}

func generateColourFunc(pointAmount int) func(i int) color.RGBA {
	clrs := make([]colorWithLinearPosition, 0, pointAmount+1)
	clrs = append(clrs, colorWithLinearPosition{
		RGBA: color.RGBA{0, 0, 0, 255},
		pos:  0,
	})
	step := 256 / (pointAmount + 1)
	for i := 0; i < pointAmount; i++ {
		clrs = append(clrs, colorWithLinearPosition{
			RGBA: color.RGBA{
				uint8(rand.Intn(256)),
				uint8(rand.Intn(256)),
				uint8(rand.Intn(256)),
				uint8(rand.Intn(256)),
			},
			pos: (i + 1) * step,
		})
	}
	sort.Slice(clrs, func(i, j int) bool {
		return clrs[i].pos < clrs[j].pos
	})

	return func(v int) color.RGBA {
		index := pointAmount - 1
		for i := 0; i < pointAmount; i++ {
			if v < clrs[i].pos {
				index = i
				break
			}
		}
		clr1, clr2 := clrs[index-1], clrs[index]

		x := (v - clr2.pos) / (clr2.pos - clr1.pos)
		r := uint8(x)*(clr2.R-clr1.R) + clr2.R
		g := uint8(x)*(clr2.G-clr1.G) + clr2.G
		b := uint8(x)*(clr2.B-clr1.B) + clr2.B

		return color.RGBA{r, g, b, 255}
	}
}

func GeneratePic(width, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	f := generateFunc()
	colF := generateColourFunc(2)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			clr := colF(f(x, y))
			if x == y {
				fmt.Printf("x: %v, y: %v, colour: %v\n", x, y, clr.R)
			}
			img.Set(x, y, clr)
		}
	}

	return img
}
