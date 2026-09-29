package main

import zermelogo "blacklizardcode/zermelo-go"

func testCourses(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	courses, err := c.GetCourses(zermelogo.CoursesParams{
		Fields:             "id,schoolYear,subjectCode,subjectName,groupName,departmentCode,branchCode,educationTypes,yearOfEducation,students,teachers",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("courses", courses, err)
}

func testGroupsInDepartments(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	groupsInDepartments, err := c.GetGroupsInDepartments(zermelogo.GroupsInDepartmentsParams{
		Fields:             "id,departmentOfBranch,name,extendedName,yearOfEducation,isMainGroup,isMentorGroup,mentorUserCodes,studentInDepartments",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("groupindepartments", groupsInDepartments, err)
}

func testHolidays(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	holidays, err := c.GetHolidays(zermelogo.HolidaysParams{
		Fields:             "id,schoolInSchoolYear,name,start,end",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("holidays", holidays, err)
}

func testEnrollmentStatus(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	enrollmentStatus, err := c.GetEnrollmentStatus(zermelogo.EnrollmentStatusParams{
		Student:            "~me",
		Week:               currentWeek(),
		Fields:             "student,week,schoolInSchoolYear,lessonEnrollment,minimalLessonEnrollment,maximalLessonEnrollment",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("enrollmentstatus", enrollmentStatus, err)
}

func testLiveSchedule(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	liveSchedule, err := c.GetLiveSchedule(zermelogo.LiveScheduleParams{
		Student:            "~me",
		Week:               currentWeek(),
		Fields:             "week,user,appointments,status,replacements",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("liveschedule", liveSchedule, err)
}
