// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package accessanalyzer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Retrieves a list of analyzers.
type ListAnalyzersInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in the response.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" type:"integer"`

	// A token used for pagination of results returned.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The type of analyzer.
	Type Type `location:"querystring" locationName:"type" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListAnalyzersInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAnalyzersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// The response to the request.
type ListAnalyzersOutput struct {
	_ struct{} `type:"structure"`

	// The analyzers retrieved.
	//
	// Analyzers is a required field
	Analyzers []AnalyzerSummary `locationName:"analyzers" type:"list" required:"true"`

	// A token used for pagination of results returned.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListAnalyzersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAnalyzersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Analyzers != nil {
		v := s.Analyzers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "analyzers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListAnalyzers = "ListAnalyzers"

// ListAnalyzersRequest returns a request value for making API operation for
// Access Analyzer.
//
// Retrieves a list of analyzers.
//
//    // Example sending a request using ListAnalyzersRequest.
//    req := client.ListAnalyzersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/accessanalyzer-2019-11-01/ListAnalyzers
func (c *Client) ListAnalyzersRequest(input *ListAnalyzersInput) ListAnalyzersRequest {
	op := &aws.Operation{
		Name:       opListAnalyzers,
		HTTPMethod: "GET",
		HTTPPath:   "/analyzer",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListAnalyzersInput{}
	}

	req := c.newRequest(op, input, &ListAnalyzersOutput{})
	return ListAnalyzersRequest{Request: req, Input: input, Copy: c.ListAnalyzersRequest}
}

// ListAnalyzersRequest is the request type for the
// ListAnalyzers API operation.
type ListAnalyzersRequest struct {
	*aws.Request
	Input *ListAnalyzersInput
	Copy  func(*ListAnalyzersInput) ListAnalyzersRequest
}

// Send marshals and sends the ListAnalyzers API request.
func (r ListAnalyzersRequest) Send(ctx context.Context) (*ListAnalyzersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAnalyzersResponse{
		ListAnalyzersOutput: r.Request.Data.(*ListAnalyzersOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListAnalyzersRequestPaginator returns a paginator for ListAnalyzers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListAnalyzersRequest(input)
//   p := accessanalyzer.NewListAnalyzersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListAnalyzersPaginator(req ListAnalyzersRequest) ListAnalyzersPaginator {
	return ListAnalyzersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListAnalyzersInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListAnalyzersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListAnalyzersPaginator struct {
	aws.Pager
}

func (p *ListAnalyzersPaginator) CurrentPage() *ListAnalyzersOutput {
	return p.Pager.CurrentPage().(*ListAnalyzersOutput)
}

// ListAnalyzersResponse is the response type for the
// ListAnalyzers API operation.
type ListAnalyzersResponse struct {
	*ListAnalyzersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAnalyzers request.
func (r *ListAnalyzersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}