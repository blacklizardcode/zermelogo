package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type StudentsInDepartmentsParams struct {
	DeanApprovesSelection           bool               `url:"deanApprovesSelection,omitempty" json:"deanApprovesSelection,omitempty"`
	DepartmentOfBranch              int                `url:"departmentOfBranch,omitempty" json:"departmentOfBranch,omitempty"`
	DepartmentOfBranchCode          string             `url:"departmentOfBranchCode,omitempty" json:"departmentOfBranchCode,omitempty"`
	EducationalTrack                int                `url:"educationalTrack,omitempty" json:"educationalTrack,omitempty"`
	Educations                      []int              `url:"educations,omitempty" json:"educations,omitempty"`
	Email                           string             `url:"email,omitempty" json:"email,omitempty"`
	FirstName                       string             `url:"firstName,omitempty" json:"firstName,omitempty"`
	FullName                        string             `url:"fullName,omitempty" json:"fullName,omitempty"`
	Gender                          string             `url:"gender,omitempty" json:"gender,omitempty"`
	GroupInDepartments              []int              `url:"groupInDepartments,omitempty" json:"groupInDepartments,omitempty"`
	ID                              int                `url:"id,omitempty" json:"id,omitempty"`
	LastModified                    int64              `url:"lastModified,omitempty" json:"lastModified,omitempty"`
	LastName                        string             `url:"lastName,omitempty" json:"lastName,omitempty"`
	Lwoo                            bool               `url:"lwoo,omitempty" json:"lwoo,omitempty"`
	MainGroup                       int                `url:"mainGroup,omitempty" json:"mainGroup,omitempty"`
	MainGroupName                   string             `url:"mainGroupName,omitempty" json:"mainGroupName,omitempty"`
	MainGroupPreviousYear           int                `url:"mainGroupPreviousYear,omitempty" json:"mainGroupPreviousYear,omitempty"`
	MainGroupPreviousYearName       string             `url:"mainGroupPreviousYearName,omitempty" json:"mainGroupPreviousYearName,omitempty"`
	MentorGroup                     int                `url:"mentorGroup,omitempty" json:"mentorGroup,omitempty"`
	MentorGroupPreviousYear         int                `url:"mentorGroupPreviousYear,omitempty" json:"mentorGroupPreviousYear,omitempty"`
	OutflowComment                  string             `url:"outflowComment,omitempty" json:"outflowComment,omitempty"`
	OutflowPercentage               int                `url:"outflowPercentage,omitempty" json:"outflowPercentage,omitempty"`
	ParentApprovalTime              int64              `url:"parentApprovalTime,omitempty" json:"parentApprovalTime,omitempty"`
	ParentApprovesSelection         bool               `url:"parentApprovesSelection,omitempty" json:"parentApprovesSelection,omitempty"`
	ParticipationWeight             float64            `url:"participationWeight,omitempty" json:"participationWeight,omitempty"`
	Prefix                          string             `url:"prefix,omitempty" json:"prefix,omitempty"`
	PreviousMainGroup               string             `url:"previousMainGroup,omitempty" json:"previousMainGroup,omitempty"`
	Profiles                        []int              `url:"profiles,omitempty" json:"profiles,omitempty"`
	PrognosisComment                string             `url:"prognosisComment,omitempty" json:"prognosisComment,omitempty"`
	SchoolInSchoolYear              SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	SchoolPreviousYear              string             `url:"schoolPreviousYear,omitempty" json:"schoolPreviousYear,omitempty"`
	Student                         string             `url:"student,omitempty" json:"student,omitempty"`
	StudentApprovesSelection        bool               `url:"studentApprovesSelection,omitempty" json:"studentApprovesSelection,omitempty"`
	StudentHasEditedSelection       bool               `url:"studentHasEditedSelection,omitempty" json:"studentHasEditedSelection,omitempty"`
	StudentInDepartmentPreviousYear int                `url:"studentInDepartmentPreviousYear,omitempty" json:"studentInDepartmentPreviousYear,omitempty"`
	Fields                          string             `url:"fields,omitempty" json:"fields,omitempty"`
	Archived                        bool               `url:"archived,omitempty" json:"archived,omitempty"`
	BranchOfSchool                  int                `url:"branchOfSchool,omitempty" json:"branchOfSchool,omitempty"`
	DepartmentOfBranchPreviousYear  int                `url:"departmentOfBranchPreviousYear,omitempty" json:"departmentOfBranchPreviousYear,omitempty"`
	FamilyMember                    string             `url:"familyMember,omitempty" json:"familyMember,omitempty"`
	GroupInDepartmentPreviousYear   int                `url:"groupInDepartmentPreviousYear,omitempty" json:"groupInDepartmentPreviousYear,omitempty"`
	Inflow                          bool               `url:"inflow,omitempty" json:"inflow,omitempty"`
	Outflow                         bool               `url:"outflow,omitempty" json:"outflow,omitempty"`
	RelatedToDepartment             int                `url:"relatedToDepartment,omitempty" json:"relatedToDepartment,omitempty"`
	SchoolYear                      int                `url:"schoolYear,omitempty" json:"schoolYear,omitempty"`
	Teacher                         string             `url:"teacher,omitempty" json:"teacher,omitempty"`
	TeacherTeam                     int                `url:"teacherTeam,omitempty" json:"teacherTeam,omitempty"`
	VisibleToStudent                bool               `url:"visibleToStudent,omitempty" json:"visibleToStudent,omitempty"`
}

type StudentsInDepartmentsResponse struct {
	Response struct {
		Status int                           `json:"status"`
		Data   []StudentsInDepartmentsParams `json:"data"`
	} `json:"response"`
}

func (c *Client) GetStudentsInDepartments(param StudentsInDepartmentsParams) ([]StudentsInDepartmentsParams, error) {
	var Output StudentsInDepartmentsResponse
	path, err := url.Parse("/api/v3/studentsindepartments")
	if err != nil {
		return []StudentsInDepartmentsParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []StudentsInDepartmentsParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []StudentsInDepartmentsParams{}, err
	}
	return Output.Response.Data, nil
}
