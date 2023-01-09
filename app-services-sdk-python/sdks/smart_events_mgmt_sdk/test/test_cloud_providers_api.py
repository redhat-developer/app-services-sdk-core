"""
    Red Hat Openshift SmartEvents Fleet Manager V2

    The API exposed by the fleet manager of the SmartEvents service.  # noqa: E501

    The version of the OpenAPI document: 0.0.1
    Contact: openbridge-dev@redhat.com
    Generated by: https://openapi-generator.tech
"""


import unittest

import rhoas_smart_events_mgmt_sdk
from rhoas_smart_events_mgmt_sdk.api.cloud_providers_api import CloudProvidersApi  # noqa: E501


class TestCloudProvidersApi(unittest.TestCase):
    """CloudProvidersApi unit test stubs"""

    def setUp(self):
        self.api = CloudProvidersApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_cloud_provider_api_get_cloud_provider(self):
        """Test case for cloud_provider_api_get_cloud_provider

        Get Cloud Provider.  # noqa: E501
        """
        pass

    def test_cloud_provider_api_list_cloud_provider_regions(self):
        """Test case for cloud_provider_api_list_cloud_provider_regions

        List Supported Cloud Regions.  # noqa: E501
        """
        pass

    def test_cloud_provider_api_list_cloud_providers(self):
        """Test case for cloud_provider_api_list_cloud_providers

        List Supported Cloud Providers.  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()