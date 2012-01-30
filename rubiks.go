package rubiks

import (
  "fmt"
  "container/vector"
)


const MAX_DEPTH = 5


type Face      int
type Side      int
type Direction int
type Rotation  int
type Edge      [4]Side
type Transform [4]Direction
type Piece     [6]Face
type Cube      [27]Piece

type Move struct {
  side     Side
  rotation Rotation
}

// Faces
const (
  blank  Face = iota
  red    Face = iota
  green  Face = iota
  blue   Face = iota
  yellow Face = iota
  orange Face = iota
  white  Face = iota
)

// Sides
const (
  top    Side = iota
  bottom Side = iota
  front  Side = iota
  back   Side = iota
  left   Side = iota
  right  Side = iota
)

// Directions
const (
  north Direction = iota
  east  Direction = iota
  south Direction = iota
  west  Direction = iota
)

// Rotations
const (
  once   Rotation = iota
  twice  Rotation = iota
  thrice Rotation = iota
)

var (
  sides     = [...] Side     { top, bottom, front, back, left, right }
  rotations = [...] Rotation { once, twice, thrice }

  opposites = map [Side] Side {
    top:    bottom,
    bottom: top,
    front:  back,
    back:   front,
    left:   right,
    right:  left,
  }

  edges = [...] Edge {
    /*      North   East    South    West
    |*      -----   ----    -----    ---- */
    Edge {  back,   right,  front,   left   }, // Top
    Edge {  front,  left,   back,    right  }, // Bottom
    Edge {  top,    right,  bottom,  left   }, // Front
    Edge {  top,    left,   bottom,  right  }, // Back
    Edge {  top,    front,  bottom,  back   }, // Left
    Edge {  top,    back,   bottom,  front  }, // Right
  }

  transforms = [...] Transform {
    /*                   North   East    South   West
    |*                   -----   ----    -----   ---- */
    once:   Transform {  east,   south,  west,   north  },
    twice:  Transform {  south,  west,   north,  east   },
    thrice: Transform {  west,   north,  east,   south  },
  }
)


// -- Face methods ------------------------------------------------------------

func (face Face) toString() string {
  switch face {
    case red:    return "R"
    case green:  return "G"
    case blue:   return "B"
    case yellow: return "Y"
    case orange: return "O"
    case white:  return "W"
  }

  return "_"
}


// -- Piece methods -----------------------------------------------------------

func(piece *Piece) rotate(pivot Side, rotation Rotation) {
  piece[edges[pivot][transforms[rotation][north]]], piece[edges[pivot][transforms[rotation][south]]], piece[edges[pivot][transforms[rotation][east]]], piece[edges[pivot][transforms[rotation][west]]] = piece[edges[pivot][north]], piece[edges[pivot][south]], piece[edges[pivot][east]], piece[edges[pivot][west]]
}

func(piece *Piece) toString() string {
  s := ""

  for _, side := range sides {
    s += piece[side].toString()
  }

  return s
}


// -- Cube methods -----------------------------------------------------------

func (cube *Cube) toString() string {
  s := ""

  for i, side := range sides {
    if i > 0 { s += " " }
    s += cube.sideToString(side)
  }

  return s
}

func (cube *Cube) sideToString(side Side) string {
  s := ""

  for _, face := range cube.facesOn(side) {
    s += face.toString()
  }

  return s
}

func (cube *Cube) isEqual(other *Cube) bool {
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

func (cube *Cube) isSolved() bool {
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

func faceIndex(piece *Piece, side Side) int {
  i := 4

  if piece[edges[side][north]] != blank { i -= 3 }
  if piece[edges[side][east]]  != blank { i += 1 }
  if piece[edges[side][south]] != blank { i += 3 }
  if piece[edges[side][west]]  != blank { i -= 1 }

  return i
}

func (cube *Cube) facesOn(side Side) [9]Face {
  var faces [9]Face

  for n := range cube {
    if cube[n][side] != blank {

      faces[faceIndex(&cube[n], side)] = cube[n][side]

    }
  }

  return faces
}

func (cube *Cube) twist(side Side, direction Rotation) {
  for n := range cube {
    if cube[n][side] != blank {
      cube[n].rotate(side, direction)
    }
  }
}

func (cube *Cube) untwist(side Side, direction Rotation) {
  switch(direction) {
    case once:   cube.twist(side, thrice)
    case twice:  cube.twist(side, twice)
    case thrice: cube.twist(side, once)
  }
}


// -- Solver functions --------------------------------------------------------

func findRouteByForce(src Cube, dest Cube) bool {
  fmt.Println("src =", src.toString())
  fmt.Println("dst =", dest.toString())

  //stack = make(MoveList, 0, MAX_DEPTH)
  stack := new(vector.Vector)
  return doFindRouteByForce(&src, &dest, stack)
}

func doFindRouteByForce (src *Cube, dest *Cube, stack *vector.Vector) bool {

  if src.isEqual(dest) {
    fmt.Println("Solved:", stack)
    return true
  }

  if stack.Len() >= MAX_DEPTH {
    return false
  }

  for _, side := range sides {
    for _, direction := range rotations {

      stack.Push(Move { side, direction })
      src.twist(side, direction)

      if doFindRouteByForce(src, dest, stack) == true {
        return true
      }

      src.untwist(side, direction)
      stack.Pop()
    }
  }

  return false
}
