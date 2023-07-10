package password

import (
	"strings"
	"testing"
)

// Test Generate function
func TestGenerate(t *testing.T) {

	// Test case 1: Generating a password without numbers or symbols
	t.Run("generate password without numbers or symbols", func(t *testing.T) {
		if _, err := Generate(8, false, false, true); err != nil {
			t.Errorf("Expected Generate(8, false, false) to not return an error, but got %s", err)
		}

		if password, err := Generate(8, false, false, true); err != nil {
			if len(password) != 8 {
				t.Errorf("Expected password length of 8, but got %d", len(password))
			}
			if strings.ContainsAny(password, Numerals) {
				t.Errorf("Expected password to not contain numbers")
			}
			if strings.ContainsAny(password, Symbols) {
				t.Errorf("Expected password to not contain symbols")
			}

			t.Log("Generated password:", password)
		}

	})

	// Test case 2: Generating a password with numbers but without symbols
	t.Run("generate password with numbers but without symbols", func(t *testing.T) {
		if _, err := Generate(12, true, false, true); err != nil {
			t.Errorf("Expected Generate(12, true, false) to not return an error, but got %s", err)
		}

		if password, err := Generate(12, true, false, true); err != nil {
			if len(password) != 12 {
				t.Errorf("Expected password length of 12, but got %d", len(password))
			}
			if strings.ContainsAny(password, Symbols) {
				t.Errorf("Expected password to not contain symbols")
			}

			t.Log("Generated password:", password)
		}
	})

	// Test case 3: Generating a password with numbers and symbols
	t.Run("generate password with numbers and symbols", func(t *testing.T) {
		if _, err := Generate(10, true, true, true); err != nil {
			t.Errorf("Expected Generate(10, true, true) to not return an error, but got %s", err)
		}

		if password, err := Generate(10, true, true, true); err != nil {
			if len(password) != 10 {
				t.Errorf("Expected password length of 10, but got %d", len(password))
			}

			if len(password) != 10 {
				t.Errorf("Expected password length of 10, but got %d", len(password))
			}

			t.Log("Generated password:", password)
		}
	})

	// Test case 4: Generating a password with symbols but without numbers
	t.Run("generate password with symbols but without numbers", func(t *testing.T) {
		if _, err := Generate(15, false, true, true); err != nil {
			t.Errorf("Expected Generate(15, false, true) to not return an error, but got %s", err)
		}

		if password, err := Generate(15, false, true, true); err != nil {
			if len(password) != 15 {
				t.Errorf("Expected password length of 15, but got %d", len(password))
			}

			if strings.ContainsAny(password, Numerals) {
				t.Errorf("Expected password to not contain numbers")
			}

			t.Log("Generated password:", password)
		}
	})

	// Test case 5: Generating a password without uppercase letters
	t.Run("generate password without uppercase letters", func(t *testing.T) {
		if _, err := Generate(20, true, true, false); err != nil {
			t.Errorf("Expected Generate(20, true, true, false) to not return an error, but got %s", err)
		}

		if password, err := Generate(20, true, true, false); err != nil {
			if len(password) != 20 {
				t.Errorf("Expected password length of 20, but got %d", len(password))
			}

			if strings.ContainsAny(password, LettersUpper) {
				t.Errorf("Expected password to not contain uppercase letters")
			}

			t.Log("Generated password:", password)
		}
	})
}
