package zermelogo

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

// SchoolsInSchoolYearsParams holds the query filters for GetSchoolsInSchoolYears and the fields of the
// records it returns. Zermelo uses the same field names in requests and
// responses, so one struct covers both. Zero values are omitted, and Fields
// selects the columns to return.
type SchoolsInSchoolYearsParams struct {
	Archived               bool               `url:"archived,omitempty" json:"archived,omitempty"`
	ArchivedTime           int64              `url:"archivedTime,omitempty" json:"archivedTime,omitempty"`
	BranchOfSchools        []int              `url:"branchOfSchools,omitempty" json:"branchOfSchools,omitempty"`
	DesktopProjectSettings string             `url:"desktopProjectSettings,omitempty" json:"desktopProjectSettings,omitempty"`
	DesktopStorageOnline   bool               `url:"desktopStorageOnline,omitempty" json:"desktopStorageOnline,omitempty"`
	ID                     int                `url:"id,omitempty" json:"id,omitempty"`
	Name                   string             `url:"name,omitempty" json:"name,omitempty"`
	Project                string             `url:"project,omitempty" json:"project,omitempty"`
	ProjectName            string             `url:"projectName,omitempty" json:"projectName,omitempty"`
	School                 int                `url:"school,omitempty" json:"school,omitempty"`
	SchoolHrmsCode         string             `url:"schoolHrmsCode,omitempty" json:"schoolHrmsCode,omitempty"`
	SchoolName             string             `url:"schoolName,omitempty" json:"schoolName,omitempty"`
	Year                   int                `url:"year,omitempty" json:"year,omitempty"`
	Fields                 string             `url:"fields,omitempty" json:"fields,omitempty"`
	SchoolInSchoolYear     SchoolInSchoolYear `url:"schoolInSchoolYear,omitempty" json:"schoolInSchoolYear,omitempty"`
	Archivable             bool               `url:"archivable,omitempty" json:"archivable,omitempty"`
	AtlasNexusEnabled      bool               `url:"atlasNexusEnabled,omitempty" json:"atlasNexusEnabled,omitempty"`
	Branches               string             `url:"branches,omitempty" json:"branches,omitempty"`
	Removable              bool               `url:"removable,omitempty" json:"removable,omitempty"`
	SchoolYear             int                `url:"schoolYear,omitempty" json:"schoolYear,omitempty"`
	UseHrms                bool               `url:"useHrms,omitempty" json:"useHrms,omitempty"`
}

// SchoolsInSchoolYearsResponse is the envelope the Zermelo API wraps around every response.
// Data holds the decoded records and Status holds the reported status.
type SchoolsInSchoolYearsResponse struct {
	Response struct {
		Status int                          `json:"status"`
		Data   []SchoolsInSchoolYearsParams `json:"data"`
	} `json:"response"`
}

// GetSchoolsInSchoolYears calls /api/v3/schoolsinschoolyears and returns the records that match param.
// A non nil *APIError is returned when the portal rejects the request.
func (c *Client) GetSchoolsInSchoolYears(param SchoolsInSchoolYearsParams) ([]SchoolsInSchoolYearsParams, error) {
	var Output SchoolsInSchoolYearsResponse
	path, err := url.Parse("/api/v3/schoolsinschoolyears")
	if err != nil {
		return []SchoolsInSchoolYearsParams{}, err
	}

	values, err := query.Values(param)
	if err != nil {
		return []SchoolsInSchoolYearsParams{}, err
	}
	path.RawQuery = values.Encode()

	err = c.do(http.MethodGet, path.String(), nil, &Output)
	if err != nil {
		return []SchoolsInSchoolYearsParams{}, err
	}
	return Output.Response.Data, nil
}
