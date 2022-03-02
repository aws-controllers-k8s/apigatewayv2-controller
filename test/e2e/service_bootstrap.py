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
"""Bootstraps the resources required to run the APIGatewayV2 integration tests.
"""
import logging
import os
from zipfile import ZipFile
import tempfile

import boto3

from acktest.bootstrapping import Resources, BootstrapFailureException
from acktest.bootstrapping.iam import Role
from acktest.bootstrapping.vpc import VPC
from acktest.resources import random_suffix_name
from acktest.aws.identity import get_region
from e2e import bootstrap_directory
from e2e.bootstrap_resources import BootstrapResources


def service_bootstrap() -> Resources:
    logging.getLogger().setLevel(logging.INFO)

    # First create the AuthorizerRole and VPC.
    # Then use the created AuthorizerRole.arn for
    # creating Authorizer lambda function.

    resources = BootstrapResources(
        AuthorizerRole=Role("ack-apigwv2-authorizer-role",
                            "lambda.amazonaws.com",
                            managed_policies=["arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"]),
        VPC=VPC(name_prefix="apigwv2-vpc-link")
    )

    try:
        resources.bootstrap()
    except BootstrapFailureException as ex:
        exit(254)

    authorizer_function_name = random_suffix_name("ack-apigwv2-authorizer", 30)
    authorizer_function_arn = create_lambda_authorizer(authorizer_function_name, resources.AuthorizerRole.arn)

    resources.AuthorizerFunctionName = authorizer_function_name
    resources.AuthorizerFunctionArn = authorizer_function_arn
    return resources


def create_lambda_authorizer(authorizer_function_name: str, authorizer_role_arn: str) -> str:
    region = get_region()
    lambda_client = boto3.client("lambda", region)

    with tempfile.TemporaryDirectory() as tempdir:
        current_directory = os.path.dirname(os.path.realpath(__file__))
        index_zip = ZipFile(f'{tempdir}/index.zip', 'w')
        index_zip.write(f'{current_directory}/resources/index.js', 'index.js')
        index_zip.close()

        with open(f'{tempdir}/index.zip', 'rb') as f:
            b64_encoded_zip_file = f.read()

        response = lambda_client.create_function(
            FunctionName=authorizer_function_name,
            Role=authorizer_role_arn,
            Handler='index.handler',
            Runtime='nodejs12.x',
            Code={'ZipFile': b64_encoded_zip_file}
        )

    return response['FunctionArn']


if __name__ == "__main__":
    config = service_bootstrap()
    # Write config to current directory by default
    config.serialize(bootstrap_directory)
