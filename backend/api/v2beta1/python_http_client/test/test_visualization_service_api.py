# coding: utf-8

"""
    Kubeflow Pipelines API

    This file contains REST API specification for Kubeflow Pipelines. The file is autogenerated from the swagger definition.

    Contact: kubeflow-pipelines@google.com
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest

import kfp_server_api
from kfp_server_api.api.visualization_service_api import VisualizationServiceApi  # noqa: E501
from kfp_server_api.rest import ApiException


class TestVisualizationServiceApi(unittest.TestCase):
    """VisualizationServiceApi unit test stubs"""

    def setUp(self):
        self.api = kfp_server_api.api.visualization_service_api.VisualizationServiceApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_visualization_service_create_visualization_v1(self):
        """Test case for visualization_service_create_visualization_v1

        """
        pass


if __name__ == '__main__':
    unittest.main()
