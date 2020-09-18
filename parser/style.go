package parser

import "net/url"
import "strconv"

type Style struct {
	Country string
	Color   string
	Scale   float64
	River   bool
}

var allow_country = []string{"cn", "jp"}
var allow_color = []string{"orange", "blue"}

func newStyle(query *url.Values) *Style {
	color := query.Get("color")
	if color != "orange" {
		color = "blue"
	}
	country := query.Get("country")
	if country != "cn" {
		country = "jp"
	}
	scale, err := strconv.ParseFloat(query.Get("scale"), 64)
	if err != nil {
		scale = 1
	}
	if scale > 10 {
		scale = 10
	}
	river := query.Get("river") == "true"

	s := &Style{
		Color:   color,
		Country: country,
		Scale:   scale,
		River:   river,
	}
	return s
}

func (s *Style) str() string {
	return "Cou:" + s.Country + "#Col:" + s.Color + "#S:" + strconv.FormatFloat(s.Scale, 'E', -1, 64) + "#R:" + strconv.FormatBool(s.River)
}
