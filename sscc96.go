package epctds

import (
	"errors"
	"fmt"

	"github.com/mzeiher/epctds/pkg/partition"
	"github.com/mzeiher/epctds/pkg/utils"
)

const sscc96_header = 0x31

var sscc96_partition = [7][2]partition.Partition{
	{{Start: 13, Length: 40, Digits: 12}, {Start: 53, Length: 18, Digits: 5}},
	{{Start: 13, Length: 37, Digits: 11}, {Start: 50, Length: 21, Digits: 6}},
	{{Start: 13, Length: 34, Digits: 10}, {Start: 47, Length: 24, Digits: 7}},
	{{Start: 13, Length: 30, Digits: 9}, {Start: 43, Length: 28, Digits: 8}},
	{{Start: 13, Length: 27, Digits: 8}, {Start: 40, Length: 31, Digits: 9}},
	{{Start: 13, Length: 24, Digits: 7}, {Start: 37, Length: 34, Digits: 10}},
	{{Start: 13, Length: 20, Digits: 6}, {Start: 33, Length: 38, Digits: 11}},
}

type SSCC96 struct {
	EPCTag
	CompanyPrefix int64
	Serial        int64
	Filter        int64
}

func (sscc SSCC96) ToTagURI() string {
	return fmt.Sprintf("urn:epc:tag:sscc-96:%d.%d.%d", sscc.Filter, sscc.CompanyPrefix, sscc.Serial)
}
func (sscc SSCC96) ToPureIdentityURI() string {
	return fmt.Sprintf("urn:epc:tag:sscc:%d.%d", sscc.CompanyPrefix, sscc.Serial)
}

func sscc69FromByes(epcBytes []byte) (SSCC96, error) {
	partitionNumber, err := utils.GetInt64FromBytes(epcBytes, partition.Partition{Start: 10, Length: 3, Digits: 3})
	if err != nil {
		return SSCC96{}, err
	}
	filter, err := utils.GetInt64FromBytes(epcBytes, partition.Partition{Start: 7, Length: 3, Digits: 3})
	if err != nil {
		return SSCC96{}, err
	}
	if partitionNumber >= int64(len(sscc96_partition)) {
		return SSCC96{}, errors.Join(partition.ErrInvalidPartition, fmt.Errorf("got partition %d", partitionNumber))
	}
	companyPrefix, err := utils.GetInt64FromBytes(epcBytes, sscc96_partition[partitionNumber][0])
	if err != nil {
		return SSCC96{}, err
	}
	serial, err := utils.GetInt64FromBytes(epcBytes, sscc96_partition[partitionNumber][1])
	if err != nil {
		return SSCC96{}, err
	}
	return SSCC96{CompanyPrefix: companyPrefix, Serial: serial, Filter: filter}, nil
}