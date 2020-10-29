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
	for i := 0; i < len(Classic); i++{
		for j := 0; j < len(Classic); j++{
			fmt.Println("|" Classic[i][j] "|")
	}
}

// tentative de determination de la type de piece a un endroit néccessaire pour le mouvement de la piece 
// PieceAt retrievves piece at give coordinates.
// Returns nil if no piece was found.
func (c *Classic) PieceAt(at coord.ChessCoordinates) piece.Piece {
	for i := 0; i < len(Classic); i++{
		for j := 0; j < len(Classic); j++{
			if Classic[i][j] == coord.ChessCoordinates {
				return (" la piece est la " i , j)
			}
			else {
				return nil
			}
		}
}

// MovePiece moves a piece from given coordinates to
// given coordinates.
// Returns an error if destination was occupied.
func (c *Classic) MovePiece(from coord.ChessCoordinates, to coord.ChessCoordinates) error {
	//tentative de determination des coordonées de départ 
	depart := map[int]int{
		map = PieceAt(from)[0] , PieceAt(from)[1]
	}

	//tentative détermination des coordonées d'arrivé 
	arrive := map[int]int{
		map = PieceAt(to)[0] , PieceAt(to)[1]
	}
	
	depart[0][1] = arrive[0][1]
	depart[0][0] = [0][0]
	return fmt.Println("done")

}

// PlacePieceAt places a given piece at given location.
// Returns an error if destination was occupied.
func (c *Classic) PlacePieceAt(p piece.Piece, at coord.ChessCoordinates) error {
	panic("not implemented") // TODO: Implement
}
