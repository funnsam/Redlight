package redlight

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var boolToIntLookup = map[bool]int8{false: 0, true: 1}

var RedstoneLights [2]*ebiten.Image
var RedstoneLightSize int = 4

func Render(img [][]bool) *ebiten.Image {
	tempimg := ebiten.NewImage(len(img[0])*RedstoneLightSize, len(img)*RedstoneLightSize)
	for y := 0; y < len(img[0]); y++ {
		tmp := &ebiten.DrawImageOptions{}
		tmp.GeoM.Translate(0, float64(y*RedstoneLightSize))
		for x := 0; x < len(img); x++ {
			tempimg.DrawImage(RedstoneLights[boolToIntLookup[img[y][x]]], tmp)
			tmp.GeoM.Translate(float64(RedstoneLightSize), 0)
		}
	}
	return tempimg
}
