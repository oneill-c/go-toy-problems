package main

import (
	"fmt"
	"regexp"
	"strings"
)

type User struct {
	Email string
	Phone string
}

type Stats struct {
	Total        int
	Valid        int
	Skipped      int
	InvalidEmail int
	InvalidPhone int
}

var numRegex = regexp.MustCompile(`\d+`)

func normalizeEmail(e string) string {
	return strings.TrimSpace(strings.ToLower(e))
}

func normalizePhone(p string) string {
	d := strings.Join(numRegex.FindAllString(p, -1), "")
	if len(d) == 11 && strings.HasPrefix(d, "1") {
		d = d[1:]
	}
	return d
}

func validateEmail(e string) bool {
	at := strings.IndexByte(e, '@')
	if at < 0 {
		return false
	}
	return strings.Contains(e[at+1:], ".")
}

func validatePhone(p string) bool {
	return len(p) == 10
}

func normalizeUsers(users []User) ([]User, Stats) {
	var out []User
	st := Stats{
		Total:        len(users),
		Valid:        0,
		Skipped:      0,
		InvalidEmail: 0,
		InvalidPhone: 0,
	}

	for _, u := range users {
		np := normalizePhone(u.Phone)
		ne := normalizeEmail(u.Email)
		vp := validatePhone(np)
		ve := validateEmail(ne)

		if !vp {
			st.InvalidPhone++
		}
		if !ve {
			st.InvalidEmail++
		}

		if vp && ve {
			out = append(out, User{Phone: np, Email: ne})
		} else {
			st.Skipped++
		}
	}
	st.Valid = len(out)
	return out, st
}

func main() {

	var users = []User{
		{Email: " ADA@Example.COM ", Phone: "(312) 555-1212"},
		{Email: "Alan@Org", Phone: "1-312-555-3434"},
		{Email: "   GRACE.Hopper@navy.mil", Phone: ""},
		{Email: " linus@kernel.org ", Phone: "+1 (773) 444 5566"},
		{Email: "BOB@", Phone: "555555555"}, // invalid email + short phone
		{Email: "mary.jane@example.com", Phone: "  312.666.7777 "},
		{Email: " jdoe@company.com ", Phone: "001-847-222-1111"}, // too many digits
		{Email: "eve@example.com", Phone: "(773)555-8888"},
		{Email: " frank@EXAMPLE.NET ", Phone: "773555999x"}, // weird trailing char
	}

	nu, st := normalizeUsers(users)

	fmt.Printf("a total of %d users were processed. %d were valid and %d were skipped completely. Of those skipped, %d had invalid phone numbers and %d had invalid email addresses\n", st.Total, st.Valid, st.Skipped, st.InvalidPhone, st.InvalidEmail)
	for _, u := range nu {
		fmt.Printf("email=%s phone=%s\n", u.Email, u.Phone)
	}
}
