package GUI

import (
	"fmt"
	paintColor "github.com/fatih/color"
	"image"
	"image/color"
	"math"
)

const(
	BOX = "█"
	)

var (

	COLORS = map[color.RGBA]paintColor.Attribute{
		/* A12830 - FF584F*/
		color.RGBA{R:0xFF, G:0x00, B:0x00, A:0xFF}:paintColor.FgRed, 	// 	Red Range
		color.RGBA{R:0x99, G:0x00, B:0x00, A:0xFF}:paintColor.FgHiRed,	// Red Range

		color.RGBA{R:0x00, G:0x00, B:0xFF, A:0xFF}:paintColor.FgBlue,
		color.RGBA{R:0x99, G:0xCC, B:0xFF, A:0xFF}:paintColor.FgHiBlue,

		color.RGBA{R:0xFF, G:0xFF, B:0x00, A:0xFF}:paintColor.FgHiYellow,//paintColor.FgHiGreen,
		color.RGBA{R:0xFF, G:0x99, B:0x33, A:0xFF}:paintColor.FgYellow, 		// Orange

		color.RGBA{R:0x00, G:0x00, B:0x00, A:0xFF}:paintColor.FgBlack,
		color.RGBA{R:0xFF, G:0xFF, B:0xFF, A:0xFF}:paintColor.FgHiWhite,

		color.RGBA{R:0x00, G:0xFF, B:0xFF, A:0xFF}:paintColor.FgCyan,
		color.RGBA{R:0x00, G:0xFF, B:0xFF, A:0xFF}:paintColor.FgHiCyan,

		color.RGBA{R:0x99, G:0x00, B:0x4c, A:0xFF}:paintColor.FgYellow,		// Purple M
		color.RGBA{R:0x99, G:0x00, B:0x99, A:0xFF}:paintColor.FgHiMagenta,		// Magenta M

//		color.RGBA{R:0x20, G:0x20, B:0x20, A:0xFF}:paintColor.FgHiBlack,		// Gray M
//		color.RGBA{R:0xE0, G:0xE0, B:0xE0, A:0xFF}:paintColor.FgWhite,			// light gray
	}
)

func init() {
	paintColor.New(paintColor.FgRed).Println("RED",BOX)
	paintColor.New(paintColor.FgBlue).Println("BLUE",BOX)
	paintColor.New(paintColor.FgHiGreen).Println("GREEN",BOX)
	paintColor.New(paintColor.FgHiYellow).Println("YELLOW",BOX)
	paintColor.New(paintColor.FgBlack).Println("BLACK",BOX)
	paintColor.New(paintColor.FgHiWhite).Println("WHITE",BOX)
	paintColor.New(paintColor.FgCyan).Println("CYAN",BOX)
	paintColor.New(paintColor.FgMagenta).Println("PURPLE",BOX)
	paintColor.New(paintColor.FgHiMagenta).Println("MAGENTA",BOX)
	paintColor.New(paintColor.FgHiBlack).Println("GRAY",BOX)
	paintColor.New(paintColor.FgYellow).Println("ORANGE",BOX)
	paintColor.New(paintColor.FgHiCyan).Println("LIGHT BLUE",BOX)
	paintColor.New(paintColor.FgWhite).Println("LIGHT GRAY",BOX)
	paintColor.New(paintColor.FgHiRed).Println("LIGHT RED",BOX)
	paintColor.New(paintColor.FgGreen).Println("DARKER GREEN",BOX)
}

func PaintMeme(img image.Image) {


	/*
	xStep := img.Bounds().Max.X / 50
	yStep := img.Bounds().Max.Y / 50
	for y := 0;y < img.Bounds().Max.Y; y += yStep {
		for x := 0; x < img.Bounds().Max.X; x += xStep {
			var rSum, bSum, gSum, aSum uint32
			var nR,nG,nB, nA uint32
			for i := 0; i <= xStep; i++ {
				nR, nG, nB, nA = img.At(x+xStep-1, y).RGBA()
				rSum += nR
				gSum += nG
				bSum += nB
				aSum += nA
			}

			var diff uint32 = 0xFFFFFFFF
			var c paintColor.Attribute

			for color := range COLORS {
				cR,cG,cB,cA := color.RGBA()
				nDiff := uint32(math.Abs(float64(aSum/4) - float64(cA)) + math.Abs(float64(rSum/4) - float64(cR)) + math.Abs(float64(gSum/4) - float64(cG)) + math.Abs(float64(bSum/4) - float64(cB)))
				if diff > nDiff {
					diff = nDiff
					c = COLORS[color]
				}
			}
			paintColor.New(c).Print(BOX)
		}
		fmt.Print("\n")

	}

*/

	for y :=0;y<img.Bounds().Max.Y;y+=img.Bounds().Bounds().Max.Y/140 {
		fmt.Print("\r")
		for x := 0; x<img.Bounds().Max.X; x+=img.Bounds().Max.X/140 {
			var diff uint32 = 0xFFFFFFFF
			r,g,b,a := img.At(x,y).RGBA()
			var c paintColor.Attribute

			for color := range COLORS {
				cR,cG,cB,cA := color.RGBA()
				nDiff := uint32(math.Abs(float64(a) - float64(cA)) + math.Abs(float64(r) - float64(cR)) + math.Abs(float64(g) - float64(cG)) + math.Abs(float64(b) - float64(cB)))
				if diff > nDiff {
					diff = nDiff
					c = COLORS[color]
				}
			}
			paintColor.New(c).Print(BOX)
		}
		fmt.Print("\n")
	}

}

