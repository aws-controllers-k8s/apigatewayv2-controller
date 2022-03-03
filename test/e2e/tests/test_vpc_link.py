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

"""Integration tests for the API Gateway V2 VPCLink resource
"""

import logging
import time

import boto3
import pytest

from acktest.k8s import resource as k8s
from acktest.resources import random_suffix_name

from e2e import service_marker, load_apigatewayv2_resource
from e2e.bootstrap_resources import get_bootstrap_resources
from e2e.replacement_values import REPLACEMENT_VALUES
import e2e.tests.helper as helper
from e2e.tests.helper import ApiGatewayValidator

DELETE_WAIT_AFTER_SECONDS = 30
UPDATE_WAIT_AFTER_SECONDS = 10

apigw_validator = ApiGatewayValidator(boto3.client('apigatewayv2'))
test_resource_values = REPLACEMENT_VALUES.copy()


@service_marker
@pytest.mark.canary
class TestApiGatewayV2:

    def test_crud_vpc_link(self):
        test_data = REPLACEMENT_VALUES.copy()
        vpc_link_name = random_suffix_name("ack-test", 15)
        subnet_id_1 = get_bootstrap_resources().VPC.public_subnets.subnet_ids[0]
        subnet_id_2 = get_bootstrap_resources().VPC.public_subnets.subnet_ids[1]
        test_data['VPC_LINK_NAME'] = vpc_link_name
        test_data['VPC_LINK_RES_NAME'] = vpc_link_name
        test_data['SUBNET_ID_1'] = subnet_id_1
        test_data['SUBNET_ID_2'] = subnet_id_2
        vpc_link_ref, vpc_link_data = helper.vpc_link_ref_and_data(vpc_link_resource_name=vpc_link_name,
                                                                   replacement_values=test_data)
        logging.debug(f"vpc link resource. name: {vpc_link_name}, data: {vpc_link_data}")

        # test create
        k8s.create_custom_resource(vpc_link_ref, vpc_link_data)
        cr = k8s.wait_resource_consumed_by_controller(vpc_link_ref)

        assert cr is not None
        assert k8s.wait_on_condition(vpc_link_ref, "ACK.ResourceSynced", "True", wait_periods=10)

        vpc_link_id = cr['status']['vpcLinkID']

        # Let's check that the VPC Link appears in Amazon API Gateway
        apigw_validator.assert_vpc_link_is_present(vpc_link_id=vpc_link_id)

        apigw_validator.assert_vpc_link_name(
            vpc_link_id=vpc_link_id,
            expected_vpc_link_name=vpc_link_name
        )

        # test update
        updated_vpc_link_name = 'updated-' + vpc_link_name
        test_data['VPC_LINK_NAME'] = updated_vpc_link_name
        updated_vpc_link_resource_data = load_apigatewayv2_resource(
            "vpc-link",
            additional_replacements=test_data,
        )
        logging.debug(f"updated vpcLink resource: {updated_vpc_link_resource_data}")

        # Update the k8s resource
        k8s.patch_custom_resource(vpc_link_ref, updated_vpc_link_resource_data)
        time.sleep(UPDATE_WAIT_AFTER_SECONDS)

        assert k8s.wait_on_condition(vpc_link_ref, "ACK.ResourceSynced", "True", wait_periods=10)
        # Let's check that the VPCLink appears in Amazon API Gateway with updated name
        apigw_validator.assert_vpc_link_name(
            vpc_link_id=vpc_link_id,
            expected_vpc_link_name=updated_vpc_link_name
        )

        # test delete
        k8s.delete_custom_resource(vpc_link_ref)
        time.sleep(DELETE_WAIT_AFTER_SECONDS)
        assert not k8s.get_resource_exists(vpc_link_ref)
        # VPCLink should no longer appear in Amazon API Gateway
        apigw_validator.assert_vpc_link_is_deleted(vpc_link_id=vpc_link_id)
