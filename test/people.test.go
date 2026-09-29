package main

import zermelogo "github.com/blacklizardcode/zermelogo"

func testContracts(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	contracts, err := c.GetContracts(zermelogo.ContractsParams{
		Fields:             "id,employee,schoolInSchoolYear,start,end,isMainContract,salaryScale,salaryStep",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("contracts", contracts, err)
}

func testEmployeeAbsences(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	employeeAbsences, err := c.GetEmployeeAbsences(zermelogo.EmployeeAbsencesParams{
		Fields:             "id,contract,employee,start,end,absenceType,absenceTypeCode,definitive",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("employeeabsences", employeeAbsences, err)
}

func testEmployees(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	employees, err := c.GetEmployees(zermelogo.EmployeesParams{
		Fields:             "userCode,firstName,lastName,prefix,email,employeeNumber,created,lastModified,lastModifiedBy,commencementSchool,commencementTeaching",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("employees", employees, err)
}

func testFamilyMembers(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	familyMembers, err := c.GetFamilyMembers(zermelogo.FamilyMembersParams{
		Fields:             "userCode,firstName,lastName,prefix,email,students",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("familymembers", familyMembers, err)
}

func testStudents(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	students, err := c.GetStudents(zermelogo.StudentsParams{
		Fields:             "userCode,firstName,lastName,prefix,email,familyMembers",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("students", students, err)
}

func testStudentsInDepartments(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	studentsInDepartments, err := c.GetStudentsInDepartments(zermelogo.StudentsInDepartmentsParams{
		Fields:             "id,schoolInSchoolYear,student,departmentOfBranch,departmentOfBranchCode,mainGroup,mainGroupName,mentorGroup,educationalTrack",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("studentsindepartments", studentsInDepartments, err)
}

func testJobs(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	jobs, err := c.GetJobs(zermelogo.JobsParams{
		Fields:             "id,contract,employee,start,end,employmentType,functionCategory,fte,fteReal,durationManually",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("jobs", jobs, err)
}
