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

package vpc_link

import (
	"fmt"

	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
)

var (
	waitForAvailableRequeue = ackrequeue.NeededAfter(
		fmt.Errorf("VPCLink not in '%s' state, cannot be modified",
			svcsdk.VpcLinkStatusAvailable),
		ackrequeue.DefaultRequeueAfterDuration,
	)
)
