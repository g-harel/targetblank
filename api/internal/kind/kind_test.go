package kind

import (
	"strings"
	"testing"
)

type testCase struct {
	kind   kind
	value  string
	result bool
}

func checkCases(t *testing.T, tcs ...*testCase) {
	for _, tc := range tcs {
		err := Of(tc.value).Is(tc.kind)
		if (err == nil) != tc.result {
			if tc.result {
				t.Fatalf("Test case failed \"%v\" is not of kind %v", tc.value, tc.kind)
			} else {
				t.Fatalf("Test case failed \"%v\" is of kind %v", tc.value, tc.kind)
			}
		}
	}
}

func TestEMAIL(t *testing.T) {
	t.Run("should produce a useful error", func(t *testing.T) {
		err := Of("not an email").Is(EMAIL)
		if err == nil {
			t.Fatal("Non-email test value declared valid")
		}

		if strings.Index(err.Error(), "email") < 0 {
			t.Fatal("Generated error does not contain the \"email\" substring")
		}
	})

	t.Run("should correctly match emails", func(t *testing.T) {
		checkCases(t, []*testCase{
			{EMAIL, "address@domain.tld", true},
			{EMAIL, "firstname.lastname@ex.ample", true},
			{EMAIL, "whoops@email", false},
			{EMAIL, "", false},
			{EMAIL, "not an email", false},
		}...)
	})
}

func TestPASSWORD(t *testing.T) {
	t.Run("should produce a useful error", func(t *testing.T) {
		err := Of("pass").Is(PASSWORD)
		if err == nil {
			t.Fatal("Bad test value declared valid")
		}

		if strings.Index(err.Error(), "password") < 0 || strings.Index(err.Error(), "eight") < 0 {
			t.Fatal("Generated error does not contain \"password\" and \"eight\"")
		}
	})

	t.Run("should correctly match passwords", func(t *testing.T) {
		checkCases(t, []*testCase{
			{PASSWORD, "2short", false},
			{PASSWORD, "?", false},
			{PASSWORD, "super duper long password", true},
			{PASSWORD, "longenuf", true},
		}...)
	})
}
