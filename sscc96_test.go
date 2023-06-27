package epctds

import (
	"testing"

	"github.com/mzeiher/epctds/pkg/utils"
)

func TestSSCC96Parsing(t *testing.T) {
	testSscc96(t, "3118E511C46699F387000000", 234567, 18901234567)
	testSscc96(t, "31148F2B3871528987000000", 2345678, 1901234567)
	testSscc96(t, "3110B2F60A8608B787000000", 23456789, 101234567)
	testSscc96(t, "310CDFB38D20AB6D07000000", 234567890, 11234567)
	testSscc96(t, "31088BD0383512D687000000", 2345678901, 1234567)
	testSscc96(t, "3104AEC44642820DA7000000", 23456789012, 134567)
	testSscc96(t, "3100DA7557D32C38E7000000", 234567890123, 14567)
}

func testSscc96(t *testing.T, epcString string, expectedCompanyPrefix int64, expectedSerial int64) {
	epcBytes, err := utils.GetEpcBytes(epcString)
	if err != nil {
		t.Fatal(err)
	}
	sscc96, err := sscc69FromByes(epcBytes)
	if err != nil {
		t.Fatal(err)
	}
	if sscc96.CompanyPrefix != expectedCompanyPrefix && sscc96.Serial != expectedSerial {
		t.Fatalf("invalid values")
	}
}
