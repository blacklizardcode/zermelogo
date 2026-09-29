package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type TalkTimeBlockAbsenceParentParams struct {
	ID                      int                `url:"id,omitempty" json:"id,omitempty"`
	Overruled               bool               `url:"overruled,omitempty" json:"overruled,omitempty"`
	SchoolInSchoolYear      SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	TalkDay                 int                `url:"talkDay,omitempty" json:"talkDay,omitempty"`
	TalkTimeBlock           int                `url:"talkTimeBlock,omitempty" json:"talkTimeBlock,omitempty"`
	TalkTimeBlockDefinition int                `url:"talkTimeBlockDefinition,omitempty" json:"talkTimeBlockDefinition,omitempty"`
	User                    string             `url:"user,omitempty" json:"user,omitempty"`
	Fields                  string             `url:"fields,omitempty" json:"fields,omitempty"`
	ParentTeacherNight      int                `url:"parentTeacherNight,omitempty" json:"parentTeacherNight,omitempty"`
}

type TalkTimeBlockAbsenceParentResponse struct {
	Response struct {
		Status int                                `json:"status"`
		Data   []TalkTimeBlockAbsenceParentParams `json:"data"`
	} `json:"response"`
}

func (c *Client) GetTalkTimeBlockAbsenceParent(param TalkTimeBlockAbsenceParentParams) ([]TalkTimeBlockAbsenceParentParams, error) {
	var Output TalkTimeBlockAbsenceParentResponse
	path, err := url.Parse("/api/v3/talktimeblockabsenceparent")
	if err != nil {
		return []TalkTimeBlockAbsenceParentParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []TalkTimeBlockAbsenceParentParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []TalkTimeBlockAbsenceParentParams{}, err
	}
	return Output.Response.Data, nil
}
