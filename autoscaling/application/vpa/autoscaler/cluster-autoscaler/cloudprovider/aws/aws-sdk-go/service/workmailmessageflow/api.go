// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmailmessageflow

import (
	"fmt"
	"io"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/awsutil"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol/restjson"
)

const opGetRawMessageContent = "GetRawMessageContent"

// GetRawMessageContentRequest generates a "aws/request.Request" representing the
// client's request for the GetRawMessageContent operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetRawMessageContent for more information on using the GetRawMessageContent
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//	// Example sending a request using the GetRawMessageContentRequest method.
//	req, resp := client.GetRawMessageContentRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/workmailmessageflow-2019-05-01/GetRawMessageContent
func (c *WorkMailMessageFlow) GetRawMessageContentRequest(input *GetRawMessageContentInput) (req *request.Request, output *GetRawMessageContentOutput) {
	op := &request.Operation{
		Name:       opGetRawMessageContent,
		HTTPMethod: "GET",
		HTTPPath:   "/messages/{messageId}",
	}

	if input == nil {
		input = &GetRawMessageContentInput{}
	}

	output = &GetRawMessageContentOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetRawMessageContent API operation for Amazon WorkMail Message Flow.
//
// Retrieves the raw content of an in-transit email message, in MIME format.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon WorkMail Message Flow's
// API operation GetRawMessageContent for usage and error information.
//
// Returned Error Types:
//   - ResourceNotFoundException
//     The requested email message is not found.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/workmailmessageflow-2019-05-01/GetRawMessageContent
func (c *WorkMailMessageFlow) GetRawMessageContent(input *GetRawMessageContentInput) (*GetRawMessageContentOutput, error) {
	req, out := c.GetRawMessageContentRequest(input)
	return out, req.Send()
}

// GetRawMessageContentWithContext is the same as GetRawMessageContent with the addition of
// the ability to pass a context and additional request options.
//
// See GetRawMessageContent for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WorkMailMessageFlow) GetRawMessageContentWithContext(ctx aws.Context, input *GetRawMessageContentInput, opts ...request.Option) (*GetRawMessageContentOutput, error) {
	req, out := c.GetRawMessageContentRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opPutRawMessageContent = "PutRawMessageContent"

// PutRawMessageContentRequest generates a "aws/request.Request" representing the
// client's request for the PutRawMessageContent operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See PutRawMessageContent for more information on using the PutRawMessageContent
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//	// Example sending a request using the PutRawMessageContentRequest method.
//	req, resp := client.PutRawMessageContentRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/workmailmessageflow-2019-05-01/PutRawMessageContent
func (c *WorkMailMessageFlow) PutRawMessageContentRequest(input *PutRawMessageContentInput) (req *request.Request, output *PutRawMessageContentOutput) {
	op := &request.Operation{
		Name:       opPutRawMessageContent,
		HTTPMethod: "POST",
		HTTPPath:   "/messages/{messageId}",
	}

	if input == nil {
		input = &PutRawMessageContentInput{}
	}

	output = &PutRawMessageContentOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Swap(restjson.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	return
}

// PutRawMessageContent API operation for Amazon WorkMail Message Flow.
//
// Updates the raw content of an in-transit email message, in MIME format.
//
// This example describes how to update in-transit email message. For more information
// and examples for using this API, see Updating message content with AWS Lambda
// (https://docs.aws.amazon.com/workmail/latest/adminguide/update-with-lambda.html).
//
// Updates to an in-transit message only appear when you call PutRawMessageContent
// from an AWS Lambda function configured with a synchronous Run Lambda (https://docs.aws.amazon.com/workmail/latest/adminguide/lambda.html#synchronous-rules)
// rule. If you call PutRawMessageContent on a delivered or sent message, the
// message remains unchanged, even though GetRawMessageContent (https://docs.aws.amazon.com/workmail/latest/APIReference/API_messageflow_GetRawMessageContent.html)
// returns an updated message.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon WorkMail Message Flow's
// API operation PutRawMessageContent for usage and error information.
//
// Returned Error Types:
//
//   - ResourceNotFoundException
//     The requested email message is not found.
//
//   - InvalidContentLocation
//     WorkMail could not access the updated email content. Possible reasons:
//
//   - You made the request in a region other than your S3 bucket region.
//
//   - The S3 bucket owner (https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-owner-condition.html)
//     is not the same as the calling AWS account.
//
//   - You have an incomplete or missing S3 bucket policy. For more information
//     about policies, see Updating message content with AWS Lambda (https://docs.aws.amazon.com/workmail/latest/adminguide/update-with-lambda.html)
//     in the WorkMail Administrator Guide.
//
//   - MessageRejected
//     The requested email could not be updated due to an error in the MIME content.
//     Check the error message for more information about what caused the error.
//
//   - MessageFrozen
//     The requested email is not eligible for update. This is usually the case
//     for a redirected email.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/workmailmessageflow-2019-05-01/PutRawMessageContent
func (c *WorkMailMessageFlow) PutRawMessageContent(input *PutRawMessageContentInput) (*PutRawMessageContentOutput, error) {
	req, out := c.PutRawMessageContentRequest(input)
	return out, req.Send()
}

// PutRawMessageContentWithContext is the same as PutRawMessageContent with the addition of
// the ability to pass a context and additional request options.
//
// See PutRawMessageContent for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WorkMailMessageFlow) PutRawMessageContentWithContext(ctx aws.Context, input *PutRawMessageContentInput, opts ...request.Option) (*PutRawMessageContentOutput, error) {
	req, out := c.PutRawMessageContentRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetRawMessageContentInput struct {
	_ struct{} `type:"structure" nopayload:"true"`

	// The identifier of the email message to retrieve.
	//
	// MessageId is a required field
	MessageId *string `location:"uri" locationName:"messageId" min:"1" type:"string" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s GetRawMessageContentInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s GetRawMessageContentInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRawMessageContentInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetRawMessageContentInput"}
	if s.MessageId == nil {
		invalidParams.Add(request.NewErrParamRequired("MessageId"))
	}
	if s.MessageId != nil && len(*s.MessageId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("MessageId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetMessageId sets the MessageId field's value.
func (s *GetRawMessageContentInput) SetMessageId(v string) *GetRawMessageContentInput {
	s.MessageId = &v
	return s
}

type GetRawMessageContentOutput struct {
	_ struct{} `type:"structure" payload:"MessageContent"`

	// The raw content of the email message, in MIME format.
	//
	// MessageContent is a required field
	MessageContent io.ReadCloser `locationName:"messageContent" type:"blob" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s GetRawMessageContentOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s GetRawMessageContentOutput) GoString() string {
	return s.String()
}

// SetMessageContent sets the MessageContent field's value.
func (s *GetRawMessageContentOutput) SetMessageContent(v io.ReadCloser) *GetRawMessageContentOutput {
	s.MessageContent = v
	return s
}

// WorkMail could not access the updated email content. Possible reasons:
//
//   - You made the request in a region other than your S3 bucket region.
//
//   - The S3 bucket owner (https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-owner-condition.html)
//     is not the same as the calling AWS account.
//
//   - You have an incomplete or missing S3 bucket policy. For more information
//     about policies, see Updating message content with AWS Lambda (https://docs.aws.amazon.com/workmail/latest/adminguide/update-with-lambda.html)
//     in the WorkMail Administrator Guide.
type InvalidContentLocation struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InvalidContentLocation) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InvalidContentLocation) GoString() string {
	return s.String()
}

func newErrorInvalidContentLocation(v protocol.ResponseMetadata) error {
	return &InvalidContentLocation{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *InvalidContentLocation) Code() string {
	return "InvalidContentLocation"
}

// Message returns the exception's message.
func (s *InvalidContentLocation) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *InvalidContentLocation) OrigErr() error {
	return nil
}

func (s *InvalidContentLocation) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *InvalidContentLocation) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *InvalidContentLocation) RequestID() string {
	return s.RespMetadata.RequestID
}

// The requested email is not eligible for update. This is usually the case
// for a redirected email.
type MessageFrozen struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s MessageFrozen) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s MessageFrozen) GoString() string {
	return s.String()
}

func newErrorMessageFrozen(v protocol.ResponseMetadata) error {
	return &MessageFrozen{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *MessageFrozen) Code() string {
	return "MessageFrozen"
}

// Message returns the exception's message.
func (s *MessageFrozen) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *MessageFrozen) OrigErr() error {
	return nil
}

func (s *MessageFrozen) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *MessageFrozen) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *MessageFrozen) RequestID() string {
	return s.RespMetadata.RequestID
}

// The requested email could not be updated due to an error in the MIME content.
// Check the error message for more information about what caused the error.
type MessageRejected struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s MessageRejected) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s MessageRejected) GoString() string {
	return s.String()
}

func newErrorMessageRejected(v protocol.ResponseMetadata) error {
	return &MessageRejected{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *MessageRejected) Code() string {
	return "MessageRejected"
}

// Message returns the exception's message.
func (s *MessageRejected) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *MessageRejected) OrigErr() error {
	return nil
}

func (s *MessageRejected) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *MessageRejected) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *MessageRejected) RequestID() string {
	return s.RespMetadata.RequestID
}

type PutRawMessageContentInput struct {
	_ struct{} `type:"structure"`

	// Describes the raw message content of the updated email message.
	//
	// Content is a required field
	Content *RawMessageContent `locationName:"content" type:"structure" required:"true"`

	// The identifier of the email message being updated.
	//
	// MessageId is a required field
	MessageId *string `location:"uri" locationName:"messageId" min:"1" type:"string" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PutRawMessageContentInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PutRawMessageContentInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutRawMessageContentInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutRawMessageContentInput"}
	if s.Content == nil {
		invalidParams.Add(request.NewErrParamRequired("Content"))
	}
	if s.MessageId == nil {
		invalidParams.Add(request.NewErrParamRequired("MessageId"))
	}
	if s.MessageId != nil && len(*s.MessageId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("MessageId", 1))
	}
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			invalidParams.AddNested("Content", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetContent sets the Content field's value.
func (s *PutRawMessageContentInput) SetContent(v *RawMessageContent) *PutRawMessageContentInput {
	s.Content = v
	return s
}

// SetMessageId sets the MessageId field's value.
func (s *PutRawMessageContentInput) SetMessageId(v string) *PutRawMessageContentInput {
	s.MessageId = &v
	return s
}

type PutRawMessageContentOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PutRawMessageContentOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PutRawMessageContentOutput) GoString() string {
	return s.String()
}

// Provides the MIME content of the updated email message as an S3 object. All
// MIME content must meet the following criteria:
//
//   - Each part of a multipart MIME message must be formatted properly.
//
//   - Attachments must be of a content type that Amazon SES supports. For
//     more information, see Unsupported Attachment Types (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/mime-types-appendix.html).
//
//   - If any of the MIME parts in a message contain content that is outside
//     of the 7-bit ASCII character range, we recommend encoding that content.
//
//   - Per RFC 5321 (https://tools.ietf.org/html/rfc5321#section-4.5.3.1.6),
//     the maximum length of each line of text, including the <CRLF>, must not
//     exceed 1,000 characters.
//
//   - The message must contain all the required header fields. Check the returned
//     error message for more information.
//
//   - The value of immutable headers must remain unchanged. Check the returned
//     error message for more information.
//
//   - Certain unique headers can only appear once. Check the returned error
//     message for more information.
type RawMessageContent struct {
	_ struct{} `type:"structure"`

	// The S3 reference of an email message.
	//
	// S3Reference is a required field
	S3Reference *S3Reference `locationName:"s3Reference" type:"structure" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s RawMessageContent) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s RawMessageContent) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RawMessageContent) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RawMessageContent"}
	if s.S3Reference == nil {
		invalidParams.Add(request.NewErrParamRequired("S3Reference"))
	}
	if s.S3Reference != nil {
		if err := s.S3Reference.Validate(); err != nil {
			invalidParams.AddNested("S3Reference", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetS3Reference sets the S3Reference field's value.
func (s *RawMessageContent) SetS3Reference(v *S3Reference) *RawMessageContent {
	s.S3Reference = v
	return s
}

// The requested email message is not found.
type ResourceNotFoundException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ResourceNotFoundException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ResourceNotFoundException) GoString() string {
	return s.String()
}

func newErrorResourceNotFoundException(v protocol.ResponseMetadata) error {
	return &ResourceNotFoundException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ResourceNotFoundException) Code() string {
	return "ResourceNotFoundException"
}

// Message returns the exception's message.
func (s *ResourceNotFoundException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ResourceNotFoundException) OrigErr() error {
	return nil
}

func (s *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ResourceNotFoundException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ResourceNotFoundException) RequestID() string {
	return s.RespMetadata.RequestID
}

// Amazon S3 object representing the updated message content, in MIME format.
//
// The region for the S3 bucket containing the S3 object must match the region
// used for WorkMail operations. Also, for WorkMail to process an S3 object,
// it must have permission to access that object. For more information, see
// Updating message content with AWS Lambda (https://docs.aws.amazon.com/workmail/latest/adminguide/update-with-lambda.html).
type S3Reference struct {
	_ struct{} `type:"structure"`

	// The S3 bucket name.
	//
	// Bucket is a required field
	Bucket *string `locationName:"bucket" min:"3" type:"string" required:"true"`

	// The S3 key object name.
	//
	// Key is a required field
	Key *string `locationName:"key" min:"1" type:"string" required:"true"`

	// If you enable versioning for the bucket, you can specify the object version.
	ObjectVersion *string `locationName:"objectVersion" min:"1" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s S3Reference) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s S3Reference) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *S3Reference) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "S3Reference"}
	if s.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}
	if s.Bucket != nil && len(*s.Bucket) < 3 {
		invalidParams.Add(request.NewErrParamMinLen("Bucket", 3))
	}
	if s.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Key", 1))
	}
	if s.ObjectVersion != nil && len(*s.ObjectVersion) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("ObjectVersion", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBucket sets the Bucket field's value.
func (s *S3Reference) SetBucket(v string) *S3Reference {
	s.Bucket = &v
	return s
}

// SetKey sets the Key field's value.
func (s *S3Reference) SetKey(v string) *S3Reference {
	s.Key = &v
	return s
}

// SetObjectVersion sets the ObjectVersion field's value.
func (s *S3Reference) SetObjectVersion(v string) *S3Reference {
	s.ObjectVersion = &v
	return s
}
