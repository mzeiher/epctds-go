package partition

import "errors"

var ErrInvalidPartition = errors.New("invalid partition")

type Partition struct {
	Start  int
	Length int
	Digits int
}
