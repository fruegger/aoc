package main

import (
	"advent/aoc/common"
	"advent/aoc/ds"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type gateOperation func(in1, in2 int) int

type Operation struct {
	name string
	op   gateOperation
}

type Gate struct {
	name   string //synthetic; for debugging
	input1 *Signal
	input2 *Signal
	output *Signal
	opType *Operation
}

type Signal struct {
	name    string
	value   int
	wiredTo []*Gate
	isInput bool
}

type Circuit struct {
	signals       map[string]*Signal
	gates         ds.Set[*Gate]
	gatesByOutput map[*Signal]*Gate
	gatesByName   map[string]*Gate
}

const AND = "AND"
const OR = "OR"
const XOR = "XOR"

var operations = map[string]Operation{
	AND: {
		name: AND,
		op: func(in1, in2 int) int {
			return in1 & in2
		},
	},
	OR: {
		name: OR,
		op: func(in1, in2 int) int {
			return in1 | in2
		},
	},
	XOR: {
		name: XOR,
		op: func(in1, in2 int) int {
			return in1 ^ in2
		},
	},
}

func main() {
	lines := common.StartDay(24, "input")
	var c = &Circuit{
		signals:       map[string]*Signal{},
		gates:         ds.Set[*Gate]{},
		gatesByOutput: map[*Signal]*Gate{},
		gatesByName:   map[string]*Gate{},
	}

	c.initialize(lines)
	c.printCircuit()
	c.propagateAll()

	var numeric int64 = 0
	in1 := c.calcValueForLabel("x")

	numeric, _ = strconv.ParseInt(in1, 2, 64)
	in2 := c.calcValueForLabel("y")
	numeric, _ = strconv.ParseInt(in2, 2, 64)

	output := c.calcValueForLabel("z")
	numeric, _ = strconv.ParseInt(output, 2, 64)

	fmt.Println("Part1: ", numeric, "(", output, ")")

	for i := 0; i < 45; i++ {
		c.signals[signalName("x", i)].value = 0
		c.signals[signalName("y", i)].value = 0
	}

	var broken ds.Set[*Gate]

	// this is insufficient, since it does not consider carry - bit errors that may cancel out.
	// analysis done by hand, because swapping gates may induce cycles (feedback loops or short circuits)
	for i := 0; i < 45; i++ {
		var labels = [2]string{"x", "y"}
		for _, label := range labels {
			affected := c.probeSignal(label, i)
			//			sig := c.signals[signalName(label, i)]
			//			gates := c.propagateOne(sig)
			if !affected.IsEmpty() {
				broken = broken.Union(affected)
			}
		}
	}

	c = &Circuit{
		signals:       map[string]*Signal{},
		gates:         ds.Set[*Gate]{},
		gatesByOutput: map[*Signal]*Gate{},
		gatesByName:   map[string]*Gate{},
	}
	c.initialize(lines)
	c.propagateAll()

	diff, _ := iterate(c)
	fmt.Println("Before: ", diff)

	c.swapOutputs("G159", "G107")
	c.swapOutputs("G196", "G15")
	c.swapOutputs("G102", "G91")
	c.swapOutputs("G208", "G213")

	c.propagateAll()
	diff, _ = iterate(c)
	fmt.Println("After: ", diff)

	var sigNames = []string{
		c.gatesByName["G159"].output.name,
		c.gatesByName["G107"].output.name,
		c.gatesByName["G196"].output.name,
		c.gatesByName["G15"].output.name,
		c.gatesByName["G102"].output.name,
		c.gatesByName["G91"].output.name,

		c.gatesByName["G208"].output.name,
		c.gatesByName["G213"].output.name,
	}
	slices.Sort(sigNames)
	fmt.Print("Part2 :")
	for i, n := range sigNames {
		fmt.Print(n)
		if i < len(sigNames)-1 {
			fmt.Print(",")
		}
	}
	fmt.Println()
}

func (c *Circuit) swapOutputs(name1, name2 string) {
	g1 := c.gatesByName[name1]
	g2 := c.gatesByName[name2]
	save := g1.output
	g1.output = g2.output
	g2.output = save
}

func (c *Circuit) probeSignal(name string, nr int) ds.Set[*Gate] {
	signal := c.signals[signalName(name, nr)]
	signal.value = 1

	var affected ds.Set[*Gate]

	c.propagateAll()
	diff, bitLen := iterate(c)
	affected = c.affectedGates(signal, diff, bitLen)

	if diff != 0 {
		fmt.Printf("%s%d - affected; gates - ", name, nr)
		for _, g := range *affected.Elements() {
			fmt.Print(" ", g.name)
		}
		fmt.Println()
	}
	c.signals[signalName(name, nr)].value = 0
	return affected
}

func iterate(c *Circuit) (int, int) {
	in1 := c.calcValueForLabel("x")
	in2 := c.calcValueForLabel("y")

	num1, _ := strconv.ParseInt(in1, 2, 64)
	num2, _ := strconv.ParseInt(in2, 2, 64)

	expectedVal := num1 + num2
	actual := c.calcValueForLabel("z")
	actualVal, _ := strconv.ParseInt(actual, 2, 64)
	diff := actualVal ^ expectedVal

	//	fmt.Printf("x:   %s (%d)\n", in1, num1)
	//	fmt.Printf("y    %s (%d)\n", in2, num2)
	//fmt.Printf("Exp: %b (%d)\n", expectexVal, expectexVal)
	//fmt.Printf("Act: %s (%d)\n", actual, actualVal)
	//fmt.Printf("Dif: %b (%d)\n", diff, diff)

	return int(diff), len(actual)
}

func (c *Circuit) affectedGates(input *Signal, diff, bitLen int) ds.Set[*Gate] {
	var result ds.Set[*Gate]
	for i := 0; i < bitLen; i++ {
		if diff&(1<<i) != 0 {
			sigName := signalName("z", i)
			signal := c.signals[sigName]
			var affected ds.Set[*Gate]
			c.findRoute(input, signal, &affected)
			result = result.Union(affected)
		}
	}
	return result
}

var sigQ = ds.Queue[*Signal]{}
var gateQ = ds.Queue[*Gate]{}

func (c *Circuit) propagateAll() ds.Set[*Gate] {
	c.pushInputs()
	return c.propagate()
}

func (c *Circuit) propagateOne(signal *Signal) ds.Set[*Gate] {
	sigQ.Push(signal)
	return c.propagate()
}

func (c *Circuit) propagate() ds.Set[*Gate] {
	var results ds.Set[*Gate]
	var signal *Signal
	for sigQ.Pull(&signal) {
		for _, gate := range signal.wiredTo {
			op := gate.opType.op
			gate.output.value = op(gate.input1.value, gate.input2.value)
			results.Add(gate)
			sigQ.Push(gate.output)
		}
	}
	return results
}

func (c *Circuit) findRoute(in *Signal, out *Signal, gates *ds.Set[*Gate]) bool {
	var gate *Gate
	gate = c.gatesByOutput[out]
	if gate == nil {
		return false
	}
	if gate.input1 == in || gate.input2 == in {
		gates.Add(gate)
		return true
	} else {
		found := c.findRoute(in, gate.input1, gates)
		if found {
			gates.Add(gate)
		} else {
			found = c.findRoute(in, gate.input2, gates)
			if found {
				gates.Add(gate)
			}
		}
		return found
	}
}

func (c *Circuit) pushInputs() {
	for _, s := range c.signals {
		if s.isInput {
			sigQ.Push(s)
		}
	}
}

func (c *Circuit) initialize(lines []string) {
	for _, line := range lines {
		if strings.Contains(line, ": ") {
			parts := strings.Split(line, ": ")
			name := parts[0]
			value := parts[1]
			c.signals[name] = &Signal{
				name:    name,
				value:   common.StringToNum(value),
				wiredTo: []*Gate{},
				isInput: true,
			}
		} else {
			if strings.Contains(line, " -> ") {
				equation := strings.Split(line, " -> ")
				leftPart := strings.Split(equation[0], " ")
				signal1 := c.findOrCreateSignal(leftPart[0])
				signal2 := c.findOrCreateSignal(leftPart[2])
				signal3 := c.findOrCreateSignal(equation[1])
				operation := operations[leftPart[1]]
				gate := Gate{
					name:   "G" + strconv.Itoa(len(*c.gates.Elements())+1),
					input1: signal1,
					input2: signal2,
					output: signal3,
					opType: &operation,
				}
				*c.gates.Elements() = append(*c.gates.Elements(), &gate)
				c.gatesByName[gate.name] = &gate
				c.gatesByOutput[signal3] = &gate
			}
		}
	}

	// now wire all signals to their gate
	for _, gate := range *c.gates.Elements() {
		gate.input1.wiredTo = append(gate.input1.wiredTo, gate)
		gate.input2.wiredTo = append(gate.input2.wiredTo, gate)
	}
}

func (c *Circuit) findOrCreateSignal(name string) *Signal {
	var signal *Signal
	signal, found := c.signals[name]
	if !found {
		signal = &Signal{
			name:    name,
			isInput: false,
		}
		c.signals[signal.name] = signal
	}
	return signal
}

func (c *Circuit) printCircuit() {
	fmt.Println("Signals:")
	for _, s := range c.signals {
		fmt.Print(s.name, ": ", s.value, " (")
		for i, gate := range s.wiredTo {
			fmt.Print(gate.name)
			if i+1 < len(s.wiredTo) {
				fmt.Print(", ")
			}
		}
		fmt.Println(")")
	}
	fmt.Println("Gates:")
	for _, g := range *c.gates.Elements() {
		fmt.Println(g.name, ":", g.input1.name, " ", g.opType.name, " ", g.input2.name, " -> ", g.output.name)
	}
}

func signalName(label string, i int) string {
	sigName := fmt.Sprintf("%d", i)
	if i < 10 {
		sigName = "0" + sigName
	}
	sigName = label + sigName
	return sigName
}

func (c *Circuit) setValueForLabel(label, val string) {
	for i := 0; i < len(val); i++ {
		sigName := signalName(label, i)
		s, found := c.signals[sigName]
		if found {
			s.value = int(val[i] - '0')
		}
	}
}

func (c *Circuit) calcValueForLabel(label string) string {
	result := ""
	done := false
	for i := 0; !done; i++ {
		sigName := signalName(label, i)
		s, found := c.signals[sigName]
		if found {
			result = string(uint8(s.value)+'0') + result
		} else {
			done = true
		}
	}
	return result
}
