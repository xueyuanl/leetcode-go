package weekly108

import "strings"

func numUniqueEmails(emails []string) int {

	emailMap := map[string]bool{}

	for _,v := range emails {
		email := strings.Split(v, "@")
		local, domain := email[0], email[1]

		local = filterLocalName(local)
		newEmail := local + domain

		if !emailMap[newEmail] {
			emailMap[newEmail] = true
		}
	}

	return len(emailMap)
}

func filterLocalName(local string) string {

	plus := strings.Split(local, "+")

	local = plus[0]

	dots := strings.Split(local, ".")


	res := ""
	for _,v := range dots {
		res += v
	}
	return res
}