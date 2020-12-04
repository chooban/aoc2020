package passports

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var requiredFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
	// "cid",
}

var optionalFields = []string{
	"cid",
}

func validateByr(input string) bool {
	return validateNumberRange(input, 1920, 2002)
}

func validateIyr(input string) bool {
	return validateNumberRange(input, 2010, 2020)
}

func validateEyr(input string) bool {
	return validateNumberRange(input, 2020, 2030)
}

func validateHgt(input string) bool {
	re := regexp.MustCompile("^(\\d*)(in|cm)$")
	parts := re.FindStringSubmatch(input)
	if parts == nil || len(parts) != 3 {
		return false
	}
	if parts[2] == "cm" {
		return validateNumberRange(parts[1], 150, 193)
	} else if parts[2] == "in" {
		return validateNumberRange(parts[1], 59, 76)
	}
	return false
}

func validateHcl(input string) bool {
	re := regexp.MustCompile("^#[a-f0-9]{6}$")
	return re.MatchString(input)
}

var eyeColours = map[string]bool{
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

func validateEcl(input string) bool {
	return eyeColours[input] || false
}

func validatePid(input string) bool {
	re := regexp.MustCompile("^\\d{9}$")
	return re.MatchString(input)
}

func validateNumberRange(input string, min, max int) bool {
	val, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Could not convert %s to number", input)
		return false
	}

	return val >= min && val <= max
}

func ValidatePassport(passport []string) bool {
	// Get all the supplied field names
	fieldNames := []string{}
	for _, field := range passport {
		fieldParts := strings.Split(field, ":")
		fieldNames = append(fieldNames, fieldParts[0])
	}

	for _, requiredField := range requiredFields {
		if !contains(fieldNames, requiredField) {
			return false
		}
	}

	// Now validate the fields
	for _, field := range passport {
		fieldParts := strings.Split(field, ":")
		valid := true
		switch fieldParts[0] {
		case "byr":
			valid = validateByr(fieldParts[1])
		case "iyr":
			valid = validateIyr(fieldParts[1])
		case "eyr":
			valid = validateEyr(fieldParts[1])
		case "hgt":
			valid = validateHgt(fieldParts[1])
		case "hcl":
			valid = validateHcl(fieldParts[1])
		case "ecl":
			valid = validateEcl(fieldParts[1])
		case "pid":
			valid = validatePid(fieldParts[1])
		}
		if !valid {
			fmt.Printf("%s was invalid. %s:%s field invalid\n", passport, fieldParts[0], fieldParts[1])
			return false
		}
	}

	return true
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
