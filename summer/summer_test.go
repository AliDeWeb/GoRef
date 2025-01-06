package summer

// we can run tests using go test command
import "testing"

/*
t from testing package has

Fail() => mark test as failed
Error(string) => mark test as failed with a message
FailNow() => mark test as failed and testing will be stopped
Fatalf(string) => mark test as failed with a message and testing will be stopped
Logf() => it is like printf but only runs when test fails
*/
func TestSummer(t *testing.T) {
	tests := []struct {
		a, b   int
		result int
	}{{a: 1, b: 2, result: 3}, {a: 2, b: 2, result: 4}}

	for _, test := range tests {
		result := summer(test.a, test.b)
		if result != test.result {
			t.Error("summer error")
		}
	}
}
