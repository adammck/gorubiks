package rubiks

import (
  "testing"
)


// -- Helpers -----------------------------------------------------------------

func testPiece() Piece {
  return Piece {
    top:    red,
    bottom: green,
    front:  blue,
    back:   yellow,
    left:   orange,
    right:  white,
  }
}

func testCube() Cube {
  return Cube {
    //       TOP,   BOTTM, FRONT, BACK,   LEFT, RIGHT
    //       ---    -----  -----  ----    ----  -----
    Piece {  red,   blank, blank, yellow, orange, blank },
    Piece {  red,   blank, blank, yellow, blank,  blank },
    Piece {  red,   blank, blank, yellow, blank,  white },
    Piece {  red,   blank, blank, blank,  orange, blank },
    Piece {  red,   blank, blank, blank,  blank,  blank },
    Piece {  red,   blank, blank, blank,  blank,  white },
    Piece {  red,   blank, blue,  blank,  orange, blank },
    Piece {  red,   blank, blue,  blank,  blank,  blank },
    Piece {  red,   blank, blue,  blank,  blank,  white },
    Piece {  blank, blank, blank, yellow, orange, blank },
    Piece {  blank, blank, blank, yellow, blank,  blank },
    Piece {  blank, blank, blank, yellow, blank,  white },
    Piece {  blank, blank, blank, blank,  orange, blank },
    Piece {  blank, blank, blank, blank,  blank,  blank },
    Piece {  blank, blank, blank, blank,  blank,  white },
    Piece {  blank, blank, blue,  blank,  orange, blank },
    Piece {  blank, blank, blue,  blank,  blank,  blank },
    Piece {  blank, blank, blue,  blank,  blank,  white },
    Piece {  blank, green, blank, yellow, orange, blank },
    Piece {  blank, green, blank, yellow, blank,  blank },
    Piece {  blank, green, blank, yellow, blank,  white },
    Piece {  blank, green, blank, blank,  orange, blank },
    Piece {  blank, green, blank, blank,  blank,  blank },
    Piece {  blank, green, blank, blank,  blank,  white },
    Piece {  blank, green, blue,  blank,  orange, blank },
    Piece {  blank, green, blue,  blank,  blank,  blank },
    Piece {  blank, green, blue,  blank,  blank,  white },
  }
}


// -- Piece tests -------------------------------------------------------------

func TestPieceToString(t *testing.T) {
  corner := Piece { top: red, front: blue, left: orange }
  edge   := Piece { top: red, front: blue }
  middle := Piece { top: red }
  center := Piece { }

  if corner.toString() != "R_B_O_" { t.Error("CORNER should be 'R_B_O_'.") }
  if edge.toString()   != "R_B___" { t.Error("EDGE should be 'R_B___'.")   }
  if middle.toString() != "R_____" { t.Error("MIDDLE should be 'R_____'.") }
  if center.toString() != "______" { t.Error("CENTER should be '______'.") }
}

func TestPieceRotateTopClockwise(t *testing.T) {
  piece := testPiece()
  piece.rotate(top, clockwise)

  if piece[top]    != red    { t.Error("TOP should be RED.")      }
  if piece[bottom] != green  { t.Error("BOTTOM should be GREEN.") }
  if piece[front]  != white  { t.Error("FRONT should be WHITE.")  }
  if piece[back]   != orange { t.Error("BACK should be ORANGE.")  }
  if piece[left]   != blue   { t.Error("LEFT should be BLUE.")    }
  if piece[right]  != yellow { t.Error("RIGHT should be YELLOW.") }
}

func TestPieceRotateTopAnticlockwise(t *testing.T) {
  piece := testPiece()
  piece.rotate(top, anticlockwise)

  if piece[top]    != red    { t.Error("TOP should be RED.")      }
  if piece[bottom] != green  { t.Error("BOTTOM should be GREEN.") }
  if piece[front]  != orange { t.Error("FRONT should be ORANGE.") }
  if piece[back]   != white  { t.Error("BACK should be WHITE.")   }
  if piece[left]   != yellow { t.Error("LEFT should be YELLOW.")  }
  if piece[right]  != blue   { t.Error("RIGHT should be BLUE.")   }
}

func TestPieceRotateFrontClockwise(t *testing.T) {
  piece := testPiece()
  piece.rotate(front, clockwise)

  if piece[top]    != orange { t.Error("TOP should be ORANGE.")   }
  if piece[bottom] != white  { t.Error("BOTTOM should be WHITE.") }
  if piece[front]  != blue   { t.Error("FRONT should be BLUE.")   }
  if piece[back]   != yellow { t.Error("BACK should be YELLOW.")  }
  if piece[left]   != green  { t.Error("LEFT should be GREEN.")   }
  if piece[right]  != red    { t.Error("RIGHT should be RED.")    }
}

func TestPieceRotateFrontAnticlockwise(t *testing.T) {
  piece := testPiece()
  piece.rotate(front, anticlockwise)

  if piece[top]    != white  { t.Error("TOP should be WHITE.")     }
  if piece[bottom] != orange { t.Error("BOTTOM should be ORANGE.") }
  if piece[front]  != blue   { t.Error("FRONT should be BLUE.")    }
  if piece[back]   != yellow { t.Error("BACK should be YELLOW.")   }
  if piece[left]   != red    { t.Error("LEFT should be RED.")      }
  if piece[right]  != green  { t.Error("RIGHT should be GREEN.")   }
}

func TestPieceRotateLeftClockwise(t *testing.T) {
  piece := testPiece()
  piece.rotate(left, clockwise)
  if piece[top]    != yellow { t.Error("TOP should be YELLOW.")   }
  if piece[bottom] != blue   { t.Error("BOTTOM should be BLUE.")  }
  if piece[front]  != red    { t.Error("FRONT should be RED.")    }
  if piece[back]   != green  { t.Error("BACK should be GREEN.")   }
  if piece[left]   != orange { t.Error("LEFT should be ORANGE.")  }
  if piece[right]  != white  { t.Error("RIGHT should be WHITE.")  }
}

func TestPieceRotateLeftAnticlockwise(t *testing.T) {
  piece := testPiece()
  piece.rotate(left, anticlockwise)
  if piece[top]    != blue   { t.Error("TOP should be BLUE.")      }
  if piece[bottom] != yellow { t.Error("BOTTOM should be YELLOW.") }
  if piece[front]  != green  { t.Error("FRONT should be GREEN.")   }
  if piece[back]   != red    { t.Error("BACK should be RED.")      }
  if piece[left]   != orange { t.Error("LEFT should be ORANGE.")   }
  if piece[right]  != white  { t.Error("RIGHT should be WHITE.")   }
}


// -- Cube tests --------------------------------------------------------------

func TestCubeToString(t *testing.T) {
  test_cube := testCube()

  if test_cube.toString() != "RRRRRRRRR GGGGGGGGG BBBBBBBBB YYYYYYYYY OOOOOOOOO WWWWWWWWW" {
    t.Error("Cube.toString was wrong.")
  }
}

func TestCubeSideToString(t *testing.T) {
  test_cube := testCube()

  if test_cube.sideToString(top)   != "RRRRRRRRR" { t.Error("T != RRRRRRRRR") }
  if test_cube.sideToString(front) != "BBBBBBBBB" { t.Error("F != BBBBBBBBB") }
  if test_cube.sideToString(left)  != "OOOOOOOOO" { t.Error("L != OOOOOOOOO") }
}

func TestCubeEquality(t *testing.T) {
  test_cube_one := testCube()
  test_cube_two := testCube()

  if test_cube_one.isEqual(&test_cube_two) != true {
    t.Error("Cube.isEqual was wrong.")
  }
}

func TestCubeFacesOn(t *testing.T) {
  test_cube := testCube()

  top_faces := test_cube.facesOn(top)
  if top_faces[0] != red { t.Error("T0 != R") }
  if top_faces[1] != red { t.Error("T1 != R") }
  if top_faces[2] != red { t.Error("T2 != R") }

  front_faces := test_cube.facesOn(front)
  if front_faces[0] != blue { t.Error("F0 != B") }
  if front_faces[1] != blue { t.Error("F1 != B") }
  if front_faces[2] != blue { t.Error("F2 != B") }

  left_faces := test_cube.facesOn(left)
  if left_faces[0] != orange { t.Error("T0 != O") }
  if left_faces[1] != orange { t.Error("T1 != O") }
  if left_faces[2] != orange { t.Error("L2 != O") }
}


// -- Cube twisting tests -----------------------------------------------------

func TestCubeIsSolved(t *testing.T) {
  test_cube := testCube()

  if test_cube.isSolved() != true {
    t.Error("Cube should be solved.")
  }

  test_cube.twist(top, clockwise)

  if test_cube.isSolved() != false {
    t.Error("Cube should not be solved.")
  }
}

func TestCubeTwistTopClockwise(t *testing.T) {
  test_cube := testCube()

  test_cube.twist(top, clockwise)
  if test_cube.sideToString(top)   != "RRRRRRRRR" { t.Error("T != RRRRRRRRR") }
  if test_cube.sideToString(front) != "WWWBBBBBB" { t.Error("F != WWWBBBBBB") }
  if test_cube.sideToString(left)  != "BBBOOOOOO" { t.Error("L != BBBOOOOOO") }

  test_cube.twist(top, clockwise)
  if test_cube.sideToString(top)   != "RRRRRRRRR" { t.Error("T != RRRRRRRRR") }
  if test_cube.sideToString(front) != "YYYBBBBBB" { t.Error("F != YYYRRRRRR") }
  if test_cube.sideToString(left)  != "WWWOOOOOO" { t.Error("L != WWWOOOOOO") }

  test_cube.twist(top, clockwise)
  if test_cube.sideToString(top)   != "RRRRRRRRR" { t.Error("T != RRRRRRRRR") }
  if test_cube.sideToString(front) != "OOOBBBBBB" { t.Error("F != OOORRRRRR") }
  if test_cube.sideToString(left)  != "YYYOOOOOO" { t.Error("L != YYYOOOOOO") }
}

func TestCubeTwistFrontClockwise(t *testing.T) {
  test_cube := testCube()

  test_cube.twist(front, clockwise)
  if test_cube.sideToString(top)   != "RRRRRROOO" { t.Error("T != RRRRRROOO") }
  if test_cube.sideToString(front) != "BBBBBBBBB" { t.Error("F != BBBBBBBBB") }
  if test_cube.sideToString(left)  != "OOGOOGOOG" { t.Error("L != OOGOOGOOG") }

  test_cube.twist(front, clockwise)
  if test_cube.sideToString(top)   != "RRRRRRGGG" { t.Error("T != RRRRRRGGG") }
  if test_cube.sideToString(front) != "BBBBBBBBB" { t.Error("F != BBBBBBBBB") }
  if test_cube.sideToString(left)  != "OOWOOWOOW" { t.Error("L != OOWOOWOOW") }

  test_cube.twist(front, clockwise)
  if test_cube.sideToString(top)   != "RRRRRRWWW" { t.Error("T != RRRRRRWWW") }
  if test_cube.sideToString(front) != "BBBBBBBBB" { t.Error("F != BBBBBBBBB") }
  if test_cube.sideToString(left)  != "OOROOROOR" { t.Error("L != OOROOROOR") }
}

func TestCubeTwistLeftClockwise(t *testing.T) {
  test_cube := testCube()

  test_cube.twist(left, clockwise)
  if test_cube.sideToString(top)   != "YRRYRRYRR" { t.Error("T != YRRYRRYRR") }
  if test_cube.sideToString(front) != "RBBRBBRBB" { t.Error("F != RBBRBBRBB") }
  if test_cube.sideToString(left)  != "OOOOOOOOO" { t.Error("L != OOOOOOOOO") }

  test_cube.twist(left, clockwise)
  if test_cube.sideToString(top)   != "GRRGRRGRR" { t.Error("T != GRRGRRGRR") }
  if test_cube.sideToString(front) != "YBBYBBYBB" { t.Error("F != YBBYBBYBB") }
  if test_cube.sideToString(left)  != "OOOOOOOOO" { t.Error("L != OOOOOOOOO") }

  test_cube.twist(left, clockwise)
  if test_cube.sideToString(top)   != "BRRBRRBRR" { t.Error("T != BRRBRRBRR") }
  if test_cube.sideToString(front) != "GBBGBBGBB" { t.Error("F != GBBGBBGBB") }
  if test_cube.sideToString(left)  != "OOOOOOOOO" { t.Error("L != OOOOOOOOO") }
}


// -- Solver tests ------------------------------------------------------------

func TestFindRouteByForce(t *testing.T) {
  scrambledCube := testCube()

  scrambledCube.twist(top, clockwise)
  scrambledCube.twist(left, anticlockwise)
  scrambledCube.twist(bottom, clockwise)
  scrambledCube.twist(right, anticlockwise)

  scrambledCube.twist(top, clockwise)
  scrambledCube.twist(left, anticlockwise)
  scrambledCube.twist(bottom, clockwise)
  scrambledCube.twist(right, anticlockwise)

  if findRouteByForce(scrambledCube, testCube()) != true {
    t.Error("Couldn't find route.")
  }
}
