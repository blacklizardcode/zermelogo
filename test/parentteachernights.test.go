package main

import zermelogo "github.com/blacklizardcode/zermelogo"

func testParentTeacherNights(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	parentTeacherNights, err := c.GetParentTeacherNights(zermelogo.ParentTeacherNightsParams{
		Fields:             "id,schoolInSchoolYear,title,description,archived,published,referenceWeek,allowRequestingUntil,hasParentToSchedulerRemark,hasParentToTeacherRemark",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("parentteachernights", parentTeacherNights, err)
}

func testParentToSchedulerRemarks(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	parentToSchedulerRemarks, err := c.GetParentToSchedulerRemarks(zermelogo.ParentToSchedulerRemarksParams{
		Fields:             "id,schoolInSchoolYear,familyMember,parentTeacherNight,remark,processed",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("parenttoschedulerremarks", parentToSchedulerRemarks, err)
}

func testTalkDays(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	talkDays, err := c.GetTalkDays(zermelogo.TalkDaysParams{
		Fields:             "id,parentTeacherNight,date,title,stringDate",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("talkdays", talkDays, err)
}

func testTalkTimeBlocks(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	talkTimeBlocks, err := c.GetTalkTimeBlocks(zermelogo.TalkTimeBlocksParams{
		Fields:             "id,parentTeacherNight,title,span,talkTimeSlots",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("talktimeblocks", talkTimeBlocks, err)
}

func testTalkTimeBlockAbsenceParent(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	talkTimeBlockAbsenceParent, err := c.GetTalkTimeBlockAbsenceParent(zermelogo.TalkTimeBlockAbsenceParentParams{
		Fields:             "id,schoolInSchoolYear,talkDay,talkTimeBlock,talkTimeBlockDefinition,user,overruled",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("talktimeblockabsenceparent", talkTimeBlockAbsenceParent, err)
}
