// Code generated by goa v3.14.1, DO NOT EDIT.
//
// auth HTTP client CLI support package
//
// Command:
// $ goa gen backend/design

package client

import (
	auth "backend/gen/auth"
	"encoding/json"
	"fmt"
)

// BuildLoginPayload builds the payload for the auth login endpoint from CLI
// flags.
func BuildLoginPayload(authLoginBody string) (*auth.LoginPayload, error) {
	var err error
	var body LoginRequestBody
	{
		err = json.Unmarshal([]byte(authLoginBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"password\": \"Nesciunt quas corrupti explicabo voluptate aut.\",\n      \"username\": \"Voluptas deleniti earum praesentium non dolorum laboriosam.\"\n   }'")
		}
	}
	v := &auth.LoginPayload{
		Username: body.Username,
		Password: body.Password,
	}

	return v, nil
}

// BuildRegisterPayload builds the payload for the auth register endpoint from
// CLI flags.
func BuildRegisterPayload(authRegisterBody string) (*auth.RegisterPayload, error) {
	var err error
	var body RegisterRequestBody
	{
		err = json.Unmarshal([]byte(authRegisterBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"password\": \"Sed vitae inventore laboriosam consequatur odio.\",\n      \"username\": \"Dolores vel aliquam a voluptas.\"\n   }'")
		}
	}
	v := &auth.RegisterPayload{
		Username: body.Username,
		Password: body.Password,
	}

	return v, nil
}

// BuildLogoutPayload builds the payload for the auth logout endpoint from CLI
// flags.
func BuildLogoutPayload(authLogoutBody string) (*auth.LogoutPayload, error) {
	var err error
	var body LogoutRequestBody
	{
		err = json.Unmarshal([]byte(authLogoutBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"token\": \"Voluptates nihil repellat impedit.\"\n   }'")
		}
	}
	v := &auth.LogoutPayload{
		Token: body.Token,
	}

	return v, nil
}
