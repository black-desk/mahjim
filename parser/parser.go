package parser

import (
	"errors"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"image/png"
	"net/url"
	"os"
	"strconv"
)

var maj_height int = 100
var maj_width int = 70

type Parser struct {
	imgs  *[]image.Image
	lexer *Lexer
	look  *Word
	style *Style
}

type Pair [2]int

func GetParser(maj_string *string, maj_style_config *url.Values) *Parser {
	return newParser(maj_string, maj_style_config) // TODO parser 重复初始化?多线程? 已经解析过的字符串不解析第二次
}

func newParser(maj_string *string, maj_style_config *url.Values) *Parser {
	return &Parser{
		imgs:  &[]image.Image{},
		lexer: newLexer(maj_string),
		look:  nil,
		style: newStyle(maj_style_config),
	}
}

func (p *Parser) move() error {
	look, err := p.lexer.scan()
	p.look = look
	return err
}

func (p *Parser) notfound(object string) {
	panic("expect a " + object + ", but found a " + string(p.look.text))
}

func (p *Parser) Parse() (*[]image.Image, error) {
	err := p.move()
	if err != nil {
		return nil, err
	}
	err = p.majs()
	if err != nil {
		return nil, err
	}
	return p.imgs, nil
}

type Position uint8

const (
	normal  Position = 0
	rotated Position = 1
	double  Position = 2
)

type nodeP struct {
	pos   Position
	num   int //-1 means is a flower
	class nodeClass
	style *Style
}

var pos2string = map[Position]string{
	normal:  "",
	rotated: "_",
	double:  "^",
}

func (p *nodeP) getFileName() string {
	if p.num >= 0 {
		return strconv.Itoa(p.num) + string(p.class)
	} else {
		if p.class != "+" {
			return string(p.class)
		} else {
			return p.style.Color
		}
	}
}

func (p *nodeP) getImg(l *Lexer) (image.Image, error) {
	if p.class == "|" {
		return image.NewNRGBA(image.Rect(0, 0, int(float64(maj_width/10)*p.style.Scale), int(float64(maj_height)*p.style.Scale))), nil
	}
	file, err := os.Open("files/" + p.getFileName() + ".png") // cache
	if err != nil {
		file, err = os.Open("files/" + p.getFileName() + p.style.Country + ".png") // cache
	}
	src, err := png.Decode(file)
	var res image.Image
	switch p.pos {
	case rotated:
		res = imaging.Rotate90(src)
	case normal:
		res = src
	case double:
		res = imaging.New(src.Bounds().Dy(), src.Bounds().Dx()*2, color.Black)
		rotatedSrc := imaging.Rotate90(src)
		res = imaging.Paste(res, rotatedSrc, image.Point{0, 0})
		res = imaging.Paste(res, rotatedSrc, image.Point{0, src.Bounds().Dx()})
	}
	res = imaging.Thumbnail(res, int(float64(res.Bounds().Dx())*p.style.Scale), int(float64(res.Bounds().Dy())*p.style.Scale), imaging.Box)
	return res, err
}

func (p *Parser) p() (*nodeP, error) {
	pos, err := p.pre()
	if err != nil {
		return nil, err
	}
	num, err := p.num()
	if err != nil {
		return nil, err
	}
	node := nodeP{
		pos:   pos,
		num:   num,
		style: p.style,
	}
	return &node, nil
}

type nodeClass string

func (p *Parser) majs() error {
	if p.look.tag != End {
		err := p.group()
		if err != nil {
			return err
		}
		err = p.majs()
		if err != nil {
			return err
		}
	}
	return nil
}

type nodeGroup struct {
	ps []*nodeP
}

func (n nodeGroup) output(imgs *[]image.Image, l *Lexer) error {
	for _, p := range n.ps {
		img, err := p.getImg(l)
		if err != nil {
			return err
		}
		(*imgs) = append((*imgs), img)
	}
	return nil
}

func (p *Parser) group() error {
	var node nodeGroup
	for p.look.tag == Pre || p.look.tag == Num {
		tmp, err := p.p()
		if err != nil {
			return err
		}
		node.ps = append(node.ps, tmp)
	}
	switch p.look.tag {
	case F:
		err := p.f(&node.ps)
		if err != nil {
			return err
		}
	case Class:
		err := p.class(&node.ps)
		if err != nil {
			return err
		}
	}
	err := node.output(p.imgs, p.lexer)
	return err
}

func (p *Parser) pre() (Position, error) {
	if p.look.tag == Pre {
		var res Position
		switch string(p.look.text) {
		case "_":
			res = rotated
		case "^":
			res = double
		}
		err := p.move()
		if err != nil {
			return normal, err
		}
		return res, nil
	}
	return normal, nil
}

func (p *Parser) num() (int, error) {
	if p.look.tag == Num {
		res, _ := strconv.Atoi(string(p.look.text))
		err := p.move()
		if err != nil {
			return -10, err
		}
		return res, nil
	} else {
		return 1, nil
	}
}
func (p *Parser) class(ps *[]*nodeP) error {
	if p.look.tag == Class {
		for _, ap := range *ps {
			ap.class = nodeClass(p.look.text)
		}
		err := p.move()
		return err
	} else {
		return errors.New("expect a class but found " + string(p.look.text))
	}
}

func (p *Parser) f(ps *[]*nodeP) error {
	var newps = []*nodeP{}
	if p.look.tag == F {
		switch len(*ps) {
		case 0:
			*ps = append(*ps, &nodeP{pos: normal, num: 1, style: p.style})
		case 1:
			break
		default:
			return errors.New("expect a class but found " + string(p.look.text))
		}
		i := (*(*ps)[0]).num
		(*(*ps)[0]).num = -1
		(*(*ps)[0]).class = nodeClass(p.look.text)
		for ; i > 0; i-- {
			newps = append(newps, (*ps)[0])
		}
		*ps = newps
		err := p.move()
		return err
	} else {
		return errors.New("expect a flower but found " + string(p.look.text))
	}
}
