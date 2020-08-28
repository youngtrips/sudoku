package main

import (
	"bufio"
	"fmt"
	"os"
)

type SudokuSolver struct {
	pos  [][2]int
	mask []int
	rows []int
	cols []int
	subs []int
	ok   bool
	n    int
}

func (s *SudokuSolver) toSubNo(i int, j int) int {
	return i/3*3 + j/3
}

func (s *SudokuSolver) check(i int, j int, no int, k int) bool {
	mask := s.mask[k]
	if s.rows[i]&mask != 0 || s.cols[j]&mask != 0 || s.subs[no]&mask != 0 {
		return false
	}
	return true
}

func (s *SudokuSolver) search(board [][]byte, i int) bool {
	if i >= len(s.pos) {
		return true
	}
	for k := 1; k <= s.n; k++ {
		x := s.pos[i][0]
		y := s.pos[i][1]
		no := s.toSubNo(x, y)
		if s.check(x, y, no, k) {
			mask := s.mask[k]
			s.rows[x] |= mask
			s.cols[y] |= mask
			s.subs[no] |= mask
			board[x][y] = byte(k + '0')
			if s.search(board, i+1) {
				return true
			}
			s.rows[x] ^= mask
			s.cols[y] ^= mask
			s.subs[no] ^= mask
			board[x][y] = '.'
		}
	}
	return false
}

func build(board [][]byte) *SudokuSolver {
	n := 9
	s := &SudokuSolver{
		pos:  make([][2]int, 0),
		mask: make([]int, n+1),
		rows: make([]int, n),
		cols: make([]int, n),
		subs: make([]int, n),
		n:    n,
	}

	for i := 1; i <= n; i++ {
		s.mask[i] = 1 << (i - 1)
		s.mask[0] |= s.mask[i]
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ch := board[i][j]
			if ch == '.' {
				s.pos = append(s.pos, [2]int{i, j})
			} else {
				k := int(ch - '0')
				s.rows[i] |= s.mask[k]
				s.cols[j] |= s.mask[k]
				s.subs[s.toSubNo(i, j)] |= s.mask[k]
			}
		}
	}
	return s
}

func solveSudoku(board [][]byte) {
	sudoko := build(board)
	sudoko.search(board, 0)
}

func main() {
	board := make([][]byte, 9)
	r := bufio.NewReader(os.Stdin)
	for i := 0; i < 9; i++ {
		board[i] = make([]byte, 9)
		line, _, _ := r.ReadLine()
		for j, ch := range line {
			board[i][j] = ch
		}
	}

	show(board)
	solveSudoku(board)
	show(board)
}

func show(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}
