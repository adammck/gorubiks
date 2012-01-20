package rubiks

import (
  "testing"
)

var (
  test_piece = Piece {
    top:    red,
    bottom: green,
    front:  blue,
    back:   yellow,
    left:   orange,
    right:  white,
  }

  test_cube = Cube {
    Piece {  top: red,       back: yellow,  left: orange  },
    Piece {  top: red,       back: yellow                 },
    Piece {  top: red,       back: yellow,  right: white  },
    Piece {  top: red,                      left: orange  },
    Piece {  top: red                                     },
    Piece {  top: red,                      right: white  },
    Piece {  top: red,       front: blue,   left: orange  },
    Piece {  top: red,       front: blue                  },
    Piece {  top: red,       front: blue,   right: white  },
    Piece {                  back: yellow,  left: orange  },
    Piece {                  back: yellow                 },
    Piece {                  back: yellow,  right: white  },
    Piece {                                 left: orange  },
    Piece {                                               },
    Piece {                                 right: white  },
    Piece {                  front: blue,   left: orange  },
    Piece {                  front: blue                  },
    Piece {                  front: blue,   right: white  },
    Piece {  bottom: green,  back: yellow,  left: orange  },
    Piece {  bottom: green,  back: yellow                 },
    Piece {  bottom: green,  back: yellow,  right: white  },
    Piece {  bottom: green,                 left: orange  },
    Piece {  bottom: green                                },
    Piece {  bottom: green,                 right: white  },
    Piece {  bottom: green,  front: blue,   left: orange  },
    Piece {  bottom: green,  front: blue                  },
    Piece {  bottom: green,  front: blue,   right: white  },
  }
)


// -- Piece tests -------------------------------------------------------------

func TestPieceToString(t *testing.T) {
  corner := Piece { top: red, front: blue, left: orange }
  edge   := Piece { top: red, front: blue }
  middle := Piece { top: red }
  center := Piece { }

  if corner.toString() != "R_B_O_" { t.Error("CORNER should be 'R_B_O_'.") }
  if edge.toString()   != "R_B___" { t.Error("EDGE should be 'R_B___'.") }
  if middle.toString() != "R_____" { t.Error("MIDDLE should be 'R_____'.") }
  if center.toString() != "______" { t.Error("CENTER should be '______'.") }
}

func TestPieceRotateTopClockwise(t *testing.T) {
  piece := test_piece.rotate(top, clockwise)
  if piece[top]    != red    { t.Error("TOP should be RED.")      }
  if piece[bottom] != green  { t.Error("BOTTOM should be GREEN.") }
  if piece[front]  != white  { t.Error("FRONT should be WHITE.")  }
  if piece[back]   != orange { t.Error("BACK should be ORANGE.")  }
  if piece[left]   != blue   { t.Error("LEFT should be BLUE.")    }
  if piece[right]  != yellow { t.Error("RIGHT should be YELLOW.") }
}

func TestPieceRotateTopAnticlockwise(t *testing.T) {
  piece := test_piece.rotate(top, anticlockwise)
  if piece[top]    != red    { t.Error("TOP should be RED.")      }
  if piece[bottom] != green  { t.Error("BOTTOM should be GREEN.") }
  if piece[front]  != orange { t.Error("FRONT should be ORANGE.") }
  if piece[back]   != white  { t.Error("BACK should be WHITE.")   }
  if piece[left]   != yellow { t.Error("LEFT should be YELLOW.")  }
  if piece[right]  != blue   { t.Error("RIGHT should be BLUE.")   }
}

func TestPieceRotateFrontClockwise(t *testing.T) {
  piece := test_piece.rotate(front, clockwise)
  if piece[top]    != orange { t.Error("TOP should be ORANGE.")   }
  if piece[bottom] != white  { t.Error("BOTTOM should be WHITE.") }
  if piece[front]  != blue   { t.Error("FRONT should be BLUE.")   }
  if piece[back]   != yellow { t.Error("BACK should be YELLOW.")  }
  if piece[left]   != green  { t.Error("LEFT should be GREEN.")   }
  if piece[right]  != red    { t.Error("RIGHT should be RED.")    }
}

func TestPieceRotateFrontAnticlockwise(t *testing.T) {
  piece := test_piece.rotate(front, anticlockwise)
  if piece[top]    != white  { t.Error("TOP should be WHITE.")     }
  if piece[bottom] != orange { t.Error("BOTTOM should be ORANGE.") }
  if piece[front]  != blue   { t.Error("FRONT should be BLUE.")    }
  if piece[back]   != yellow { t.Error("BACK should be YELLOW.")   }
  if piece[left]   != red    { t.Error("LEFT should be RED.")      }
  if piece[right]  != green  { t.Error("RIGHT should be GREEN.")   }
}

func TestPieceRotateLeftClockwise(t *testing.T) {
  piece := test_piece.rotate(left, clockwise)
  if piece[top]    != yellow { t.Error("TOP should be YELLOW.")   }
  if piece[bottom] != blue   { t.Error("BOTTOM should be BLUE.")  }
  if piece[front]  != red    { t.Error("FRONT should be RED.")    }
  if piece[back]   != green  { t.Error("BACK should be GREEN.")   }
  if piece[left]   != orange { t.Error("LEFT should be ORANGE.")  }
  if piece[right]  != white  { t.Error("RIGHT should be WHITE.")  }
}

func TestPieceRotateLeftAnticlockwise(t *testing.T) {
  piece := test_piece.rotate(left, anticlockwise)
  if piece[top]    != blue   { t.Error("TOP should be BLUE.")      }
  if piece[bottom] != yellow { t.Error("BOTTOM should be YELLOW.") }
  if piece[front]  != green  { t.Error("FRONT should be GREEN.")   }
  if piece[back]   != red    { t.Error("BACK should be RED.")      }
  if piece[left]   != orange { t.Error("LEFT should be ORANGE.")   }
  if piece[right]  != white  { t.Error("RIGHT should be WHITE.")   }
}


// -- Cube tests --------------------------------------------------------------

func TestCubePiecesOn(t *testing.T) {
  top_pieces := test_cube.piecesOn(top)
  if top_pieces[0].toString() != "R__YO_" { t.Error("TP[0] should be R__YO_") }
  if top_pieces[1].toString() != "R__Y__" { t.Error("TP[1] should be R__Y__") }
  if top_pieces[2].toString() != "R__Y_W" { t.Error("TP[2] should be R__Y_W") }

  front_pieces := test_cube.piecesOn(front)
  if front_pieces[3].toString() != "__B_O_" { t.Error("FP[0] should be __B_O_") }
  if front_pieces[4].toString() != "__B___" { t.Error("FP[1] should be __B___") }
  if front_pieces[5].toString() != "__B__W" { t.Error("FP[2] should be __B__W") }

  left_pieces := test_cube.piecesOn(left)
  if left_pieces[6].toString() != "_G_YO_" { t.Error("LP[0] should be _G_YO_") }
  if left_pieces[7].toString() != "_G__O_" { t.Error("LP[1] should be _G__O_") }
  if left_pieces[8].toString() != "_GB_O_" { t.Error("LP[2] should be _GB_O_") }
}

func TestCubeFacesOn(t *testing.T) {
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

func TestCubeSideToString(t *testing.T) {
  if test_cube.sideToString(top)   != "RRRRRRRRR" { t.Error("TOP   should be RRRRRRRRR") }
  if test_cube.sideToString(front) != "BBBBBBBBB" { t.Error("FRONT should be BBBBBBBBB") }
  if test_cube.sideToString(left)  != "OOOOOOOOO" { t.Error("LEFT  should be OOOOOOOOO") }
}
