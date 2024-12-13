package main

import (
	"advent/aoc/common"
	"advent/aoc/pos"
	"fmt"
)

func main() {

	lines := common.StartDay(12, "test")

	regions := findRegions(lines)

	for _, r := range regions {
		fmt.Println("Type: ", string(r.plantType), " Plots: ", r.plots)
	}

}

type Region struct {
	plantType          uint8
	plots              []pos.Position
	markedForMergeInto *Region
}

func (r *Region) belongs(plot pos.Position, t uint8) bool {
	if r.plantType != t {
		return false
	}
	for _, p := range r.plots {
		if plotsAreNeighbours(p, plot) {
			return true
		}
	}
	return false
}

func plotsAreNeighbours(plot1 pos.Position, plot2 pos.Position) bool {
	dist := plot1.DistanceTo(plot2)
	if dist.Dx == 0 && (dist.Dy == 1 || dist.Dy == -1) {
		return true
	}
	if dist.Dy == 0 && (dist.Dx == 1 || dist.Dx == -1) {
		return true
	}
	return false
}

func findRegions(lines []string) []Region {
	var result []Region
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			p := pos.Position{X: x, Y: y}
			plantType := lines[y][x]
			found := false

			// find a region neighboring this plot
			var region *Region
			for i := 0; i < len(result) && !found; i++ {
				found = result[i].belongs(p, plantType)
				if found {
					region = &result[i]
				}
			}
			if found {
				// if there is a neighbouring region, then add this plot to it
				region.plots = append(region.plots, p)
			} else {
				// otherwise create a new region with this plot in it
				region = &Region{
					plantType:          plantType,
					plots:              []pos.Position{},
					markedForMergeInto: nil,
				}
				region.plots = append(region.plots, p)
				result = append(result, *region)
			}
		}
	}
	return mergeRegions(result)
}

func mergeRegions(r []Region) []Region {
	for i := 0; i < len(r); i++ {
		region1 := &(r)[i]
		for j := 0; j < len(r); j++ {
			region2 := &(r)[j]
			if region1 != region2 && region1.plantType == region2.plantType {
				isNeighbour := false
				for i := 0; i < len(region1.plots) && !isNeighbour; i++ {
					for j := 0; j < len(region2.plots) && !isNeighbour; j++ {
						isNeighbour = plotsAreNeighbours(region1.plots[i], region2.plots[j])
					}
				}
				if isNeighbour && region2.markedForMergeInto == nil {
					region1.markedForMergeInto = region2
				}
			}
		}
	}
	var result2 []Region
	for i := 0; i < len(r); i++ {
		region1 := &r[i]
		if region1.markedForMergeInto == nil {
			for _, region2 := range r {
				if region2.markedForMergeInto == region1 {
					for j := 0; j < len(region2.plots); j++ {
						region1.plots = append(region1.plots, region2.plots[j])
					}
				}
			}
			result2 = append(result2, *region1)
		}
	}
	return result2
}
