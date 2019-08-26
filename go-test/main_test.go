package main

import (
	"fmt"
	"testing"
)


func TestCalculate(t *testing.T) {
	expected := 4
	result := Calculate(2)
	if expected != result {
		t.Error("Failed")
	}
}

func benchmarkCalculate(input int, b *testing.B){
	for i := 0; i < b.N; i++ {
		Calculate(input)
	}
}

func BenchmarkCalculate100(b *testing.B) { benchmarkCalculate(100, b)}
func BenchmarkCalculateNegative100(b *testing.B) { benchmarkCalculate(-100, b)}
func BenchmarkCalculateNegative1(b *testing.B) { benchmarkCalculate(-1, b)}


func TestTableCalculate(t *testing.T) {
	var tests = []struct{
		input int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}

func TestOther(t *testing.T){
	fmt.Println("Testing something else")
	fmt.Println("This shouldn't run with -run=calc")
}