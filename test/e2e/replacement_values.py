# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
#	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

"""Stores the values used by each of the integration tests for replacing the
APIGatewayV2-specific test variables.
"""

REPLACEMENT_VALUES = {
    "API_NAME": "ack-test-api",
    "API_TITLE": "ack-test-api",
    "IMPORT_API_NAME": "ack-test-import-api",
    "IMPORT_API_TITLE": "ack-test-import-api",
    "API_ID": "api_id",
    "INTEGRATION_NAME": "ack-test-integration",
    "INTEGRATION_URI": "https://httpbin.org/get",
    "AUTHORIZER_NAME": "ack-test-authorizer",
    "AUTHORIZER_TITLE": "ack-test-authorizer",
    "IDENTITY_SOURCE": "$request.header.Authorization",
    "AUTHORIZER_URI": "authorizer_uri",
    "ROUTE_NAME": "ack-test-route",
    "ROUTE_PATH": "httpbin",
    "ROUTE_KEY": "GET /httpbin",
    "INTEGRATION_ID": "integration_id",
    "AUTHORIZER_ID": "authorizer_id",
    "STAGE_NAME": "test",
    "STAGE_DESCRIPTION": "ack-test-stage",
    "VPC_LINK_RES_NAME": "ack-test",
    "VPC_LINK_NAME": "ack-test",
    "SUBNET_ID_1": "subnet-id-1",
    "SUBNET_ID_2": "subnet-id-2",
    "REF_API_NAME": "ack-test-ref-api",
    "REF_INTEGRATION_NAME": "ack-test-ref-integration",
    "REF_ROUTE_NAME": "ack-test-ref-route",
    "REF_STAGE_NAME": "ack-test-ref-stage",
    "REF_STAGE_DESCRIPTION": "Stage for ack resource reference testing"
}
