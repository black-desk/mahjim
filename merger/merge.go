package merger

import (
	"image"
	"image/draw"
)

func Merge(array *[][]image.Image) image.Image {
	height := 0
	recs := []*image.Rectangle{}
	max_width := 0
	for _, row := range *array {
		_height := 0
		width := 0
		for _, img := range row {
			_height = Max(_height, img.Bounds().Max.Y)
		}
		height += _height
		for _, img := range row {
			sp := image.Point{width, height - img.Bounds().Dy()}
			recs = append(recs, &image.Rectangle{sp, sp.Add(img.Bounds().Size())})
			width += img.Bounds().Dx()
			max_width = Max(max_width, width)
		}
	}
	res := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{max_width, height}})
	draw.Draw(res, res.Bounds(), image.Transparent, image.Point{0, 0}, draw.Src)
	cnt := 0
	for _, row := range *array {
		for _, img := range row {
			draw.Draw(res, *recs[cnt], img, image.Point{0, 0}, draw.Src)
			cnt++
		}
	}
	return res
}
func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
