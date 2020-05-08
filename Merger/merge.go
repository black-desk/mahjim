package Merger

import (
	"image"
	"image/draw"
)

func Merge(imgs *[]*image.Image) image.Image {
	height := 0
	width := 0
	recs := []*image.Rectangle{}
	for _, img := range *imgs {
		height = Max(height, (*img).Bounds().Max.Y)
	}
	for _, img := range *imgs {
		sp := image.Point{width, height - (*img).Bounds().Max.Y}
		recs = append(recs, &image.Rectangle{sp, sp.Add((*img).Bounds().Size())})
		width+=(*img).Bounds().Max.X
	}
	rec := image.Rectangle{image.Point{0, 0}, (*recs[len(recs)-1]).Max}
	rgba := image.NewRGBA(rec)
	draw.Draw(rgba, rgba.Bounds(), image.Transparent, image.Point{0,0}, draw.Src)
	for i,img := range *imgs{
		draw.Draw(rgba, *recs[i], *img, image.Point{0, 0}, draw.Src)
	}
	return rgba
}
func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

