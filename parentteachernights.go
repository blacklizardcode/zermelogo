package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type ParentTeacherNightsParams struct {
	AllowRequestingUntil            int64              `url:"allowRequestingUntil,omitempty" json:"allowRequestingUntil,omitempty"`
	Archived                        bool               `url:"archived,omitempty" json:"archived,omitempty"`
	Description                     string             `url:"description,omitempty" json:"description,omitempty"`
	EmployeeScheduleFill            int                `url:"employeeScheduleFill,omitempty" json:"employeeScheduleFill,omitempty"`
	EmployeeScheduleSpanWidthFactor int                `url:"employeeScheduleSpanWidthFactor,omitempty" json:"employeeScheduleSpanWidthFactor,omitempty"`
	EmployeeScheduleWidth           int                `url:"employeeScheduleWidth,omitempty" json:"employeeScheduleWidth,omitempty"`
	HasParentToSchedulerRemark      bool               `url:"hasParentToSchedulerRemark,omitempty" json:"hasParentToSchedulerRemark,omitempty"`
	HasParentToTeacherRemark        bool               `url:"hasParentToTeacherRemark,omitempty" json:"hasParentToTeacherRemark,omitempty"`
	ID                              int                `url:"id,omitempty" json:"id,omitempty"`
	MaxNumberOfRequests             int                `url:"maxNumberOfRequests,omitempty" json:"maxNumberOfRequests,omitempty"`
	ParentCanRequestLongTalk        bool               `url:"parentCanRequestLongTalk,omitempty" json:"parentCanRequestLongTalk,omitempty"`
	ParentScheduleFill              int                `url:"parentScheduleFill,omitempty" json:"parentScheduleFill,omitempty"`
	ParentScheduleSpanWidthFactor   int                `url:"parentScheduleSpanWidthFactor,omitempty" json:"parentScheduleSpanWidthFactor,omitempty"`
	ParentScheduleWidth             int                `url:"parentScheduleWidth,omitempty" json:"parentScheduleWidth,omitempty"`
	Published                       bool               `url:"published,omitempty" json:"published,omitempty"`
	ReferenceWeek                   string             `url:"referenceWeek,omitempty" json:"referenceWeek,omitempty"`
	SchedulePublished               bool               `url:"schedulePublished,omitempty" json:"schedulePublished,omitempty"`
	SchoolInSchoolYear              SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	TalkDays                        []string           `url:"talkDays,omitempty" json:"talkDays,omitempty"`
	TalkDaysText                    []string           `url:"talkDaysText,omitempty" json:"talkDaysText,omitempty"`
	Title                           string             `url:"title,omitempty" json:"title,omitempty"`
	Fields                          string             `url:"fields,omitempty" json:"fields,omitempty"`
	FamilyMember                    []int              `url:"familyMember,omitempty" json:"familyMember,omitempty"`
}

type ParentTeacherNightsResponse struct {
	Response struct {
		Status int                         `json:"status"`
		Data   []ParentTeacherNightsParams `json:"data"`
	} `json:"response"`
}

func (c *Client) GetParentTeacherNights(param ParentTeacherNightsParams) ([]ParentTeacherNightsParams, error) {
	var Output ParentTeacherNightsResponse
	path, err := url.Parse("/api/v3/parentteachernights")
	if err != nil {
		return []ParentTeacherNightsParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []ParentTeacherNightsParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []ParentTeacherNightsParams{}, err
	}
	return Output.Response.Data, nil
}
