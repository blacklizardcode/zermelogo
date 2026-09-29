package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type FamilyMembersParams struct {
	Archived                bool               `url:"archived,omitempty" json:"archived,omitempty"`
	City                    string             `url:"city,omitempty" json:"city,omitempty"`
	Email                   string             `url:"email,omitempty" json:"email,omitempty"`
	FirstName               string             `url:"firstName,omitempty" json:"firstName,omitempty"`
	HouseNumber             string             `url:"houseNumber,omitempty" json:"houseNumber,omitempty"`
	IncludedForRequests     bool               `url:"includedForRequests,omitempty" json:"includedForRequests,omitempty"`
	MagisterUUID            string             `url:"magisterUUID,omitempty" json:"magisterUUID,omitempty"`
	PostalCode              string             `url:"postalCode,omitempty" json:"postalCode,omitempty"`
	Prefix                  string             `url:"prefix,omitempty" json:"prefix,omitempty"`
	ScheduleFill            int                `url:"scheduleFill,omitempty" json:"scheduleFill,omitempty"`
	ScheduleSpanWidthFactor int                `url:"scheduleSpanWidthFactor,omitempty" json:"scheduleSpanWidthFactor,omitempty"`
	ScheduleWidth           int                `url:"scheduleWidth,omitempty" json:"scheduleWidth,omitempty"`
	SomUUID                 string             `url:"somUUID,omitempty" json:"somUUID,omitempty"`
	Street                  string             `url:"street,omitempty" json:"street,omitempty"`
	Students                []string           `url:"students,omitempty" json:"students,omitempty"`
	UserCode                string             `url:"userCode,omitempty" json:"userCode,omitempty"`
	Fields                  string             `url:"fields,omitempty" json:"fields,omitempty"`
	SchoolInSchoolYear      SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	HasPassword             bool               `url:"hasPassword,omitempty" json:"hasPassword,omitempty"`
	Ldap                    bool               `url:"ldap,omitempty" json:"ldap,omitempty"`
	ParentTeacherNight      int                `url:"parentTeacherNight,omitempty" json:"parentTeacherNight,omitempty"`
}

type FamilyMembersResponse struct {
	Response struct {
		Status int                   `json:"status"`
		Data   []FamilyMembersParams `json:"data"`
	} `json:"response"`
}

func (c *Client) GetFamilyMembers(param FamilyMembersParams) ([]FamilyMembersParams, error) {
	var Output FamilyMembersResponse
	path, err := url.Parse("/api/v3/familymembers")
	if err != nil {
		return []FamilyMembersParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []FamilyMembersParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []FamilyMembersParams{}, err
	}
	return Output.Response.Data, nil
}
