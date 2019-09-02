package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Row keeps a row of shapes
type Row struct {
	cells       []*Shape
	first, last *Position // positions of centers of first and last shape
}

// Board keeps the whole board of rows
type Board struct {
	rows     []*Row
	baseSide float64
}

func (b *Board) draw(renderer *sdl.Renderer) {
	for _, row := range b.rows {
		for _, cell := range row.cells {
			cell.draw(renderer)
		}
	}
}

func convertStrToInts(str string) []int {
	rv := []int{}
	for _, ch := range str {
		if ch-'0' >= 0 && ch-'0' <= 5 {
			rv = append(rv, int(ch-'0'))
		} else {
			rv = append(rv, e)
		}
	}
	return rv
}

func makeSquareBoard(offset *Position, baseSide float64, rows []string) *Board {
	rv := &Board{rows: []*Row{}, baseSide: baseSide}
	currentRowPos := offset
	for _, r := range rows {
		ints := convertStrToInts(r)
		row := &Row{cells: []*Shape{}, first: currentRowPos, last: currentRowPos}
		row.squaresRow(baseSide, ints)
		currentRowPos = downShapePos(square, currentRowPos, baseSide)
		rv.rows = append(rv.rows, row)
	}
	return rv
}

func (r *Row) squaresRow(baseSide float64, nums []int) {
	for _, num := range nums {
		sq := makeSquare(r.last, 0, baseSide, num)
		r.cells = append(r.cells, sq)
		r.last = nextShapePos(square, r.last, baseSide)
	}
}
