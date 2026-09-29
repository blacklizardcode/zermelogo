package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type StudentsParams struct {
	City                        string             `url:"city,omitempty" json:"city,omitempty"`
	DateOfBirth                 string             `url:"dateOfBirth,omitempty" json:"dateOfBirth,omitempty"`
	DepartmentOfBranch          []int              `url:"departmentOfBranch,omitempty" json:"departmentOfBranch,omitempty"`
	Email                       string             `url:"email,omitempty" json:"email,omitempty"`
	FamilyMembers               []string           `url:"familyMembers,omitempty" json:"familyMembers,omitempty"`
	FirstName                   string             `url:"firstName,omitempty" json:"firstName,omitempty"`
	Gender                      string             `url:"gender,omitempty" json:"gender,omitempty"`
	GroupInDepartments          []int              `url:"groupInDepartments,omitempty" json:"groupInDepartments,omitempty"`
	HouseNumber                 string             `url:"houseNumber,omitempty" json:"houseNumber,omitempty"`
	IndividualRequestsParents   bool               `url:"individualRequestsParents,omitempty" json:"individualRequestsParents,omitempty"`
	IntakeProjects              []int              `url:"intakeProjects,omitempty" json:"intakeProjects,omitempty"`
	LastName                    string             `url:"lastName,omitempty" json:"lastName,omitempty"`
	Lwoo                        bool               `url:"lwoo,omitempty" json:"lwoo,omitempty"`
	NonarchivedFamilyMembers    []string           `url:"nonarchivedFamilyMembers,omitempty" json:"nonarchivedFamilyMembers,omitempty"`
	PostalCode                  string             `url:"postalCode,omitempty" json:"postalCode,omitempty"`
	Prefix                      string             `url:"prefix,omitempty" json:"prefix,omitempty"`
	Street                      string             `url:"street,omitempty" json:"street,omitempty"`
	StudentInDepartments        []string           `url:"studentInDepartments,omitempty" json:"studentInDepartments,omitempty"`
	UserCode                    string             `url:"userCode,omitempty" json:"userCode,omitempty"`
	Fields                      string             `url:"fields,omitempty" json:"fields,omitempty"`
	SchoolInSchoolYear          SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	Archived                    bool               `url:"archived,omitempty" json:"archived,omitempty"`
	PlacedIntakeForProject      int                `url:"placedIntakeForProject,omitempty" json:"placedIntakeForProject,omitempty"`
	PossibleIntakeForProject    int                `url:"possibleIntakeForProject,omitempty" json:"possibleIntakeForProject,omitempty"`
	UnplacedIntakeForSchoolYear int                `url:"unplacedIntakeForSchoolYear,omitempty" json:"unplacedIntakeForSchoolYear,omitempty"`
}

type StudentsResponse struct {
	Response struct {
		Status int              `json:"status"`
		Data   []StudentsParams `json:"data"`
	} `json:"response"`
}

func (c *Client) GetStudents(param StudentsParams) ([]StudentsParams, error) {
	var Output StudentsResponse
	path, err := url.Parse("/api/v3/students")
	if err != nil {
		return []StudentsParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []StudentsParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []StudentsParams{}, err
	}
	return Output.Response.Data, nil
}
