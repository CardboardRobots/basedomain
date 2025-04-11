package domain

import "github.com/cardboardrobots/baseerror"

var (
	ErrEventNoMatch   = baseerror.ErrFailedPrecondition.Wrap("event not matched")
	ErrPreviouslySeen = baseerror.ErrFailedPrecondition.Wrap("previously seen")
)

func NewErrEventNoMatch(name string) error {
	return ErrEventNoMatch.Wrap(name)
}
