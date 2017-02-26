package main

import "fmt"

const N = 40
const LINE_FEED = 5
const MAX_COUNT = 100

var pattern_1 = [N]int{
	-1, -1, 1, -1, -1,
	-1, 1, 1, -1, -1,
	-1, -1, 1, -1, -1,
	-1, -1, 1, -1, -1,
	-1, -1, 1, -1, -1,
	-1, -1, 1, -1, -1,
	-1, -1, 1, -1, -1,
	-1, -1, -1, -1, -1}

var pattern_2 = [N]int{
	-1, 1, 1, 1, -1,
	1, -1, -1, -1, 1,
	-1, -1, -1, -1, 1,
	-1, -1, 1, 1, -1,
	-1, 1, -1, -1, -1,
	1, -1, -1, -1, -1,
	1, 1, 1, 1, 1,
	-1, -1, -1, -1, -1}

var pattern_4 = [N]int{
	1, -1, -1, 1, -1,
	1, -1, -1, 1, -1,
	1, -1, -1, 1, -1,
	1, 1, 1, 1, 1,
	-1, -1, -1, 1, -1,
	-1, -1, -1, 1, -1,
	-1, -1, -1, 1, -1,
	-1, -1, -1, -1, -1}

func weights(w *[N][N]int, x0 [N]int, x1 [N]int, x2 [N]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			w[i][j] = x0[i]*x0[j] + x1[i]*x1[j] + x2[i]*x2[j]
		}
	}
	for k := 0; k < N; k++ {
		w[k][k] = 0
	}
}

func mul(w [N][N]int, s [N]int, h *[N]int) {
	for i := 0; i < N; i++ {
		sum := 0
		for j := 0; j < N; j++ {
			sum += w[i][j] * s[j]
		}
		h[i] = sum
	}
}

func sign(y int) int {
	if y > 0 {
		return 1
	}
	return -1
}

func energy(w [N][N]int, s [N]int) int {
	e := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			e += w[i][j] * s[i] * s[j]
		}
	}
	return -1 * e
}

func update_start_configuration(s *[N]int, h [N]int, old_s [N]int) {
	for i := 0; i < N; i++ {
		if h[i] != 0 {
			s[i] = sign(h[i])
		} else {
			s[i] = old_s[i]
		}
	}
}

func print_pattern(pattern [N]int, line_feed int) {
	for i := 0; i < N; i++ {
		if pattern[i] == -1 {
			fmt.Print("0")
		} else {
			fmt.Print("1")
		}
		if ((i + 1) % line_feed) == 0 {
			fmt.Printf("\n")
		}
	}
}

func main() {
	var w [N][N]int

	var s = [N]int{
		1, 1, -1, 1, -1,
		-1, 1, -1, 1, 1,
		1, -1, -1, 1, -1,
		-1, 1, 1, 1, 1,
		-1, -1, -1, 1, -1,
		-1, -1, -1, 1, -1,
		-1, -1, -1, 1, -1,
		-1, -1, -1, -1, -1}

	weights(&w, pattern_1, pattern_2, pattern_4)
	fmt.Printf("Energy of initial configuration: %d\n", energy(w, s))

	var h [N]int
	var old_s [N]int
	count := 0
	for ; count < MAX_COUNT && old_s != s; count++ {
		copy(old_s[0:], s[0:])
		mul(w, s, &h)
		update_start_configuration(&s, h, old_s)
		fmt.Printf("count:%d\n", count)
	}
	fmt.Printf("\nNumber of iterations: %d\n", count)
	print_pattern(s, LINE_FEED)
	fmt.Printf("Energy of end configuration: %d\n", energy(w, s))
}
