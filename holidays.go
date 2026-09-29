package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type HolidaysParams struct {
	End                string             `url:"end,omitempty" json:"end,omitempty"`
	ID                 int                `url:"id,omitempty" json:"id,omitempty"`
	Name               string             `url:"name,omitempty" json:"name,omitempty"`
	SchoolInSchoolYear SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	Start              string             `url:"start,omitempty" json:"start,omitempty"`
	Fields             string             `url:"fields,omitempty" json:"fields,omitempty"`
}

type HolidaysResponse struct {
	Response struct {
		Status int              `json:"status"`
		Data   []HolidaysParams `json:"data"`
	} `json:"response"`
}

func (c *Client) GetHolidays(param HolidaysParams) ([]HolidaysParams, error) {
	var Output HolidaysResponse
	path, err := url.Parse("/api/v3/holidays")
	if err != nil {
		return []HolidaysParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []HolidaysParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []HolidaysParams{}, err
	}
	return Output.Response.Data, nil
}
