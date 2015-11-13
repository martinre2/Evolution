package Stage

import (
	b "github.com/martinre2/Evolution/Board"
)

type Stage interface {
	Execute(board *b.BlackBoard) bool
	String() string
}
