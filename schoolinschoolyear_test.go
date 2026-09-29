package zermelogo

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-querystring/query"
)

func TestSchoolInSchoolYearUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected SchoolInSchoolYear
	}{
		{"number", `{"schoolInSchoolYear":452}`, "452"},
		{"list", `{"schoolInSchoolYear":"452,502"}`, "452,502"},
		{"null", `{"schoolInSchoolYear":null}`, ""},
		{"absent", `{}`, ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var params HolidaysParams
			if err := json.Unmarshal([]byte(test.input), &params); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if params.SchoolInSchoolYear != test.expected {
				t.Errorf("got %q, want %q", params.SchoolInSchoolYear, test.expected)
			}
		})
	}
}

func TestSchoolInSchoolYearMarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    SchoolInSchoolYear
		expected string
	}{
		{"number", "452", `452`},
		{"list", "452,502", `"452,502"`},
		{"empty", "", `""`},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			encoded, err := json.Marshal(test.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if string(encoded) != test.expected {
				t.Errorf("got %s, want %s", encoded, test.expected)
			}
		})
	}
}

func TestSchoolInSchoolYearQueryEncoding(t *testing.T) {
	values, err := query.Values(HolidaysParams{SchoolInSchoolYear: "452,502", Fields: "id"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := values.Get("schoolInSchoolYear"); got != "452,502" {
		t.Errorf("got %q, want %q", got, "452,502")
	}
}

func TestSchoolInSchoolYearOmittedWhenEmpty(t *testing.T) {
	values, err := query.Values(HolidaysParams{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := values["schoolInSchoolYear"]; ok {
		t.Error("schoolInSchoolYear should be omitted when empty")
	}
}

func TestParseAPIErrorReadsNestedResponse(t *testing.T) {
	body := `{"response":{"status":403,"message":"Er zijn geen beveiligingsregels gevonden die uw verzoek toestaan.","details":"SecurityRuleException: denied","data":[]}}`
	resp := &http.Response{
		StatusCode: 403,
		Body:       io.NopCloser(strings.NewReader(body)),
	}

	apiErr, ok := parseAPIError(resp).(*APIError)
	if !ok {
		t.Fatal("expected *APIError")
	}
	if apiErr.StatusCode != 403 {
		t.Errorf("got status %d, want 403", apiErr.StatusCode)
	}
	if apiErr.Message != "Er zijn geen beveiligingsregels gevonden die uw verzoek toestaan." {
		t.Errorf("got message %q", apiErr.Message)
	}
}

func TestParseAPIErrorFallsBackToHTTPStatus(t *testing.T) {
	resp := &http.Response{
		StatusCode: 500,
		Body:       io.NopCloser(strings.NewReader(`not json`)),
	}

	apiErr, ok := parseAPIError(resp).(*APIError)
	if !ok {
		t.Fatal("expected *APIError")
	}
	if apiErr.StatusCode != 500 {
		t.Errorf("got status %d, want 500", apiErr.StatusCode)
	}
}
