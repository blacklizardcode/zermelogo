package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type ContractsParams struct {
	Employee                                         string             `url:"employee,omitempty" json:"employee,omitempty"`
	End                                              string             `url:"end,omitempty" json:"end,omitempty"`
	ID                                               int                `url:"id,omitempty" json:"id,omitempty"`
	SchoolInSchoolYear                               SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	Start                                            string             `url:"start,omitempty" json:"start,omitempty"`
	AvailabilityManually                             bool               `url:"availabilityManually,omitempty" json:"availabilityManually,omitempty"`
	AvailableDayParts                                int                `url:"availableDayParts,omitempty" json:"availableDayParts,omitempty"`
	AvailableDays                                    int                `url:"availableDays,omitempty" json:"availableDays,omitempty"`
	ClockHoursBalance                                int                `url:"clockHoursBalance,omitempty" json:"clockHoursBalance,omitempty"`
	ClockHoursBalancePreviousYear                    int                `url:"clockHoursBalancePreviousYear,omitempty" json:"clockHoursBalancePreviousYear,omitempty"`
	ClockHoursGeneralTasks                           int                `url:"clockHoursGeneralTasks,omitempty" json:"clockHoursGeneralTasks,omitempty"`
	ClockHoursGeneralTasksManually                   bool               `url:"clockHoursGeneralTasksManually,omitempty" json:"clockHoursGeneralTasksManually,omitempty"`
	ClockHoursGross                                  int                `url:"clockHoursGross,omitempty" json:"clockHoursGross,omitempty"`
	ClockHoursLessonRelief                           int                `url:"clockHoursLessonRelief,omitempty" json:"clockHoursLessonRelief,omitempty"`
	ClockHoursLessonReliefManually                   bool               `url:"clockHoursLessonReliefManually,omitempty" json:"clockHoursLessonReliefManually,omitempty"`
	ClockHoursLessons                                int                `url:"clockHoursLessons,omitempty" json:"clockHoursLessons,omitempty"`
	ClockHoursLessonsManually                        bool               `url:"clockHoursLessonsManually,omitempty" json:"clockHoursLessonsManually,omitempty"`
	ClockHoursLessonsMax                             int                `url:"clockHoursLessonsMax,omitempty" json:"clockHoursLessonsMax,omitempty"`
	ClockHoursNet                                    int                `url:"clockHoursNet,omitempty" json:"clockHoursNet,omitempty"`
	ClockHoursNetManually                            bool               `url:"clockHoursNetManually,omitempty" json:"clockHoursNetManually,omitempty"`
	ClockHoursProfessionalDevelopment                int                `url:"clockHoursProfessionalDevelopment,omitempty" json:"clockHoursProfessionalDevelopment,omitempty"`
	ClockHoursProfessionalDevelopmentManually        bool               `url:"clockHoursProfessionalDevelopmentManually,omitempty" json:"clockHoursProfessionalDevelopmentManually,omitempty"`
	ClockHoursTaskRelief                             int                `url:"clockHoursTaskRelief,omitempty" json:"clockHoursTaskRelief,omitempty"`
	ClockHoursTaskReliefManually                     bool               `url:"clockHoursTaskReliefManually,omitempty" json:"clockHoursTaskReliefManually,omitempty"`
	ClockHoursTasks                                  int                `url:"clockHoursTasks,omitempty" json:"clockHoursTasks,omitempty"`
	ClockHoursToAuxiliaryProjects                    int                `url:"clockHoursToAuxiliaryProjects,omitempty" json:"clockHoursToAuxiliaryProjects,omitempty"`
	CommencementSchool                               string             `url:"commencementSchool,omitempty" json:"commencementSchool,omitempty"`
	CommencementTeaching                             string             `url:"commencementTeaching,omitempty" json:"commencementTeaching,omitempty"`
	CompensationAbsenceHours                         int                `url:"compensationAbsenceHours,omitempty" json:"compensationAbsenceHours,omitempty"`
	DateOfBirth                                      string             `url:"dateOfBirth,omitempty" json:"dateOfBirth,omitempty"`
	DefaultFunctionCategory                          string             `url:"defaultFunctionCategory,omitempty" json:"defaultFunctionCategory,omitempty"`
	DevelopmentTime                                  int                `url:"developmentTime,omitempty" json:"developmentTime,omitempty"`
	DevelopmentTimeBudget                            int                `url:"developmentTimeBudget,omitempty" json:"developmentTimeBudget,omitempty"`
	DevelopmentTimeManually                          bool               `url:"developmentTimeManually,omitempty" json:"developmentTimeManually,omitempty"`
	DurationManually                                 bool               `url:"durationManually,omitempty" json:"durationManually,omitempty"`
	Email                                            string             `url:"email,omitempty" json:"email,omitempty"`
	EmployeeNumber                                   string             `url:"employeeNumber,omitempty" json:"employeeNumber,omitempty"`
	EmployerCostBasisMaintenance                     int                `url:"employerCostBasisMaintenance,omitempty" json:"employerCostBasisMaintenance,omitempty"`
	EmployerCostBasisPlanning                        int                `url:"employerCostBasisPlanning,omitempty" json:"employerCostBasisPlanning,omitempty"`
	EmployerCostLeaveReductionMaintenance            int                `url:"employerCostLeaveReductionMaintenance,omitempty" json:"employerCostLeaveReductionMaintenance,omitempty"`
	EmployerCostLeaveReductionPlanning               int                `url:"employerCostLeaveReductionPlanning,omitempty" json:"employerCostLeaveReductionPlanning,omitempty"`
	EmployerCostPostingBalanceMaintenance            int                `url:"employerCostPostingBalanceMaintenance,omitempty" json:"employerCostPostingBalanceMaintenance,omitempty"`
	EmployerCostPostingBalancePlanning               int                `url:"employerCostPostingBalancePlanning,omitempty" json:"employerCostPostingBalancePlanning,omitempty"`
	EmployerCostTotalMaintenance                     int                `url:"employerCostTotalMaintenance,omitempty" json:"employerCostTotalMaintenance,omitempty"`
	EmployerCostTotalPlanning                        int                `url:"employerCostTotalPlanning,omitempty" json:"employerCostTotalPlanning,omitempty"`
	EmployerCostsPercentage                          float64            `url:"employerCostsPercentage,omitempty" json:"employerCostsPercentage,omitempty"`
	ExtensionsDevelopmentTime                        int                `url:"extensionsDevelopmentTime,omitempty" json:"extensionsDevelopmentTime,omitempty"`
	ExtensionsGeneralTasks                           int                `url:"extensionsGeneralTasks,omitempty" json:"extensionsGeneralTasks,omitempty"`
	ExtensionsGeneralTasksManually                   bool               `url:"extensionsGeneralTasksManually,omitempty" json:"extensionsGeneralTasksManually,omitempty"`
	ExtensionsProfessionalDevelopment                int                `url:"extensionsProfessionalDevelopment,omitempty" json:"extensionsProfessionalDevelopment,omitempty"`
	ExtensionsProfessionalDevelopmentManually        bool               `url:"extensionsProfessionalDevelopmentManually,omitempty" json:"extensionsProfessionalDevelopmentManually,omitempty"`
	ExtraWork                                        int                `url:"extraWork,omitempty" json:"extraWork,omitempty"`
	ExtraWorkAllowance                               int                `url:"extraWorkAllowance,omitempty" json:"extraWorkAllowance,omitempty"`
	ExtraWorkAllowanceBasis                          int                `url:"extraWorkAllowanceBasis,omitempty" json:"extraWorkAllowanceBasis,omitempty"`
	FirstName                                        string             `url:"firstName,omitempty" json:"firstName,omitempty"`
	FteNet                                           float64            `url:"fteNet,omitempty" json:"fteNet,omitempty"`
	FtePermanent                                     float64            `url:"ftePermanent,omitempty" json:"ftePermanent,omitempty"`
	FteTemporary                                     float64            `url:"fteTemporary,omitempty" json:"fteTemporary,omitempty"`
	FteToAuxiliaryProjects                           float64            `url:"fteToAuxiliaryProjects,omitempty" json:"fteToAuxiliaryProjects,omitempty"`
	FteYearLeave                                     float64            `url:"fteYearLeave,omitempty" json:"fteYearLeave,omitempty"`
	FullMonthlySalary                                int                `url:"fullMonthlySalary,omitempty" json:"fullMonthlySalary,omitempty"`
	Gender                                           string             `url:"gender,omitempty" json:"gender,omitempty"`
	HasLessonReliefStarter                           bool               `url:"hasLessonReliefStarter,omitempty" json:"hasLessonReliefStarter,omitempty"`
	HourlySalary                                     float64            `url:"hourlySalary,omitempty" json:"hourlySalary,omitempty"`
	HourlySalaryManually                             bool               `url:"hourlySalaryManually,omitempty" json:"hourlySalaryManually,omitempty"`
	IsMainContract                                   bool               `url:"isMainContract,omitempty" json:"isMainContract,omitempty"`
	LastName                                         string             `url:"lastName,omitempty" json:"lastName,omitempty"`
	LeavesDevelopmentTime                            int                `url:"leavesDevelopmentTime,omitempty" json:"leavesDevelopmentTime,omitempty"`
	LeavesGeneralTasks                               int                `url:"leavesGeneralTasks,omitempty" json:"leavesGeneralTasks,omitempty"`
	LeavesGeneralTasksManually                       bool               `url:"leavesGeneralTasksManually,omitempty" json:"leavesGeneralTasksManually,omitempty"`
	LeavesProfessionalDevelopment                    int                `url:"leavesProfessionalDevelopment,omitempty" json:"leavesProfessionalDevelopment,omitempty"`
	LeavesProfessionalDevelopmentManually            bool               `url:"leavesProfessionalDevelopmentManually,omitempty" json:"leavesProfessionalDevelopmentManually,omitempty"`
	LessonHoursSectionAssignments                    float64            `url:"lessonHoursSectionAssignments,omitempty" json:"lessonHoursSectionAssignments,omitempty"`
	LessonHoursSurplusAllowance                      int                `url:"lessonHoursSurplusAllowance,omitempty" json:"lessonHoursSurplusAllowance,omitempty"`
	LessonHoursSurplusAllowanceManually              bool               `url:"lessonHoursSurplusAllowanceManually,omitempty" json:"lessonHoursSurplusAllowanceManually,omitempty"`
	LessonHoursSurplusAllowanceViaSectionAssignments int                `url:"lessonHoursSurplusAllowanceViaSectionAssignments,omitempty" json:"lessonHoursSurplusAllowanceViaSectionAssignments,omitempty"`
	LessonHoursTeachingAssignments                   float64            `url:"lessonHoursTeachingAssignments,omitempty" json:"lessonHoursTeachingAssignments,omitempty"`
	LessonReducingTasks                              int                `url:"lessonReducingTasks,omitempty" json:"lessonReducingTasks,omitempty"`
	LessonReplacement                                int                `url:"lessonReplacement,omitempty" json:"lessonReplacement,omitempty"`
	LessonSpace                                      int                `url:"lessonSpace,omitempty" json:"lessonSpace,omitempty"`
	LessonsMax                                       float64            `url:"lessonsMax,omitempty" json:"lessonsMax,omitempty"`
	LessonsMaxManually                               bool               `url:"lessonsMaxManually,omitempty" json:"lessonsMaxManually,omitempty"`
	LessonsToPlan                                    float64            `url:"lessonsToPlan,omitempty" json:"lessonsToPlan,omitempty"`
	MainBranchOfSchool                               int                `url:"mainBranchOfSchool,omitempty" json:"mainBranchOfSchool,omitempty"`
	MainContract                                     int                `url:"mainContract,omitempty" json:"mainContract,omitempty"`
	MaintenanceJobClockHoursBudget                   int                `url:"maintenanceJobClockHoursBudget,omitempty" json:"maintenanceJobClockHoursBudget,omitempty"`
	MaintenanceJobFteBudget                          float64            `url:"maintenanceJobFteBudget,omitempty" json:"maintenanceJobFteBudget,omitempty"`
	MaintenanceLeaveClockHoursBudget                 int                `url:"maintenanceLeaveClockHoursBudget,omitempty" json:"maintenanceLeaveClockHoursBudget,omitempty"`
	MaintenanceLeaveFteBudget                        float64            `url:"maintenanceLeaveFteBudget,omitempty" json:"maintenanceLeaveFteBudget,omitempty"`
	PersonalBudgetStatus                             string             `url:"personalBudgetStatus,omitempty" json:"personalBudgetStatus,omitempty"`
	PostedBranches                                   []int              `url:"postedBranches,omitempty" json:"postedBranches,omitempty"`
	Postings                                         []int              `url:"postings,omitempty" json:"postings,omitempty"`
	Prefix                                           string             `url:"prefix,omitempty" json:"prefix,omitempty"`
	PreviousSalaryScale                              int                `url:"previousSalaryScale,omitempty" json:"previousSalaryScale,omitempty"`
	PreviousSalaryStep                               int                `url:"previousSalaryStep,omitempty" json:"previousSalaryStep,omitempty"`
	PrivateComment                                   string             `url:"privateComment,omitempty" json:"privateComment,omitempty"`
	PublicComment                                    string             `url:"publicComment,omitempty" json:"publicComment,omitempty"`
	ReplacementBalance                               int                `url:"replacementBalance,omitempty" json:"replacementBalance,omitempty"`
	ReservationReplacement                           int                `url:"reservationReplacement,omitempty" json:"reservationReplacement,omitempty"`
	SalaryScale                                      int                `url:"salaryScale,omitempty" json:"salaryScale,omitempty"`
	SalaryStep                                       int                `url:"salaryStep,omitempty" json:"salaryStep,omitempty"`
	School                                           int                `url:"school,omitempty" json:"school,omitempty"`
	SchoolInSchoolYearName                           string             `url:"schoolInSchoolYearName,omitempty" json:"schoolInSchoolYearName,omitempty"`
	SchoolName                                       string             `url:"schoolName,omitempty" json:"schoolName,omitempty"`
	SchoolYear                                       int                `url:"schoolYear,omitempty" json:"schoolYear,omitempty"`
	SubstitutionQuota                                int                `url:"substitutionQuota,omitempty" json:"substitutionQuota,omitempty"`
	SubstitutionQuotaManually                        bool               `url:"substitutionQuotaManually,omitempty" json:"substitutionQuotaManually,omitempty"`
	TaskBalance                                      int                `url:"taskBalance,omitempty" json:"taskBalance,omitempty"`
	TaskLessonSpace                                  int                `url:"taskLessonSpace,omitempty" json:"taskLessonSpace,omitempty"`
	TaskReplacement                                  int                `url:"taskReplacement,omitempty" json:"taskReplacement,omitempty"`
	TaskSpace                                        int                `url:"taskSpace,omitempty" json:"taskSpace,omitempty"`
	TaskSurplus                                      int                `url:"taskSurplus,omitempty" json:"taskSurplus,omitempty"`
	TeacherTeam                                      int                `url:"teacherTeam,omitempty" json:"teacherTeam,omitempty"`
	Fields                                           string             `url:"fields,omitempty" json:"fields,omitempty"`
	BranchOfSchool                                   int                `url:"branchOfSchool,omitempty" json:"branchOfSchool,omitempty"`
	RelatedToEmployee                                string             `url:"relatedToEmployee,omitempty" json:"relatedToEmployee,omitempty"`
	SchoolFunction                                   string             `url:"schoolFunction,omitempty" json:"schoolFunction,omitempty"`
	Section                                          int                `url:"section,omitempty" json:"section,omitempty"`
	SectionOfBranch                                  int                `url:"sectionOfBranch,omitempty" json:"sectionOfBranch,omitempty"`
	Year                                             int                `url:"year,omitempty" json:"year,omitempty"`
}

type ContractsResponse struct {
	Response struct {
		Status int               `json:"status"`
		Data   []ContractsParams `json:"data"`
	} `json:"response"`
}

func (c *Client) GetContracts(param ContractsParams) ([]ContractsParams, error) {
	var Output ContractsResponse
	path, err := url.Parse("/api/v3/contracts")
	if err != nil {
		return []ContractsParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []ContractsParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []ContractsParams{}, err
	}
	return Output.Response.Data, nil
}
