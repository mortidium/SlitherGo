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

func (r *Row) makeRow(shape int, baseSide float64, nums []int) {
	switch shape {
	case square:
		for _, num := range nums {
			sq := makeSquare(r.last, 0, baseSide, num)
			r.cells = append(r.cells, sq)
			r.last = pos(rightShapePos(square, r.last, baseSide), r.last.y)
		}
	case hex:
		for _, num := range nums {
			if num >= e {
				h := makeHex(r.last, 0, baseSide, num)
				r.cells = append(r.cells, h)
				r.last = pos(rightShapePos(hex, r.last, baseSide), r.last.y)
			}
		}

	}

}

func makeBoard(shape int, center *Position, baseSide float64, rows []string) *Board {
	rv := &Board{rows: []*Row{}, baseSide: baseSide}
	centerX := center.x
	centerY := center.y - float64(len(rows)-1)*downShapePos(shape, pos(0, 0), baseSide)/2
	for i, r := range rows {
		ints := convertStrToInts(r)
		if len(ints) == 0 {
			continue
		}

		thisRowPosLeft := pos(
			centerX-float64(len(r)-1)*rightShapePos(shape, pos(0, 0), baseSide)/2,
			centerY+float64(i)*downShapePos(shape, pos(0, 0), baseSide),
		)

		row := &Row{cells: []*Shape{}, first: thisRowPosLeft, last: thisRowPosLeft}
		row.makeRow(shape, baseSide, ints)

		rv.rows = append(rv.rows, row)
	}
	return rv
}
