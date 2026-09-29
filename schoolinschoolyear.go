package zermelogo

import (
	"encoding/json"
	"strconv"
)

// SchoolInSchoolYear identifies a project, which is a school in a single school
// year. As a query filter it accepts one id or a comma separated list of ids. As
// a response field the API returns a single numeric id, which is accepted too.
type SchoolInSchoolYear string

// UnmarshalJSON accepts both a number and a quoted string, and treats null as
// empty.
func (s *SchoolInSchoolYear) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		*s = ""
		return nil
	}

	if data[0] == '"' {
		var value string
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		*s = SchoolInSchoolYear(value)
		return nil
	}

	var value int64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SchoolInSchoolYear(strconv.FormatInt(value, 10))
	return nil
}

// MarshalJSON writes a single id as a number and a list of ids as a string.
func (s SchoolInSchoolYear) MarshalJSON() ([]byte, error) {
	if _, err := strconv.ParseInt(string(s), 10, 64); err == nil {
		return []byte(s), nil
	}
	return json.Marshal(string(s))
}
