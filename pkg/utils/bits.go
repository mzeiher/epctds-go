package utils

import (
	"errors"

	"github.com/mzeiher/epctds-go/pkg/partition"
)

var (
	ErrInvalidLength = errors.New("invalid length, partition to big for input")
)

func GetInt64FromBytes(input []byte, partition partition.Partition) (int64, error) {
	if (partition.Start+partition.Length)/8 > len(input) {
		return 0, ErrInvalidLength
	}
	var data int64
	currentOffsetBit := partition.Start
	remainingBits := partition.Length
	for {
		currentByte := currentOffsetBit / 8
		bitsInCurrentByte := 8 - (currentOffsetBit % 8)
		if remainingBits == 0 {
			break
		} else if remainingBits < 8 {
			mask := byte(0xFF << (bitsInCurrentByte - remainingBits))
			shift := bitsInCurrentByte - remainingBits
			data |= int64(input[currentByte]&mask) >> shift
			break
		} else {
			mask := byte((0xFF >> (8 - bitsInCurrentByte)))
			shift := (remainingBits - bitsInCurrentByte)
			data |= int64(input[currentByte]&mask) << shift
		}
		remainingBits = remainingBits - bitsInCurrentByte
		currentOffsetBit = currentOffsetBit + bitsInCurrentByte

	}
	return data, nil

}
