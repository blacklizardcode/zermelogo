package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type TalkDaysParams struct {
	Date               string             `url:"date,omitempty" json:"date,omitempty"`
	ID                 int                `url:"id,omitempty" json:"id,omitempty"`
	ParentTeacherNight int                `url:"parentTeacherNight,omitempty" json:"parentTeacherNight,omitempty"`
	StringDate         string             `url:"stringDate,omitempty" json:"stringDate,omitempty"`
	Title              string             `url:"title,omitempty" json:"title,omitempty"`
	Fields             string             `url:"fields,omitempty" json:"fields,omitempty"`
	SchoolInSchoolYear SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
}

type TalkDaysResponse struct {
	Response struct {
		Status int              `json:"status"`
		Data   []TalkDaysParams `json:"data"`
	} `json:"response"`
}

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
