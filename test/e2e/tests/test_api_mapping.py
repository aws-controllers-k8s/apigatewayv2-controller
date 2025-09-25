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

"""Integration tests for the API Gateway V2 APIMapping resource
"""

import logging
import time

import boto3
import pytest

from acktest.k8s import resource as k8s, condition
from acktest.k8s import condition
from acktest.resources import random_suffix_name

from e2e import service_marker, load_apigatewayv2_resource
from e2e.bootstrap_resources import get_bootstrap_resources
from e2e.replacement_values import REPLACEMENT_VALUES
import e2e.tests.helper as helper
from e2e.tests.helper import ApiGatewayValidator

CREATE_WAIT_AFTER_SECONDS = 60

apigw_validator = ApiGatewayValidator(boto3.client('apigatewayv2'))
test_resource_values = REPLACEMENT_VALUES.copy()

@service_marker
@pytest.mark.canary
class TestApiGatewayV2:

    def test_crud_api_mapping(self):
        test_data = REPLACEMENT_VALUES.copy()
        api_mapping_name = random_suffix_name("api-mapping", 24)
        test_data["API_MAPPING_RES_NAME"] = api_mapping_name
        test_data["API_MAPPING_NAME"] = api_mapping_name
        test_data["DOMAIN_NAME"] = "ack-for-the-win.com"
        api_mapping_ref, api_mapping_data = helper.api_mapping_ref_and_data(api_mapping_resource_name=api_mapping_name,
                                                                            replacement_values=test_data)
        
        logging.debug(f"api mapping ref is {api_mapping_ref}, data: {api_mapping_data}")
        
        # Attempting Create expecting BadRequestException
        k8s.create_custom_resource(api_mapping_ref, api_mapping_data)
        k8s.wait_resource_consumed_by_controller(api_mapping_ref)

        expected_msg = "BadRequestException"
        condition.assert_terminal(api_mapping_ref, expected_msg)


        