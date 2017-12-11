package vendorTestPack2

import (
	"github.com/zhanglp92/vendorTestPack1"
)

// VendorTestPack2 the type for test vendor,
// create type1.
type VendorTestPack2 struct {
	*vendorTestPack1.VendorTestPack1
}

// NewVendorTestPack2 is the function for create VendorTestPack2.
func NewVendorTestPack2(n int) *VendorTestPack2 {
	v := VendorTestPack2{
		vendorTestPack1.NewVendorTestPack1(n),
	}
	return &v
}

func (v VendorTestPack2) String() string {
	return "2"
	// return fmt.Sprintf("%v", v)
}

// TestType .
func TestType(v vendorTestPack1.VendorTestPack1) {
	vendorTestPack1.TestType(v)
}
