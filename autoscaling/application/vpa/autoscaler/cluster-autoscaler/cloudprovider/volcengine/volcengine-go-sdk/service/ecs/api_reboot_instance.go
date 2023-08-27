// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine/response"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRebootInstanceCommon = "RebootInstance"

// RebootInstanceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RebootInstanceCommon operation. The "output" return
// value will be populated with the RebootInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RebootInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after RebootInstanceCommon Send returns without error.
//
// See RebootInstanceCommon for more information on using the RebootInstanceCommon
// API call, and error handling.
//
//	// Example sending a request using the RebootInstanceCommonRequest method.
//	req, resp := client.RebootInstanceCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ECS) RebootInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRebootInstanceCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// RebootInstanceCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation RebootInstanceCommon for usage and error information.
func (c *ECS) RebootInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RebootInstanceCommonRequest(input)
	return out, req.Send()
}

// RebootInstanceCommonWithContext is the same as RebootInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RebootInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) RebootInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RebootInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRebootInstance = "RebootInstance"

// RebootInstanceRequest generates a "volcengine/request.Request" representing the
// client's request for the RebootInstance operation. The "output" return
// value will be populated with the RebootInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RebootInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after RebootInstanceCommon Send returns without error.
//
// See RebootInstance for more information on using the RebootInstance
// API call, and error handling.
//
//	// Example sending a request using the RebootInstanceRequest method.
//	req, resp := client.RebootInstanceRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ECS) RebootInstanceRequest(input *RebootInstanceInput) (req *request.Request, output *RebootInstanceOutput) {
	op := &request.Operation{
		Name:       opRebootInstance,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RebootInstanceInput{}
	}

	output = &RebootInstanceOutput{}
	req = c.newRequest(op, input, output)

	return
}

// RebootInstance API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation RebootInstance for usage and error information.
func (c *ECS) RebootInstance(input *RebootInstanceInput) (*RebootInstanceOutput, error) {
	req, out := c.RebootInstanceRequest(input)
	return out, req.Send()
}

// RebootInstanceWithContext is the same as RebootInstance with the addition of
// the ability to pass a context and additional request options.
//
// See RebootInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) RebootInstanceWithContext(ctx volcengine.Context, input *RebootInstanceInput, opts ...request.Option) (*RebootInstanceOutput, error) {
	req, out := c.RebootInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RebootInstanceInput struct {
	_ struct{} `type:"structure"`

	ForceStop *bool `type:"boolean"`

	InstanceId *string `type:"string"`
}

// String returns the string representation
func (s RebootInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RebootInstanceInput) GoString() string {
	return s.String()
}

// SetForceStop sets the ForceStop field's value.
func (s *RebootInstanceInput) SetForceStop(v bool) *RebootInstanceInput {
	s.ForceStop = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *RebootInstanceInput) SetInstanceId(v string) *RebootInstanceInput {
	s.InstanceId = &v
	return s
}

type RebootInstanceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s RebootInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RebootInstanceOutput) GoString() string {
	return s.String()
}
