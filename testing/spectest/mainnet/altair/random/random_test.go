package random

import (
	"testing"

	"github.com/prysmaticlabs/prysm/testing/spectest/shared/altair/sanity"
)

func TestMainnet_Altair_Random(t *testing.T) {
	sanity.RunBlockProcessingTest(t, "mainnet", "random/random/pyspec_tests")
}
