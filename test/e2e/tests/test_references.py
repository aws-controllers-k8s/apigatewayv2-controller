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

"""Integration tests for the API Gateway V2 resource references
"""

import logging
import time

import boto3
import pytest

from acktest.k8s import resource as k8s

from e2e import service_marker, load_apigatewayv2_resource
from e2e.replacement_values import REPLACEMENT_VALUES
import e2e.tests.helper as helper
from e2e.tests.helper import ApiGatewayValidator

DELETE_WAIT_AFTER_SECONDS = 20

apigw_validator = ApiGatewayValidator(boto3.client('apigatewayv2'))


@service_marker
@pytest.mark.canary
class TestApiGatewayV2References:

    def test_references(self):
        test_data = REPLACEMENT_VALUES.copy()
        api_name = test_data["REF_API_NAME"]
        integration_name = test_data["REF_INTEGRATION_NAME"]
        route_name = test_data["REF_ROUTE_NAME"]
        stage_name = test_data["REF_STAGE_NAME"]

        api_ref, api_data = helper.api_ref_and_data(api_resource_name=api_name,
                                                    replacement_values=test_data,
                                                    file_name="httpapi-ref")
        logging.debug(f"api resource. name: {api_name}, data: {api_data}")

        integration_ref, integration_data = helper.integration_ref_and_data(integration_resource_name=integration_name,
                                                                            replacement_values=test_data,
                                                                            file_name="integration-ref")
        logging.debug(f"integration resource. name: {integration_name}, data: {integration_data}")

        route_ref, route_data = helper.route_ref_and_data(route_resource_name=route_name,
                                                          replacement_values=test_data,
                                                          file_name="route-ref")
        logging.debug(f"route resource. name: {route_name}, data: {route_data}")

        stage_ref, stage_data = helper.stage_ref_and_data(stage_resource_name=stage_name,
                                                          replacement_values=test_data,
                                                          file_name="stage-ref")
        logging.debug(f"stage resource. name: {stage_name}, data: {stage_data}")

        # create the resources in order that initially the reference resolution fails and
        # then when the referenced resource gets created, then all resolutions eventually
        # pass and resources get synced.

        # Create stage. Needs API reference
        k8s.create_custom_resource(stage_ref, stage_data)
        stage_cr = k8s.wait_resource_consumed_by_controller(stage_ref)
        assert stage_cr is not None

        # Create route. Needs API, Integration reference
        k8s.create_custom_resource(route_ref, route_data)
        route_cr = k8s.wait_resource_consumed_by_controller(route_ref)
        assert route_cr is not None

        # Create integration. Needs API reference
        k8s.create_custom_resource(integration_ref, integration_data)
        integration_cr = k8s.wait_resource_consumed_by_controller(integration_ref)
        assert integration_cr is not None

        # Create API. Needs no reference
        k8s.create_custom_resource(api_ref, api_data)
        api_cr = k8s.wait_resource_consumed_by_controller(api_ref)
        assert api_cr is not None

        assert k8s.wait_on_condition(api_ref, "ACK.ResourceSynced", "True", wait_periods=10)
        assert k8s.wait_on_condition(integration_ref, "ACK.ResourceSynced", "True", wait_periods=10)
        assert k8s.wait_on_condition(route_ref, "ACK.ResourceSynced", "True", wait_periods=10)
        assert k8s.wait_on_condition(stage_ref, "ACK.ResourceSynced", "True", wait_periods=10)

        assert k8s.wait_on_condition(integration_ref, "ACK.ReferencesResolved", "True", wait_periods=10)
        assert k8s.wait_on_condition(route_ref, "ACK.ReferencesResolved", "True", wait_periods=10)
        assert k8s.wait_on_condition(stage_ref, "ACK.ReferencesResolved", "True", wait_periods=10)

        api_cr = k8s.get_resource(api_ref)
        api_id = api_cr['status']['apiID']

        integration_cr = k8s.get_resource(integration_ref)
        integration_id = integration_cr['status']['integrationID']

        route_cr = k8s.get_resource(route_ref)
        route_id = route_cr['status']['routeID']

        # check that the resources in Amazon API Gateway
        apigw_validator.assert_api_is_present(api_id=api_id)
        apigw_validator.assert_integration_is_present(api_id=api_id, integration_id=integration_id)
        apigw_validator.assert_route_is_present(api_id=api_id, route_id=route_id)
        apigw_validator.assert_stage_is_present(api_id=api_id, stage_name=stage_name)

        # DELETE
        k8s.delete_custom_resource(stage_ref)
        time.sleep(DELETE_WAIT_AFTER_SECONDS)

        k8s.delete_custom_resource(route_ref)
        time.sleep(DELETE_WAIT_AFTER_SECONDS)

        k8s.delete_custom_resource(integration_ref)
        time.sleep(DELETE_WAIT_AFTER_SECONDS)

        k8s.delete_custom_resource(api_ref)
        time.sleep(DELETE_WAIT_AFTER_SECONDS)

        assert not k8s.get_resource_exists(stage_ref)
        assert not k8s.get_resource_exists(route_ref)
        assert not k8s.get_resource_exists(integration_ref)
        assert not k8s.get_resource_exists(api_ref)

        # check that the resources does not exist in Amazon API Gateway
        apigw_validator.assert_stage_is_deleted(api_id=api_id, stage_name=stage_name)
        apigw_validator.assert_route_is_deleted(api_id=api_id, route_id=route_id)
        apigw_validator.assert_integration_is_deleted(api_id=api_id, integration_id=integration_id)
        apigw_validator.assert_api_is_deleted(api_id=api_id)
