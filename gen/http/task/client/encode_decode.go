// Code generated by goa v3.14.1, DO NOT EDIT.
//
// task HTTP client encoders and decoders
//
// Command:
// $ goa gen backend/design

package client

import (
	task "backend/gen/task"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	goahttp "goa.design/goa/v3/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "task" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListTaskPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("task", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the task list
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*task.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("task", "list", "*task.ListPayload", v)
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		values := req.URL.Query()
		values.Add("limit", fmt.Sprintf("%v", p.Limit))
		values.Add("offset", fmt.Sprintf("%v", p.Offset))
		values.Add("name", p.Name)
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeListResponse returns a decoder for responses returned by the task list
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
// DecodeListResponse may return the following errors:
//   - "token_verification_failed" (type *task.AuthFailed): http.StatusUnauthorized
//   - error: internal error
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("task", "list", err)
			}
			err = ValidateListResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("task", "list", err)
			}
			res := NewListResultOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body ListTokenVerificationFailedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("task", "list", err)
			}
			err = ValidateListTokenVerificationFailedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("task", "list", err)
			}
			return nil, NewListTokenVerificationFailed(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("task", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildShowRequest instantiates a HTTP request object with method and path set
// to call the "task" service "show" endpoint
func (c *Client) BuildShowRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id uint32
	)
	{
		p, ok := v.(*task.ShowPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("task", "show", "*task.ShowPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ShowTaskPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("task", "show", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeShowRequest returns an encoder for requests sent to the task show
// server.
func EncodeShowRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*task.ShowPayload)
		if !ok {
			return goahttp.ErrInvalidType("task", "show", "*task.ShowPayload", v)
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeShowResponse returns a decoder for responses returned by the task show
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
// DecodeShowResponse may return the following errors:
//   - "token_verification_failed" (type *task.AuthFailed): http.StatusUnauthorized
//   - "task_not_found" (type *task.TaskNotFound): http.StatusNotFound
//   - error: internal error
func DecodeShowResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ShowResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("task", "show", err)
			}
			err = ValidateShowResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("task", "show", err)
			}
			res := NewShowResultOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body ShowTokenVerificationFailedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("task", "show", err)
			}
			err = ValidateShowTokenVerificationFailedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("task", "show", err)
			}
			return nil, NewShowTokenVerificationFailed(&body)
		case http.StatusNotFound:
			var (
				body ShowTaskNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("task", "show", err)
			}
			err = ValidateShowTaskNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("task", "show", err)
			}
			return nil, NewShowTaskNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("task", "show", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "task" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateTaskPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("task", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the task create
// server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*task.CreatePayload)
		if !ok {
			return goahttp.ErrInvalidType("task", "create", "*task.CreatePayload", v)
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewCreateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("task", "create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the task
// create endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeCreateResponse may return the following errors:
//   - "token_verification_failed" (type *task.AuthFailed): http.StatusUnauthorized
//   - error: internal error
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body CreateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("task", "create", err)
			}
			err = ValidateCreateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("task", "create", err)
			}
			res := NewCreateResultCreated(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body CreateTokenVerificationFailedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("task", "create", err)
			}
			err = ValidateCreateTokenVerificationFailedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("task", "create", err)
			}
			return nil, NewCreateTokenVerificationFailed(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("task", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateRequest instantiates a HTTP request object with method and path
// set to call the "task" service "update" endpoint
func (c *Client) BuildUpdateRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id uint32
	)
	{
		p, ok := v.(*task.UpdatePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("task", "update", "*task.UpdatePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateTaskPath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("task", "update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRequest returns an encoder for requests sent to the task update
// server.
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*task.UpdatePayload)
		if !ok {
			return goahttp.ErrInvalidType("task", "update", "*task.UpdatePayload", v)
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewUpdateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("task", "update", err)
		}
		return nil
	}
}

// DecodeUpdateResponse returns a decoder for responses returned by the task
// update endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeUpdateResponse may return the following errors:
//   - "token_verification_failed" (type *task.AuthFailed): http.StatusUnauthorized
//   - "task_not_found" (type *task.TaskNotFound): http.StatusNotFound
//   - error: internal error
func DecodeUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("task", "update", err)
			}
			err = ValidateUpdateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("task", "update", err)
			}
			res := NewUpdateResultOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body UpdateTokenVerificationFailedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("task", "update", err)
			}
			err = ValidateUpdateTokenVerificationFailedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("task", "update", err)
			}
			return nil, NewUpdateTokenVerificationFailed(&body)
		case http.StatusNotFound:
			var (
				body UpdateTaskNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("task", "update", err)
			}
			err = ValidateUpdateTaskNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("task", "update", err)
			}
			return nil, NewUpdateTaskNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("task", "update", resp.StatusCode, string(body))
		}
	}
}

// BuildDoneRequest instantiates a HTTP request object with method and path set
// to call the "task" service "done" endpoint
func (c *Client) BuildDoneRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id uint32
	)
	{
		p, ok := v.(*task.DonePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("task", "done", "*task.DonePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DoneTaskPath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("task", "done", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDoneRequest returns an encoder for requests sent to the task done
// server.
func EncodeDoneRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*task.DonePayload)
		if !ok {
			return goahttp.ErrInvalidType("task", "done", "*task.DonePayload", v)
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewDoneRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("task", "done", err)
		}
		return nil
	}
}

// DecodeDoneResponse returns a decoder for responses returned by the task done
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
// DecodeDoneResponse may return the following errors:
//   - "token_verification_failed" (type *task.AuthFailed): http.StatusUnauthorized
//   - error: internal error
func DecodeDoneResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body DoneResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("task", "done", err)
			}
			err = ValidateDoneResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("task", "done", err)
			}
			res := NewDoneResultOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body DoneTokenVerificationFailedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("task", "done", err)
			}
			err = ValidateDoneTokenVerificationFailedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("task", "done", err)
			}
			return nil, NewDoneTokenVerificationFailed(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("task", "done", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteRequest instantiates a HTTP request object with method and path
// set to call the "task" service "delete" endpoint
func (c *Client) BuildDeleteRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id uint32
	)
	{
		p, ok := v.(*task.DeletePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("task", "delete", "*task.DeletePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteTaskPath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("task", "delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteRequest returns an encoder for requests sent to the task delete
// server.
func EncodeDeleteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*task.DeletePayload)
		if !ok {
			return goahttp.ErrInvalidType("task", "delete", "*task.DeletePayload", v)
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeDeleteResponse returns a decoder for responses returned by the task
// delete endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeDeleteResponse may return the following errors:
//   - "token_verification_failed" (type *task.AuthFailed): http.StatusUnauthorized
//   - error: internal error
func DecodeDeleteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		case http.StatusUnauthorized:
			var (
				body DeleteTokenVerificationFailedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("task", "delete", err)
			}
			err = ValidateDeleteTokenVerificationFailedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("task", "delete", err)
			}
			return nil, NewDeleteTokenVerificationFailed(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("task", "delete", resp.StatusCode, string(body))
		}
	}
}

// unmarshalBackendStoredTaskResponseBodyToTaskBackendStoredTask builds a value
// of type *task.BackendStoredTask from a value of type
// *BackendStoredTaskResponseBody.
func unmarshalBackendStoredTaskResponseBodyToTaskBackendStoredTask(v *BackendStoredTaskResponseBody) *task.BackendStoredTask {
	if v == nil {
		return nil
	}
	res := &task.BackendStoredTask{
		ID:          *v.ID,
		Name:        *v.Name,
		Description: *v.Description,
		Done:        *v.Done,
		CreatedAt:   *v.CreatedAt,
		CreatedBy:   *v.CreatedBy,
	}
	if v.DoneAt != nil {
		res.DoneAt = *v.DoneAt
	}
	if v.DoneBy != nil {
		res.DoneBy = *v.DoneBy
	}
	if v.DoneAt == nil {
		res.DoneAt = ""
	}
	if v.DoneBy == nil {
		res.DoneBy = ""
	}

	return res
}
