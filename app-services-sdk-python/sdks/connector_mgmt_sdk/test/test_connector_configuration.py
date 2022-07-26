"""
    Connector Management API

    Connector Management API is a REST API to manage connectors.  # noqa: E501

    The version of the OpenAPI document: 0.1.0
    Contact: rhosak-support@redhat.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import rhoas_connector_mgmt_sdk
from rhoas_connector_mgmt_sdk.model.kafka_connection_settings import KafkaConnectionSettings
from rhoas_connector_mgmt_sdk.model.schema_registry_connection_settings import SchemaRegistryConnectionSettings
from rhoas_connector_mgmt_sdk.model.service_account import ServiceAccount
globals()['KafkaConnectionSettings'] = KafkaConnectionSettings
globals()['SchemaRegistryConnectionSettings'] = SchemaRegistryConnectionSettings
globals()['ServiceAccount'] = ServiceAccount
from rhoas_connector_mgmt_sdk.model.connector_configuration import ConnectorConfiguration


class TestConnectorConfiguration(unittest.TestCase):
    """ConnectorConfiguration unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testConnectorConfiguration(self):
        """Test ConnectorConfiguration"""
        # FIXME: construct object with mandatory attributes with example values
        # model = ConnectorConfiguration()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
