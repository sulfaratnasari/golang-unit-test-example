package calculation_test

import (
	"golangunittestexample/calculation"
	"testing"
)

// Simple unit test
func TestAdd(t *testing.T){
	actualResult := calculation.Add(1,2)
	expectedResult := 3
	
	if actualResult != expectedResult{
		t.Fail()
	}
}

// Unit test with Test Table Driven
func TestAddWithTestTableDriven(t *testing.T){
	testTable := []struct{
		a,b int
		expectedResult int
	}{
		{
			a:1, b:2, expectedResult: 3,
		},
		{
			a:-1, b:2, expectedResult: 1,
		},
		{
			a:1, b:-2, expectedResult: -1,
		},
	}

	for _, test := range testTable {
		actualResult := calculation.Add(test.a, test.b)
		if actualResult != test.expectedResult {
			t.Log()
			t.Fail()
		}
	}
}

// Unit Test with subtest
func TestAddWithSubtest(t *testing.T){
	testTable := []struct{
		name string
		a,b int
		expectedResult int
	}{
		{
			name: "a positive, b positive", a:1, b:2, expectedResult: 3,
		},
		{
			name: "a negative, b positive", a:-1, b:2, expectedResult: 1,
		},
		{
			name: "a positive, b negative", a:1, b:-2, expectedResult: -1,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			actualResult := calculation.Add(test.a, test.b)
			if actualResult != test.expectedResult {
				t.Log()
				t.Fail()
			}
		})
	}
}