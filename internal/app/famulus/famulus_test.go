package famulus

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	testOutput := "testdata/test_output.pdf"

	defer os.Remove(testOutput)

	var err error
	if err = Generate("testdata/input.kys", testOutput); err != nil {
		t.Error(err)
	}

	refStat, err := os.Stat("testdata/output.pdf")
	if err != nil {
		t.Error(err)
	}

	testStat, err := os.Stat(testOutput)
	if err != nil {
		t.Error(err)
	}

	refSize, testSize := refStat.Size(), testStat.Size()
	if refSize != testSize {
		t.Errorf("File sizes are not equal: %v != %v", refSize, testSize)
	}
}

func TestGenerateEmptyArgument(t *testing.T) {
	testCases := []struct {
		input, output string
		expected      error
	}{
		{
			"testdata/input.kys",
			"",
			errors.New("OUTPUT PATH MUST BE SPECIFIED"),
		},
		{
			"",
			"testdata/output.pdf",
			errors.New("INPUT PATH MUST BE SPECIFIED"),
		},
	}

	var err error
	for _, tc := range testCases {
		err = Generate(tc.input, tc.output)
		if !reflect.DeepEqual(tc.expected, err) {
			fmt.Println(tc.expected, err)
			t.Errorf("Expected: %v, found: %v", tc.expected, err)
		}
	}
}
