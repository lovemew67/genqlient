package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// OmitEmptyQueryResponse is returned by OmitEmptyQuery on success.
type OmitEmptyQueryResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User         OmitEmptyQueryUser      `json:"user"`
	Users        OmitEmptyQueryUsersUser `json:"users"`
	MaybeConvert time.Time               `json:"maybeConvert"`
	Convert2     time.Time               `json:"convert2"`
}

// OmitEmptyQueryUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type OmitEmptyQueryUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id testutil.ID `json:"id"`
}

// OmitEmptyQueryUsersUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type OmitEmptyQueryUsersUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id testutil.ID `json:"id"`
}

// Role is a type a user may have.
type Role string

const (
	// What is a student?
	//
	// A student is primarily a person enrolled in a school or other educational institution and who is under learning with goals of acquiring knowledge, developing professions and achieving employment at desired field. In the broader sense, a student is anyone who applies themselves to the intensive intellectual engagement with some matter necessary to master it as part of some practical affair in which such mastery is basic or decisive.
	//
	// (from [Wikipedia](https://en.wikipedia.org/wiki/Student))
	RoleStudent Role = "STUDENT"
	// Teacher is a teacher, who teaches the students.
	RoleTeacher Role = "TEACHER"
)

// UserQueryInput is the argument to Query.users.
//
// Ideally this would support anything and everything!
// Or maybe ideally it wouldn't.
// Really I'm just talking to make this documentation longer.
type UserQueryInput struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	// id looks the user up by ID.  It's a great way to look up users.
	Id    testutil.ID `json:"id"`
	Role  Role        `json:"role"`
	Names []string    `json:"names"`
}

func OmitEmptyQuery(
	client graphql.Client,
	query UserQueryInput,
	queries []UserQueryInput,
	dt time.Time,
	tz string,
	tzNoOmitEmpty string,
) (*OmitEmptyQueryResponse, error) {
	variables := map[string]interface{}{
		"tzNoOmitEmpty": tzNoOmitEmpty,
	}

	var zero_query UserQueryInput
	if query != zero_query {
		variables["query"] = query
	}

	if len(queries) > 0 {
		variables["queries"] = queries
	}

	var zero_dt time.Time
	if dt != zero_dt {
		variables["dt"] = dt
	}

	var zero_tz string
	if tz != zero_tz {
		variables["tz"] = tz
	}

	var retval OmitEmptyQueryResponse
	err := client.MakeRequest(
		nil,
		"OmitEmptyQuery",
		`
query OmitEmptyQuery ($query: UserQueryInput, $queries: [UserQueryInput], $dt: DateTime, $tz: String, $tzNoOmitEmpty: String) {
	user(query: $query) {
		id
	}
	users(query: $queries) {
		id
	}
	maybeConvert(dt: $dt, tz: $tz)
	convert2: maybeConvert(dt: $dt, tz: $tzNoOmitEmpty)
}
`,
		&retval,
		variables,
	)
	return &retval, err
}
