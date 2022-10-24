package redlight

import (
	"bytes"
	_ "embed"
	"fmt"
	"image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var boolToIntLookup = map[bool]int8{false: 0, true: 1}

var RedstoneLights [2]*ebiten.Image // 0 is off, 1 is on
var RedstoneLightSize int = 4

func init() {
	//go:embed DefaultOff.png
	var rawRedstoneLightOff []byte
	ti, err := png.Decode(bytes.NewBuffer(rawRedstoneLightOff))
	fmt.Println(err.Error())

	RedstoneLights[0] = ebiten.NewImageFromImage(ti)

	//go:embed DefaultOn.png
	var rawRedstoneLightOn []byte
	ti, err = png.Decode(bytes.NewBuffer(rawRedstoneLightOn))
	fmt.Println(err.Error())

	RedstoneLights[1] = ebiten.NewImageFromImage(ti)
}

func Render(img [][]bool) *ebiten.Image {
	tempimg := ebiten.NewImage(len(img[0])*RedstoneLightSize, len(img)*RedstoneLightSize)
	for y := 0; y < len(img[0]); y++ {
		var tmp *ebiten.DrawImageOptions
		tmp.GeoM.Translate(0, float64(y*RedstoneLightSize))
		for x := 0; x < len(img); x++ {
			tempimg.DrawImage(RedstoneLights[boolToIntLookup[img[y][x]]], tmp)
			tmp.GeoM.Translate(float64(RedstoneLightSize), 0)
		}
	}
	return tempimg
}
