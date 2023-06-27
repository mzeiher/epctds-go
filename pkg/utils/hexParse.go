package utils

import "strconv"

func GetEpcBytes(epc string) ([]byte, error) {
	epcBytes := make([]byte, 0)
	for i := 0; i < len(epc); i = i + 2 {
		tempParse, err := strconv.ParseUint(epc[i:i+2], 16, 8)
		if err != nil {
			return epcBytes, err
		}
		epcBytes = append(epcBytes, byte(tempParse)&0xFF)
	}
	return epcBytes, nil
}
