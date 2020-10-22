// Package model contains the gameplay logic for the game of chess
package model

// Implement type Board

func (b Board) NewTray() {

	board8x8 := [8][8]int{

		{21, 22, 23, 24, 25, 26, 27, 28},
		{31, 32, 33, 34, 35, 36, 37, 38},
		{41, 42, 43, 44, 45, 46, 47, 48},
		{51, 52, 53, 54, 55, 56, 57, 58},
		{61, 62, 63, 64, 65, 66, 67, 68},
		{71, 72, 73, 74, 75, 76, 77, 78},
		{81, 82, 83, 84, 85, 86, 87, 88},
		{91, 92, 93, 94, 95, 96, 97, 98},
	}

	piece.MoveRoiNoir(25, 1)
	piece.MoveRoiBlanc(95, 1)
	piece.MoveDameNoir(94, 1)
	piece.MoveDameBlanche(93, 1)

}

func (b *Board) MovePiece(from, to int) {

}

type Board interface {
}
