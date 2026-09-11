package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type AppointmentsParams struct {
	AppointmentInstance                    int      `url:"appointmentInstance,omitempty" json:"appointmentInstance,omitempty"`
	BranchOfSchool                         int      `url:"branchOfSchool,omitempty" json:"branchOfSchool,omitempty"`
	Cancelled                              bool     `url:"cancelled,omitempty" json:"cancelled,omitempty"`
	ChangeDescription                      string   `url:"changeDescription,omitempty" json:"changeDescription,omitempty"`
	Created                                int      `url:"created,omitempty" json:"created,omitempty"`
	End                                    int64    `url:"end" json:"end"`
	EndTimeSlot                            int      `url:"endTimeSlot,omitempty" json:"endTimeSlot,omitempty"`
	EndTimeSlotName                        string   `url:"endTimeSlotName,omitempty" json:"endTimeSlotName,omitempty"`
	GroupChanged                           bool     `url:"groupChanged,omitempty" json:"groupChanged,omitempty"`
	Hidden                                 bool     `url:"hidden,omitempty" json:"hidden,omitempty"`
	ID                                     int      `url:"id,omitempty" json:"id,omitempty"`
	LocationChanged                        bool     `url:"locationChanged,omitempty" json:"locationChanged,omitempty"`
	Optional                               bool     `url:"optional,omitempty" json:"optional,omitempty"`
	SchedulerRemark                        string   `url:"schedulerRemark,omitempty" json:"schedulerRemark,omitempty"`
	Start                                  int64    `url:"start" json:"start"`
	StartTimeSlot                          int      `url:"startTimeSlot,omitempty" json:"startTimeSlot,omitempty"`
	StartTimeSlotName                      string   `url:"startTimeSlotName,omitempty" json:"startTimeSlotName,omitempty"`
	Subjects                               []string `url:"subjects,omitempty" json:"subjects,omitempty"`
	TeacherChanged                         bool     `url:"teacherChanged,omitempty" json:"teacherChanged,omitempty"`
	Teachers                               []string `url:"teachers,omitempty" json:"teachers,omitempty"`
	TimeChanged                            bool     `url:"timeChanged,omitempty" json:"timeChanged,omitempty"`
	Type                                   string   `url:"type,omitempty" json:"type,omitempty"`
	Valid                                  bool     `url:"valid,omitempty" json:"valid,omitempty"`
	Changed                                string   `url:"changed,omitempty" json:"changed,omitempty"`
	Fields                                 string   `url:"fields,omitempty" json:"fields,omitempty"`
	IncludeHidden                          string   `url:"includeHidden,omitempty" json:"includeHidden,omitempty"`
	ModifiedSince                          string   `url:"modifiedSince,omitempty" json:"modifiedSince,omitempty"`
	SchoolInSchoolYear                     string   `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	Base                                   string   `url:"base,omitempty" json:"base,omitempty"`
	Branch                                 string   `url:"branch,omitempty" json:"branch,omitempty"`
	ChoosableInDepartments                 string   `url:"choosableInDepartments,omitempty" json:"choosableInDepartments,omitempty"`
	Creator                                string   `url:"creator,omitempty" json:"creator,omitempty"`
	ExtraStudentSource                     string   `url:"extraStudentSource,omitempty" json:"extraStudentSource,omitempty"`
	GroupsInDepartments                    []int    `url:"groupsInDepartments,omitempty" json:"groupsInDepartments,omitempty"`
	IgnoreSubstitutions                    string   `url:"ignoreSubstitutions,omitempty" json:"ignoreSubstitutions,omitempty"`
	LocationsOfBranch                      []int    `url:"locationsOfBranch,omitempty" json:"locationsOfBranch,omitempty"`
	Modified                               bool     `url:"modified,omitempty" json:"modified,omitempty"`
	Moved                                  bool     `url:"moved,omitempty" json:"moved,omitempty"`
	New                                    bool     `url:"new,omitempty" json:"new,omitempty"`
	AppointmentModifiedSince               string   `url:"appointmentModifiedSince,omitempty" json:"appointmentModifiedSince,omitempty"`
	ContainsStudentsFromDepartmentOfBranch string   `url:"containsStudentsFromDepartmentOfBranch,omitempty" json:"containsStudentsFromDepartmentOfBranch,omitempty"`
	ContainsStudentsFromGroupInDepartment  string   `url:"containsStudentsFromGroupInDepartment,omitempty" json:"containsStudentsFromGroupInDepartment,omitempty"`
	ContainsTeachersFromTeacherTeam        string   `url:"containsTeachersFromTeacherTeam,omitempty" json:"containsTeachersFromTeacherTeam,omitempty"`
	DepartmentOfBranch                     string   `url:"departmentOfBranch,omitempty" json:"departmentOfBranch,omitempty"`
	EndWeekOffset                          string   `url:"endWeekOffset,omitempty" json:"endWeekOffset,omitempty"`
	ExtraStudentSourceForDepartment        string   `url:"extraStudentSourceForDepartment,omitempty" json:"extraStudentSourceForDepartment,omitempty"`
	HasCreator                             string   `url:"hasCreator,omitempty" json:"hasCreator,omitempty"`
	PossibleStudents                       string   `url:"possibleStudents,omitempty" json:"possibleStudents,omitempty"`
	StartWeekOffset                        string   `url:"startWeekOffset,omitempty" json:"startWeekOffset,omitempty"`
	User                                   string   `url:"user,omitempty" json:"user,omitempty"`
	Locations                              []string `url:"locations,omitempty" json:"locations,omitempty"`
}

type AppointmentsResponse struct {
	Response struct {
		Status int                  `json:"status"`
		Data   []AppointmentsParams `json:"data"`
	} `json:"response"`
}

func (c *Client) GetAppointments(param AppointmentsParams) ([]AppointmentsParams, error) {
	var Output AppointmentsResponse
	path, err := url.Parse("/api/v3/appointments")
	if err != nil {
		return []AppointmentsParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []AppointmentsParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []AppointmentsParams{}, err
	}
	return Output.Response.Data, nil
}
