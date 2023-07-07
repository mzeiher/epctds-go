package epctds

import (
	"errors"
	"fmt"

	"github.com/mzeiher/epctds-go/pkg/partition"
	"github.com/mzeiher/epctds-go/pkg/utils"
)

const SSCC96Header = 0x31

var (
	sscc96_headerPartition          = partition.Partition{Start: 0, Length: 8, Digits: 5}
	sscc96_partitionNumberPartition = partition.Partition{Start: 11, Length: 3, Digits: 1}
	sscc96_filterPartition          = partition.Partition{Start: 8, Length: 3, Digits: 1}
	sscc96_dataPartitions           = [7][2]partition.Partition{
		{{Start: 14, Length: 40, Digits: 12}, {Start: 54, Length: 18, Digits: 5}},
		{{Start: 14, Length: 37, Digits: 11}, {Start: 51, Length: 21, Digits: 6}},
		{{Start: 14, Length: 34, Digits: 10}, {Start: 48, Length: 24, Digits: 7}},
		{{Start: 14, Length: 30, Digits: 9}, {Start: 44, Length: 28, Digits: 8}},
		{{Start: 14, Length: 27, Digits: 8}, {Start: 41, Length: 31, Digits: 9}},
		{{Start: 14, Length: 24, Digits: 7}, {Start: 38, Length: 34, Digits: 10}},
		{{Start: 14, Length: 20, Digits: 6}, {Start: 34, Length: 38, Digits: 11}},
	}
)

type SSCC96 struct {
	EPCTag
	CompanyPrefix int64
	Serial        int64

	filter    int64
	partition int64
}

func (sscc SSCC96) ToTagURI() string {
	return fmt.Sprintf("urn:epc:tag:sscc-96:%d.%0*d.%0*d", sscc.filter, sscc96_dataPartitions[sscc.partition][0].Digits, sscc.CompanyPrefix, sscc96_dataPartitions[sscc.partition][1].Digits, sscc.Serial)
}
func (sscc SSCC96) ToPureIdentityURI() string {
	return fmt.Sprintf("urn:epc:id:sscc:%0*d.%0*d", sscc96_dataPartitions[sscc.partition][0].Digits, sscc.CompanyPrefix, sscc96_dataPartitions[sscc.partition][1].Digits, sscc.Serial)
}
func (sscc SSCC96) ToHex() (string, error) {
	sscc96Bytes := make([]byte, 12)
	sscc96Bytes, err := utils.PutInt64InBytes(SSCC96Header, sscc96Bytes, sscc96_headerPartition)
	if err != nil {
		return "", err
	}
	sscc96Bytes, err = utils.PutInt64InBytes(sscc.filter, sscc96Bytes, sscc96_filterPartition)
	if err != nil {
		return "", err
	}
	sscc96Bytes, err = utils.PutInt64InBytes(sscc.partition, sscc96Bytes, sscc96_partitionNumberPartition)
	if err != nil {
		return "", err
	}
	sscc96Bytes, err = utils.PutInt64InBytes(sscc.CompanyPrefix, sscc96Bytes, sscc96_dataPartitions[sscc.partition][0])
	if err != nil {
		return "", err
	}
	sscc96Bytes, err = utils.PutInt64InBytes(sscc.Serial, sscc96Bytes, sscc96_dataPartitions[sscc.partition][1])
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%X", sscc96Bytes), nil
}

func sscc96FromBytes(epcBytes []byte) (SSCC96, error) {
	partitionNumber, err := utils.GetInt64FromBytes(epcBytes, sscc96_partitionNumberPartition)
	if err != nil {
		return SSCC96{}, err
	}
	filter, err := utils.GetInt64FromBytes(epcBytes, sscc96_filterPartition)
	if err != nil {
		return SSCC96{}, err
	}
	if partitionNumber >= int64(len(sscc96_dataPartitions)) {
		return SSCC96{}, errors.Join(partition.ErrInvalidPartition, fmt.Errorf("got partition %d", partitionNumber))
	}
	companyPrefix, err := utils.GetInt64FromBytes(epcBytes, sscc96_dataPartitions[partitionNumber][0])
	if err != nil {
		return SSCC96{}, err
	}
	serial, err := utils.GetInt64FromBytes(epcBytes, sscc96_dataPartitions[partitionNumber][1])
	if err != nil {
		return SSCC96{}, err
	}
	return SSCC96{CompanyPrefix: companyPrefix, Serial: serial, filter: filter, partition: partitionNumber}, nil
}
