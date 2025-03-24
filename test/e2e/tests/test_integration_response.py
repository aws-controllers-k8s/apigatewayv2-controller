"""Integration tests for the API Gateway V2 IntegrationResponse
"""

import logging
import time
import json

import boto3
import pytest

from acktest.k8s import resource as k8s
from acktest.k8s import condition
from acktest.resources import random_suffix_name
from e2e import service_marker, load_apigatewayv2_resource
from e2e.replacement_values import REPLACEMENT_VALUES
from e2e.tests.helper import ApiGatewayValidator
import e2e.tests.helper as helper

CREATE_WAIT_AFTER_SECONDS = 10
UPDATE_WAIT_AFTER_SECONDS = 10
DELETE_WAIT_AFTER_SECONDS = 10

INTEGRATION_RESPONSE_RESOURCE_PLURAL = 'integrationresponses'

apigw_validator = ApiGatewayValidator(boto3.client('apigatewayv2'))
test_resource_values = REPLACEMENT_VALUES.copy()


def sanitize_dict_for_yaml(replacement_values: dict) -> dict:
    modified_replacements = replacement_values.copy()
    for key, value in replacement_values.items():
        if isinstance(value, dict):
            modified_replacements[key] = json.dumps(value)
    return modified_replacements


def integration_response_ref_and_data(resource_name: str, replacement_values: dict, file_name: str = "integration-response"):
    ref = k8s.CustomResourceReference(
        "apigatewayv2.services.k8s.aws", "v1alpha1", INTEGRATION_RESPONSE_RESOURCE_PLURAL,
        resource_name, namespace="default",
    )

    sanitized_values = sanitize_dict_for_yaml(replacement_values)

    resource_data = load_apigatewayv2_resource(
        file_name,
        additional_replacements=sanitized_values,
    )
    return ref, resource_data


@pytest.fixture(scope="module")
def api_resource():
    api_resource_name = random_suffix_name("test-api", 32)
    test_resource_values['API_NAME'] = api_resource_name
    api_ref, api_data = helper.api_ref_and_data(
        api_resource_name=api_resource_name, replacement_values=test_resource_values)

    k8s.create_custom_resource(api_ref, api_data)
    time.sleep(CREATE_WAIT_AFTER_SECONDS)
    assert k8s.wait_on_condition(
        api_ref, "ACK.ResourceSynced", "True", wait_periods=10)

    cr = k8s.get_resource(api_ref)
    assert cr is not None
    api_id = cr['status']['apiID']
    test_resource_values['API_ID'] = api_id

    yield api_ref, cr

    k8s.delete_custom_resource(api_ref)


@pytest.fixture(scope="module")
def integration_resource(api_resource):
    integration_resource_name = random_suffix_name("test-integration", 32)
    test_resource_values['INTEGRATION_NAME'] = integration_resource_name
    integration_ref, integration_data = helper.integration_ref_and_data(
        integration_resource_name=integration_resource_name,
        replacement_values=test_resource_values)

    k8s.create_custom_resource(integration_ref, integration_data)
    time.sleep(CREATE_WAIT_AFTER_SECONDS)
    assert k8s.wait_on_condition(
        integration_ref, "ACK.ResourceSynced", "True", wait_periods=10)

    cr = k8s.get_resource(integration_ref)
    assert cr is not None
    integration_id = cr['status']['integrationID']
    test_resource_values['INTEGRATION_ID'] = integration_id

    yield integration_ref, cr

    k8s.delete_custom_resource(integration_ref)


@service_marker
@pytest.mark.canary
class TestIntegrationResponse:
    def test_crud_integration_response(self, api_resource, integration_resource):
        api_ref, api_cr = api_resource
        api_id = api_cr['status']['apiID']
        integration_ref, integration_cr = integration_resource
        integration_id = integration_cr['status']['integrationID']

        test_data = test_resource_values.copy()
        response_name = random_suffix_name("test-integration-response", 32)
        test_data['INTEGRATION_RESPONSE_NAME'] = response_name
        test_data['API_ID'] = api_id
        test_data['INTEGRATION_ID'] = integration_id
        test_data['INTEGRATION_RESPONSE_KEY'] = "/2\\d{2}/"
        test_data['CONTENT_HANDLING_STRATEGY'] = "CONVERT_TO_TEXT"

        # Map keys are strings, values must also be strings
        response_params = {
            "method.response.header.Content-Type": "$integration.response.header.Content-Type"
        }
        test_data['RESPONSE_PARAMETERS'] = response_params

        # Map keys are strings, values must also be strings
        response_templates = {
            "application/json": "{\"message\": \"Hello, $input.params('name')\"}"
        }
        test_data['RESPONSE_TEMPLATES'] = response_templates

        response_ref, response_data = integration_response_ref_and_data(
            resource_name=response_name,
            replacement_values=test_data
        )

        k8s.create_custom_resource(response_ref, response_data)
        time.sleep(CREATE_WAIT_AFTER_SECONDS)
        assert k8s.wait_on_condition(
            response_ref, "ACK.ResourceSynced", "True", wait_periods=10)

        cr = k8s.get_resource(response_ref)
        assert cr is not None
        response_id = cr['status']['integrationResponseID']

        aws_res = apigw_validator.apigatewayv2_client.get_integration_response(
            ApiId=api_id,
            IntegrationId=integration_id,
            IntegrationResponseId=response_id
        )
        assert aws_res is not None
        assert aws_res['IntegrationResponseKey'] == test_data['INTEGRATION_RESPONSE_KEY']
        assert aws_res['ContentHandlingStrategy'] == test_data['CONTENT_HANDLING_STRATEGY']
        assert aws_res['ResponseParameters'] == test_data['RESPONSE_PARAMETERS']
        assert aws_res['ResponseTemplates'] == test_data['RESPONSE_TEMPLATES']

        # Test update
        updated_key = "/2\\d{2}/"
        test_data['INTEGRATION_RESPONSE_KEY'] = updated_key
        updated_data = load_apigatewayv2_resource(
            "integration-response",
            additional_replacements=sanitize_dict_for_yaml(test_data),
        )
        k8s.patch_custom_resource(response_ref, updated_data)
        time.sleep(UPDATE_WAIT_AFTER_SECONDS)
        condition.assert_synced(response_ref)

        # Verify update in AWS
        aws_res = apigw_validator.apigatewayv2_client.get_integration_response(
            ApiId=api_id,
            IntegrationId=integration_id,
            IntegrationResponseId=response_id
        )
        assert aws_res['IntegrationResponseKey'] == updated_key

        # Test delete
        k8s.delete_custom_resource(response_ref)
        time.sleep(DELETE_WAIT_AFTER_SECONDS)
        assert not k8s.get_resource_exists(response_ref)

        # Verify deletion in AWS
        try:
            apigw_validator.apigatewayv2_client.get_integration_response(
                ApiId=api_id,
                IntegrationId=integration_id,
                IntegrationResponseId=response_id
            )
            assert False, "IntegrationResponse should have been deleted"
        except apigw_validator.apigatewayv2_client.exceptions.NotFoundException:
            pass

    def test_integration_response_validation(self, api_resource, integration_resource):
        api_ref, api_cr = api_resource
        api_id = api_cr['status']['apiID']
        integration_ref, integration_cr = integration_resource
        integration_id = integration_cr['status']['integrationID']

        test_data = test_resource_values.copy()
        response_name = random_suffix_name("test-integration-response", 32)
        test_data['INTEGRATION_RESPONSE_NAME'] = response_name
        test_data['API_ID'] = api_id
        test_data['INTEGRATION_ID'] = integration_id
        test_data['INTEGRATION_RESPONSE_KEY'] = "/2\\d{2}/"
        test_data['CONTENT_HANDLING_STRATEGY'] = "INVALID_STRATEGY"
        test_data['RESPONSE_PARAMETERS'] = {}
        test_data['RESPONSE_TEMPLATES'] = {}

        response_ref, response_data = integration_response_ref_and_data(
            resource_name=response_name,
            replacement_values=test_data
        )

        k8s.create_custom_resource(response_ref, response_data)
        time.sleep(CREATE_WAIT_AFTER_SECONDS)

        assert not k8s.wait_on_condition(
            response_ref, "ACK.ResourceSynced", "True", wait_periods=3)

        k8s.delete_custom_resource(response_ref)
