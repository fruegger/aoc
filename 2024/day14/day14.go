package main

import (
	"advent/aoc/common"
	"advent/aoc/pos"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Robot struct {
	p pos.Position
	v pos.Direction
}

func (r *Robot) read(line string) {
	parts := strings.Split(line, " ")

	ps := strings.Split(parts[0], "=")
	ps = strings.Split(ps[1], ",")
	px := common.StringToNum(ps[0])
	py := common.StringToNum(ps[1])

	vel := strings.Split(parts[1], "=")
	vel = strings.Split(vel[1], ",")
	dx := common.StringToNum(vel[0])
	dy := common.StringToNum(vel[1])
	r.p = pos.Position{X: px, Y: py}
	r.v = pos.Direction{Dx: dx, Dy: dy}
}

func (r *Robot) move(space pos.Dimension) {
	r.p.X = (r.p.X + r.v.Dx) % space.X
	if (r.p.X) < 0 {
		r.p.X = space.X + r.p.X
	}
	r.p.Y = (r.p.Y + r.v.Dy) % space.Y
	if (r.p.Y) < 0 {
		r.p.Y = space.Y + r.p.Y
	}
}

func main() {
	spaces := map[string]pos.Dimension{
		"test":  pos.Dimension{X: 11, Y: 7},
		"test2": pos.Dimension{X: 11, Y: 7},
		"input": pos.Dimension{X: 101, Y: 103},
	}

	fileType := "input"
	space := spaces[fileType]

	buf := bufio.NewReader(os.Stdin)

	lines := common.StartDay(14, fileType)
	robots := readRobots(lines)
	printRobots(robots, space)

	for j := 1; j <= 100; j++ {
		for i := 0; i < len(robots); i++ {
			robots[i].move(space)
		}
	}

	q1_TL := pos.Position{0, 0}
	q1_BR := pos.Position{space.X/2 - 1, space.Y/2 - 1}

	c1 := countRobotsInQuadrant(robots, q1_TL, q1_BR)
	q2_TL := pos.Position{0, space.Y/2 + 1}
	q2_BR := pos.Position{space.X/2 - 1, space.Y - 1}
	c2 := countRobotsInQuadrant(robots, q2_TL, q2_BR)

	q3_TL := pos.Position{space.X/2 + 1, 0}
	q3_BR := pos.Position{space.X - 1, space.Y/2 - 1}
	c3 := countRobotsInQuadrant(robots, q3_TL, q3_BR)

	q4_TL := pos.Position{space.X/2 + 1, space.Y/2 + 1}
	q4_BR := pos.Position{space.X - 1, space.Y - 1}
	c4 := countRobotsInQuadrant(robots, q4_TL, q4_BR)

	fmt.Println("Part 1: ", c1*c2*c3*c4)

	tipDetected := false
	for k := 1; ; k++ {

		for i := 0; i < len(robots); i++ {
			robots[i].move(space)
		}
		tipDetected = detectTip(robots, space)
		if tipDetected {
			fmt.Printf("after %d sec", 100+k)
			fmt.Println()
			printRobots(robots, space)

			_, err := buf.ReadBytes('\n')
			if err != nil {

			}
		}
	}

}

func detectTip(robots []Robot, space pos.Position) bool {
	size := 4
	for y := 0; y < space.Y-size; y++ {
		for x := size; x < space.X-size; x++ {
			tip := countRobotsOnTile(robots, x, y) > 0
			for j := 0; j < size; j++ {
				tip = tip && countRobotsOnTile(robots, x-j, y+j) > 0
				tip = tip && countRobotsOnTile(robots, x+j, y+j) > 0
			}
			if tip {
				return true
			}
		}
	}
	return false
}

func countRobotsInQuadrant(robots []Robot, q1_TL pos.Position, q1_BR pos.Position) int {
	total := 0
	for y := q1_TL.Y; y <= q1_BR.Y; y++ {
		for x := q1_TL.X; x <= q1_BR.X; x++ {
			total += countRobotsOnTile(robots, x, y)
		}
	}
	return total
}

func readRobots(lines []string) []Robot {
	var robots []Robot
	for _, line := range lines {
		var robot Robot
		robot.read(line)
		robots = append(robots, robot)
	}
	return robots
}

func printRobots(robots []Robot, space pos.Dimension) {
	for y := 0; y < space.Y; y++ {
		for x := 0; x < space.X; x++ {
			cnt := countRobotsOnTile(robots, x, y)
			if cnt == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(string(cnt + '0'))
			}
		}
		fmt.Println()
	}
}

func countRobotsOnTile(robots []Robot, x int, y int) int {
	total := 0
	for _, r := range robots {
		if r.p.X == x && r.p.Y == y {
			total++
		}
	}
	return total
}
