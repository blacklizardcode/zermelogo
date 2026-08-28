package zermelogo

import "time"


type appointmentRequest struct {
	AppointmentInstance int64
	AppointmentLastModified time.Time
	BranchOfSchool int
	Cancelled bool
	ChangeDescription string
	ChoosableInDepartmentCodes []string
	Content string
	Created time.Time
	End time.Time
	EndTimeSlot int
	EndTimeSlotName string
	GroupChanged bool
	Groups []string
	Hidden bool
	Id int64
	LastModified time.Time
	LocationChanged bool
	Locations []string
	Optional bool
	SchedulerRemark string
	Start time.Time
	StartTimeSlot int
	StartTimeSlotName string
	Students []string
	Subjects []string
	TeacherChanged bool
	Teachers []string
	TimeChanged bool
	Type string
	Valid bool
	/**/
}