// Code generated by goa v3.14.1, DO NOT EDIT.
//
// todo HTTP client CLI support package
//
// Command:
// $ goa gen backend/design

package cli

import (
	authc "backend/gen/http/auth/client"
	todoc "backend/gen/http/todo/client"
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `auth (login|register|logout)
todo (list|show)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` auth login --body '{
      "password": "Eaque nihil dolore corrupti odit enim.",
      "username": "Molestias assumenda ea quia rerum enim."
   }'` + "\n" +
		os.Args[0] + ` todo list --limit 2423154054 --offset 3269417336` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		authFlags = flag.NewFlagSet("auth", flag.ContinueOnError)

		authLoginFlags    = flag.NewFlagSet("login", flag.ExitOnError)
		authLoginBodyFlag = authLoginFlags.String("body", "REQUIRED", "")

		authRegisterFlags    = flag.NewFlagSet("register", flag.ExitOnError)
		authRegisterBodyFlag = authRegisterFlags.String("body", "REQUIRED", "")

		authLogoutFlags    = flag.NewFlagSet("logout", flag.ExitOnError)
		authLogoutBodyFlag = authLogoutFlags.String("body", "REQUIRED", "")

		todoFlags = flag.NewFlagSet("todo", flag.ContinueOnError)

		todoListFlags      = flag.NewFlagSet("list", flag.ExitOnError)
		todoListLimitFlag  = todoListFlags.String("limit", "20", "")
		todoListOffsetFlag = todoListFlags.String("offset", "", "")

		todoShowFlags  = flag.NewFlagSet("show", flag.ExitOnError)
		todoShowIDFlag = todoShowFlags.String("id", "REQUIRED", "ID of todo to show")
	)
	authFlags.Usage = authUsage
	authLoginFlags.Usage = authLoginUsage
	authRegisterFlags.Usage = authRegisterUsage
	authLogoutFlags.Usage = authLogoutUsage

	todoFlags.Usage = todoUsage
	todoListFlags.Usage = todoListUsage
	todoShowFlags.Usage = todoShowUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "auth":
			svcf = authFlags
		case "todo":
			svcf = todoFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "auth":
			switch epn {
			case "login":
				epf = authLoginFlags

			case "register":
				epf = authRegisterFlags

			case "logout":
				epf = authLogoutFlags

			}

		case "todo":
			switch epn {
			case "list":
				epf = todoListFlags

			case "show":
				epf = todoShowFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "auth":
			c := authc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "login":
				endpoint = c.Login()
				data, err = authc.BuildLoginPayload(*authLoginBodyFlag)
			case "register":
				endpoint = c.Register()
				data, err = authc.BuildRegisterPayload(*authRegisterBodyFlag)
			case "logout":
				endpoint = c.Logout()
				data, err = authc.BuildLogoutPayload(*authLogoutBodyFlag)
			}
		case "todo":
			c := todoc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data, err = todoc.BuildListPayload(*todoListLimitFlag, *todoListOffsetFlag)
			case "show":
				endpoint = c.Show()
				data, err = todoc.BuildShowPayload(*todoShowIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// authUsage displays the usage of the auth command and its subcommands.
func authUsage() {
	fmt.Fprintf(os.Stderr, `The auth service manages authentication
Usage:
    %[1]s [globalflags] auth COMMAND [flags]

COMMAND:
    login: Login to the system
    register: Register a new user
    logout: Logout of the system

Additional help:
    %[1]s auth COMMAND --help
`, os.Args[0])
}
func authLoginUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] auth login -body JSON

Login to the system
    -body JSON: 

Example:
    %[1]s auth login --body '{
      "password": "Eaque nihil dolore corrupti odit enim.",
      "username": "Molestias assumenda ea quia rerum enim."
   }'
`, os.Args[0])
}

func authRegisterUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] auth register -body JSON

Register a new user
    -body JSON: 

Example:
    %[1]s auth register --body '{
      "password": "Iure dolore.",
      "username": "Officia vero fugiat cumque."
   }'
`, os.Args[0])
}

func authLogoutUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] auth logout -body JSON

Logout of the system
    -body JSON: 

Example:
    %[1]s auth logout --body '{
      "token": "Quia et est expedita."
   }'
`, os.Args[0])
}

// todoUsage displays the usage of the todo command and its subcommands.
func todoUsage() {
	fmt.Fprintf(os.Stderr, `The todo service manages todo lists
Usage:
    %[1]s [globalflags] todo COMMAND [flags]

COMMAND:
    list: List all todos
    show: Show a todo

Additional help:
    %[1]s todo COMMAND --help
`, os.Args[0])
}
func todoListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todo list -limit UINT32 -offset UINT32

List all todos
    -limit UINT32: 
    -offset UINT32: 

Example:
    %[1]s todo list --limit 2423154054 --offset 3269417336
`, os.Args[0])
}

func todoShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todo show -id UINT32

Show a todo
    -id UINT32: ID of todo to show

Example:
    %[1]s todo show --id 1455715248
`, os.Args[0])
}
