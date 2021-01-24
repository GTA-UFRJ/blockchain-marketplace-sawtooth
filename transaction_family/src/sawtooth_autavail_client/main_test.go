package main

import (
	"testing"
)

func Test_Sha512Compute(t *testing.T) {
	sha512OfAutavail := "d7ad2cfd56c07c038bbc5c3c2c23c33effb1d44ec59cc98ed507cc8350df74dc3b240ea39e8c4189752cd23c571e157ac3bf27b409001abd354d8c5d178a7091"
	computedSha512OfAutavail := Sha512HashValue(FAMILY_NAME)
	if sha512OfAutavail != computedSha512OfAutavail {
		t.Errorf("Sha512 computed doesn't match \nExpected: %s \nGot: %s \n",
			sha512OfAutavail, computedSha512OfAutavail)
	}
}
