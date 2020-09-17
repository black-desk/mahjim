package parser

import "net/url"
import "strconv"

type Style struct {
	Country string
	Color   string
	Scale   float64
}

var allow_country = []string{"cn", "jp"}
var allow_color = []string{"orange", "blue"}

func newStyle(query *url.Values) *Style {
	color := query.Get("color")
	if color == "" {
		color = "blue"
	}
	country := query.Get("country")
	if country == "" {
		country = "jp"
	}
	scale, err := strconv.ParseFloat(query.Get("scale"), 64)
	if err != nil {
		scale = 1
	}

	s := &Style{
		Color:   color,
		Country: country,
		Scale:   scale,
	}
	if is_legal(s) {
		return s
	} else {
		return nil
	}
}

func is_legal(s *Style) bool {
	return color_is_legal(s) && country_is_legal(s)
}

func color_is_legal(s *Style) bool {
	for _, c := range allow_color {
		if s.Color == c {
			return true
		}
	}
	return false
}

func country_is_legal(s *Style) bool {
	for _, c := range allow_country {
		if s.Country == c {
			return true
		}
	}
	return false
}
