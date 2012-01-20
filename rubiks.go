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

  North Direction = "N"
  East  Direction = "E"
  South Direction = "S"
  West  Direction = "W"

  Clockwise     Rotation = "CW"
  Anticlockwise Rotation = "ACW"

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


type Piece struct {
  faces [NUM_SIDES]Face
}


func(piece *Piece) to_s() string {
  var s = ""

  for n := 0; n < 6; n++ {
    s += string(piece.faces[n])
  }

  return s
}


func main() {
  piece := Piece{[6]Face{Red, Blank, Green, Blank, Blank, Yellow}}
  fmt.Println(piece.to_s())
}
