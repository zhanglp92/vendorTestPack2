package vendorTestPack2

import "testing"

func TestPack2(t *testing.T) {
	t.Logf(NewVendorTestPack2(2).String())
}
