package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

// EmployeesParams holds the query filters for GetEmployees and the fields of the
// records it returns. Zermelo uses the same field names in requests and
// responses, so one struct covers both. Zero values are omitted, and Fields
// selects the columns to return.
type EmployeesParams struct {
	CommencementSchool   string             `url:"commencementSchool,omitempty" json:"commencementSchool,omitempty"`
	CommencementTeaching string             `url:"commencementTeaching,omitempty" json:"commencementTeaching,omitempty"`
	City                 string             `url:"city,omitempty" json:"city,omitempty"`
	Created              int64              `url:"created,omitempty" json:"created,omitempty"`
	DateOfBirth          string             `url:"dateOfBirth,omitempty" json:"dateOfBirth,omitempty"`
	Email                string             `url:"email,omitempty" json:"email,omitempty"`
	EmployeeNumber       string             `url:"employeeNumber,omitempty" json:"employeeNumber,omitempty"`
	FirstName            string             `url:"firstName,omitempty" json:"firstName,omitempty"`
	Gender               string             `url:"gender,omitempty" json:"gender,omitempty"`
	HouseNumber          string             `url:"houseNumber,omitempty" json:"houseNumber,omitempty"`
	LastModified         int64              `url:"lastModified,omitempty" json:"lastModified,omitempty"`
	LastModifiedBy       string             `url:"lastModifiedBy,omitempty" json:"lastModifiedBy,omitempty"`
	LastName             string             `url:"lastName,omitempty" json:"lastName,omitempty"`
	PostalCode           string             `url:"postalCode,omitempty" json:"postalCode,omitempty"`
	Prefix               string             `url:"prefix,omitempty" json:"prefix,omitempty"`
	Roles                []string           `url:"roles,omitempty" json:"roles,omitempty"`
	Street               string             `url:"street,omitempty" json:"street,omitempty"`
	User                 string             `url:"user,omitempty" json:"user,omitempty"`
	UserCode             string             `url:"userCode,omitempty" json:"userCode,omitempty"`
	Fields               string             `url:"fields,omitempty" json:"fields,omitempty"`
	SchoolInSchoolYear   SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	BranchOfSchool       int                `url:"branchOfSchool,omitempty" json:"branchOfSchool,omitempty"`
	Archived             bool               `url:"archived,omitempty" json:"archived,omitempty"`
	HasPassword          bool               `url:"hasPassword,omitempty" json:"hasPassword,omitempty"`
	Ldap                 bool               `url:"ldap,omitempty" json:"ldap,omitempty"`
	TeacherTeam          int                `url:"teacherTeam,omitempty" json:"teacherTeam,omitempty"`
}

// EmployeesResponse is the envelope the Zermelo API wraps around every response.
// Data holds the decoded records and Status holds the reported status.
type EmployeesResponse struct {
	Response struct {
		Status int               `json:"status"`
		Data   []EmployeesParams `json:"data"`
	} `json:"response"`
}

// GetEmployees calls /api/v3/employees and returns the records that match param.
// A non nil *APIError is returned when the portal rejects the request.
func (c *Client) GetEmployees(param EmployeesParams) ([]EmployeesParams, error) {
	var Output EmployeesResponse
	path, err := url.Parse("/api/v3/employees")
	if err != nil {
		return []EmployeesParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []EmployeesParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []EmployeesParams{}, err
	}
	return Output.Response.Data, nil
}
