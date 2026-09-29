package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

// GroupsInDepartmentsParams holds the query filters for GetGroupsInDepartments and the fields of the
// records it returns. Zermelo uses the same field names in requests and
// responses, so one struct covers both. Zero values are omitted, and Fields
// selects the columns to return.
type GroupsInDepartmentsParams struct {
	DepartmentOfBranch   int                `url:"departmentOfBranch,omitempty" json:"departmentOfBranch,omitempty"`
	Educations           []int              `url:"educations,omitempty" json:"educations,omitempty"`
	ExtendedName         string             `url:"extendedName,omitempty" json:"extendedName,omitempty"`
	ID                   int                `url:"id,omitempty" json:"id,omitempty"`
	IsMainGroup          bool               `url:"isMainGroup,omitempty" json:"isMainGroup,omitempty"`
	IsMentorGroup        bool               `url:"isMentorGroup,omitempty" json:"isMentorGroup,omitempty"`
	MentorUserCodes      []string           `url:"mentorUserCodes,omitempty" json:"mentorUserCodes,omitempty"`
	Name                 string             `url:"name,omitempty" json:"name,omitempty"`
	StudentInDepartments []int              `url:"studentInDepartments,omitempty" json:"studentInDepartments,omitempty"`
	YearOfEducation      int                `url:"yearOfEducation,omitempty" json:"yearOfEducation,omitempty"`
	Fields               string             `url:"fields,omitempty" json:"fields,omitempty"`
	SchoolInSchoolYear   SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	BranchOfSchool       int                `url:"branchOfSchool,omitempty" json:"branchOfSchool,omitempty"`
	IsMentorOrMainGroup  bool               `url:"isMentorOrMainGroup,omitempty" json:"isMentorOrMainGroup,omitempty"`
	SchoolYear           int                `url:"schoolYear,omitempty" json:"schoolYear,omitempty"`
}

// GroupsInDepartmentsResponse is the envelope the Zermelo API wraps around every response.
// Data holds the decoded records and Status holds the reported status.
type GroupsInDepartmentsResponse struct {
	Response struct {
		Status int                         `json:"status"`
		Data   []GroupsInDepartmentsParams `json:"data"`
	} `json:"response"`
}

// GetGroupsInDepartments calls /api/v3/groupindepartments and returns the records that match param.
// A non nil *APIError is returned when the portal rejects the request.
func (c *Client) GetGroupsInDepartments(param GroupsInDepartmentsParams) ([]GroupsInDepartmentsParams, error) {
	var Output GroupsInDepartmentsResponse
	path, err := url.Parse("/api/v3/groupindepartments")
	if err != nil {
		return []GroupsInDepartmentsParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []GroupsInDepartmentsParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []GroupsInDepartmentsParams{}, err
	}
	return Output.Response.Data, nil
}
