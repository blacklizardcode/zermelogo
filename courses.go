package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

// CoursesParams holds the query filters for GetCourses and the fields of the
// records it returns. Zermelo uses the same field names in requests and
// responses, so one struct covers both. Zero values are omitted, and Fields
// selects the columns to return.
type CoursesParams struct {
	BranchCode            string             `url:"branchCode,omitempty" json:"branchCode,omitempty"`
	ChoosableInDepartment int                `url:"choosableInDepartment,omitempty" json:"choosableInDepartment,omitempty"`
	DepartmentCode        string             `url:"departmentCode,omitempty" json:"departmentCode,omitempty"`
	EducationTypes        []string           `url:"educationTypes,omitempty" json:"educationTypes,omitempty"`
	GroupName             string             `url:"groupName,omitempty" json:"groupName,omitempty"`
	ID                    int                `url:"id,omitempty" json:"id,omitempty"`
	ILTCode               int                `url:"iltCode,omitempty" json:"iltCode,omitempty"`
	SchoolYear            int                `url:"schoolYear,omitempty" json:"schoolYear,omitempty"`
	Students              []string           `url:"students,omitempty" json:"students,omitempty"`
	SubjectCode           string             `url:"subjectCode,omitempty" json:"subjectCode,omitempty"`
	SubjectName           string             `url:"subjectName,omitempty" json:"subjectName,omitempty"`
	Teachers              []string           `url:"teachers,omitempty" json:"teachers,omitempty"`
	YearOfEducation       int                `url:"yearOfEducation,omitempty" json:"yearOfEducation,omitempty"`
	Fields                string             `url:"fields,omitempty" json:"fields,omitempty"`
	SchoolInSchoolYear    SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	Student               string             `url:"student,omitempty" json:"student,omitempty"`
	Teacher               string             `url:"teacher,omitempty" json:"teacher,omitempty"`
}

// CoursesResponse is the envelope the Zermelo API wraps around every response.
// Data holds the decoded records and Status holds the reported status.
type CoursesResponse struct {
	Response struct {
		Status int             `json:"status"`
		Data   []CoursesParams `json:"data"`
	} `json:"response"`
}

// GetCourses calls /api/v3/courses and returns the records that match param.
// A non nil *APIError is returned when the portal rejects the request.
func (c *Client) GetCourses(param CoursesParams) ([]CoursesParams, error) {
	var Output CoursesResponse
	path, err := url.Parse("/api/v3/courses")
	if err != nil {
		return []CoursesParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []CoursesParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []CoursesParams{}, err
	}
	return Output.Response.Data, nil
}
