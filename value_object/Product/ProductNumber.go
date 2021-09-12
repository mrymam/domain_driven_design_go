package Product

import "fmt"

type SerialNumber struct {
	productCode string
	branch      string
	lot         string
}

func NewSerialNumber(productCode string, branch string, lot string) *SerialNumber {
	return &SerialNumber{
		productCode: productCode,
		branch:      branch,
		lot:         lot,
	}
}

func (s SerialNumber) ToString() string {
	return fmt.Sprintf("%s-%s-%s", s.productCode, s.branch, s.lot)
}
