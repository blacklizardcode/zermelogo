package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type DepartmentsOfBranchesParams struct {
	BranchOfSchool                        int                `url:"branchOfSchool,omitempty" json:"branchOfSchool,omitempty"`
	BranchOfSchoolCode                    string             `url:"branchOfSchoolCode,omitempty" json:"branchOfSchoolCode,omitempty"`
	ClockHoursPerLesson                   float64            `url:"clockHoursPerLesson,omitempty" json:"clockHoursPerLesson,omitempty"`
	ClockHoursPerLessonManually           bool               `url:"clockHoursPerLessonManually,omitempty" json:"clockHoursPerLessonManually,omitempty"`
	Code                                  string             `url:"code,omitempty" json:"code,omitempty"`
	DefaultEnrollmentValidationType       string             `url:"defaultEnrollmentValidationType,omitempty" json:"defaultEnrollmentValidationType,omitempty"`
	DefaultMaximalLessonEnrollmentPerWeek int                `url:"defaultMaximalLessonEnrollmentPerWeek,omitempty" json:"defaultMaximalLessonEnrollmentPerWeek,omitempty"`
	DefaultMaximalLessonsPerWeek          int                `url:"defaultMaximalLessonsPerWeek,omitempty" json:"defaultMaximalLessonsPerWeek,omitempty"`
	DefaultMinimalLessonEnrollmentPerWeek int                `url:"defaultMinimalLessonEnrollmentPerWeek,omitempty" json:"defaultMinimalLessonEnrollmentPerWeek,omitempty"`
	DefaultMinimalLessonsPerWeek          int                `url:"defaultMinimalLessonsPerWeek,omitempty" json:"defaultMinimalLessonsPerWeek,omitempty"`
	EducationType                         string             `url:"educationType,omitempty" json:"educationType,omitempty"`
	Educations                            []int              `url:"educations,omitempty" json:"educations,omitempty"`
	ExcludedSegments                      []int              `url:"excludedSegments,omitempty" json:"excludedSegments,omitempty"`
	ExpectedInflow                        float64            `url:"expectedInflow,omitempty" json:"expectedInflow,omitempty"`
	ExpectedInflowManually                bool               `url:"expectedInflowManually,omitempty" json:"expectedInflowManually,omitempty"`
	ExpectedMainGroupCount                int                `url:"expectedMainGroupCount,omitempty" json:"expectedMainGroupCount,omitempty"`
	ExpectedMainGroupCountManually        bool               `url:"expectedMainGroupCountManually,omitempty" json:"expectedMainGroupCountManually,omitempty"`
	ExpectedStudentCount                  float64            `url:"expectedStudentCount,omitempty" json:"expectedStudentCount,omitempty"`
	ExpectedStudentCountManually          bool               `url:"expectedStudentCountManually,omitempty" json:"expectedStudentCountManually,omitempty"`
	ID                                    int                `url:"id,omitempty" json:"id,omitempty"`
	MainGroupAuthority                    string             `url:"mainGroupAuthority,omitempty" json:"mainGroupAuthority,omitempty"`
	MaxGroupSize                          int                `url:"maxGroupSize,omitempty" json:"maxGroupSize,omitempty"`
	MaxGroupSizeManually                  bool               `url:"maxGroupSizeManually,omitempty" json:"maxGroupSizeManually,omitempty"`
	MaximalExtraEnrollmentPerSubject      int                `url:"maximalExtraEnrollmentPerSubject,omitempty" json:"maximalExtraEnrollmentPerSubject,omitempty"`
	MinimalExtraEnrollmentPerSubject      int                `url:"minimalExtraEnrollmentPerSubject,omitempty" json:"minimalExtraEnrollmentPerSubject,omitempty"`
	ParticipationThreshold                float64            `url:"participationThreshold,omitempty" json:"participationThreshold,omitempty"`
	Profiles                              []int              `url:"profiles,omitempty" json:"profiles,omitempty"`
	PrognosticInflow                      float64            `url:"prognosticInflow,omitempty" json:"prognosticInflow,omitempty"`
	PrognosticStudentCount                float64            `url:"prognosticStudentCount,omitempty" json:"prognosticStudentCount,omitempty"`
	SchoolInSchoolYear                    SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	SchoolInSchoolYearID                  int                `url:"schoolInSchoolYearId,omitempty" json:"schoolInSchoolYearId,omitempty"`
	SchoolInSchoolYearName                string             `url:"schoolInSchoolYearName,omitempty" json:"schoolInSchoolYearName,omitempty"`
	ScroungeSegments                      []int              `url:"scroungeSegments,omitempty" json:"scroungeSegments,omitempty"`
	StudentCount                          int                `url:"studentCount,omitempty" json:"studentCount,omitempty"`
	StudentCountCorrection                float64            `url:"studentCountCorrection,omitempty" json:"studentCountCorrection,omitempty"`
	StudentCountCorrectionManually        bool               `url:"studentCountCorrectionManually,omitempty" json:"studentCountCorrectionManually,omitempty"`
	StudentsMayViewSubjectSelections      bool               `url:"studentsMayViewSubjectSelections,omitempty" json:"studentsMayViewSubjectSelections,omitempty"`
	SubjectSelectionDeadline              string             `url:"subjectSelectionDeadline,omitempty" json:"subjectSelectionDeadline,omitempty"`
	TeacherTeams                          []int              `url:"teacherTeams,omitempty" json:"teacherTeams,omitempty"`
	TeachingLevel                         string             `url:"teachingLevel,omitempty" json:"teachingLevel,omitempty"`
	WeekTimeTable                         int                `url:"weekTimeTable,omitempty" json:"weekTimeTable,omitempty"`
	YearOfEducation                       int                `url:"yearOfEducation,omitempty" json:"yearOfEducation,omitempty"`
	Fields                                string             `url:"fields,omitempty" json:"fields,omitempty"`
	SchoolYear                            int                `url:"schoolYear,omitempty" json:"schoolYear,omitempty"`
	TeacherTeam                           int                `url:"teacherTeam,omitempty" json:"teacherTeam,omitempty"`
}

type DepartmentsOfBranchesResponse struct {
	Response struct {
		Status int                           `json:"status"`
		Data   []DepartmentsOfBranchesParams `json:"data"`
	} `json:"response"`
}

func (c *Client) GetDepartmentsOfBranches(param DepartmentsOfBranchesParams) ([]DepartmentsOfBranchesParams, error) {
	var Output DepartmentsOfBranchesResponse
	path, err := url.Parse("/api/v3/departmentsofbranches")
	if err != nil {
		return []DepartmentsOfBranchesParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []DepartmentsOfBranchesParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []DepartmentsOfBranchesParams{}, err
	}
	return Output.Response.Data, nil
}
