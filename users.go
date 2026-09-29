package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

// UsersParams holds the query filters for GetUsers and the fields of the
// records it returns. Zermelo uses the same field names in requests and
// responses, so one struct covers both. Zero values are omitted, and Fields
// selects the columns to return.
type UsersParams struct {
	Admin                               bool               `url:"admin,omitempty" json:"admin,omitempty"`
	Archived                            bool               `url:"archived,omitempty" json:"archived,omitempty"`
	ArchivedSchoolInSchoolYears         []int              `url:"archivedSchoolInSchoolYears,omitempty" json:"archivedSchoolInSchoolYears,omitempty"`
	ArchivedTime                        int64              `url:"archivedTime,omitempty" json:"archivedTime,omitempty"`
	City                                string             `url:"city,omitempty" json:"city,omitempty"`
	Code                                string             `url:"code,omitempty" json:"code,omitempty"`
	DateOfBirth                         string             `url:"dateOfBirth,omitempty" json:"dateOfBirth,omitempty"`
	DisplayName                         string             `url:"displayName,omitempty" json:"displayName,omitempty"`
	Email                               string             `url:"email,omitempty" json:"email,omitempty"`
	EmployeeSchoolInSchoolYears         []int              `url:"employeeSchoolInSchoolYears,omitempty" json:"employeeSchoolInSchoolYears,omitempty"`
	FamilyMemberSchoolInSchoolYears     []int              `url:"familyMemberSchoolInSchoolYears,omitempty" json:"familyMemberSchoolInSchoolYears,omitempty"`
	FirstName                           string             `url:"firstName,omitempty" json:"firstName,omitempty"`
	Gender                              string             `url:"gender,omitempty" json:"gender,omitempty"`
	HasPassword                         bool               `url:"hasPassword,omitempty" json:"hasPassword,omitempty"`
	HouseNumber                         string             `url:"houseNumber,omitempty" json:"houseNumber,omitempty"`
	IsApplicationManager                bool               `url:"isApplicationManager,omitempty" json:"isApplicationManager,omitempty"`
	IsBranchLeader                      bool               `url:"isBranchLeader,omitempty" json:"isBranchLeader,omitempty"`
	IsDean                              bool               `url:"isDean,omitempty" json:"isDean,omitempty"`
	IsEmployee                          bool               `url:"isEmployee,omitempty" json:"isEmployee,omitempty"`
	IsFamilyMember                      bool               `url:"isFamilyMember,omitempty" json:"isFamilyMember,omitempty"`
	IsMentor                            bool               `url:"isMentor,omitempty" json:"isMentor,omitempty"`
	IsParentTeacherNightScheduler       bool               `url:"isParentTeacherNightScheduler,omitempty" json:"isParentTeacherNightScheduler,omitempty"`
	IsSchoolLeader                      bool               `url:"isSchoolLeader,omitempty" json:"isSchoolLeader,omitempty"`
	IsSchoolScheduler                   bool               `url:"isSchoolScheduler,omitempty" json:"isSchoolScheduler,omitempty"`
	IsSectionLeader                     bool               `url:"isSectionLeader,omitempty" json:"isSectionLeader,omitempty"`
	IsStudent                           bool               `url:"isStudent,omitempty" json:"isStudent,omitempty"`
	IsStudentAdministrator              bool               `url:"isStudentAdministrator,omitempty" json:"isStudentAdministrator,omitempty"`
	IsTeamLeader                        bool               `url:"isTeamLeader,omitempty" json:"isTeamLeader,omitempty"`
	Ldap                                bool               `url:"ldap,omitempty" json:"ldap,omitempty"`
	MagisterUUID                        string             `url:"magisterUUID,omitempty" json:"magisterUUID,omitempty"`
	MayLogin                            bool               `url:"mayLogin,omitempty" json:"mayLogin,omitempty"`
	ParentTeacherNightOfFamilyIDs       []int              `url:"parentTeacherNightOfFamilyIds,omitempty" json:"parentTeacherNightOfFamilyIds,omitempty"`
	ParentTeacherNightOfTeacherIDs      []int              `url:"parentTeacherNightOfTeacherIds,omitempty" json:"parentTeacherNightOfTeacherIds,omitempty"`
	PostalCode                          string             `url:"postalCode,omitempty" json:"postalCode,omitempty"`
	Prefix                              string             `url:"prefix,omitempty" json:"prefix,omitempty"`
	Roles                               []string           `url:"roles,omitempty" json:"roles,omitempty"`
	SchoolInSchoolYears                 []int              `url:"schoolInSchoolYears,omitempty" json:"schoolInSchoolYears,omitempty"`
	SomUUID                             string             `url:"somUUID,omitempty" json:"somUUID,omitempty"`
	Street                              string             `url:"street,omitempty" json:"street,omitempty"`
	StudentSchoolInSchoolYears          []int              `url:"studentSchoolInSchoolYears,omitempty" json:"studentSchoolInSchoolYears,omitempty"`
	SubjectSelectionSchoolInSchoolYears []int              `url:"subjectSelectionSchoolInSchoolYears,omitempty" json:"subjectSelectionSchoolInSchoolYears,omitempty"`
	TfaMandatory                        bool               `url:"tfaMandatory,omitempty" json:"tfaMandatory,omitempty"`
	TotpEnabled                         bool               `url:"totpEnabled,omitempty" json:"totpEnabled,omitempty"`
	UserKind                            string             `url:"userKind,omitempty" json:"userKind,omitempty"`
	UserPrincipalName                   string             `url:"userPrincipalName,omitempty" json:"userPrincipalName,omitempty"`
	Username                            string             `url:"username,omitempty" json:"username,omitempty"`
	Fields                              string             `url:"fields,omitempty" json:"fields,omitempty"`
	SchoolInSchoolYear                  SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	Archivable                          bool               `url:"archivable,omitempty" json:"archivable,omitempty"`
	FamilyMember                        bool               `url:"familyMember,omitempty" json:"familyMember,omitempty"`
	NoSchoolInSchoolYear                bool               `url:"noSchoolInSchoolYear,omitempty" json:"noSchoolInSchoolYear,omitempty"`
	Removable                           bool               `url:"removable,omitempty" json:"removable,omitempty"`
	SchoolYear                          int                `url:"schoolYear,omitempty" json:"schoolYear,omitempty"`
	Type                                string             `url:"type,omitempty" json:"type,omitempty"`
}

// UsersResponse is the envelope the Zermelo API wraps around every response.
// Data holds the decoded records and Status holds the reported status.
type UsersResponse struct {
	Response struct {
		Status int           `json:"status"`
		Data   []UsersParams `json:"data"`
	} `json:"response"`
}

// GetUsers calls /api/v3/users and returns the records that match param.
// A non nil *APIError is returned when the portal rejects the request.
func (c *Client) GetUsers(param UsersParams) ([]UsersParams, error) {
	var Output UsersResponse
	path, err := url.Parse("/api/v3/users")
	if err != nil {
		return []UsersParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []UsersParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []UsersParams{}, err
	}
	return Output.Response.Data, nil
}
