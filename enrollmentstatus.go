package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type EnrollmentStatusParams struct {
	ActivityEnrollment          int                `url:"activityEnrollment,omitempty" json:"activityEnrollment,omitempty"`
	ConcurrentPlannedAttendance int                `url:"concurrentPlannedAttendance,omitempty" json:"concurrentPlannedAttendance,omitempty"`
	DepartmentOfBranch          int                `url:"departmentOfBranch,omitempty" json:"departmentOfBranch,omitempty"`
	EnrollmentValidationType    string             `url:"enrollmentValidationType,omitempty" json:"enrollmentValidationType,omitempty"`
	ExamEnrollment              int                `url:"examEnrollment,omitempty" json:"examEnrollment,omitempty"`
	LessonEnrollment            int                `url:"lessonEnrollment,omitempty" json:"lessonEnrollment,omitempty"`
	MaximalLessonEnrollment     int                `url:"maximalLessonEnrollment,omitempty" json:"maximalLessonEnrollment,omitempty"`
	MaximalLessons              int                `url:"maximalLessons,omitempty" json:"maximalLessons,omitempty"`
	MentorGroup                 int                `url:"mentorGroup,omitempty" json:"mentorGroup,omitempty"`
	MinimalLessonEnrollment     int                `url:"minimalLessonEnrollment,omitempty" json:"minimalLessonEnrollment,omitempty"`
	MinimalLessons              int                `url:"minimalLessons,omitempty" json:"minimalLessons,omitempty"`
	Name                        string             `url:"name,omitempty" json:"name,omitempty"`
	PlannedActivityAttendance   int                `url:"plannedActivityAttendance,omitempty" json:"plannedActivityAttendance,omitempty"`
	PlannedAttendance           int                `url:"plannedAttendance,omitempty" json:"plannedAttendance,omitempty"`
	PlannedExamAttendance       int                `url:"plannedExamAttendance,omitempty" json:"plannedExamAttendance,omitempty"`
	PlannedLessonAttendance     int                `url:"plannedLessonAttendance,omitempty" json:"plannedLessonAttendance,omitempty"`
	SchoolInSchoolYear          SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	Status                      []string           `url:"status,omitempty" json:"status,omitempty"`
	Student                     string             `url:"student,omitempty" json:"student,omitempty"`
	StudentInDepartment         int                `url:"studentInDepartment,omitempty" json:"studentInDepartment,omitempty"`
	Week                        string             `url:"week,omitempty" json:"week,omitempty"`
	Fields                      string             `url:"fields,omitempty" json:"fields,omitempty"`
}

type EnrollmentStatusResponse struct {
	Response struct {
		Status int                      `json:"status"`
		Data   []EnrollmentStatusParams `json:"data"`
	} `json:"response"`
}

func (c *Client) GetEnrollmentStatus(param EnrollmentStatusParams) ([]EnrollmentStatusParams, error) {
	var Output EnrollmentStatusResponse
	path, err := url.Parse("/api/v3/enrollmentstatus")
	if err != nil {
		return []EnrollmentStatusParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []EnrollmentStatusParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []EnrollmentStatusParams{}, err
	}
	return Output.Response.Data, nil
}
