package main

import (
	"advent/aoc/common"
	"fmt"
)

type BlinkStoneScanner struct {
	input      string
	pos        int
	Eof        bool
	LastNumber string
}

type ScanToken uint8

const (
	Token_Zero ScanToken = iota
	Token_NonZero
	Token_Space
	Token_Other
	NrToken
)

func (sc *BlinkStoneScanner) init(st string) {
	sc.input = st
	sc.pos = 0
	sc.Eof = false
}

func (sc *BlinkStoneScanner) nextToken() ScanToken {
	if sc.pos >= len(sc.input) {
		sc.Eof = true
		return Token_Space
	}
	result := Token_Other
	switch sc.input[sc.pos] {
	case '0':
		sc.LastNumber = sc.LastNumber + "0"
		result = Token_Zero
	case ' ':
		result = Token_Space
	default:
		if sc.input[sc.pos] > '0' && sc.input[sc.pos] <= '9' {
			sc.LastNumber = sc.LastNumber + string(sc.input[sc.pos])
			result = Token_NonZero
		}
	}
	sc.pos++
	return result
}

func (sc *BlinkStoneScanner) consumeNumber() string {
	result := sc.LastNumber
	sc.LastNumber = ""
	return result
}

type ParseState uint8

const (
	State_Initial ParseState = iota
	State_Zero
	State_Even
	State_Odd
	NrStates
)

type Production struct {
	nextState ParseState
	action    func(p *BlinkStoneCompiler)
}

type BlinkStoneGenerator struct {
	output string
}

func (g *BlinkStoneGenerator) init() {
	g.output = ""
}

func (g *BlinkStoneGenerator) emitNumber(st string) {
	g.output = g.output + fmt.Sprintf("%d", common.StringToNum(st))
}

func (g *BlinkStoneGenerator) emitSpace() {
	g.output = g.output + " "
}

type BlinkStoneCompiler struct {
	state       ParseState
	productions [NrStates][NrToken]Production
	scanner     BlinkStoneScanner
	backend     BlinkStoneGenerator
}

func (p *BlinkStoneCompiler) init() {
	// error transitions not defined
	p.productions[State_Initial][Token_Zero] = Production{nextState: State_Zero}
	p.productions[State_Initial][Token_NonZero] = Production{nextState: State_Odd}
	p.productions[State_Initial][Token_Space] = Production{nextState: State_Initial}

	p.productions[State_Zero][Token_Space] = Production{nextState: State_Initial, action: func(p *BlinkStoneCompiler) {
		p.backend.emitNumber("1")
		p.scanner.consumeNumber()
		p.backend.emitSpace()
	}}

	p.productions[State_Even][Token_Zero] = Production{nextState: State_Odd}
	p.productions[State_Even][Token_NonZero] = Production{nextState: State_Odd}
	p.productions[State_Even][Token_Space] = Production{nextState: State_Initial, action: func(p *BlinkStoneCompiler) {
		lnr := p.scanner.consumeNumber()
		l := len(lnr)
		left := lnr[:l>>1]
		right := lnr[l>>1 : l]
		p.backend.emitNumber(left)
		p.backend.emitSpace()
		p.backend.emitNumber(right)
		p.backend.emitSpace()
	}}

	p.productions[State_Odd][Token_Zero] = Production{nextState: State_Even}
	p.productions[State_Odd][Token_NonZero] = Production{nextState: State_Even}
	p.productions[State_Odd][Token_Space] = Production{nextState: State_Initial, action: func(p *BlinkStoneCompiler) {
		value := common.StringToNum(p.scanner.consumeNumber()) * 2024
		str := fmt.Sprintf("%d", value)
		p.backend.emitNumber(str)
		p.backend.emitSpace()
	}}
}

func (p *BlinkStoneCompiler) start(st string) {
	p.state = State_Initial
	p.scanner.init(st)
	p.backend.init()
}

func (c *BlinkStoneCompiler) parse() {
	tk := c.scanner.nextToken()
	action := c.productions[c.state][tk]

	if action.action != nil {
		action.action(c)
	}
	c.state = action.nextState
}
