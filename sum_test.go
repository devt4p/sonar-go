package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2, 3)

	if result != 5 {
		t.Error("The Result must be 5")
	}
}

func TestSub(t *testing.T) {
	result := sub(4, 3)

	if result != 1 {
		t.Error("The Result must be 1")
	}
}

func TestTimes(t *testing.T) {
	result := times(4, 3)

	if result != 12 {
		t.Error("The Result must be 12")
	}
}

func TestSumX(t *testing.T) {
	result := sumX(4, 3)

	if result != 11 {
		t.Error("The Result must be 11")
	}
}
