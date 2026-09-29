package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type ParentToSchedulerRemarksParams struct {
	FamilyMember       string             `url:"familyMember,omitempty" json:"familyMember,omitempty"`
	ID                 int                `url:"id,omitempty" json:"id,omitempty"`
	ParentTeacherNight int                `url:"parentTeacherNight,omitempty" json:"parentTeacherNight,omitempty"`
	Processed          bool               `url:"processed,omitempty" json:"processed,omitempty"`
	Remark             string             `url:"remark,omitempty" json:"remark,omitempty"`
	SchoolInSchoolYear SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	Fields             string             `url:"fields,omitempty" json:"fields,omitempty"`
}

type ParentToSchedulerRemarksResponse struct {
	Response struct {
		Status int                              `json:"status"`
		Data   []ParentToSchedulerRemarksParams `json:"data"`
	} `json:"response"`
}

func (c *Client) GetParentToSchedulerRemarks(param ParentToSchedulerRemarksParams) ([]ParentToSchedulerRemarksParams, error) {
	var Output ParentToSchedulerRemarksResponse
	path, err := url.Parse("/api/v3/parenttoschedulerremarks")
	if err != nil {
		return []ParentToSchedulerRemarksParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []ParentToSchedulerRemarksParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []ParentToSchedulerRemarksParams{}, err
	}
	return Output.Response.Data, nil
}
