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

// Test GenerateCustom function
func TestGenerateCustom(t *testing.T) {
	// Test case 1: Valid inputs
	t.Run("generate password with valid inputs", func(t *testing.T) {
		password, err := GenerateCustom(12, 3, 2, 2)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		t.Log("Generated password:", password)
	})

	// Verify the length of the generated password
	t.Run("verify the length of the generated password", func(t *testing.T) {
		password, err := GenerateCustom(12, 3, 2, 2)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		t.Log("Generated password:", password)

		if len(password) != 12 {
			t.Errorf("Expected password length to be 12, got: %d", len(password))
		}
	})

	// Verify the number of numbers in the password
	t.Run("verify the number of numbers in the password", func(t *testing.T) {
		password, err := GenerateCustom(12, 3, 2, 2)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		t.Log("Generated password:", password)

		if strings.ContainsAny(password, Numerals) {
			nums := 0
			for _, v := range password {
				if strings.Contains(Numerals, string(v)) {
					nums++
				}
			}

			if nums != 3 {
				t.Errorf("Expected 3 numbers in the password, got: %d", nums)
			}
		} else {
			t.Errorf("Expected password to contain numbers")
		}
	})

	// Verify the number of symbols in the password
	t.Run("verify the number of symbols in the password", func(t *testing.T) {
		password, err := GenerateCustom(12, 3, 2, 2)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		t.Log("Generated password:", password)

		if strings.ContainsAny(password, Symbols) {
			symbols := 0
			for _, v := range password {
				if strings.Contains(Symbols, string(v)) {
					symbols++
				}
			}

			if symbols != 2 {
				t.Errorf("Expected 2 symbols in the password, got: %d", symbols)
			}
		} else {
			t.Errorf("Expected password to contain symbols")
		}
	})

	// Verify the number of uppercase letters in the password
	t.Run("verify the number of uppercase letters in the password", func(t *testing.T) {
		password, err := GenerateCustom(12, 3, 2, 2)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		t.Log("Generated password:", password)

		if strings.ContainsAny(password, LettersUpper) {
			upper := 0
			for _, v := range password {
				if strings.Contains(LettersUpper, string(v)) {
					upper++
				}
			}

			if upper != 2 {
				t.Errorf("Expected 2 uppercase letters in the password, got: %d", upper)
			}
		} else {
			t.Errorf("Expected password to contain uppercase letters")
		}

	})

	// Verify that the password contains only valid characters
	t.Run("verify that the password contains only valid characters", func(t *testing.T) {
		password, err := GenerateCustom(12, 3, 2, 2)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		t.Log("Generated password:", password)

		if !strings.ContainsAny(password, LettersLower+LettersUpper+Numerals+Symbols) {
			t.Errorf("Expected password to contain only valid characters")
		} else {
			t.Log("Password contains only valid characters")
		}
	})

	// Test case 2: Invalid length
	t.Run("generate password with invalid length", func(t *testing.T) {
		_, err := GenerateCustom(0, 3, 2, 2)
		if err != ErrorLength {
			t.Errorf("Expected error ErrorLength, got: %v", err)
		}

		if err != ErrorLength {
			t.Errorf("Expected error ErrorLength, got: %v", err)
		}

	})

	// Test case 3: Invalid count of numbers
	t.Run("generate password with invalid count of numbers", func(t *testing.T) {
		_, err := GenerateCustom(10, 12, 2, 2)
		if err != ErrorNumbersMoreThanLength {
			t.Errorf("Expected error ErrorNumbersMoreThanLength, got: %v", err)
		}
	})

	// Test case 4: Invalid count of symbols
	t.Run("generate password with invalid count of symbols", func(t *testing.T) {
		_, err := GenerateCustom(10, 3, 15, 2)
		if err != ErrorSymbolsMoreThanLength {
			t.Errorf("Expected error ErrorSymbolsMoreThanLength, got: %v", err)
		}
	})

	// Test case 5: Invalid count of uppercase letters
	t.Run("generate password with invalid count of uppercase letters", func(t *testing.T) {
		_, err := GenerateCustom(10, 3, 2, 20)
		if err != ErrorUpperCaseMoreThanLength {
			t.Errorf("Expected error ErrorUpperCaseMoreThanLength, got: %v", err)
		}
	})
}
