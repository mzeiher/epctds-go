package utils

import (
	"errors"

	"github.com/mzeiher/epctds/pkg/partition"
)

var (
	ErrInvalidLength = errors.New("invalid length, partition to big for input")
)

func GetInt64FromBytes(input []byte, partition partition.Partition) (int64, error) {
	if (partition.Start+partition.Length)/8 > len(input) {
		return 0, ErrInvalidLength
	}
	var data int64
	currentOffsetBit := partition.Start + 1
	remainingBits := partition.Length
	for {
		currentByte := currentOffsetBit / 8
		bitsInByte := 8 - (currentOffsetBit % 8)
		if remainingBits == 0 {
			break
		} else if remainingBits < 8 {
			mask := byte(0xFF << (bitsInByte - remainingBits))
			shift := bitsInByte - remainingBits
			data |= int64(input[currentByte]&mask) >> shift
			break
		} else {
			mask := byte((0xFF >> (8 - bitsInByte)))
			shift := (remainingBits - bitsInByte)
			data |= int64(input[currentByte]&mask) << shift
		}
		remainingBits = remainingBits - bitsInByte
		currentOffsetBit = currentOffsetBit + bitsInByte

	}
	return data, nil

}
