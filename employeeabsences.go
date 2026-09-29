package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

// EmployeeAbsencesParams holds the query filters for GetEmployeeAbsences and the fields of the
// records it returns. Zermelo uses the same field names in requests and
// responses, so one struct covers both. Zero values are omitted, and Fields
// selects the columns to return.
type EmployeeAbsencesParams struct {
	AbsenceDuration         float64            `url:"absenceDuration,omitempty" json:"absenceDuration,omitempty"`
	AbsenceDurationManually bool               `url:"absenceDurationManually,omitempty" json:"absenceDurationManually,omitempty"`
	AbsenceType             int                `url:"absenceType,omitempty" json:"absenceType,omitempty"`
	AbsenceTypeCode         string             `url:"absenceTypeCode,omitempty" json:"absenceTypeCode,omitempty"`
	AbsenceTypePriority     int                `url:"absenceTypePriority,omitempty" json:"absenceTypePriority,omitempty"`
	Contract                int                `url:"contract,omitempty" json:"contract,omitempty"`
	Created                 int64              `url:"created,omitempty" json:"created,omitempty"`
	Definitive              bool               `url:"definitive,omitempty" json:"definitive,omitempty"`
	Employee                string             `url:"employee,omitempty" json:"employee,omitempty"`
	End                     int64              `url:"end,omitempty" json:"end,omitempty"`
	FirstName               string             `url:"firstName,omitempty" json:"firstName,omitempty"`
	ID                      int                `url:"id,omitempty" json:"id,omitempty"`
	LastModified            int64              `url:"lastModified,omitempty" json:"lastModified,omitempty"`
	LastModifiedBy          string             `url:"lastModifiedBy,omitempty" json:"lastModifiedBy,omitempty"`
	LastName                string             `url:"lastName,omitempty" json:"lastName,omitempty"`
	Leave                   int                `url:"leave,omitempty" json:"leave,omitempty"`
	LinkedToStructuralLeave bool               `url:"linkedToStructuralLeave,omitempty" json:"linkedToStructuralLeave,omitempty"`
	Prefix                  string             `url:"prefix,omitempty" json:"prefix,omitempty"`
	PrivateComment          string             `url:"privateComment,omitempty" json:"privateComment,omitempty"`
	Start                   int64              `url:"start,omitempty" json:"start,omitempty"`
	Fields                  string             `url:"fields,omitempty" json:"fields,omitempty"`
	SchoolInSchoolYear      SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	BranchOfSchool          int                `url:"branchOfSchool,omitempty" json:"branchOfSchool,omitempty"`
	EndDate                 string             `url:"endDate,omitempty" json:"endDate,omitempty"`
	EndWeek                 string             `url:"endWeek,omitempty" json:"endWeek,omitempty"`
	RelatedToEmployee       string             `url:"relatedToEmployee,omitempty" json:"relatedToEmployee,omitempty"`
	SectionOfBranch         int                `url:"sectionOfBranch,omitempty" json:"sectionOfBranch,omitempty"`
	StartDate               string             `url:"startDate,omitempty" json:"startDate,omitempty"`
	StartWeek               string             `url:"startWeek,omitempty" json:"startWeek,omitempty"`
	TeacherTeam             int                `url:"teacherTeam,omitempty" json:"teacherTeam,omitempty"`
	Unlinked                bool               `url:"unlinked,omitempty" json:"unlinked,omitempty"`
}

// EmployeeAbsencesResponse is the envelope the Zermelo API wraps around every response.
// Data holds the decoded records and Status holds the reported status.
type EmployeeAbsencesResponse struct {
	Response struct {
		Status int                      `json:"status"`
		Data   []EmployeeAbsencesParams `json:"data"`
	} `json:"response"`
}

// GetEmployeeAbsences calls /api/v3/employeeabsences and returns the records that match param.
// A non nil *APIError is returned when the portal rejects the request.
func (c *Client) GetEmployeeAbsences(param EmployeeAbsencesParams) ([]EmployeeAbsencesParams, error) {
	var Output EmployeeAbsencesResponse
	path, err := url.Parse("/api/v3/employeeabsences")
	if err != nil {
		return []EmployeeAbsencesParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []EmployeeAbsencesParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []EmployeeAbsencesParams{}, err
	}
	return Output.Response.Data, nil
}
