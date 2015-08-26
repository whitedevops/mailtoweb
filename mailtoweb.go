// Package mailtoweb returns the webmail URL related to a an email address.
//
// It can be used in web apps for a better user experience, by providing a direct link to user's webmail from his email address, after a subscription for example.
package mailtoweb

import "strings"

var relations = map[string]string{
	"@gmail.": "https://mail.google.com/",
}

// From returns the webmail URL related to the email address.
// If unknown, result is empty.
func From(email string) string {
	for e, u := range relations {
		if strings.Contains(email, e) {
			return u
		}
	}
	return ""
}
