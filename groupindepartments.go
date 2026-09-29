package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

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

type GroupsInDepartmentsResponse struct {
	Response struct {
		Status int                         `json:"status"`
		Data   []GroupsInDepartmentsParams `json:"data"`
	} `json:"response"`
}

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
