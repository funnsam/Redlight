package redlight

import (
	"bytes"
	_ "embed"
	"fmt"
	"image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed DefaultOff.png
var rawRedstoneLightOff []byte

//go:embed DefaultOn.png
var rawRedstoneLightOn []byte

var boolToIntLookup = map[bool]int8{false: 0, true: 1}

var RedstoneLights [2]*ebiten.Image // 0 is off, 1 is on
var RedstoneLightSize int = 4

func init() {
	tm1 := bytes.NewReader(rawRedstoneLightOff)
	tm2 := bytes.NewReader(rawRedstoneLightOn)

	ti1, err := png.Decode(tm1)
	fmt.Println(err.Error())

	ti2, err := png.Decode(tm2)
	fmt.Println(err.Error())

	RedstoneLights[0] = ebiten.NewImageFromImage(ti1)
	RedstoneLights[1] = ebiten.NewImageFromImage(ti2)
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
