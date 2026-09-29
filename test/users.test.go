package main

import zermelogo "github.com/blacklizardcode/zermelogo"

func testUsers(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	users, err := c.GetUsers(zermelogo.UsersParams{
		Code:               "~me",
		Fields:             "code,firstName,lastName,prefix,email",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("users", users, err)
}
