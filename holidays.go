package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

// HolidaysParams holds the query filters for GetHolidays and the fields of the
// records it returns. Zermelo uses the same field names in requests and
// responses, so one struct covers both. Zero values are omitted, and Fields
// selects the columns to return.
type HolidaysParams struct {
	End                string             `url:"end,omitempty" json:"end,omitempty"`
	ID                 int                `url:"id,omitempty" json:"id,omitempty"`
	Name               string             `url:"name,omitempty" json:"name,omitempty"`
	SchoolInSchoolYear SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	Start              string             `url:"start,omitempty" json:"start,omitempty"`
	Fields             string             `url:"fields,omitempty" json:"fields,omitempty"`
}

// HolidaysResponse is the envelope the Zermelo API wraps around every response.
// Data holds the decoded records and Status holds the reported status.
type HolidaysResponse struct {
	Response struct {
		Status int              `json:"status"`
		Data   []HolidaysParams `json:"data"`
	} `json:"response"`
}

// GetHolidays calls /api/v3/holidays and returns the records that match param.
// A non nil *APIError is returned when the portal rejects the request.
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
