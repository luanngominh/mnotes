package model

import "regexp"

var (
	uuidRegex = "^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"
)

//IsUUID check a string is uuid
func IsUUID(s string) bool {
	r := regexp.MustCompile(uuidRegex)
	return r.MatchString(s)
}
