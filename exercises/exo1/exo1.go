package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	longueur float64
	largeur  float64
}

type Cercle struct {
	rayon float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.largeur * r.longueur
}
func (c Cercle) Area() float64 {
	return 3.14 * math.Pow(c.rayon, 2)
}
func (c Cercle) String() string {
	return fmt.Sprintf("%v cm rayon ", c.rayon)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("%v cm longueur %v cm largeur", r.longueur, r.largeur)
}

func main() {
	r := &Rectangle{4, 5}
	c := &Cercle{4}

	shapes := []Shape{r, c}
	for _, s := range shapes {
		fmt.Println(s)
		fmt.Println("l'air", s.Area())
	}
}
