package utils

import (
	"bytes"
	"testing"

	"github.com/mzeiher/epctds-go/pkg/partition"
)

func TestBitPartitioningGet(t *testing.T) {
	test := []byte{0x31, 0x14, 0x8f, 0x2b, 0x38, 0x71, 0x52, 0x89, 0x87, 0x0, 0x0, 0x0}
	expectedHeader := int64(0x31)
	expectedSerial := int64(1901234567)
	expectedCompanyPrefix := int64(2345678)

	output, err := GetInt64FromBytes(test, partition.Partition{Length: 24, Start: 14, Digits: 7})
	if err != nil {
		t.Fatal(err)
	}
	if output != expectedCompanyPrefix {
		t.Fatalf("wrong data, got %d, expected %d", output, expectedCompanyPrefix)
	}

	output, err = GetInt64FromBytes(test, partition.Partition{Length: 34, Start: 38, Digits: 10})
	if err != nil {
		t.Fatal(err)
	}
	if output != expectedSerial {
		t.Fatalf("wrong data, got %d, expected %d", output, expectedSerial)
	}

	output, err = GetInt64FromBytes(test, partition.Partition{Length: 8, Start: 0, Digits: 8})
	if err != nil {
		t.Fatal(err)
	}
	if output != expectedHeader {
		t.Fatalf("wrong data, got %d, expected %d", output, expectedSerial)
	}
}

func TestBitPartitioningSet(t *testing.T) {
	test := []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	expectedBytes := []byte{0x31, 0x14, 0x8f, 0x2b, 0x38, 0x71, 0x52, 0x89, 0x87, 0x0, 0x0, 0x0}
	header := int64(0x31)
	serial := int64(1901234567)
	companyPrefix := int64(2345678)
	partitionNumber := int64(0x5)

	output, err := PutInt64InBytes(header, test, partition.Partition{Length: 8, Start: 0, Digits: 8})
	if err != nil {
		t.Fatal(err)
	}
	output, err = PutInt64InBytes(partitionNumber, output, partition.Partition{Start: 11, Length: 3, Digits: 3})
	if err != nil {
		t.Fatal(err)
	}
	output, err = PutInt64InBytes(companyPrefix, output, partition.Partition{Length: 24, Start: 14, Digits: 7})
	if err != nil {
		t.Fatal(err)
	}
	output, err = PutInt64InBytes(serial, output, partition.Partition{Length: 34, Start: 38, Digits: 10})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(output, expectedBytes) {
		t.Fatalf("invalid data: expected %x, got %x", expectedBytes, test)
	}
}

func BenchmarkBitPartitioningGet(b *testing.B) {
	test := []byte{0x31, 0x14, 0x8f, 0x2b, 0x38, 0x71, 0x52, 0x89, 0x87, 0x0, 0x0, 0x0}
	for n := 0; n < b.N; n++ {
		GetInt64FromBytes(test, partition.Partition{Length: 24, Start: 14, Digits: 7})
	}
}

func BenchmarkBitPartitioningSet(b *testing.B) {
	test := []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	for n := 0; n < b.N; n++ {
		GetInt64FromBytes(test, partition.Partition{Length: 24, Start: 14, Digits: 7})
	}
}
