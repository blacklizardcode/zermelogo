# zermelogo

A Go client for the Zermelo API.

## Requirements

- Go 1.24.4 or later
- A Zermelo school portal and an account with access to it

## Installation

```bash
go get github.com/blacklizardcode/zermelogo
```

The package name is `zermelogo`, so import it like this:

```go
import zermelogo "github.com/blacklizardcode/zermelogo"
```

## Documentation

All public types, methods, and functions are documented with Go doc comments. Read them
with the standard tool:

```bash
go doc github.com/blacklizardcode/zermelogo
```

To read the documentation for a single symbol:

```bash
go doc github.com/blacklizardcode/zermelogo Client
go doc github.com/blacklizardcode/zermelogo Client.GetHolidays
go doc github.com/blacklizardcode/zermelogo HolidaysParams
```

## Usage

Create a client with the base URL of your portal, exchange an authentication code for a
bearer token, and call the endpoint you need.

```go
package main

import (
	"fmt"
	"log"

	zermelogo "github.com/blacklizardcode/zermelogo"
)

func main() {
	client := zermelogo.New("https://example.zportal.nl")

	bearer, err := client.Auth().CodeToBearer(123456)
	if err != nil {
		log.Fatal(err)
	}
	client.Auth().SetBearerToken(bearer)

	holidays, err := client.GetHolidays(zermelogo.HolidaysParams{
		SchoolInSchoolYear: "452",
		Fields:             "id,name,start,end",
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, holiday := range holidays {
		fmt.Println(holiday.Name, holiday.Start, holiday.End)
	}
}
```

If you already have a bearer token, set it directly and skip the code exchange:

```go
client := zermelogo.New("https://example.zportal.nl")
client.Auth().SetBearerToken("your-bearer-token")
```

## Filtering

Every endpoint takes a typed parameter struct. Zero values are omitted from the request,
so you only set the fields you need. Use `Fields` to limit the returned columns.

```go
data, err := client.GetStudents(zermelogo.StudentsParams{
	SchoolInSchoolYear: "452,502",
	Fields:             "id,firstName,lastName",
})
```

`SchoolInSchoolYear` accepts a single project id or a comma separated list of project ids.
The type also decodes the numeric form that the API returns inside a response.

## Endpoints

| Method | Path |
| --- | --- |
| `GetAppointments` | `/api/v3/appointments` |
| `GetBranchesOfSchools` | `/api/v3/branchesofschools` |
| `GetContracts` | `/api/v3/contracts` |
| `GetCourses` | `/api/v3/courses` |
| `GetDepartmentsOfBranches` | `/api/v3/departmentsofbranches` |
| `GetEmployeeAbsences` | `/api/v3/employeeabsences` |
| `GetEmployees` | `/api/v3/employees` |
| `GetEnrollmentStatus` | `/api/v3/enrollmentstatus` |
| `GetFamilyMembers` | `/api/v3/familymembers` |
| `GetGroupsInDepartments` | `/api/v3/groupindepartments` |
| `GetHolidays` | `/api/v3/holidays` |
| `GetJobs` | `/api/v3/jobs` |
| `GetLiveSchedule` | `/api/v3/liveschedule` |
| `GetLocationsOfBranches` | `/api/v3/locationofbranches` |
| `GetParentTeacherNights` | `/api/v3/parentteachernights` |
| `GetParentToSchedulerRemarks` | `/api/v3/parenttoschedulerremarks` |
| `GetSchoolsInSchoolYears` | `/api/v3/schoolsinschoolyears` |
| `GetStudents` | `/api/v3/students` |
| `GetStudentsInDepartments` | `/api/v3/studentsindepartments` |
| `GetTalkDays` | `/api/v3/talkdays` |
| `GetTalkTimeBlockAbsenceParent` | `/api/v3/talktimeblockabsenceparent` |
| `GetTalkTimeBlocks` | `/api/v3/talktimeblocks` |
| `GetUsers` | `/api/v3/users` |

Run `go doc` on any method above for its parameter struct and fields.

## Errors

Failed requests return an `*APIError` that carries the HTTP status code and the message
reported by the portal. Use `errors.As` to inspect it.

```go
import (
	"errors"
	"fmt"
	"log"
	"net/http"

	zermelogo "github.com/blacklizardcode/zermelogo"
)

data, err := client.GetStudents(zermelogo.StudentsParams{SchoolInSchoolYear: "452"})
if err != nil {
	var apiErr *zermelogo.APIError
	if errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusForbidden {
		log.Println("no permission for this endpoint:", apiErr.Message)
		return
	}
	log.Fatal(err)
}
fmt.Println(data)
```

A status of 403 usually means the account lacks the security rule for that endpoint, so
check the permissions granted in Zermelo before changing the request.

## Live smoke test

The `test` directory contains a small program that calls every endpoint and reports the
result. Pass a base URL plus either a bearer token or an authentication code, and
optionally a `schoolInSchoolYear` value.

```bash
go run ./test https://example.zportal.nl "" "$BEARER_TOKEN" "452,502"
```

## License

No license file is included in this repository.
