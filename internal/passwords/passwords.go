package passwords

import (
	"strconv"
	"strings"
)

func ValidateSledRentalPassword(pattern, password string) bool {
	patternParts := strings.Split(pattern, " ")
	patternLimits := strings.Split(patternParts[0], "-")

	character := patternParts[1]
	lower, _ := strconv.Atoi(patternLimits[0])
	upper, _ := strconv.Atoi(patternLimits[1])

	characterCount := strings.Count(password, character)

	return characterCount >= lower && characterCount <= upper
}

func ValidateTobogganPassword(pattern, password string) bool {
	patternParts := strings.Split(pattern, " ")
	patternLocs := strings.Split(patternParts[0], "-")

	character := patternParts[1]
	firstIdx, _ := strconv.Atoi(patternLocs[0])
	secondIdx, _ := strconv.Atoi(patternLocs[1])

	firstChar := string(password[firstIdx])
	secondChar := string(password[secondIdx])

	return (firstChar == character && secondChar != character) || (firstChar != character && secondChar == character)
}
