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

package domain_name

import (
	svcsdk "github.com/aws/aws-sdk-go-v2/service/apigatewayv2"

	svcapitypes "github.com/aws-controllers-k8s/apigatewayv2-controller/apis/v1alpha1"
)

// setDomainNameEndpointConfigurations copies the service-assigned endpoint
// values from a GetDomainName response into Status. They sit on an SDK shape
// shared with the Create/Update input, so they stay ignored during generation
// to keep them out of Spec and are projected here instead.
func setDomainNameEndpointConfigurations(
	ko *svcapitypes.DomainName,
	resp *svcsdk.GetDomainNameOutput,
) {
	if resp == nil || resp.DomainNameConfigurations == nil {
		ko.Status.DomainNameConfigurations = nil
		return
	}
	cfgs := make([]*svcapitypes.DomainNameEndpointConfiguration, 0, len(resp.DomainNameConfigurations))
	for _, elem := range resp.DomainNameConfigurations {
		cfgs = append(cfgs, &svcapitypes.DomainNameEndpointConfiguration{
			APIGatewayDomainName: elem.ApiGatewayDomainName,
			HostedZoneID:         elem.HostedZoneId,
		})
	}
	ko.Status.DomainNameConfigurations = cfgs
}
