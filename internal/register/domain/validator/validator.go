package validator

import "regexp"

func Id(email string) bool {
	return true
}

func Name(email string) bool {
	return true
}

func Email(email string) bool {
	emailRegexp := regexp.MustCompile("[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?")
	return emailRegexp.MatchString(email)
}

func Password(email string) bool {
	return true
}
