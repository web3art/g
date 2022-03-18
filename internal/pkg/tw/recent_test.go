package tw_test

import (
	"testing"

	"github.com/web3art/g/internal/pkg/tw"
)

func TestExtractEthAddress(t *testing.T) {
	address, err := tw.ExtractEthAddress("I'm eth address is 0x714df076992f95E452A345cD8289882CEc6ab82Fasdasxczxcasda")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if *address != "0x714df076992f95E452A345cD8289882CEc6ab82F" {
		t.Fatalf("unexpected address: %v", *address)
	}
}
