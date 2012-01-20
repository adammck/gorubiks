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
