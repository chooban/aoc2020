package passwords

import (
	"testing"
)

func TestInvalidPassword(t *testing.T) {
	valid := ValidateSledRentalPassword("1-3 b", "cdefg")

	if valid {
		t.Errorf("abcde should not be valid")
	}
}

func TestValidPasswords(t *testing.T) {
	passwords := map[string]string{
		"1-3 a": "abcde",
		"2-9 c": "ccccccccc",
	}

	for pattern, password := range passwords {
		valid := ValidateSledRentalPassword(pattern, password)
		if !valid {
			t.Errorf("%s should be valid", password)
		}
	}
}

func BenchmarkValidPassword(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ValidateSledRentalPassword("1-3 a", "abcde")
	}
}
