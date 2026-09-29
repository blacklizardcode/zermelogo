package zermelogo

import (
	"encoding/json"
	"strconv"
)

type SchoolInSchoolYear string

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

func (s SchoolInSchoolYear) MarshalJSON() ([]byte, error) {
	if _, err := strconv.ParseInt(string(s), 10, 64); err == nil {
		return []byte(s), nil
	}
	return json.Marshal(string(s))
}
