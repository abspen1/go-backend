package alp

import "testing"

func TestFloatToString(t *testing.T) {
	var tests = []struct {
		input    float64
		expected string
	}{
		{4.37, "4.370000"},
		{0.0031, "0.003100"},
		{0, "0.000000"},
		{999.69, "999.690000"},
	}
	for _, test := range tests {
		if output := floatToString(test.input); output != test.expected {
			t.Errorf("Test Failed: %g inputted, %s expected, received: %s", test.input, test.expected, output)
		}
	}
}

func TestGetCurrentPrice(t *testing.T) {
	resp := GetCurrentPrice("TSLA")
	if resp == "500" || resp == "404" {
		t.Errorf("Test failed to get current price. Received %s error.", resp)
	}
	resp = GetCurrentPrice("TLOASO")
	if resp != "500" && resp != "404" {
		t.Errorf("Test failed to get error. Received %s.", resp)
	}
}
