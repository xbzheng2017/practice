package main

import (
	"encoding/json"
)

type CPU struct {
	Usage int `json:",string"`
}

func DecodeCPU(cpu string) (*CPU, error) {
	res := new(CPU)
	err := json.Unmarshal([]byte(cpu), res)
	return res, err
}
