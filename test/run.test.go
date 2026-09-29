package main

import zermelogo "github.com/blacklizardcode/zermelogo"

func runAllTests(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	testAppointments(c, schoolInSchoolYear)
	testBranchesOfSchools(c, schoolInSchoolYear)
	testContracts(c, schoolInSchoolYear)
	testCourses(c, schoolInSchoolYear)
	testDepartmentsOfBranches(c, schoolInSchoolYear)
	testEmployeeAbsences(c, schoolInSchoolYear)
	testEmployees(c, schoolInSchoolYear)
	testEnrollmentStatus(c, schoolInSchoolYear)
	testFamilyMembers(c, schoolInSchoolYear)
	testGroupsInDepartments(c, schoolInSchoolYear)
	testHolidays(c, schoolInSchoolYear)
	testJobs(c, schoolInSchoolYear)
	testLiveSchedule(c, schoolInSchoolYear)
	testLocationsOfBranches(c, schoolInSchoolYear)
	testParentTeacherNights(c, schoolInSchoolYear)
	testParentToSchedulerRemarks(c, schoolInSchoolYear)
	testSchoolsInSchoolYears(c, schoolInSchoolYear)
	testStudentsInDepartments(c, schoolInSchoolYear)
	testStudents(c, schoolInSchoolYear)
	testTalkDays(c, schoolInSchoolYear)
	testTalkTimeBlockAbsenceParent(c, schoolInSchoolYear)
	testTalkTimeBlocks(c, schoolInSchoolYear)
	testUsers(c, schoolInSchoolYear)
}
