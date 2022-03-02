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

"""Cleans up the resources created by the bootstrapping process.
"""

import boto3
import logging

from acktest.bootstrapping import Resources
from acktest.aws.identity import get_region

from e2e import bootstrap_directory


def service_cleanup():
    logging.getLogger().setLevel(logging.INFO)

    resources = Resources.deserialize(bootstrap_directory)
    delete_authorizer_function(function_name=resources.AuthorizerFunctionName)
    resources.cleanup()


def delete_authorizer_function(function_name: str):
    region = get_region()
    lambda_client = boto3.client("lambda", region_name=region)

    try:
        lambda_client.delete_function(FunctionName=function_name)
    except lambda_client.exceptions.ResourceNotFoundException:
        pass


if __name__ == "__main__":
    service_cleanup()
