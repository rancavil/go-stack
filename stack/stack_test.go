package stack

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPop(t *testing.T) {
	testCases := []struct {
		expected Stack
		given    Stack
	}{
		{
			Stack{1, 2},
			Stack{1, 2, 3},
		},
		{
			Stack{"Fred", "Frida", "Francis"},
			Stack{"Fred", "Frida", "Francis", "Frederick"},
		},
		{
			Stack{1.4, 7.0},
			Stack{1.4, 7.0, 3.14},
		},
	}

	for i, testCase := range testCases {
		testName := fmt.Sprintf("Test-Pop-%d", i)
		t.Run(testName, func(t *testing.T) {
			actual := testCase.given
			actual.Pop()
			if !reflect.DeepEqual(testCase.expected, actual) {
				t.Errorf("Failed: expected %v, actual %v", testCase.expected, actual)
			}
		})
	}
}

func TestPush(t *testing.T) {
	testCases := []struct {
		expected Stack
		given    []interface{}
	}{
		{Stack{1, 2, 3, 4}, []interface{}{1, 2, 3, 4}},
		{Stack{"Sonny", "Freddo", "Michael", "Connie"}, []interface{}{"Sonny", "Freddo", "Michael", "Connie"}},
	}
	for i, testCase := range testCases {
		testName := fmt.Sprintf("Test-Push-%d", i)
		t.Run(testName, func(t *testing.T) {
			actual := Stack{}
			for _, value := range testCase.given {
				actual.Push(value)
			}
			if !reflect.DeepEqual(testCase.expected, actual) {
				t.Errorf("Failed: expected %v, actual %v", testCase.expected, actual)
			}
		})
	}
}
