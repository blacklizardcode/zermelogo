package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

// TalkDaysParams holds the query filters for GetTalkDays and the fields of the
// records it returns. Zermelo uses the same field names in requests and
// responses, so one struct covers both. Zero values are omitted, and Fields
// selects the columns to return.
type TalkDaysParams struct {
	Date               string             `url:"date,omitempty" json:"date,omitempty"`
	ID                 int                `url:"id,omitempty" json:"id,omitempty"`
	ParentTeacherNight int                `url:"parentTeacherNight,omitempty" json:"parentTeacherNight,omitempty"`
	StringDate         string             `url:"stringDate,omitempty" json:"stringDate,omitempty"`
	Title              string             `url:"title,omitempty" json:"title,omitempty"`
	Fields             string             `url:"fields,omitempty" json:"fields,omitempty"`
	SchoolInSchoolYear SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
}

// TalkDaysResponse is the envelope the Zermelo API wraps around every response.
// Data holds the decoded records and Status holds the reported status.
type TalkDaysResponse struct {
	Response struct {
		Status int              `json:"status"`
		Data   []TalkDaysParams `json:"data"`
	} `json:"response"`
}

// GetTalkDays calls /api/v3/talkdays and returns the records that match param.
// A non nil *APIError is returned when the portal rejects the request.
func (c *Client) GetTalkDays(param TalkDaysParams) ([]TalkDaysParams, error) {
	var Output TalkDaysResponse
	path, err := url.Parse("/api/v3/talkdays")
	if err != nil {
		return []TalkDaysParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []TalkDaysParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []TalkDaysParams{}, err
	}
	return Output.Response.Data, nil
}
