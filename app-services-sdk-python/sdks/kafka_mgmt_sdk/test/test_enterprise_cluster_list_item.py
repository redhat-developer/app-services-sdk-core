"""
    Kafka Management API

    Kafka Management API is a REST API to manage Kafka instances  # noqa: E501

    The version of the OpenAPI document: 1.15.0
    Contact: rhosak-support@redhat.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import rhoas_kafka_mgmt_sdk
from rhoas_kafka_mgmt_sdk.model.enterprise_cluster_list_item_all_of import EnterpriseClusterListItemAllOf
from rhoas_kafka_mgmt_sdk.model.object_reference import ObjectReference
globals()['EnterpriseClusterListItemAllOf'] = EnterpriseClusterListItemAllOf
globals()['ObjectReference'] = ObjectReference
from rhoas_kafka_mgmt_sdk.model.enterprise_cluster_list_item import EnterpriseClusterListItem


class TestEnterpriseClusterListItem(unittest.TestCase):
    """EnterpriseClusterListItem unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testEnterpriseClusterListItem(self):
        """Test EnterpriseClusterListItem"""
        # FIXME: construct object with mandatory attributes with example values
        # model = EnterpriseClusterListItem()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()