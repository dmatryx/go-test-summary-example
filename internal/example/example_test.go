package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNumbers(t *testing.T) {
	type testCase struct {
		name   string
		value1 int
		value2 int
		result int
	}
	testCases := []testCase{
		{"1+1", 1, 1, 2},
		{"1+2", 1, 2, 3},
		{"2+1", 2, 1, 3},
		{"5+0", 5, 0, 5},
		{"0+0", 0, 0, 0},
		{"-1+-1", -1, -1, -2},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			assert.True(t, AddNumbers(test.value1, test.value2) == test.result)
		})
	}
}

func TestAddNumbersAgainButFailTwotests(t *testing.T) {
	type testCase struct {
		name   string
		value1 int
		value2 int
		result int
	}
	testCases := []testCase{
		{"1+1", 1, 1, 2},
		{"1+2", 1, 2, 5},
		{"2+1", 2, 1, 5},
		{"5+0", 5, 0, 5},
		{"0+0", 0, 0, 0},
		{"0+0", 0, 0, 0},
		{"0+0", 0, 0, 0},
		{"0+0", 0, 0, 0},
		{"-1+-1", -1, -1, -2},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			assert.True(t, AddNumbers(test.value1, test.value2) == test.result)
		})
	}
}

func TestMirrorString(t *testing.T) {
	assert.True(t, MirrorString("test") == "test")
}
