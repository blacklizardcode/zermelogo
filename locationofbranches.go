package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type LocationsOfBranchesParams struct {
	AllowMeetings                  bool               `url:"allowMeetings,omitempty" json:"allowMeetings,omitempty"`
	Branch                         int                `url:"branch,omitempty" json:"branch,omitempty"`
	BranchOfSchool                 int                `url:"branchOfSchool,omitempty" json:"branchOfSchool,omitempty"`
	CourseCapacity                 int                `url:"courseCapacity,omitempty" json:"courseCapacity,omitempty"`
	ID                             int                `url:"id,omitempty" json:"id,omitempty"`
	Name                           string             `url:"name,omitempty" json:"name,omitempty"`
	ParentTeacherNightCapacity     int                `url:"parentteachernightCapacity,omitempty" json:"parentteachernightCapacity,omitempty"`
	SchoolInSchoolYear             SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	SecondaryBranches              []int              `url:"secondaryBranches,omitempty" json:"secondaryBranches,omitempty"`
	SupportsConcurrentAppointments bool               `url:"supportsConcurrentAppointments,omitempty" json:"supportsConcurrentAppointments,omitempty"`
	Fields                         string             `url:"fields,omitempty" json:"fields,omitempty"`
}

type LocationsOfBranchesResponse struct {
	Response struct {
		Status int                         `json:"status"`
		Data   []LocationsOfBranchesParams `json:"data"`
	} `json:"response"`
}

func (c *Client) GetLocationsOfBranches(param LocationsOfBranchesParams) ([]LocationsOfBranchesParams, error) {
	var Output LocationsOfBranchesResponse
	path, err := url.Parse("/api/v3/locationofbranches")
	if err != nil {
		return []LocationsOfBranchesParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []LocationsOfBranchesParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []LocationsOfBranchesParams{}, err
	}
	return Output.Response.Data, nil
}
