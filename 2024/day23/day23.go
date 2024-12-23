package main

import (
	"advent/aoc/common"
	"fmt"
	"maps"
	"slices"
	"strings"
)

type Name = string

var adjacency = map[Name][]Name{}
var triplets []string

func main() {
	lines := common.StartDay(23, "input")

	for _, line := range lines {
		ends := strings.Split(line, "-")
		addLink(ends[0], ends[1])
	}
	formTriplets()
	total := 0
	for _, t := range triplets {
		names := strings.Split(t, ",")
		if names[0][0] == 't' || names[1][0] == 't' || names[2][0] == 't' {
			total++
		}
	}
	fmt.Println("Part1: ", total)

	var nodes []string
	for v := range maps.Keys(adjacency) {
		nodes = append(nodes, v)
	}

	var r []string
	var x []string
	slices.Sort(nodes)
	var p = nodes
	fmt.Println(nodes)
	cliques := []string{}
	BronKerbosch(r, p, x, func(r []string) {
		if len(r) > len(cliques) {
			cliques = r
		}
	})
	//	slices.Sort(result)
	fmt.Print("Part 2: ")
	for i, r := range cliques {
		fmt.Print(r)
		if i+1 < len(cliques) {
			fmt.Print(",")
		} else {
			fmt.Println()
		}
	}

	// part 2
}

func addLink(start, end Name) {
	edges, _ := adjacency[start]
	edges = append(edges, end)
	adjacency[start] = edges
	edges, _ = adjacency[end]
	edges = append(edges, start)
	adjacency[end] = edges
}

// every node only has three edges!
func formTriplets() {
	// pairs of edges that are connected
	for name1, edge := range adjacency {
		for i := 1; i < len(edge); i++ {
			for j := 0; j < i; j++ {
				name2 := edge[i]
				name3 := edge[j]
				edges2 := adjacency[name2]
				if slices.Contains(edges2, name3) {
					addTriplet(name1, name2, name3)
				}
			}
		}
	}
}

func addTriplet(name1, name2, name3 string) {
	var swap = ""
	if name2 > name3 {
		swap = name2
		name2 = name3
		name3 = swap
	}
	if name1 > name3 {
		swap = name1
		name1 = name3
		name3 = swap
	}
	if name1 > name2 {
		swap = name1
		name1 = name2
		name2 = swap
	}
	triplet := name1 + "," + name2 + "," + name3
	if !slices.Contains(triplets, triplet) {
		triplets = append(triplets, triplet)
	}
}

func BronKerbosch(r, p, x []Name, collect func([]string)) {
	if len(p) == 0 && len(x) == 0 {
		checkDouble("r", r)
		collect(r)
		fmt.Println("---", r)
	} else {
		for _, v := range p {
			neighbours := adjacency[v]
			checkDouble("neighbours", neighbours)
			p2 := intersection(p, neighbours)
			checkDouble("p2", p2)
			x2 := intersection(x, neighbours)
			checkDouble("x2", x2)
			r2 := union(r, v)
			checkDouble("r2", r2)

			BronKerbosch(r2, p2, x2, collect)

			p = remove(p, v)
			checkDouble("p", p)
			x = union(x, v)
			checkDouble("x", x)

		}
	}
}

func intersection(a, b []Name) []Name {
	var r []Name
	for _, v := range a {
		if slices.Contains(b, v) {
			r = append(r, v)
		}
	}
	return r
}

func union(a []Name, el Name) []Name {
	var r = a
	if !slices.Contains(a, el) {
		r = append(r, el)
	}
	return r
}

func remove(set []Name, el Name) []Name {
	var p2 []Name
	i := slices.Index(set, el)
	if i >= 0 {
		p2 = deleteElement(set, i)
	} else {
		p2 = set
	}
	return p2
}

func checkDouble(msg string, slice []string) {
	for i, v1 := range slice {
		for j, v2 := range slice {
			if v1 == v2 && i != j {
				fmt.Println(msg, slice)
			}
		}
	}
}
func deleteElement(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}
