package epctds

import (
	"testing"

	"github.com/mzeiher/epctds-go/pkg/utils"
)

func TestSGLN96Parsing(t *testing.T) {
	testSgln96(t, "321878901815060000BC614E", 123456, 789123, 12345678)
	testSgln96(t, "32144B5A1EB8460000BC614E", 1234567, 89123, 12345678)
	testSgln96(t, "32105E30A747460000BC614E", 12345678, 9123, 12345678)
	testSgln96(t, "320C75BCD150F60000BC614E", 123456789, 123, 12345678)
	testSgln96(t, "3208499602D32E0000BC614E", 1234567891, 23, 12345678)
	testSgln96(t, "32045BFB8388060000BC614E", 12345678912, 3, 12345678)
	testSgln96(t, "320072FA646A0C0000BC614E", 123456789123, 0, 12345678)
}

func testSgln96(t *testing.T, epcString string, expectedCompanyPrefix int64, expectedLocationReference int64, expectedExtension int64) {
	epcBytes, err := utils.GetEpcBytes(epcString)
	if err != nil {
		t.Fatal(err)
	}
	sgln96, err := sgln96FromBytes(epcBytes)
	if err != nil {
		t.Fatal(err)
	}
	if sgln96.CompanyPrefix != expectedCompanyPrefix && sgln96.LocationReference != expectedLocationReference && sgln96.Extension != expectedExtension {
		t.Fatalf("invalid values for string %s", epcString)
	}
}

func TestSGLN96Serialization(t *testing.T) {
	expectedPureIdentityUri := "urn:epc:id:sgln:00000010.0000.1100"
	expectedTagUri := "urn:epc:tag:sgln-96:0.00000010.0000.1100"
	hexString := "32100000050000000000044C"

	epcBytes, err := utils.GetEpcBytes(hexString)
	if err != nil {
		t.Fatal(err)
	}
	sgln96, err := sgln96FromBytes(epcBytes)
	if err != nil {
		t.Fatal(err)
	}
	if expectedPureIdentityUri != sgln96.ToPureIdentityURI() {
		t.Fatalf("invalid pure identity URI: expected: %s, got: %s", expectedPureIdentityUri, sgln96.ToPureIdentityURI())
	}
	if expectedTagUri != sgln96.ToTagURI() {
		t.Fatalf("invalid tag URI: expected: %s, got: %s", expectedTagUri, sgln96.ToTagURI())
	}
	generatedHexString, err := sgln96.ToHex()
	if err != nil {
		t.Fatal(err)
	}
	if generatedHexString != hexString {
		t.Fatalf("invalid hex string: expected %s, got %s", hexString, generatedHexString)
	}
}

func BenchmarkSGNL96Parsing(b *testing.B) {
	epcBytes, err := utils.GetEpcBytes("32144B5A1EB8460000BC614E")
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		sgln96FromBytes(epcBytes)
	}
}
