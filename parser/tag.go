package parser

type Tag uint8

const (
	F     Tag = 0
	Pre   Tag = 1
	Num   Tag = 2
	Class Tag = 3
	End   Tag = 4
)
