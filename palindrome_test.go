package main

import "testing"

func TestIsP_Success(t *testing.T) {
	testCases := map[string]bool{
		"lol":   true,
		"Lol":   true,
		"":      true,
		"Lo l":  true,
		"lo":    false,
		"s":     true,
		"l o l": true,
		"A man, a plan, a canal, Panama!": true,
		"Was it a car or a cat I saw?":    true,
		"No 'x' in Nixon":                 true,
	}

	for word, expectedResult := range testCases {
		actualResult := isPalindrome(word)
		if actualResult != expectedResult {
			t.Errorf(
				"Case '%s'. Expected result is %v but actual is %v",
				word,
				expectedResult,
				actualResult,
			)
		}
	}
}
