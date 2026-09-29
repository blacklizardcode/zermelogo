package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type LiveScheduleParams struct {
	Appointments       []string           `url:"appointments,omitempty" json:"appointments,omitempty"`
	Replacements       []string           `url:"replacements,omitempty" json:"replacements,omitempty"`
	Status             []string           `url:"status,omitempty" json:"status,omitempty"`
	User               string             `url:"user,omitempty" json:"user,omitempty"`
	Week               string             `url:"week,omitempty" json:"week,omitempty"`
	Fields             string             `url:"fields,omitempty" json:"fields,omitempty"`
	SchoolInSchoolYear SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	Student            string             `url:"student,omitempty" json:"student,omitempty"`
	Teacher            string             `url:"teacher,omitempty" json:"teacher,omitempty"`
}

type LiveScheduleResponse struct {
	Response struct {
		Status int                  `json:"status"`
		Data   []LiveScheduleParams `json:"data"`
	} `json:"response"`
}

func (c *Client) GetLiveSchedule(param LiveScheduleParams) ([]LiveScheduleParams, error) {
	var Output LiveScheduleResponse
	path, err := url.Parse("/api/v3/liveschedule")
	if err != nil {
		return []LiveScheduleParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []LiveScheduleParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []LiveScheduleParams{}, err
	}
	return Output.Response.Data, nil
}
