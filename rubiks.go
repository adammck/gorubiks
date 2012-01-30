package rubiks

import (
  "fmt"
)


/* The maximum depth which the solver will search to.
|* Some very rough benchmarks:
|*
|*  1 (12)     =   0.00
|*  2 (144)    =   0.05
|*  3 (1728)   =   1.11403
|*  4 (20736)  =  12.27251
|*  5 (248832) = 151.75544 */
const MAX_DEPTH = 4


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
  return Piece {
    edges[pivot][transforms[rotation][north]]: oldPiece[edges[pivot][north]],
    edges[pivot][transforms[rotation][south]]: oldPiece[edges[pivot][south]],
    edges[pivot][transforms[rotation][east]]:  oldPiece[edges[pivot][east]],
    edges[pivot][transforms[rotation][west]]:  oldPiece[edges[pivot][west]],
    opposites[pivot]:                          oldPiece[opposites[pivot]],
    pivot:                                     oldPiece[pivot],
  }
}

func(piece Piece) toString() string {
  s := ""

  for _, side := range sides {
    face := piece[side]

    if face != "" {
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

  if piece[edges[side][north]] != "" { i -= 3 }
  if piece[edges[side][east]]  != "" { i += 1 }
  if piece[edges[side][south]] != "" { i += 3 }
  if piece[edges[side][west]]  != "" { i -= 1 }

  return i
}

func (cube Cube) piecesOn(side Side) [9]Piece {
  var pieces [9]Piece

  for _, piece := range cube {
    if piece[side] != "" {
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
    if piece[side] != "" {
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
