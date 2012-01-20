package main

import (
  "fmt"
)

type Face      string 
type Side      string
type Direction string
type Rotation  string
type Edge      map[Direction] Side
type Transform map[Direction] Direction
type Piece     map[Side] Face

var (

  Blank  Face = "_"
  red    Face = "R"
  green  Face = "G"
  blue   Face = "B"
  yellow Face = "Y"
  orange Face = "O"
  white  Face = "W"

  top    Side = "top"
  bottom Side = "bottom"
  front  Side = "front"
  back   Side = "back"
  left   Side = "left"
  right  Side = "right"
  sides = [...]Side { top, bottom, front, back, left, right }

  north Direction = "N"
  east  Direction = "E"
  south Direction = "S"
  west  Direction = "W"

  clockwise     Rotation = "CW"
  anticlockwise Rotation = "ACW"
  
  opposites = map [Side] Side {
    top:    bottom,
    bottom: top,
    front:  back,
    back:   front,
    left:   right,
    right:  left,
  }

  edges = map [Side] Edge {
    top:    Edge {  north: back,   south: front,   east: right,  west: left   },
    bottom: Edge {  north: front,  south: back,    east: left,   west: right  },
    front:  Edge {  north: top,    south: bottom,  east: right,  west: left   },
    back:   Edge {  north: top,    south: bottom,  east: left,   west: right  },
    left:   Edge {  north: top,    south: bottom,  east: front,  west: back   },
    right:  Edge {  north: top,    south: bottom,  east: back,   west: front  },
  }

  transforms = map [Rotation] Transform {
    clockwise:     Transform {  west: north,  north: east,  east: south,  south: west  },
    anticlockwise: Transform {  east: north,  south: east,  west: south,  north: west  },
  }
)

func(oldPiece Piece) rotate(pivot Side, rotation Rotation) Piece {
  newPiece := Piece { }

  // Copy the face which we are pivoting around, and its opposite.
  opp := opposites[pivot]
  s, ok := oldPiece[pivot]; if ok { newPiece[pivot] = s }
  s, ok  = oldPiece[opp];   if ok { newPiece[opp]   = s }

  // Copy the other sides (those which are actually changing).
  for src, dest := range transforms[rotation] {
    face, ok := oldPiece[edges[pivot][src]]

    if ok {
      newPiece[edges[pivot][dest]] = face
    }
  }

  return newPiece
}

func main() {
  piece1 := Piece {  top: red,  front: green,  left: blue  }
  piece2 := piece1.rotate(top, clockwise)
  fmt.Println(piece1, "->", piece2)
}
