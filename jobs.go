package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type JobsParams struct {
	ClockHoursGross          float64            `url:"clockHoursGross,omitempty" json:"clockHoursGross,omitempty"`
	Contract                 int                `url:"contract,omitempty" json:"contract,omitempty"`
	Created                  int64              `url:"created,omitempty" json:"created,omitempty"`
	DurationManually         bool               `url:"durationManually,omitempty" json:"durationManually,omitempty"`
	Employee                 string             `url:"employee,omitempty" json:"employee,omitempty"`
	EmploymentType           string             `url:"employmentType,omitempty" json:"employmentType,omitempty"`
	EmployerCost             int                `url:"employerCost,omitempty" json:"employerCost,omitempty"`
	End                      string             `url:"end,omitempty" json:"end,omitempty"`
	ExternalID               string             `url:"externalId,omitempty" json:"externalId,omitempty"`
	FTE                      float64            `url:"fte,omitempty" json:"fte,omitempty"`
	FTEManually              bool               `url:"fteManually,omitempty" json:"fteManually,omitempty"`
	FTEReal                  float64            `url:"fteReal,omitempty" json:"fteReal,omitempty"`
	FirstName                string             `url:"firstName,omitempty" json:"firstName,omitempty"`
	FunctionCategory         string             `url:"functionCategory,omitempty" json:"functionCategory,omitempty"`
	FunctionCategoryManually bool               `url:"functionCategoryManually,omitempty" json:"functionCategoryManually,omitempty"`
	ID                       int                `url:"id,omitempty" json:"id,omitempty"`
	LastModified             int64              `url:"lastModified,omitempty" json:"lastModified,omitempty"`
	LastModifiedBy           string             `url:"lastModifiedBy,omitempty" json:"lastModifiedBy,omitempty"`
	LastName                 string             `url:"lastName,omitempty" json:"lastName,omitempty"`
	Prefix                   string             `url:"prefix,omitempty" json:"prefix,omitempty"`
	PrivateComment           string             `url:"privateComment,omitempty" json:"privateComment,omitempty"`
	ProjectBudget            int                `url:"projectBudget,omitempty" json:"projectBudget,omitempty"`
	PublicComment            string             `url:"publicComment,omitempty" json:"publicComment,omitempty"`
	Start                    string             `url:"start,omitempty" json:"start,omitempty"`
	Type                     string             `url:"type,omitempty" json:"type,omitempty"`
	Fields                   string             `url:"fields,omitempty" json:"fields,omitempty"`
	SchoolInSchoolYear       SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	TeacherTeam              int                `url:"teacherTeam,omitempty" json:"teacherTeam,omitempty"`
	Week                     string             `url:"week,omitempty" json:"week,omitempty"`
}

type JobsResponse struct {
	Response struct {
		Status int          `json:"status"`
		Data   []JobsParams `json:"data"`
	} `json:"response"`
}

func (c *Client) GetJobs(param JobsParams) ([]JobsParams, error) {
	var Output JobsResponse
	path, err := url.Parse("/api/v3/jobs")
	if err != nil {
		return []JobsParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []JobsParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []JobsParams{}, err
	}
	return Output.Response.Data, nil
}
