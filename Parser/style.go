package Parser

import "strings"

type Style struct {
	country string
	color string
}

func NewStyle(source string) *Style{
	var style Style
	if strings.Index(source,"cn")>=0{
		style.country="cn"
	}else{
		style.country="jp"
	}
	//if strings.Index(source,"B")>=0{
	//	style.color="B"
	//} else {
	//	style.color="O"
	//}
	return &style
}

func (s *Style) toString()string{
	return s.country
}
