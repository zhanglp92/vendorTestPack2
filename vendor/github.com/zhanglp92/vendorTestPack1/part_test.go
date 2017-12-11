package vendorTestPack1

import "testing"

func TestPack1(t *testing.T) {
	t.Logf(NewVendorTestPack1(1).String())
}
