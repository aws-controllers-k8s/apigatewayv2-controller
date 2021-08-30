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

package authorizer

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.APIID, b.ko.Spec.APIID) {
		delta.Add("Spec.APIID", a.ko.Spec.APIID, b.ko.Spec.APIID)
	} else if a.ko.Spec.APIID != nil && b.ko.Spec.APIID != nil {
		if *a.ko.Spec.APIID != *b.ko.Spec.APIID {
			delta.Add("Spec.APIID", a.ko.Spec.APIID, b.ko.Spec.APIID)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.AuthorizerCredentialsARN, b.ko.Spec.AuthorizerCredentialsARN) {
		delta.Add("Spec.AuthorizerCredentialsARN", a.ko.Spec.AuthorizerCredentialsARN, b.ko.Spec.AuthorizerCredentialsARN)
	} else if a.ko.Spec.AuthorizerCredentialsARN != nil && b.ko.Spec.AuthorizerCredentialsARN != nil {
		if *a.ko.Spec.AuthorizerCredentialsARN != *b.ko.Spec.AuthorizerCredentialsARN {
			delta.Add("Spec.AuthorizerCredentialsARN", a.ko.Spec.AuthorizerCredentialsARN, b.ko.Spec.AuthorizerCredentialsARN)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.AuthorizerPayloadFormatVersion, b.ko.Spec.AuthorizerPayloadFormatVersion) {
		delta.Add("Spec.AuthorizerPayloadFormatVersion", a.ko.Spec.AuthorizerPayloadFormatVersion, b.ko.Spec.AuthorizerPayloadFormatVersion)
	} else if a.ko.Spec.AuthorizerPayloadFormatVersion != nil && b.ko.Spec.AuthorizerPayloadFormatVersion != nil {
		if *a.ko.Spec.AuthorizerPayloadFormatVersion != *b.ko.Spec.AuthorizerPayloadFormatVersion {
			delta.Add("Spec.AuthorizerPayloadFormatVersion", a.ko.Spec.AuthorizerPayloadFormatVersion, b.ko.Spec.AuthorizerPayloadFormatVersion)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.AuthorizerResultTtlInSeconds, b.ko.Spec.AuthorizerResultTtlInSeconds) {
		delta.Add("Spec.AuthorizerResultTtlInSeconds", a.ko.Spec.AuthorizerResultTtlInSeconds, b.ko.Spec.AuthorizerResultTtlInSeconds)
	} else if a.ko.Spec.AuthorizerResultTtlInSeconds != nil && b.ko.Spec.AuthorizerResultTtlInSeconds != nil {
		if *a.ko.Spec.AuthorizerResultTtlInSeconds != *b.ko.Spec.AuthorizerResultTtlInSeconds {
			delta.Add("Spec.AuthorizerResultTtlInSeconds", a.ko.Spec.AuthorizerResultTtlInSeconds, b.ko.Spec.AuthorizerResultTtlInSeconds)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.AuthorizerType, b.ko.Spec.AuthorizerType) {
		delta.Add("Spec.AuthorizerType", a.ko.Spec.AuthorizerType, b.ko.Spec.AuthorizerType)
	} else if a.ko.Spec.AuthorizerType != nil && b.ko.Spec.AuthorizerType != nil {
		if *a.ko.Spec.AuthorizerType != *b.ko.Spec.AuthorizerType {
			delta.Add("Spec.AuthorizerType", a.ko.Spec.AuthorizerType, b.ko.Spec.AuthorizerType)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.AuthorizerURI, b.ko.Spec.AuthorizerURI) {
		delta.Add("Spec.AuthorizerURI", a.ko.Spec.AuthorizerURI, b.ko.Spec.AuthorizerURI)
	} else if a.ko.Spec.AuthorizerURI != nil && b.ko.Spec.AuthorizerURI != nil {
		if *a.ko.Spec.AuthorizerURI != *b.ko.Spec.AuthorizerURI {
			delta.Add("Spec.AuthorizerURI", a.ko.Spec.AuthorizerURI, b.ko.Spec.AuthorizerURI)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EnableSimpleResponses, b.ko.Spec.EnableSimpleResponses) {
		delta.Add("Spec.EnableSimpleResponses", a.ko.Spec.EnableSimpleResponses, b.ko.Spec.EnableSimpleResponses)
	} else if a.ko.Spec.EnableSimpleResponses != nil && b.ko.Spec.EnableSimpleResponses != nil {
		if *a.ko.Spec.EnableSimpleResponses != *b.ko.Spec.EnableSimpleResponses {
			delta.Add("Spec.EnableSimpleResponses", a.ko.Spec.EnableSimpleResponses, b.ko.Spec.EnableSimpleResponses)
		}
	}
	if !ackcompare.SliceStringPEqual(a.ko.Spec.IdentitySource, b.ko.Spec.IdentitySource) {
		delta.Add("Spec.IdentitySource", a.ko.Spec.IdentitySource, b.ko.Spec.IdentitySource)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.IdentityValidationExpression, b.ko.Spec.IdentityValidationExpression) {
		delta.Add("Spec.IdentityValidationExpression", a.ko.Spec.IdentityValidationExpression, b.ko.Spec.IdentityValidationExpression)
	} else if a.ko.Spec.IdentityValidationExpression != nil && b.ko.Spec.IdentityValidationExpression != nil {
		if *a.ko.Spec.IdentityValidationExpression != *b.ko.Spec.IdentityValidationExpression {
			delta.Add("Spec.IdentityValidationExpression", a.ko.Spec.IdentityValidationExpression, b.ko.Spec.IdentityValidationExpression)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.JWTConfiguration, b.ko.Spec.JWTConfiguration) {
		delta.Add("Spec.JWTConfiguration", a.ko.Spec.JWTConfiguration, b.ko.Spec.JWTConfiguration)
	} else if a.ko.Spec.JWTConfiguration != nil && b.ko.Spec.JWTConfiguration != nil {
		if !ackcompare.SliceStringPEqual(a.ko.Spec.JWTConfiguration.Audience, b.ko.Spec.JWTConfiguration.Audience) {
			delta.Add("Spec.JWTConfiguration.Audience", a.ko.Spec.JWTConfiguration.Audience, b.ko.Spec.JWTConfiguration.Audience)
		}
		if ackcompare.HasNilDifference(a.ko.Spec.JWTConfiguration.Issuer, b.ko.Spec.JWTConfiguration.Issuer) {
			delta.Add("Spec.JWTConfiguration.Issuer", a.ko.Spec.JWTConfiguration.Issuer, b.ko.Spec.JWTConfiguration.Issuer)
		} else if a.ko.Spec.JWTConfiguration.Issuer != nil && b.ko.Spec.JWTConfiguration.Issuer != nil {
			if *a.ko.Spec.JWTConfiguration.Issuer != *b.ko.Spec.JWTConfiguration.Issuer {
				delta.Add("Spec.JWTConfiguration.Issuer", a.ko.Spec.JWTConfiguration.Issuer, b.ko.Spec.JWTConfiguration.Issuer)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Name, b.ko.Spec.Name) {
		delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
	} else if a.ko.Spec.Name != nil && b.ko.Spec.Name != nil {
		if *a.ko.Spec.Name != *b.ko.Spec.Name {
			delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
		}
	}

	return delta
}
