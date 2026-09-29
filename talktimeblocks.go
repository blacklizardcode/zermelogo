package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

// TalkTimeBlocksParams holds the query filters for GetTalkTimeBlocks and the fields of the
// records it returns. Zermelo uses the same field names in requests and
// responses, so one struct covers both. Zero values are omitted, and Fields
// selects the columns to return.
type TalkTimeBlocksParams struct {
	ID                 int                `url:"id,omitempty" json:"id,omitempty"`
	ParentTeacherNight int                `url:"parentTeacherNight,omitempty" json:"parentTeacherNight,omitempty"`
	Span               string             `url:"span,omitempty" json:"span,omitempty"`
	TalkTimeSlots      []int              `url:"talkTimeSlots,omitempty" json:"talkTimeSlots,omitempty"`
	Title              string             `url:"title,omitempty" json:"title,omitempty"`
	Fields             string             `url:"fields,omitempty" json:"fields,omitempty"`
	SchoolInSchoolYear SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
}

// TalkTimeBlocksResponse is the envelope the Zermelo API wraps around every response.
// Data holds the decoded records and Status holds the reported status.
type TalkTimeBlocksResponse struct {
	Response struct {
		Status int                    `json:"status"`
		Data   []TalkTimeBlocksParams `json:"data"`
	} `json:"response"`
}

// GetTalkTimeBlocks calls /api/v3/talktimeblocks and returns the records that match param.
// A non nil *APIError is returned when the portal rejects the request.
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
