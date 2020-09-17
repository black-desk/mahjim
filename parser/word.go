package parser

import (
	"errors"
	"strings"
)

type Word struct {
	text rune
	tag  Tag
}

var wordTable = map[rune]*Word{}

func getWord(s rune) (*Word, error) {
	if w, exist := wordTable[s]; exist == true {
		return w, nil
	} else {
		return nil, errors.New("unrecognized character" + string(s))
	}
}
func SetWord(s rune, tag Tag) {
	wordTable[s] = &Word{
		text: s,
		tag:  tag,
	}
}
func init() {
	wordLists := map[Tag]string{
		Pre:   "_,^",
		Num:   "0,1,2,3,4,5,6,7,8,9",
		Class: "p,s,m,z,n",
		F:     "中,发,白,春,夏,秋,冬,东,南,西,北,梅,兰,竹,菊,|,+",
		End:   "$",
	}
	for tag, wordList := range wordLists {
		words := strings.Split(wordList, ",")
		for _, word := range words {
			SetWord([]rune(word)[0], tag)
		}
	}
}
