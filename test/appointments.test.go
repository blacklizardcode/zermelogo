package main

import (
	zermelogo "github.com/blacklizardcode/zermelogo"
	"time"
)

func testAppointments(c *zermelogo.Client, schoolInSchoolYear zermelogo.SchoolInSchoolYear) {
	start := time.Now().AddDate(0, 0, -7).Unix()
	end := time.Now().AddDate(0, 0, 7).Unix()

	appointments, err := c.GetAppointments(zermelogo.AppointmentsParams{
		User:               "~me",
		Start:              start,
		End:                end,
		Fields:             "id,appointmentInstance,start,end,type,subjects,teachers,locations,groups,students,cancelled,valid,hidden,schedulerRemark",
		SchoolInSchoolYear: schoolInSchoolYear,
	})
	reportResult("appointments", appointments, err)
}
