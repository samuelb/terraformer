// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type ContinueDeploymentInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of a blue/green deployment for which you want to start rerouting
	// traffic to the replacement environment.
	DeploymentId *string `locationName:"deploymentId" type:"string"`

	// The status of the deployment's waiting period. READY_WAIT indicates the deployment
	// is ready to start shifting traffic. TERMINATION_WAIT indicates the traffic
	// is shifted, but the original target is not terminated.
	DeploymentWaitType DeploymentWaitType `locationName:"deploymentWaitType" type:"string" enum:"true"`
}

// String returns the string representation
func (s ContinueDeploymentInput) String() string {
	return awsutil.Prettify(s)
}

type ContinueDeploymentOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ContinueDeploymentOutput) String() string {
	return awsutil.Prettify(s)
}

const opContinueDeployment = "ContinueDeployment"

// ContinueDeploymentRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// For a blue/green deployment, starts the process of rerouting traffic from
// instances in the original environment to instances in the replacement environment
// without waiting for a specified wait time to elapse. (Traffic rerouting,
// which is achieved by registering instances in the replacement environment
// with the load balancer, can start as soon as all instances have a status
// of Ready.)
//
//    // Example sending a request using ContinueDeploymentRequest.
//    req := client.ContinueDeploymentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/ContinueDeployment
func (c *Client) ContinueDeploymentRequest(input *ContinueDeploymentInput) ContinueDeploymentRequest {
	op := &aws.Operation{
		Name:       opContinueDeployment,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ContinueDeploymentInput{}
	}

	req := c.newRequest(op, input, &ContinueDeploymentOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return ContinueDeploymentRequest{Request: req, Input: input, Copy: c.ContinueDeploymentRequest}
}

// ContinueDeploymentRequest is the request type for the
// ContinueDeployment API operation.
type ContinueDeploymentRequest struct {
	*aws.Request
	Input *ContinueDeploymentInput
	Copy  func(*ContinueDeploymentInput) ContinueDeploymentRequest
}

// Send marshals and sends the ContinueDeployment API request.
func (r ContinueDeploymentRequest) Send(ctx context.Context) (*ContinueDeploymentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ContinueDeploymentResponse{
		ContinueDeploymentOutput: r.Request.Data.(*ContinueDeploymentOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ContinueDeploymentResponse is the response type for the
// ContinueDeployment API operation.
type ContinueDeploymentResponse struct {
	*ContinueDeploymentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ContinueDeployment request.
func (r *ContinueDeploymentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}