package intmath

func Gcd(a int, b int) int {
	var t int
	for b != 0 {
		t = b
		b = a % b
		a = t
	}
	return a
}

//returns gcd, x0 and y0)
func GcdExtended(a int, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}
	g, x1, y1 := GcdExtended(b%a, a)
	return g, y1 - (b/a)*x1, x1
}
