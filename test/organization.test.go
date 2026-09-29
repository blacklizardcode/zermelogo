package main

import zermelogo "blacklizardcode/zermelo-go"

func testBranchesOfSchools(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	branchesOfSchools, err := c.GetBranchesOfSchools(zermelogo.BranchesOfSchoolsParams{
		Fields:             "id,schoolInSchoolYear,branch,name",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("branchesofschools", branchesOfSchools, err)
}

func testDepartmentsOfBranches(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	departmentsOfBranches, err := c.GetDepartmentsOfBranches(zermelogo.DepartmentsOfBranchesParams{
		Fields:             "id,schoolInSchoolYear,branchOfSchool,code,yearOfEducation,educationType,teachingLevel",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("departmentsofbranches", departmentsOfBranches, err)
}

func testLocationsOfBranches(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	locationsOfBranches, err := c.GetLocationsOfBranches(zermelogo.LocationsOfBranchesParams{
		Fields:             "id,schoolInSchoolYear,branchOfSchool,name,allowMeetings,courseCapacity,parentteachernightCapacity,secondaryBranches,supportsConcurrentAppointments",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("locationofbranches", locationsOfBranches, err)
}

func testSchoolsInSchoolYears(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	schoolsInSchoolYears, err := c.GetSchoolsInSchoolYears(zermelogo.SchoolsInSchoolYearsParams{
		Fields:             "id,school,year,name,archived,archivedTime",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("schoolsinschoolyears", schoolsInSchoolYears, err)
}
