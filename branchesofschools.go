package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type BranchesOfSchoolsParams struct {
	Archived           bool               `url:"archived,omitempty" json:"archived,omitempty"`
	Branch             string             `url:"branch,omitempty" json:"branch,omitempty"`
	ID                 int                `url:"id,omitempty" json:"id,omitempty"`
	Name               string             `url:"name,omitempty" json:"name,omitempty"`
	SchoolInSchoolYear SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	SchoolYear         int                `url:"schoolYear,omitempty" json:"schoolYear,omitempty"`
	Fields             string             `url:"fields,omitempty" json:"fields,omitempty"`
}

type BranchesOfSchoolsResponse struct {
	Response struct {
		Status int                       `json:"status"`
		Data   []BranchesOfSchoolsParams `json:"data"`
	} `json:"response"`
}

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
