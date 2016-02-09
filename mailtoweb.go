// Package mailtoweb returns the webmail URL related to an email address.
package mailtoweb

import "strings"

var relations = map[string]string{
	"@gmail.": "https://mail.google.com/",
}

// For returns the webmail URL related to the email address.
// If unknown, result is empty.
func For(email string) string {
	for e, u := range relations {
		if strings.Contains(email, e) {
			return u
		}
	}
	return ""
}
