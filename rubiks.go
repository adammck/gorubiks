package rubiks

import (
  "fmt"
)


const MAX_DEPTH = 10


type Face      string 
type Side      string
type Direction string
type Rotation  string
type Edge      map[Direction] Side
type Transform map[Direction] Direction
type Piece     map[Side] Face
type Cube      [27]Piece

type Move struct {
  side     Side
  rotation Rotation
}

type MoveList []Move


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
  rotations = [...]Rotation { clockwise, anticlockwise }
  
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


// -- Piece methods -----------------------------------------------------------

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


// -- Cube methods -----------------------------------------------------------

func (cube Cube) toString() string {
  s := ""

  for i, side := range sides {
    if i > 0 { s += " " }
    s += cube.sideToString(side)
  }

  return s
}

func (cube Cube) sideToString(side Side) string {
  s := ""

  for _, face := range cube.facesOn(side) {
    s += string(face)
  }

  return s
}

func (cube Cube) isEqual(other Cube) bool {
  for _, side := range sides {
    faces := cube.facesOn(side)

    for i, face := range other.facesOn(side) {
      if faces[i] != face {
        return false
      }
    }
  }

  return true
}

func (cube Cube) isSolved() bool {
  for _, side := range sides {
    faces := cube.facesOn(side)
    centerFaceColor := faces[4]

    for _, face := range faces {
      if face != centerFaceColor {
        return false
      }
    }
  }

  return true
}

func pieceIndex(piece Piece, side Side) int {
  i := 4

  if _, ok := piece[edges[side][north]]; ok { i -= 3 }
  if _, ok := piece[edges[side][east]];  ok { i += 1 }
  if _, ok := piece[edges[side][south]]; ok { i += 3 }
  if _, ok := piece[edges[side][west]];  ok { i -= 1 }

  return i
}

func (cube Cube) piecesOn(side Side) [9]Piece {
  var pieces [9]Piece

  for _, piece := range cube {
    _, ok := piece[side]

    if ok {
      i := pieceIndex(piece, side)
      pieces[i] = piece
    }
  }

  return pieces
}

func (cube Cube) facesOn(side Side) [9]Face {
  var faces [9]Face

  for i, piece := range cube.piecesOn(side) {
    faces[i] = piece[side]
  }

  return faces
}

func (oldCube Cube) twist(side Side, direction Rotation) Cube {
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


// -- Solver functions --------------------------------------------------------

func findRouteByForce(src Cube, dest Cube) bool {
  fmt.Println("src =", src.toString())
  fmt.Println("dst =", dest.toString())

  return doFindRouteByForce(src, dest, make(MoveList, 0)) != nil
}

func doFindRouteByForce(src Cube, dest Cube, stack MoveList) MoveList {

  // If the cube is solved (i.e. src=dest), there's nothing for us to do.
  if src.isEqual(dest) {
    fmt.Println("Solved:", stack)
    return stack
  }

  // If the stack is too deep, abort.
  if len(stack) > MAX_DEPTH {
    return nil
  }

  // Iterate through all of the possible moves.
  for _, side := range sides {
    for _, direction := range rotations {

      // Build a new cube for this move.
      thisStack := append(stack, Move { side, direction })
      thisCube := src.twist(side, direction)

      /* Recurse, to:
      |*  (a) check if we found a solution,
      |*  (b) continue searching, if not. */
      thisMoveList := doFindRouteByForce(thisCube, dest, thisStack)

      /* We found a solution (somewhere down the stack)! Allow the return value
      |* (a list of moves) to propagate, to eventually return the solution to
      |* the original caller. */
      if thisMoveList != nil {
        return thisMoveList
      }
    }
  }

  // Didn't find a solution, so return nothing.
  return nil
}
