package main

import (
	"testing"
)

func TestDecodeJsonInt2String(t *testing.T) {
	target := `{"Usage": "50"}`
	res, err := DecodeCPU(target)
	if err != nil {
		t.Error(err)
	}
	if res == nil {
		t.Error("")
	}
	if res.Usage != 50 {
		t.Error("")
	}
}
