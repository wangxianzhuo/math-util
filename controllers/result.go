package controllers

import (
	"encoding/json"
)

type CalcResult struct {
	Value       interface{} `json:"value"`
	Description string      `json:"des"`
}

func encodeCalcResult(cr CalcResult) ([]byte, error) {
	return json.Marshal(cr)
}
