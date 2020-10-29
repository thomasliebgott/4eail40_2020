package board

import (
	"github.com/jglouis/4eail40_2020/exercises/chess/model/coord"
	"github.com/jglouis/4eail40_2020/exercises/chess/model/piece"
)

// Classic 8x8 Chess board
type Classic [8][8]piece.Piece

func (c *Classic) String() string {
	for i := 0; i < len(Classic); i++{
		for j := 0; j < len(Classic); j++{
			fmt.Println("|" Classic[i][j] "|")
	}
	
}

// PieceAt retrievves piece at give coordinates.
// Returns nil if no piece was found.
func (c *Classic) PieceAt(at coord.ChessCoordinates) piece.Piece {
	for i := 0; i < len(Classic); i++{
		for j := 0; j < len(Classic); j++{
			if Classic[i][j] == coord.ChessCoordinates {
				return (" la piece est la " i , j)
			}
		}
}

// MovePiece moves a piece from given coordinates to
// given coordinates.
// Returns an error if destination was occupied.
func (c *Classic) MovePiece(from coord.ChessCoordinates, to coord.ChessCoordinates) error {
	panic("not implemented") // TODO: Implement
}

// PlacePieceAt places a given piece at given location.
// Returns an error if destination was occupied.
func (c *Classic) PlacePieceAt(p piece.Piece, at coord.ChessCoordinates) error {
	panic("not implemented") // TODO: Implement
}
