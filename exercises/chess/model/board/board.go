package board

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jglouis/4eail40_2020/exercises/chess/model/coord"
	"github.com/jglouis/4eail40_2020/exercises/chess/model/piece"
)

// Classic 8x8 Chess board
type Classic [8][8]piece.Piece

// tentative d'affichage du tableau de jeu
func NewBoardClassic() *Classic {
	return new(Classic)
}

// tentative de determination de la type de piece a un endroit néccessaire pour le mouvement de la piece
// PieceAt retrievves piece at give coordinates.
// Returns nil if no piece was found.
func (c *Classic) String() string {
	var str strings.Builder
	for y := 7; y >= 0; y-- {
		for x := 0; x <= 7; x++ {
			p := c[x][y]
			if p == nil {
				str.WriteString(" ")
			} else {
				str.WriteString(p.String())
			}
		}
		str.WriteString("\n")
	}
	return str.String()
}

// MovePiece moves a piece from given coordinates to
// given coordinates.
// Returns an error if destination was occupied.
func (c *Classic) MovePiece(from coord.ChessCoordinates, to coord.ChessCoordinates) error {
	// check if the given coordinates exist
	if c.IsCoordinateValid(from) && c.IsCoordinateValid(to) {
		return errors.New("invalid coordinates")
	}

	pieceAtCoord := c.PieceAt(from)
	// check if there is a piece at from
	if pieceAtCoord == nil {
		return fmt.Errorf("couldn't find piece at %s", from.String())
	}
	// place the piece at to coord if possible
	if err := c.PlacePieceAt(pieceAtCoord, to); err != nil {
		return err
	}
	// remove the piece from its orignal state
	c.RemovePieceAt(from)
	return nil
}

// PlacePieceAt places a given piece at given location.
// Returns an error if destination was occupied.
func (c *Classic) PlacePieceAt(p piece.Piece, at coord.ChessCoordinates) error {
	// check if the given coordinate exists
	if !c.IsCoordinateValid(at) {
		return fmt.Errorf("invalid coordinates: %s", at.String())
	}
	pieceAtCoord := c.PieceAt(at)
	// check if pieceAtCoord is occupied
	if pieceAtCoord != nil {
		return fmt.Errorf("coord %s already occupied by %s", at.String(), pieceAtCoord.String())
	}
	// set p at at coordinates
	x, _ := at.Coord(0)
	y, _ := at.Coord(1)
	c[x][y] = p
	return nil
}

// IsCoordinateValid a bool if the given coordinates is valid or not.
func (c *Classic) IsCoordinateValid(from coord.ChessCoordinates) bool {
	x, errx := from.Coord(0)
	y, erry := from.Coord(1)
	return errx == nil && erry == nil && x >= 0 && y >= 0 && x < 8 && y < 8
}
