package epctds

import (
	"errors"
	"fmt"

	"github.com/mzeiher/epctds-go/pkg/partition"
	"github.com/mzeiher/epctds-go/pkg/utils"
)

var ErrInvalidHeader = errors.New("invalid header value")

type EPCTag interface {
	ToTagURI() string
	ToPureIdentityURI() string
}

func ParseEpcTagData(hexString string) (EPCTag, error) {
	epcBytes, err := utils.GetEpcBytes(hexString)
	if err != nil {
		return nil, err
	}
	header, err := utils.GetInt64FromBytes(epcBytes, partition.Partition{Start: 0, Length: 8, Digits: 3})
	if err != nil {
		return nil, err
	}
	switch header {
	case sscc96_header:
		return sscc69FromByes(epcBytes)
	}

	return nil, errors.Join(ErrInvalidHeader, fmt.Errorf("got header %x", header))
}
