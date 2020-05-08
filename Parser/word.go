package Parser

import "strings"

type Word struct {
	text rune
	tag  Tag
}

var wordTable = map[rune]*Word{}

func GetWord(s rune) (*Word,error) {
	if w, exist := wordTable[s]; exist == true {
		return w,nil
	} else {
		SetWord(s,F)
		return wordTable[s],nil
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
		Class: "p,s,w,z,n",
		End:   "$",
	}
	for tag, wordList := range wordLists {
		words := strings.Split(wordList, ",")
		for _, word := range words {
			SetWord([]rune(word)[0], tag)
		}
	}
}
