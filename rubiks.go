package main

import (
  "fmt"
)

type Face struct {
  caption string
}

type Direction struct {
  caption string
}

type Rotation struct {
  caption string
}

var (

  // Faces
  Blank  = Face{"_"}
  Red    = Face{"R"}
  Green  = Face{"G"}
  Blue   = Face{"B"}
  Yellow = Face{"Y"}
  Orange = Face{"O"}
  White  = Face{"W"}

  // Directions
  North = Direction{"N"}
  South = Direction{"S"}
  East  = Direction{"E"}
  West  = Direction{"W"}

  // Rotations
  Clockwise     = Rotation{"CW"}
  Anticlockwise = Rotation{"ACW"}
)

type Piece struct {
  faces [6]Face
}

func(piece *Piece) to_s() string {
  var s = ""

  for n := 0; n < 6; n++ {
    s += piece.faces[n].caption
  }

  return s
}

func main() {
  piece := Piece{[6]Face{Red, Blank, Green, Blank, Blank, Yellow}}
  fmt.Println(piece.to_s())
}
