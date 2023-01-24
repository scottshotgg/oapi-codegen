// Package param provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package param

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// Test defines model for test.
type Test struct {
	union json.RawMessage
}

// Test0 defines model for .
type Test0 struct {
	Item1 string `json:"item1"`
	Item2 string `json:"item2"`
}

// Test1 defines model for .
type Test1 struct {
	Item2 *string `json:"item2,omitempty"`
	Item3 *string `json:"item3,omitempty"`
}

// Test2 defines model for test2.
type Test2 struct {
	union json.RawMessage
}

// Test20 defines model for .
type Test20 = int

// Test21 defines model for .
type Test21 = string

// GetTestParams defines parameters for GetTest.
type GetTestParams struct {
	Test  *Test    `form:"test,omitempty" json:"test,omitempty"`
	Test2 *[]Test2 `form:"test2,omitempty" json:"test2,omitempty"`
}

// AsTest0 returns the union data inside the Test as a Test0
func (t Test) AsTest0() (Test0, error) {
	var body Test0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTest0 overwrites any union data inside the Test as the provided Test0
func (t *Test) FromTest0(v Test0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTest0 performs a merge with any union data inside the Test, using the provided Test0
func (t *Test) MergeTest0(v Test0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

// AsTest1 returns the union data inside the Test as a Test1
func (t Test) AsTest1() (Test1, error) {
	var body Test1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTest1 overwrites any union data inside the Test as the provided Test1
func (t *Test) FromTest1(v Test1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTest1 performs a merge with any union data inside the Test, using the provided Test1
func (t *Test) MergeTest1(v Test1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

func (t Test) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Test) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsTest20 returns the union data inside the Test2 as a Test20
func (t Test2) AsTest20() (Test20, error) {
	var body Test20
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTest20 overwrites any union data inside the Test2 as the provided Test20
func (t *Test2) FromTest20(v Test20) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTest20 performs a merge with any union data inside the Test2, using the provided Test20
func (t *Test2) MergeTest20(v Test20) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

// AsTest21 returns the union data inside the Test2 as a Test21
func (t Test2) AsTest21() (Test21, error) {
	var body Test21
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTest21 overwrites any union data inside the Test2 as the provided Test21
func (t *Test2) FromTest21(v Test21) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTest21 performs a merge with any union data inside the Test2, using the provided Test21
func (t *Test2) MergeTest21(v Test21) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

func (t Test2) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Test2) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetTest request
	GetTest(ctx context.Context, params *GetTestParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetTest(ctx context.Context, params *GetTestParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetTestRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetTestRequest generates requests for GetTest
func NewGetTestRequest(server string, params *GetTestParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/test")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Test != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "test", runtime.ParamLocationQuery, *params.Test); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Test2 != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "test2", runtime.ParamLocationQuery, *params.Test2); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetTest request
	GetTestWithResponse(ctx context.Context, params *GetTestParams, reqEditors ...RequestEditorFn) (*GetTestResponse, error)
}

type GetTestResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetTestResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetTestResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetTestWithResponse request returning *GetTestResponse
func (c *ClientWithResponses) GetTestWithResponse(ctx context.Context, params *GetTestParams, reqEditors ...RequestEditorFn) (*GetTestResponse, error) {
	rsp, err := c.GetTest(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetTestResponse(rsp)
}

// ParseGetTestResponse parses an HTTP response from a GetTestWithResponse call
func ParseGetTestResponse(rsp *http.Response) (*GetTestResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetTestResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
