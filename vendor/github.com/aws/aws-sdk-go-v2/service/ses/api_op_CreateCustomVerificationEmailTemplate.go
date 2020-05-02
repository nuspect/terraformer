// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Represents a request to create a custom verification email template.
type CreateCustomVerificationEmailTemplateInput struct {
	_ struct{} `type:"structure"`

	// The URL that the recipient of the verification email is sent to if his or
	// her address is not successfully verified.
	//
	// FailureRedirectionURL is a required field
	FailureRedirectionURL *string `type:"string" required:"true"`

	// The email address that the custom verification email is sent from.
	//
	// FromEmailAddress is a required field
	FromEmailAddress *string `type:"string" required:"true"`

	// The URL that the recipient of the verification email is sent to if his or
	// her address is successfully verified.
	//
	// SuccessRedirectionURL is a required field
	SuccessRedirectionURL *string `type:"string" required:"true"`

	// The content of the custom verification email. The total size of the email
	// must be less than 10 MB. The message body may contain HTML, with some limitations.
	// For more information, see Custom Verification Email Frequently Asked Questions
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/custom-verification-emails.html#custom-verification-emails-faq)
	// in the Amazon SES Developer Guide.
	//
	// TemplateContent is a required field
	TemplateContent *string `type:"string" required:"true"`

	// The name of the custom verification email template.
	//
	// TemplateName is a required field
	TemplateName *string `type:"string" required:"true"`

	// The subject line of the custom verification email.
	//
	// TemplateSubject is a required field
	TemplateSubject *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateCustomVerificationEmailTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCustomVerificationEmailTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCustomVerificationEmailTemplateInput"}

	if s.FailureRedirectionURL == nil {
		invalidParams.Add(aws.NewErrParamRequired("FailureRedirectionURL"))
	}

	if s.FromEmailAddress == nil {
		invalidParams.Add(aws.NewErrParamRequired("FromEmailAddress"))
	}

	if s.SuccessRedirectionURL == nil {
		invalidParams.Add(aws.NewErrParamRequired("SuccessRedirectionURL"))
	}

	if s.TemplateContent == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateContent"))
	}

	if s.TemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateName"))
	}

	if s.TemplateSubject == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateSubject"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateCustomVerificationEmailTemplateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateCustomVerificationEmailTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateCustomVerificationEmailTemplate = "CreateCustomVerificationEmailTemplate"

// CreateCustomVerificationEmailTemplateRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Creates a new custom verification email template.
//
// For more information about custom verification email templates, see Using
// Custom Verification Email Templates (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/custom-verification-emails.html)
// in the Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using CreateCustomVerificationEmailTemplateRequest.
//    req := client.CreateCustomVerificationEmailTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/CreateCustomVerificationEmailTemplate
func (c *Client) CreateCustomVerificationEmailTemplateRequest(input *CreateCustomVerificationEmailTemplateInput) CreateCustomVerificationEmailTemplateRequest {
	op := &aws.Operation{
		Name:       opCreateCustomVerificationEmailTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateCustomVerificationEmailTemplateInput{}
	}

	req := c.newRequest(op, input, &CreateCustomVerificationEmailTemplateOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return CreateCustomVerificationEmailTemplateRequest{Request: req, Input: input, Copy: c.CreateCustomVerificationEmailTemplateRequest}
}

// CreateCustomVerificationEmailTemplateRequest is the request type for the
// CreateCustomVerificationEmailTemplate API operation.
type CreateCustomVerificationEmailTemplateRequest struct {
	*aws.Request
	Input *CreateCustomVerificationEmailTemplateInput
	Copy  func(*CreateCustomVerificationEmailTemplateInput) CreateCustomVerificationEmailTemplateRequest
}

// Send marshals and sends the CreateCustomVerificationEmailTemplate API request.
func (r CreateCustomVerificationEmailTemplateRequest) Send(ctx context.Context) (*CreateCustomVerificationEmailTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCustomVerificationEmailTemplateResponse{
		CreateCustomVerificationEmailTemplateOutput: r.Request.Data.(*CreateCustomVerificationEmailTemplateOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCustomVerificationEmailTemplateResponse is the response type for the
// CreateCustomVerificationEmailTemplate API operation.
type CreateCustomVerificationEmailTemplateResponse struct {
	*CreateCustomVerificationEmailTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCustomVerificationEmailTemplate request.
func (r *CreateCustomVerificationEmailTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}