"""
    Service Registry API

    Service Registry Instance API  NOTE: This API cannot be called directly from the portal.  # noqa: E501

    The version of the OpenAPI document: 2.2.5.Final
    Contact: apicurio@lists.jboss.org
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.model.artifact_reference import ArtifactReference
from rhoas_registry_instance_sdk.model.artifact_state import ArtifactState
from rhoas_registry_instance_sdk.model.artifact_type import ArtifactType
from rhoas_registry_instance_sdk.model.properties import Properties
globals()['ArtifactReference'] = ArtifactReference
globals()['ArtifactState'] = ArtifactState
globals()['ArtifactType'] = ArtifactType
globals()['Properties'] = Properties
from rhoas_registry_instance_sdk.model.artifact_meta_data import ArtifactMetaData


class TestArtifactMetaData(unittest.TestCase):
    """ArtifactMetaData unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testArtifactMetaData(self):
        """Test ArtifactMetaData"""
        # FIXME: construct object with mandatory attributes with example values
        # model = ArtifactMetaData()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
