package utils

import (
	"bytes"
	"testing"
)

func TestGetEpcBytes(t *testing.T) {
	testString := "31148F2B3871528987000000"
	expected := []byte{0x31, 0x14, 0x8f, 0x2b, 0x38, 0x71, 0x52, 0x89, 0x87, 0x0, 0x0, 0x0}

	epcBytes, err := GetEpcBytes(testString)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(epcBytes, expected) {
		t.Fatalf("not equal, expected %v, got %v", expected, epcBytes)
	}
}

func BenchmarkGetEpcBytes(t *testing.B) {
	for i := 0; i < t.N; i++ {
		GetEpcBytes("31148F2B3871528987000000")
	}
}
