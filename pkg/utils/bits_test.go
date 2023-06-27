package utils

import (
	"testing"

	"github.com/mzeiher/epctds-go/pkg/partition"
)

func TestBitPartitioning(t *testing.T) {
	test := []byte{0x31, 0x14, 0x8f, 0x2b, 0x38, 0x71, 0x52, 0x89, 0x87, 0x0, 0x0, 0x0}
	expectedSerial := int64(1901234567)
	expectedCompanyPrefix := int64(2345678)

	output, err := GetInt64FromBytes(test, partition.Partition{Length: 24, Start: 13, Digits: 7})
	if err != nil {
		t.Fatal(err)
	}
	if output != expectedCompanyPrefix {
		t.Fatalf("wrong data, got %d, expected %d", output, expectedCompanyPrefix)
	}

	output, err = GetInt64FromBytes(test, partition.Partition{Length: 34, Start: 37, Digits: 10})
	if err != nil {
		t.Fatal(err)
	}
	if output != expectedSerial {
		t.Fatalf("wrong data, got %d, expected %d", output, expectedSerial)
	}

}

func BenchmarkBitPartitioning(b *testing.B) {
	test := []byte{0x31, 0x14, 0x8f, 0x2b, 0x38, 0x71, 0x52, 0x89, 0x87, 0x0, 0x0, 0x0}
	for n := 0; n < b.N; n++ {
		GetInt64FromBytes(test, partition.Partition{Length: 24, Start: 13, Digits: 7})
	}
}
