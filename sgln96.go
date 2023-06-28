package epctds

import (
	"errors"
	"fmt"

	"github.com/mzeiher/epctds-go/pkg/partition"
	"github.com/mzeiher/epctds-go/pkg/utils"
)

const (
	SGLN96Header = 0x32
)

var (
	sgln96_headerPartition          = partition.Partition{Start: 0, Length: 8, Digits: 5}
	sgln96_filterPartition          = partition.Partition{Start: 8, Length: 3, Digits: 1}
	sgln96_partitionNumberPartition = partition.Partition{Start: 11, Length: 3, Digits: 1}
	sgln96_extensionPartition       = partition.Partition{Start: 55, Length: 41, Digits: 13}
	sgln96_partition                = [7][2]partition.Partition{
		{{Start: 14, Length: 40, Digits: 12}, {Start: 54, Length: 1, Digits: 0}},
		{{Start: 14, Length: 37, Digits: 11}, {Start: 51, Length: 4, Digits: 1}},
		{{Start: 14, Length: 34, Digits: 10}, {Start: 48, Length: 7, Digits: 2}},
		{{Start: 14, Length: 30, Digits: 9}, {Start: 44, Length: 11, Digits: 3}},
		{{Start: 14, Length: 27, Digits: 8}, {Start: 41, Length: 14, Digits: 4}},
		{{Start: 14, Length: 24, Digits: 7}, {Start: 38, Length: 17, Digits: 5}},
		{{Start: 14, Length: 20, Digits: 6}, {Start: 34, Length: 21, Digits: 6}},
	}
)

type SGLN96 struct {
	EPCTag
	CompanyPrefix     int64
	LocationReference int64
	Extension         int64

	filter    int64
	partition int64
}

func (sgln SGLN96) ToTagURI() string {
	return fmt.Sprintf("urn:epc:tag:sgln-96:%d.%0*d.%0*d.%d", sgln.filter, sgln96_partition[sgln.partition][0].Digits, sgln.CompanyPrefix, sgln96_partition[sgln.partition][1].Digits, sgln.LocationReference, sgln.Extension)
}
func (sgln SGLN96) ToPureIdentityURI() string {
	return fmt.Sprintf("urn:epc:id:sgln:%0*d.%0*d.%d", sgln96_partition[sgln.partition][0].Digits, sgln.CompanyPrefix, sgln96_partition[sgln.partition][1].Digits, sgln.LocationReference, sgln.Extension)
}
func (sgln SGLN96) ToHex() (string, error) {
	sgln96Bytes := make([]byte, 12)
	sgln96Bytes, err := utils.PutInt64InBytes(SGLN96Header, sgln96Bytes, sgln96_headerPartition)
	if err != nil {
		return "", err
	}
	sgln96Bytes, err = utils.PutInt64InBytes(sgln.filter, sgln96Bytes, sgln96_filterPartition)
	if err != nil {
		return "", err
	}
	sgln96Bytes, err = utils.PutInt64InBytes(sgln.partition, sgln96Bytes, sgln96_partitionNumberPartition)
	if err != nil {
		return "", err
	}
	sgln96Bytes, err = utils.PutInt64InBytes(sgln.CompanyPrefix, sgln96Bytes, sgln96_partition[sgln.partition][0])
	if err != nil {
		return "", err
	}
	sgln96Bytes, err = utils.PutInt64InBytes(sgln.LocationReference, sgln96Bytes, sgln96_partition[sgln.partition][1])
	if err != nil {
		return "", err
	}
	sgln96Bytes, err = utils.PutInt64InBytes(sgln.Extension, sgln96Bytes, sgln96_extensionPartition)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%X", sgln96Bytes), nil
}

func sgln96FromBytes(epcBytes []byte) (SGLN96, error) {
	partitionNumber, err := utils.GetInt64FromBytes(epcBytes, sgln96_partitionNumberPartition)
	if err != nil {
		return SGLN96{}, err
	}
	filter, err := utils.GetInt64FromBytes(epcBytes, sgln96_filterPartition)
	if err != nil {
		return SGLN96{}, err
	}
	if partitionNumber >= int64(len(sgln96_partition)) {
		return SGLN96{}, errors.Join(partition.ErrInvalidPartition, fmt.Errorf("got partition %d", partitionNumber))
	}
	companyPrefix, err := utils.GetInt64FromBytes(epcBytes, sgln96_partition[partitionNumber][0])
	if err != nil {
		return SGLN96{}, err
	}
	serial, err := utils.GetInt64FromBytes(epcBytes, sgln96_partition[partitionNumber][1])
	if err != nil {
		return SGLN96{}, err
	}
	extension, err := utils.GetInt64FromBytes(epcBytes, sgln96_extensionPartition)
	if err != nil {
		return SGLN96{}, err
	}
	return SGLN96{CompanyPrefix: companyPrefix, LocationReference: serial, Extension: extension, filter: filter, partition: partitionNumber}, nil
}
