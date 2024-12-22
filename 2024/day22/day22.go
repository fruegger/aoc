package main

import (
	"advent/aoc/common"
	"fmt"
)

func main() {
	lines := common.StartDay(22, "input")
	total := 0
	var buyers []Buyer

	for _, line := range lines {
		buyer := Buyer{}
		buyer.secrets[0] = common.StringToNum(line)
		total += buyer.secret2000()
		buyers = append(buyers, buyer)
	}
	fmt.Println("Part 1:", total)
	bestSale := 0
	sequences := createSequences()

	for _, s := range sequences {
		total = 0
		for _, b := range buyers {
			total += b.sales(s)
		}
		if total > bestSale {
			bestSale = total
			fmt.Println(s, "-", bestSale)
		}
	}
	fmt.Println("Sales: ", bestSale)
}

type Buyer struct {
	secrets [2000]int
	deltas  [2000]int
}

func createSequences() [][4]int {
	var sequences [][4]int

	for s1 := -9; s1 <= 9; s1++ {
		for s2 := -9; s2 <= 9; s2++ {
			for s3 := -9; s3 <= 9; s3++ {
				for s4 := -9; s4 <= 9; s4++ {
					if s1+s2 >= -9 && s1+s2 <= 9 &&
						s2+s3 >= -9 && s2+s3 <= 9 &&
						s3+s4 >= -9 && s3+s4 <= 9 &&
						s1+s2+s3 >= -9 && s1+s2+s3 <= 9 &&
						s1+s2+s4 >= -9 && s1+s2+s4 <= 9 &&
						s1+s2+s3+s4 >= -9 && s1+s2+s3+s4 <= 9 {
						sequences = append(sequences, [4]int{s1, s2, s3, s4})
					}
				}
			}
		}
	}
	return sequences
}

func (b *Buyer) secret2000() int {
	for i := 1; i < 2000; i++ {
		b.secrets[i] = nextSecret(b.secrets[i-1])
		b.deltas[i] = (b.secrets[i] % 10) - (b.secrets[i-1] % 10)
	}
	return b.secrets[1999]
}

func (b *Buyer) sales(sequence [4]int) int {
	for p := 1; p < 1997; p++ {
		if b.deltas[p] == sequence[0] &&
			b.deltas[p+1] == sequence[1] &&
			b.deltas[p+2] == sequence[2] &&
			b.deltas[p+3] == sequence[3] {
			return b.secrets[p+3] % 10
		}
	}
	return 0
}

func nextSecret(s int) int {
	// * 64 mix, prune
	s2 := (s ^ (s << 6)) & 0xffffff
	// / 32 mix, prune
	s3 := (s2 ^ (s2 >> 5)) & 0xffffff
	// * 2048 mix, prune
	s4 := (s3 ^ (s3 << 11)) & 0xffffff
	return s4
}
