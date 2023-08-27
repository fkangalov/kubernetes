// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsdataservice

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have sufficient access to perform this action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// There is an error in the call or in a SQL statement.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeForbiddenException for service response error code
	// "ForbiddenException".
	//
	// There are insufficient privileges to make the call.
	ErrCodeForbiddenException = "ForbiddenException"

	// ErrCodeInternalServerErrorException for service response error code
	// "InternalServerErrorException".
	//
	// An internal error occurred.
	ErrCodeInternalServerErrorException = "InternalServerErrorException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// The resourceArn, secretArn, or transactionId value can't be found.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeServiceUnavailableError for service response error code
	// "ServiceUnavailableError".
	//
	// The service specified by the resourceArn parameter is not available.
	ErrCodeServiceUnavailableError = "ServiceUnavailableError"

	// ErrCodeStatementTimeoutException for service response error code
	// "StatementTimeoutException".
	//
	// The execution of the SQL statement timed out.
	ErrCodeStatementTimeoutException = "StatementTimeoutException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":        newErrorAccessDeniedException,
	"BadRequestException":          newErrorBadRequestException,
	"ForbiddenException":           newErrorForbiddenException,
	"InternalServerErrorException": newErrorInternalServerErrorException,
	"NotFoundException":            newErrorNotFoundException,
	"ServiceUnavailableError":      newErrorServiceUnavailableError,
	"StatementTimeoutException":    newErrorStatementTimeoutException,
}
