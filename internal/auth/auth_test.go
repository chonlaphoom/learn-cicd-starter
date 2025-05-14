package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type res struct {
		Text string
		Err  error
	}
	type testSuite struct {
		Name     string
		Input    http.Header
		Expected res
	}

	tests := []testSuite{
		{
			Name: "should get APIKey",
			Input: map[string][]string{
				"Authorization": {
					"ApiKey 123",
				},
			},
			Expected: res{
				Text: "123",
				Err:  nil,
			},
		},
	}

	fmt.Println(GetAPIKey(tests[0].Input))

	for _, test := range tests {
		result, err := GetAPIKey(test.Input)
		if result != test.Expected.Text || err != test.Expected.Err {

			t.Fatalf("test case '%s' failed: %v", test.Name, result)
		}
	}
}
