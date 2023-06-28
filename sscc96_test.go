package epctds

import (
	"testing"

	"github.com/mzeiher/epctds-go/pkg/utils"
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

func TestSSC96Serialization(t *testing.T) {
	expectedPureIdentityUri := "urn:epc:id:sscc:0001000.0000000100"
	expectedTagUri := "urn:epc:tag:sscc-96:0.0001000.0000000100"
	hexString := "3114000FA000000064000000"

	epcBytes, err := utils.GetEpcBytes(hexString)
	if err != nil {
		t.Fatal(err)
	}
	sscc96, err := sscc69FromBytes(epcBytes)
	if err != nil {
		t.Fatal(err)
	}
	if expectedPureIdentityUri != sscc96.ToPureIdentityURI() {
		t.Fatalf("invalid pure identity URI: expected: %s, got: %s", expectedPureIdentityUri, sscc96.ToPureIdentityURI())
	}
	if expectedTagUri != sscc96.ToTagURI() {
		t.Fatalf("invalid tag URI: expected: %s, got: %s", expectedTagUri, sscc96.ToTagURI())
	}
	generatedHexString, err := sscc96.ToHex()
	if err != nil {
		t.Fatal(err)
	}
	if generatedHexString != hexString {
		t.Fatalf("invalid hex string: expected %s, got %s", hexString, generatedHexString)
	}
}

func testSscc96(t *testing.T, epcString string, expectedCompanyPrefix int64, expectedSerial int64) {
	epcBytes, err := utils.GetEpcBytes(epcString)
	if err != nil {
		t.Fatal(err)
	}
	sscc96, err := sscc69FromBytes(epcBytes)
	if err != nil {
		t.Fatal(err)
	}
	if sscc96.CompanyPrefix != expectedCompanyPrefix && sscc96.Serial != expectedSerial {
		t.Fatalf("invalid values")
	}
}

func BenchmarkSSCC96Parsing(b *testing.B) {
	epcBytes, err := utils.GetEpcBytes("3118E511C46699F387000000")
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		sscc69FromBytes(epcBytes)
	}
}
