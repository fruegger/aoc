package main

type BlinkStoneScanner struct {
	input      string
	pos        int
	Eof        bool
	LastNumber uint64
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
		sc.LastNumber = sc.LastNumber * 10
		result = Token_Zero
	case ' ':
		result = Token_Space
	default:
		if sc.input[sc.pos] > '0' && sc.input[sc.pos] <= '9' {
			sc.LastNumber = sc.LastNumber*10 + uint64(sc.input[sc.pos]) - '0'
			result = Token_NonZero
		}
	}
	sc.pos++
	return result
}

func (sc *BlinkStoneScanner) consumeNumber() uint64 {
	result := sc.LastNumber
	sc.LastNumber = 0
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

var tenpow []uint64

func divisorHalf(n uint64) uint64 {
	var i uint
	i = 1
	for n > tenpow[i] {
		i++
	}
	return tenpow[i>>1]
}

func divisorFull(n uint64) uint64 {
	var i uint
	i = 0
	for n >= tenpow[i] {
		i++
	}
	return tenpow[i-1]
}

func initTenPow() {
	var tp uint64 = 1
	for i := 0; i < 20; i++ {
		tenpow = append(tenpow, tp)
		tp *= 10
	}
}

type Production struct {
	nextState ParseState
	action    func(p *BlinkStoneCompiler)
}

const ENOUGH = 1024 * 1024 * 24

type IntPair struct {
	left  uint64
	right uint64
}

type BlinkStoneGenerator struct {
	output     []uint8
	writePos   int
	evenValues map[uint64]IntPair
	oddValues  map[uint64]uint64
}

func (g *BlinkStoneGenerator) init() {
	g.evenValues = make(map[uint64]IntPair)
	g.oddValues = make(map[uint64]uint64)
	g.output = make([]uint8, ENOUGH)
}

func (g *BlinkStoneGenerator) start() {
	g.writePos = 0
}

func (g *BlinkStoneGenerator) emitEven(lnr uint64) {
	val, found := g.evenValues[lnr]
	if !found {
		lnr64 := lnr
		d := divisorHalf(lnr64)
		left := lnr64 / d
		right := lnr64 % d
		val = IntPair{left, right}
		g.evenValues[lnr] = val
	}
	g.emit(val.left)
	g.emit(val.right)
}

func (g *BlinkStoneGenerator) emitOdd(lnr uint64) {
	num, found := g.oddValues[lnr]
	if !found {
		num = lnr * 2024
		g.oddValues[lnr] = num
	}
	g.emit(num)
}

func (g *BlinkStoneGenerator) emitZero() {
	g.output[g.writePos] = '1'
	g.writePos++
	g.output[g.writePos] = ' '
	g.writePos++
}

func (g *BlinkStoneGenerator) emit(val uint64) {
	if val == 0 {
		g.output[g.writePos] = '0'
		g.writePos++
	} else {
		for d := divisorFull(val); d > 0; d = d / 10 {
			sym := uint8(val / d)
			g.output[g.writePos] = sym + '0'
			g.writePos++
			val = val % d
		}
	}
	g.output[g.writePos] = ' '
	g.writePos++
}

func (g BlinkStoneGenerator) getOutput() string {
	return string(g.output[:g.writePos])
}

type BlinkStoneCompiler struct {
	state   ParseState
	scanner BlinkStoneScanner
	backend BlinkStoneGenerator
	blinks  int
	stones  int
}

var productions [NrStates][NrToken]Production

func initProductions() {
	// error transitions not defined
	productions[State_Initial][Token_Zero] = Production{nextState: State_Zero}
	productions[State_Initial][Token_NonZero] = Production{nextState: State_Odd}
	productions[State_Initial][Token_Space] = Production{nextState: State_Initial}

	productions[State_Zero][Token_Space] = Production{nextState: State_Initial, action: func(p *BlinkStoneCompiler) {
		p.backend.emitZero()
		p.scanner.consumeNumber()
		if p.blinks > 0 {
			p.stones += p.recursiveDescent()
		} else {
			p.stones = 1
		}
	}}

	productions[State_Even][Token_Zero] = Production{nextState: State_Odd}
	productions[State_Even][Token_NonZero] = Production{nextState: State_Odd}
	productions[State_Even][Token_Space] = Production{nextState: State_Initial, action: func(p *BlinkStoneCompiler) {
		lnr := p.scanner.consumeNumber()
		p.backend.emitEven(lnr)
		if p.blinks > 0 {
			p.stones += p.recursiveDescent()
		} else {
			p.stones = 2
		}
	}}

	productions[State_Odd][Token_Zero] = Production{nextState: State_Even}
	productions[State_Odd][Token_NonZero] = Production{nextState: State_Even}
	productions[State_Odd][Token_Space] = Production{nextState: State_Initial, action: func(p *BlinkStoneCompiler) {
		lnr := p.scanner.consumeNumber()
		p.backend.emitOdd(lnr)
		if p.blinks > 0 {
			p.stones += p.recursiveDescent()
		} else {
			p.stones = 1
		}
	}}

}

func (c *BlinkStoneCompiler) recursiveDescent() int {
	var c2 BlinkStoneCompiler
	c2.init()
	c2.start(c.backend.getOutput())
	return c2.compile(c.blinks - 1)
}

func (c *BlinkStoneCompiler) init() {
	c.backend.init()
}

func (c *BlinkStoneCompiler) start(st string) {
	c.state = State_Initial
	c.scanner.init(st)
	c.backend.start()
	c.stones = 0
}

func (c *BlinkStoneCompiler) parse() {
	tk := c.scanner.nextToken()
	action := productions[c.state][tk]

	if action.action != nil {
		action.action(c)
	}
	c.state = action.nextState
}

func (c *BlinkStoneCompiler) compile(blinks int) int {
	c.blinks = blinks
	for !c.scanner.Eof {
		c.parse()
	}
	return c.stones
}
