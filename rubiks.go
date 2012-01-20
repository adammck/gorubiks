package rubiks


type Face      string 
type Side      string
type Direction string
type Rotation  string
type Edge      map[Direction] Side
type Transform map[Direction] Direction
type Piece     map[Side] Face
type Cube      [27]Piece


var (

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

func(piece Piece) toString() string {
  s := ""

  for _, side := range sides {
    face, ok := piece[side]

    if ok {
      s += string(face)

    } else {
      s += "_"

    }
  }

  return s
}


func(cube Cube) piecesOn(side Side) [9]Piece {
  var pieces [9]Piece
  n := 0

  for _, piece := range cube {
    _, ok := piece[side]

    if ok {
      pieces[n] = piece
      n += 1
    }
  }

  return pieces
}

func(cube Cube) facesOn(side Side) [9]Face {
  var faces [9]Face

  for i, piece := range cube.piecesOn(side) {
    faces[i] = piece[side]
  }

  return faces
}

func (cube Cube) sideToString(side Side) string {
  s := ""

  for _, face := range cube.facesOn(side) {
    s += string(face)
  }

  return s
}

func(oldCube Cube) twist(side Side, direction Rotation) Cube {
  var newCube Cube

  for i, piece := range oldCube {
    _, ok := piece[side]

    // if +piece+ is on +side+.
    if ok {
      newCube[i] = piece.rotate(side, direction)

    } else {
      newCube[i] = piece
    }
  }

  return newCube
}
