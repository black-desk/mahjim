package Parser

import (
	"errors"
	"image"
	"image/png"
	"os"
	"strconv"
)

type Parser struct {
	Imgs  *[]*image.Image
	lexer *Lexer
	look  *Word
}

type Pair [2]int

func NewParser(sources []string) Parser {
	return Parser{
		Imgs:  &[]*image.Image{},
		lexer: NewLexer(sources),
		look:  nil,
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

func (p *Parser) Parse() error {
	err := p.move()
	if err != nil {
		return err
	}
	err = p.majs()
	if err != nil {
		return err
	}
	return nil
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
}

var pos2string = map[Position]string{
	normal:  "",
	rotated: "_",
	double:  "^",
}

func (p *nodeP) toString() string {
	if p.num >= 0 {
		return pos2string[p.pos] + strconv.Itoa(p.num) + string(p.class)
	} else {
		return pos2string[p.pos] + string(p.class)
	}
}

func (p *nodeP) getImg(l *Lexer) (*image.Image, error) {
	reader, err := os.Open("files/" + l.style.toString() + p.toString() + ".png")
	if err != nil {
		return nil, err
	}
	img, err := png.Decode(reader)
	return &img, err
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
		pos: pos,
		num: num,
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

func (n nodeGroup) output(imgs *[]*image.Image, l *Lexer) error {
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
	err := node.output(p.Imgs, p.lexer)
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
			*ps = append(*ps, &nodeP{pos: normal, num: 1})
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
