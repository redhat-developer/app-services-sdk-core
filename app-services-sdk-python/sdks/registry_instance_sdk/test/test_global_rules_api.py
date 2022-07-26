"""
    Service Registry API

    Service Registry Instance API  NOTE: This API cannot be called directly from the portal.  # noqa: E501

    The version of the OpenAPI document: 2.2.5.Final
    Contact: apicurio@lists.jboss.org
    Generated by: https://openapi-generator.tech
"""


import unittest

import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.api.global_rules_api import GlobalRulesApi  # noqa: E501


class TestGlobalRulesApi(unittest.TestCase):
    """GlobalRulesApi unit test stubs"""

    def setUp(self):
        self.api = GlobalRulesApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_create_global_rule(self):
        """Test case for create_global_rule

        Create global rule  # noqa: E501
        """
        pass

    def test_delete_all_global_rules(self):
        """Test case for delete_all_global_rules

        Delete all global rules  # noqa: E501
        """
        pass

    def test_delete_global_rule(self):
        """Test case for delete_global_rule

        Delete global rule  # noqa: E501
        """
        pass

    def test_get_global_rule_config(self):
        """Test case for get_global_rule_config

        Get global rule configuration  # noqa: E501
        """
        pass

    def test_list_global_rules(self):
        """Test case for list_global_rules

        List global rules  # noqa: E501
        """
        pass

    def test_update_global_rule_config(self):
        """Test case for update_global_rule_config

        Update global rule configuration  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
