package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type TalkTimeBlocksParams struct {
	ID                 int                `url:"id,omitempty" json:"id,omitempty"`
	ParentTeacherNight int                `url:"parentTeacherNight,omitempty" json:"parentTeacherNight,omitempty"`
	Span               string             `url:"span,omitempty" json:"span,omitempty"`
	TalkTimeSlots      []int              `url:"talkTimeSlots,omitempty" json:"talkTimeSlots,omitempty"`
	Title              string             `url:"title,omitempty" json:"title,omitempty"`
	Fields             string             `url:"fields,omitempty" json:"fields,omitempty"`
	SchoolInSchoolYear SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
}

type TalkTimeBlocksResponse struct {
	Response struct {
		Status int                    `json:"status"`
		Data   []TalkTimeBlocksParams `json:"data"`
	} `json:"response"`
}

func (c *Client) GetTalkTimeBlocks(param TalkTimeBlocksParams) ([]TalkTimeBlocksParams, error) {
	var Output TalkTimeBlocksResponse
	path, err := url.Parse("/api/v3/talktimeblocks")
	if err != nil {
		return []TalkTimeBlocksParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []TalkTimeBlocksParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []TalkTimeBlocksParams{}, err
	}
	return Output.Response.Data, nil
}
