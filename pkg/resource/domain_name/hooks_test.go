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
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	svcsdktypes "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"

	svcapitypes "github.com/aws-controllers-k8s/apigatewayv2-controller/apis/v1alpha1"
)

func TestSetDomainNameEndpointConfigurations(t *testing.T) {
	tests := []struct {
		name string
		resp *svcsdk.GetDomainNameOutput
		want []*svcapitypes.DomainNameEndpointConfiguration
	}{
		{
			name: "nil response clears status",
			resp: nil,
			want: nil,
		},
		{
			name: "nil configuration list clears status",
			resp: &svcsdk.GetDomainNameOutput{},
			want: nil,
		},
		{
			name: "empty configuration list",
			resp: &svcsdk.GetDomainNameOutput{
				DomainNameConfigurations: []svcsdktypes.DomainNameConfiguration{},
			},
			want: []*svcapitypes.DomainNameEndpointConfiguration{},
		},
		{
			name: "single configuration",
			resp: &svcsdk.GetDomainNameOutput{
				DomainNameConfigurations: []svcsdktypes.DomainNameConfiguration{
					{
						ApiGatewayDomainName: aws.String("d-abc123.execute-api.eu-central-1.amazonaws.com"),
						HostedZoneId:         aws.String("Z1UJRXOUMOOFQ8"),
						CertificateArn:       aws.String("arn:aws:acm:eu-central-1:123456789012:certificate/abc"),
					},
				},
			},
			want: []*svcapitypes.DomainNameEndpointConfiguration{
				{
					APIGatewayDomainName: aws.String("d-abc123.execute-api.eu-central-1.amazonaws.com"),
					HostedZoneID:         aws.String("Z1UJRXOUMOOFQ8"),
				},
			},
		},
		{
			name: "multiple configurations preserve order",
			resp: &svcsdk.GetDomainNameOutput{
				DomainNameConfigurations: []svcsdktypes.DomainNameConfiguration{
					{
						ApiGatewayDomainName: aws.String("d-first.execute-api.eu-central-1.amazonaws.com"),
						HostedZoneId:         aws.String("Z1UJRXOUMOOFQ8"),
					},
					{
						ApiGatewayDomainName: aws.String("d-second.execute-api.us-east-1.amazonaws.com"),
						HostedZoneId:         aws.String("Z1UJRXOUMOOFQ9"),
					},
				},
			},
			want: []*svcapitypes.DomainNameEndpointConfiguration{
				{
					APIGatewayDomainName: aws.String("d-first.execute-api.eu-central-1.amazonaws.com"),
					HostedZoneID:         aws.String("Z1UJRXOUMOOFQ8"),
				},
				{
					APIGatewayDomainName: aws.String("d-second.execute-api.us-east-1.amazonaws.com"),
					HostedZoneID:         aws.String("Z1UJRXOUMOOFQ9"),
				},
			},
		},
		{
			name: "unset members stay nil",
			resp: &svcsdk.GetDomainNameOutput{
				DomainNameConfigurations: []svcsdktypes.DomainNameConfiguration{
					{HostedZoneId: aws.String("Z1UJRXOUMOOFQ8")},
				},
			},
			want: []*svcapitypes.DomainNameEndpointConfiguration{
				{HostedZoneID: aws.String("Z1UJRXOUMOOFQ8")},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Seed a stale value so the clearing cases are meaningful.
			ko := &svcapitypes.DomainName{}
			ko.Status.DomainNameConfigurations = []*svcapitypes.DomainNameEndpointConfiguration{
				{APIGatewayDomainName: aws.String("stale")},
			}

			setDomainNameEndpointConfigurations(ko, test.resp)

			got := ko.Status.DomainNameConfigurations
			if test.want == nil {
				if got != nil {
					t.Fatalf("expected nil status, got %d element(s)", len(got))
				}
				return
			}
			if got == nil {
				t.Fatal("expected non-nil status, got nil")
			}
			if len(got) != len(test.want) {
				t.Fatalf("expected %d element(s), got %d", len(test.want), len(got))
			}
			for i := range test.want {
				assertStrPtrEqual(t, "APIGatewayDomainName", test.want[i].APIGatewayDomainName, got[i].APIGatewayDomainName)
				assertStrPtrEqual(t, "HostedZoneID", test.want[i].HostedZoneID, got[i].HostedZoneID)
			}
		})
	}
}

func assertStrPtrEqual(t *testing.T, field string, want, got *string) {
	t.Helper()
	switch {
	case want == nil && got == nil:
	case want == nil:
		t.Errorf("%s: expected nil, got %q", field, *got)
	case got == nil:
		t.Errorf("%s: expected %q, got nil", field, *want)
	case *want != *got:
		t.Errorf("%s: expected %q, got %q", field, *want, *got)
	}
}
