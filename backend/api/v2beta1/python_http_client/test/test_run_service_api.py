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
from kfp_server_api.api.run_service_api import RunServiceApi  # noqa: E501
from kfp_server_api.rest import ApiException


class TestRunServiceApi(unittest.TestCase):
    """RunServiceApi unit test stubs"""

    def setUp(self):
        self.api = kfp_server_api.api.run_service_api.RunServiceApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_run_service_archive_run(self):
        """Test case for run_service_archive_run

        Archives a run in an experiment given by run ID and experiment ID.  # noqa: E501
        """
        pass

    def test_run_service_create_run(self):
        """Test case for run_service_create_run

        Creates a new run in an experiment specified by experiment ID.  If experiment ID is not specified, the run is created in the default experiment.  # noqa: E501
        """
        pass

    def test_run_service_delete_run(self):
        """Test case for run_service_delete_run

        Deletes a run in an experiment given by run ID and experiment ID.  # noqa: E501
        """
        pass

    def test_run_service_get_run(self):
        """Test case for run_service_get_run

        Finds a specific run by ID.  # noqa: E501
        """
        pass

    def test_run_service_list_runs(self):
        """Test case for run_service_list_runs

        Finds all runs in an experiment given by experiment ID.  If experiment id is not specified, finds all runs across all experiments.  # noqa: E501
        """
        pass

    def test_run_service_read_artifact(self):
        """Test case for run_service_read_artifact

        Finds artifact data in a run.  # noqa: E501
        """
        pass

    def test_run_service_retry_run(self):
        """Test case for run_service_retry_run

        Re-initiates a failed or terminated run.  # noqa: E501
        """
        pass

    def test_run_service_terminate_run(self):
        """Test case for run_service_terminate_run

        Terminates an active run.  # noqa: E501
        """
        pass

    def test_run_service_unarchive_run(self):
        """Test case for run_service_unarchive_run

        Restores an archived run in an experiment given by run ID and experiment ID.  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
