package board

import (
	"fmt"

	"github.com/jglouis/4eail40_2020/exercises/chess/model/coord"
	"github.com/jglouis/4eail40_2020/exercises/chess/model/piece"
)

// Classic 8x8 Chess board
type Classic [8][8]piece.Piece

// tentative d'affichage du tableau de jeu
func (c *Classic) String() string {
	panic("not implemented") // TODO: Implement
	}
}

// tentative de determination de la type de piece a un endroit néccessaire pour le mouvement de la piece
// PieceAt retrievves piece at give coordinates.
// Returns nil if no piece was found.
func (c *Classic) PieceAt(at coord.ChessCoordinates) piece.Piece {
	x, _ := at.Coord(0)
	y, _ := at.Coord(1)
	tab := c[x][y]
	if tab != nil {
		return tab
	} else {
		return nil
	}
}

// MovePiece moves a piece from given coordinates to
// given coordinates.
// Returns an error if destination was occupied.
func (c *Classic) MovePiece(from coord.ChessCoordinates, to coord.ChessCoordinates) error {

	if c.PieceAt(from) != nil {
		x, _ := from.Coord(0)
		y, _ := from.Coord(1)
		pos := c[x][y]
		return c.PlacePieceAt(pos, to)
	} else {
		return fmt.Errorf("pas de piece a l'emplacement %v", from.String())
	}

}

// PlacePieceAt places a given piece at given location.
// Returns an error if destination was occupied.
func (c *Classic) PlacePieceAt(p piece.Piece, at coord.ChessCoordinates) error {
	panic("not implemented") // TODO: Implement
}
