package helpers

import "regexp"

func IsValidPassword(password string) bool {
	passwordPatternRegex := "[a-zA-Z0-9!@#$%^&*()_+{}|:<>?]{8,}"

	regexpObj := regexp.MustCompile(passwordPatternRegex)

	return regexpObj.MatchString(password)
}
