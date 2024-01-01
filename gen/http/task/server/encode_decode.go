// Code generated by goa v3.14.1, DO NOT EDIT.
//
// task HTTP server encoders and decoders
//
// Command:
// $ goa gen backend/design

package server

import (
	task "backend/gen/task"
	"context"
	"net/http"
	"strconv"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeListResponse returns an encoder for responses returned by the task
// list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*task.ListResult)
		enc := encoder(ctx, w)
		body := NewListResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListRequest returns a decoder for requests sent to the task list
// endpoint.
func DecodeListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			limit  uint32
			offset uint32
			err    error
		)
		{
			limitRaw := r.URL.Query().Get("limit")
			if limitRaw == "" {
				limit = 20
			} else {
				v, err2 := strconv.ParseUint(limitRaw, 10, 32)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("limit", limitRaw, "unsigned integer"))
				}
				limit = uint32(v)
			}
		}
		{
			offsetRaw := r.URL.Query().Get("offset")
			if offsetRaw != "" {
				v, err2 := strconv.ParseUint(offsetRaw, 10, 32)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("offset", offsetRaw, "unsigned integer"))
				}
				offset = uint32(v)
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewListPayload(limit, offset)

		return payload, nil
	}
}

// marshalTaskBackendStoredTaskToBackendStoredTaskResponseBody builds a value
// of type *BackendStoredTaskResponseBody from a value of type
// *task.BackendStoredTask.
func marshalTaskBackendStoredTaskToBackendStoredTaskResponseBody(v *task.BackendStoredTask) *BackendStoredTaskResponseBody {
	if v == nil {
		return nil
	}
	res := &BackendStoredTaskResponseBody{
		ID:          v.ID,
		Name:        v.Name,
		Description: v.Description,
		Done:        v.Done,
		DoneAt:      v.DoneAt,
		DoneBy:      v.DoneBy,
		CreatedAt:   v.CreatedAt,
		CreatedBy:   v.CreatedBy,
	}
	{
		var zero string
		if res.DoneAt == zero {
			res.DoneAt = ""
		}
	}
	{
		var zero string
		if res.DoneBy == zero {
			res.DoneBy = ""
		}
	}

	return res
}
