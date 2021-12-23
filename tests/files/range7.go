package main

func main() {
	x := " abc\tdef\nghi\rjkl\fmno\vpqr\u0085stu\u00a0\n"
	for i, r := range x {
		println(i, int(r))
	}
}

// Output:
// 0 32
// 1 97
// 2 98
// 3 99
// 4 9
// 5 100
// 6 101
// 7 102
// 8 10
// 9 103
// 10 104
// 11 105
// 12 13
// 13 106
// 14 107
// 15 108
// 16 12
// 17 109
// 18 110
// 19 111
// 20 11
// 21 112
// 22 113
// 23 114
// 24 133
// 26 115
// 27 116
// 28 117
// 29 160
// 31 10
