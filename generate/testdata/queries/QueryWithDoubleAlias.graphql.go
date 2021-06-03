package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// QueryWithDoubleAliasResponse is returned by QueryWithDoubleAlias on success.
type QueryWithDoubleAliasResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User QueryWithDoubleAliasUser `json:"user"`
}

// QueryWithDoubleAliasUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type QueryWithDoubleAliasUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	ID testutil.ID `json:"ID"`
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	AlsoID testutil.ID `json:"AlsoID"`
}

func QueryWithDoubleAlias(
	client graphql.Client,
) (*QueryWithDoubleAliasResponse, error) {
	var retval QueryWithDoubleAliasResponse
	err := client.MakeRequest(
		nil,
		"QueryWithDoubleAlias",
		`
query QueryWithDoubleAlias {
	user {
		ID: id
		AlsoID: id
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}
