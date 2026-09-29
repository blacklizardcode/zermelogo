package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

// BranchesOfSchoolsParams holds the query filters for GetBranchesOfSchools and the fields of the
// records it returns. Zermelo uses the same field names in requests and
// responses, so one struct covers both. Zero values are omitted, and Fields
// selects the columns to return.
type BranchesOfSchoolsParams struct {
	Archived           bool               `url:"archived,omitempty" json:"archived,omitempty"`
	Branch             string             `url:"branch,omitempty" json:"branch,omitempty"`
	ID                 int                `url:"id,omitempty" json:"id,omitempty"`
	Name               string             `url:"name,omitempty" json:"name,omitempty"`
	SchoolInSchoolYear SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	SchoolYear         int                `url:"schoolYear,omitempty" json:"schoolYear,omitempty"`
	Fields             string             `url:"fields,omitempty" json:"fields,omitempty"`
}

// BranchesOfSchoolsResponse is the envelope the Zermelo API wraps around every response.
// Data holds the decoded records and Status holds the reported status.
type BranchesOfSchoolsResponse struct {
	Response struct {
		Status int                       `json:"status"`
		Data   []BranchesOfSchoolsParams `json:"data"`
	} `json:"response"`
}

// GetBranchesOfSchools calls /api/v3/branchesofschools and returns the records that match param.
// A non nil *APIError is returned when the portal rejects the request.
func (c *Client) GetBranchesOfSchools(param BranchesOfSchoolsParams) ([]BranchesOfSchoolsParams, error) {
	var Output BranchesOfSchoolsResponse
	path, err := url.Parse("/api/v3/branchesofschools")
	if err != nil {
		return []BranchesOfSchoolsParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []BranchesOfSchoolsParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []BranchesOfSchoolsParams{}, err
	}
	return Output.Response.Data, nil
}
