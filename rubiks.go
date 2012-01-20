package main

import (
  "fmt"
)

const NUM_SIDES = 6

type Face      string 
type Side      string
type Direction string
type Rotation  string
type Edge      map[Direction] Side
type Transform map[Direction] Direction
type Piece     map[Side] Face

var (

  Blank  Face = "_"
  Red    Face = "R"
  Green  Face = "G"
  Blue   Face = "B"
  Yellow Face = "Y"
  Orange Face = "O"
  White  Face = "W"

  Top    Side = "Top"
  Bottom Side = "Bottom"
  Front  Side = "Front"
  Back   Side = "Back"
  Left   Side = "Left"
  Right  Side = "Right"
  Sides = [...]Side { Top, Bottom, Front, Back, Left, Right }

  North Direction = "N"
  East  Direction = "E"
  South Direction = "S"
  West  Direction = "W"

  Clockwise     Rotation = "CW"
  Anticlockwise Rotation = "ACW"
  
  Opposites = map [Side] Side {
    Top:    Bottom,
    Bottom: Top,
    Front:  Back,
    Back:   Front,
    Left:   Right,
    Right:  Left,
  }

  Edges = map [Side] Edge {
    Top:    Edge {  North: Back,   South: Front,   East: Right,  West: Left   },
    Bottom: Edge {  North: Front,  South: Back,    East: Left,   West: Right  },
    Front:  Edge {  North: Top,    South: Bottom,  East: Right,  West: Left   },
    Back:   Edge {  North: Top,    South: Bottom,  East: Left,   West: Right  },
    Left:   Edge {  North: Top,    South: Bottom,  East: Front,  West: Back   },
    Right:  Edge {  North: Top,    South: Bottom,  East: Back,   West: Front  },
  }

  Transforms = map [Rotation] Transform {
    Clockwise:     Transform {  West: North,  North: East,  East: South,  South: West  },
    Anticlockwise: Transform {  East: North,  South: East,  West: South,  North: West  },
  }
)

func(old_piece Piece) rotate(pivot Side, rotation Rotation) Piece {
  new_piece := Piece { }

  // Copy the face which we are pivoting around, and its opposite.
  opp := Opposites[pivot]
  s, ok := old_piece[pivot]; if ok { new_piece[pivot] = s }
  s, ok  = old_piece[opp];   if ok { new_piece[opp]  = s }

  // Copy the other sides (those which are actually changing).
  for from_dir, to_dir := range Transforms[rotation] {
    face, ok := old_piece[Edges[pivot][from_dir]]

    if ok {
      new_piece[Edges[pivot][to_dir]] = face
    }
  }

  return new_piece
}

func main() {
  piece1 := Piece {  Top: Red,  Front: Green,  Left: Blue  }
  piece2 := piece1.rotate(Top, Clockwise)
  fmt.Println(piece1, "->", piece2)
}
