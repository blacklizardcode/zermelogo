package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

// LocationsOfBranchesParams holds the query filters for GetLocationsOfBranches and the fields of the
// records it returns. Zermelo uses the same field names in requests and
// responses, so one struct covers both. Zero values are omitted, and Fields
// selects the columns to return.
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

// LocationsOfBranchesResponse is the envelope the Zermelo API wraps around every response.
// Data holds the decoded records and Status holds the reported status.
type LocationsOfBranchesResponse struct {
	Response struct {
		Status int                         `json:"status"`
		Data   []LocationsOfBranchesParams `json:"data"`
	} `json:"response"`
}

// GetLocationsOfBranches calls /api/v3/locationofbranches and returns the records that match param.
// A non nil *APIError is returned when the portal rejects the request.
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
