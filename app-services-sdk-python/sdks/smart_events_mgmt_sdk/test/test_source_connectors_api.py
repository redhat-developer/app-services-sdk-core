"""
    Red Hat Openshift SmartEvents Fleet Manager V2

    The API exposed by the fleet manager of the SmartEvents service.  # noqa: E501

    The version of the OpenAPI document: 0.0.1
    Contact: openbridge-dev@redhat.com
    Generated by: https://openapi-generator.tech
"""


import unittest

import rhoas_smart_events_mgmt_sdk
from rhoas_smart_events_mgmt_sdk.api.source_connectors_api import SourceConnectorsApi  # noqa: E501


class TestSourceConnectorsApi(unittest.TestCase):
    """SourceConnectorsApi unit test stubs"""

    def setUp(self):
        self.api = SourceConnectorsApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_source_connectors_api_create_source_connector(self):
        """Test case for source_connectors_api_create_source_connector

        Create a Source Connector for a Bridge instance  # noqa: E501
        """
        pass

    def test_source_connectors_api_delete_source_connector(self):
        """Test case for source_connectors_api_delete_source_connector

        Delete a Source Connector of a Bridge instance  # noqa: E501
        """
        pass

    def test_source_connectors_api_get_source_connector(self):
        """Test case for source_connectors_api_get_source_connector

        Get a Source Connector of a Bridge instance  # noqa: E501
        """
        pass

    def test_source_connectors_api_get_source_connectors(self):
        """Test case for source_connectors_api_get_source_connectors

        Get the list of Sink Connectors for a Bridge  # noqa: E501
        """
        pass

    def test_source_connectors_api_update_source_connector(self):
        """Test case for source_connectors_api_update_source_connector

        Update a Source Connector instance.  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
