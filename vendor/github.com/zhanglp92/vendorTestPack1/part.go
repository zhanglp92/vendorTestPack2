package vendorTestPack1

// VendorTestPack1 the type for test vendor,
// create type1.
type VendorTestPack1 int

// NewVendorTestPack1 is the function for create VendorTestPack1.
func NewVendorTestPack1(n int) *VendorTestPack1 {
	v := VendorTestPack1(n)
	return &v
}

func (v VendorTestPack1) String() string {
	return "124"
	// return fmt.Sprintf("%v", v)
}
