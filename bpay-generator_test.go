package bpay

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"demo", "omed"},
		{"123456", "654321"},
	}

	for i := range tests {
		got := reverse(tests[i].input)
		if got != tests[i].expected {
			t.Fatalf("Expected %v, got %v", tests[i].expected, got)
		}
	}
}

func TestGenerateMOD10V1(t *testing.T) {

	tests := []struct {
		input    string
		expected string
	}{
		{"1234567890", "12345678903"},
		{"10000456", "100004563"},
	}

	for i := range tests {
		got := GenerateMOD10V1(tests[i].input)
		if got != tests[i].expected {
			t.Fatalf("Expected %v, got %v", tests[i].expected, got)
		}
	}
}

func TestGenerateMOD10V5(t *testing.T) {

	tests := []struct {
		input    string
		expected string
	}{
		{"1234567890", "12345678905"},
		{"40007923", "400079231"},
	}

	for i := range tests {
		got := GenerateMOD10V5(tests[i].input)
		if got != tests[i].expected {
			t.Fatalf("Expected %v, got %v", tests[i].expected, got)
		}
	}
}
