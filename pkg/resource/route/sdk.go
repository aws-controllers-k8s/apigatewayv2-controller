// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package route

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	svcsdktypes "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	smithy "github.com/aws/smithy-go"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/apigatewayv2-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &svcsdk.Client{}
	_ = &svcapitypes.Route{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
	_ = fmt.Sprintf("")
	_ = &ackrequeue.NoRequeue{}
	_ = &aws.Config{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer func() {
		exit(err)
	}()
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.GetRouteOutput
	resp, err = rm.sdkapi.GetRoute(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "GetRoute", err)
	if err != nil {
		var awsErr smithy.APIError
		if errors.As(err, &awsErr) && awsErr.ErrorCode() == "NotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.ApiGatewayManaged != nil {
		ko.Status.APIGatewayManaged = resp.ApiGatewayManaged
	} else {
		ko.Status.APIGatewayManaged = nil
	}
	if resp.ApiKeyRequired != nil {
		ko.Spec.APIKeyRequired = resp.ApiKeyRequired
	} else {
		ko.Spec.APIKeyRequired = nil
	}
	if resp.AuthorizationScopes != nil {
		ko.Spec.AuthorizationScopes = aws.StringSlice(resp.AuthorizationScopes)
	} else {
		ko.Spec.AuthorizationScopes = nil
	}
	if resp.AuthorizationType != "" {
		ko.Spec.AuthorizationType = aws.String(string(resp.AuthorizationType))
	} else {
		ko.Spec.AuthorizationType = nil
	}
	if resp.AuthorizerId != nil {
		ko.Spec.AuthorizerID = resp.AuthorizerId
	} else {
		ko.Spec.AuthorizerID = nil
	}
	if resp.ModelSelectionExpression != nil {
		ko.Spec.ModelSelectionExpression = resp.ModelSelectionExpression
	} else {
		ko.Spec.ModelSelectionExpression = nil
	}
	if resp.OperationName != nil {
		ko.Spec.OperationName = resp.OperationName
	} else {
		ko.Spec.OperationName = nil
	}
	if resp.RequestModels != nil {
		ko.Spec.RequestModels = aws.StringMap(resp.RequestModels)
	} else {
		ko.Spec.RequestModels = nil
	}
	if resp.RequestParameters != nil {
		f8 := map[string]*svcapitypes.ParameterConstraints{}
		for f8key, f8valiter := range resp.RequestParameters {
			f8val := &svcapitypes.ParameterConstraints{}
			if f8valiter.Required != nil {
				f8val.Required = f8valiter.Required
			}
			f8[f8key] = f8val
		}
		ko.Spec.RequestParameters = f8
	} else {
		ko.Spec.RequestParameters = nil
	}
	if resp.RouteId != nil {
		ko.Status.RouteID = resp.RouteId
	} else {
		ko.Status.RouteID = nil
	}
	if resp.RouteKey != nil {
		ko.Spec.RouteKey = resp.RouteKey
	} else {
		ko.Spec.RouteKey = nil
	}
	if resp.RouteResponseSelectionExpression != nil {
		ko.Spec.RouteResponseSelectionExpression = resp.RouteResponseSelectionExpression
	} else {
		ko.Spec.RouteResponseSelectionExpression = nil
	}
	if resp.Target != nil {
		ko.Spec.Target = resp.Target
	} else {
		ko.Spec.Target = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.APIID == nil || r.ko.Status.RouteID == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.GetRouteInput, error) {
	res := &svcsdk.GetRouteInput{}

	if r.ko.Spec.APIID != nil {
		res.ApiId = r.ko.Spec.APIID
	}
	if r.ko.Status.RouteID != nil {
		res.RouteId = r.ko.Status.RouteID
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateRouteOutput
	_ = resp
	resp, err = rm.sdkapi.CreateRoute(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateRoute", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.ApiGatewayManaged != nil {
		ko.Status.APIGatewayManaged = resp.ApiGatewayManaged
	} else {
		ko.Status.APIGatewayManaged = nil
	}
	if resp.ApiKeyRequired != nil {
		ko.Spec.APIKeyRequired = resp.ApiKeyRequired
	} else {
		ko.Spec.APIKeyRequired = nil
	}
	if resp.AuthorizationScopes != nil {
		ko.Spec.AuthorizationScopes = aws.StringSlice(resp.AuthorizationScopes)
	} else {
		ko.Spec.AuthorizationScopes = nil
	}
	if resp.AuthorizationType != "" {
		ko.Spec.AuthorizationType = aws.String(string(resp.AuthorizationType))
	} else {
		ko.Spec.AuthorizationType = nil
	}
	if resp.AuthorizerId != nil {
		ko.Spec.AuthorizerID = resp.AuthorizerId
	} else {
		ko.Spec.AuthorizerID = nil
	}
	if resp.ModelSelectionExpression != nil {
		ko.Spec.ModelSelectionExpression = resp.ModelSelectionExpression
	} else {
		ko.Spec.ModelSelectionExpression = nil
	}
	if resp.OperationName != nil {
		ko.Spec.OperationName = resp.OperationName
	} else {
		ko.Spec.OperationName = nil
	}
	if resp.RequestModels != nil {
		ko.Spec.RequestModels = aws.StringMap(resp.RequestModels)
	} else {
		ko.Spec.RequestModels = nil
	}
	if resp.RequestParameters != nil {
		f8 := map[string]*svcapitypes.ParameterConstraints{}
		for f8key, f8valiter := range resp.RequestParameters {
			f8val := &svcapitypes.ParameterConstraints{}
			if f8valiter.Required != nil {
				f8val.Required = f8valiter.Required
			}
			f8[f8key] = f8val
		}
		ko.Spec.RequestParameters = f8
	} else {
		ko.Spec.RequestParameters = nil
	}
	if resp.RouteId != nil {
		ko.Status.RouteID = resp.RouteId
	} else {
		ko.Status.RouteID = nil
	}
	if resp.RouteKey != nil {
		ko.Spec.RouteKey = resp.RouteKey
	} else {
		ko.Spec.RouteKey = nil
	}
	if resp.RouteResponseSelectionExpression != nil {
		ko.Spec.RouteResponseSelectionExpression = resp.RouteResponseSelectionExpression
	} else {
		ko.Spec.RouteResponseSelectionExpression = nil
	}
	if resp.Target != nil {
		ko.Spec.Target = resp.Target
	} else {
		ko.Spec.Target = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateRouteInput, error) {
	res := &svcsdk.CreateRouteInput{}

	if r.ko.Spec.APIID != nil {
		res.ApiId = r.ko.Spec.APIID
	}
	if r.ko.Spec.APIKeyRequired != nil {
		res.ApiKeyRequired = r.ko.Spec.APIKeyRequired
	}
	if r.ko.Spec.AuthorizationScopes != nil {
		res.AuthorizationScopes = aws.ToStringSlice(r.ko.Spec.AuthorizationScopes)
	}
	if r.ko.Spec.AuthorizationType != nil {
		res.AuthorizationType = svcsdktypes.AuthorizationType(*r.ko.Spec.AuthorizationType)
	}
	if r.ko.Spec.AuthorizerID != nil {
		res.AuthorizerId = r.ko.Spec.AuthorizerID
	}
	if r.ko.Spec.ModelSelectionExpression != nil {
		res.ModelSelectionExpression = r.ko.Spec.ModelSelectionExpression
	}
	if r.ko.Spec.OperationName != nil {
		res.OperationName = r.ko.Spec.OperationName
	}
	if r.ko.Spec.RequestModels != nil {
		res.RequestModels = aws.ToStringMap(r.ko.Spec.RequestModels)
	}
	if r.ko.Spec.RequestParameters != nil {
		f8 := map[string]svcsdktypes.ParameterConstraints{}
		for f8key, f8valiter := range r.ko.Spec.RequestParameters {
			f8val := &svcsdktypes.ParameterConstraints{}
			if f8valiter.Required != nil {
				f8val.Required = f8valiter.Required
			}
			f8[f8key] = *f8val
		}
		res.RequestParameters = f8
	}
	if r.ko.Spec.RouteKey != nil {
		res.RouteKey = r.ko.Spec.RouteKey
	}
	if r.ko.Spec.RouteResponseSelectionExpression != nil {
		res.RouteResponseSelectionExpression = r.ko.Spec.RouteResponseSelectionExpression
	}
	if r.ko.Spec.Target != nil {
		res.Target = r.ko.Spec.Target
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newUpdateRequestPayload(ctx, desired, delta)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.UpdateRouteOutput
	_ = resp
	resp, err = rm.sdkapi.UpdateRoute(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UpdateRoute", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.ApiGatewayManaged != nil {
		ko.Status.APIGatewayManaged = resp.ApiGatewayManaged
	} else {
		ko.Status.APIGatewayManaged = nil
	}
	if resp.ApiKeyRequired != nil {
		ko.Spec.APIKeyRequired = resp.ApiKeyRequired
	} else {
		ko.Spec.APIKeyRequired = nil
	}
	if resp.AuthorizationScopes != nil {
		ko.Spec.AuthorizationScopes = aws.StringSlice(resp.AuthorizationScopes)
	} else {
		ko.Spec.AuthorizationScopes = nil
	}
	if resp.AuthorizationType != "" {
		ko.Spec.AuthorizationType = aws.String(string(resp.AuthorizationType))
	} else {
		ko.Spec.AuthorizationType = nil
	}
	if resp.AuthorizerId != nil {
		ko.Spec.AuthorizerID = resp.AuthorizerId
	} else {
		ko.Spec.AuthorizerID = nil
	}
	if resp.ModelSelectionExpression != nil {
		ko.Spec.ModelSelectionExpression = resp.ModelSelectionExpression
	} else {
		ko.Spec.ModelSelectionExpression = nil
	}
	if resp.OperationName != nil {
		ko.Spec.OperationName = resp.OperationName
	} else {
		ko.Spec.OperationName = nil
	}
	if resp.RequestModels != nil {
		ko.Spec.RequestModels = aws.StringMap(resp.RequestModels)
	} else {
		ko.Spec.RequestModels = nil
	}
	if resp.RequestParameters != nil {
		f8 := map[string]*svcapitypes.ParameterConstraints{}
		for f8key, f8valiter := range resp.RequestParameters {
			f8val := &svcapitypes.ParameterConstraints{}
			if f8valiter.Required != nil {
				f8val.Required = f8valiter.Required
			}
			f8[f8key] = f8val
		}
		ko.Spec.RequestParameters = f8
	} else {
		ko.Spec.RequestParameters = nil
	}
	if resp.RouteId != nil {
		ko.Status.RouteID = resp.RouteId
	} else {
		ko.Status.RouteID = nil
	}
	if resp.RouteKey != nil {
		ko.Spec.RouteKey = resp.RouteKey
	} else {
		ko.Spec.RouteKey = nil
	}
	if resp.RouteResponseSelectionExpression != nil {
		ko.Spec.RouteResponseSelectionExpression = resp.RouteResponseSelectionExpression
	} else {
		ko.Spec.RouteResponseSelectionExpression = nil
	}
	if resp.Target != nil {
		ko.Spec.Target = resp.Target
	} else {
		ko.Spec.Target = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
	delta *ackcompare.Delta,
) (*svcsdk.UpdateRouteInput, error) {
	res := &svcsdk.UpdateRouteInput{}

	if r.ko.Spec.APIID != nil {
		res.ApiId = r.ko.Spec.APIID
	}
	if r.ko.Spec.APIKeyRequired != nil {
		res.ApiKeyRequired = r.ko.Spec.APIKeyRequired
	}
	if r.ko.Spec.AuthorizationScopes != nil {
		res.AuthorizationScopes = aws.ToStringSlice(r.ko.Spec.AuthorizationScopes)
	}
	if r.ko.Spec.AuthorizationType != nil {
		res.AuthorizationType = svcsdktypes.AuthorizationType(*r.ko.Spec.AuthorizationType)
	}
	if r.ko.Spec.AuthorizerID != nil {
		res.AuthorizerId = r.ko.Spec.AuthorizerID
	}
	if r.ko.Spec.ModelSelectionExpression != nil {
		res.ModelSelectionExpression = r.ko.Spec.ModelSelectionExpression
	}
	if r.ko.Spec.OperationName != nil {
		res.OperationName = r.ko.Spec.OperationName
	}
	if r.ko.Spec.RequestModels != nil {
		res.RequestModels = aws.ToStringMap(r.ko.Spec.RequestModels)
	}
	if r.ko.Spec.RequestParameters != nil {
		f8 := map[string]svcsdktypes.ParameterConstraints{}
		for f8key, f8valiter := range r.ko.Spec.RequestParameters {
			f8val := &svcsdktypes.ParameterConstraints{}
			if f8valiter.Required != nil {
				f8val.Required = f8valiter.Required
			}
			f8[f8key] = *f8val
		}
		res.RequestParameters = f8
	}
	if r.ko.Status.RouteID != nil {
		res.RouteId = r.ko.Status.RouteID
	}
	if r.ko.Spec.RouteKey != nil {
		res.RouteKey = r.ko.Spec.RouteKey
	}
	if r.ko.Spec.RouteResponseSelectionExpression != nil {
		res.RouteResponseSelectionExpression = r.ko.Spec.RouteResponseSelectionExpression
	}
	if r.ko.Spec.Target != nil {
		res.Target = r.ko.Spec.Target
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer func() {
		exit(err)
	}()
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteRouteOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteRoute(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteRoute", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteRouteInput, error) {
	res := &svcsdk.DeleteRouteInput{}

	if r.ko.Spec.APIID != nil {
		res.ApiId = r.ko.Spec.APIID
	}
	if r.ko.Status.RouteID != nil {
		res.RouteId = r.ko.Status.RouteID
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.Route,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}
	var termError *ackerr.TerminalError
	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
