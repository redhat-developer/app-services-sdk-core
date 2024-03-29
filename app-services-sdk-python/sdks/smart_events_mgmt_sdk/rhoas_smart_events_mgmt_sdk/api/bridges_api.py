"""
    Red Hat Openshift SmartEvents Fleet Manager V2

    The API exposed by the fleet manager of the SmartEvents service.  # noqa: E501

    The version of the OpenAPI document: 0.0.1
    Contact: openbridge-dev@redhat.com
    Generated by: https://openapi-generator.tech
"""


import re  # noqa: F401
import sys  # noqa: F401

from rhoas_smart_events_mgmt_sdk.api_client import ApiClient, Endpoint as _Endpoint
from rhoas_smart_events_mgmt_sdk.model_utils import (  # noqa: F401
    check_allowed_values,
    check_validations,
    date,
    datetime,
    file_type,
    none_type,
    validate_and_convert_types
)
from rhoas_smart_events_mgmt_sdk.model.bridge_list_response import BridgeListResponse
from rhoas_smart_events_mgmt_sdk.model.bridge_request import BridgeRequest
from rhoas_smart_events_mgmt_sdk.model.bridge_response import BridgeResponse
from rhoas_smart_events_mgmt_sdk.model.errors_list import ErrorsList
from rhoas_smart_events_mgmt_sdk.model.managed_resource_status import ManagedResourceStatus


class BridgesApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client
        self.bridges_api_create_bridge_endpoint = _Endpoint(
            settings={
                'response_type': (BridgeResponse,),
                'auth': [
                    'bearer'
                ],
                'endpoint_path': '/api/smartevents_mgmt/v2/bridges',
                'operation_id': 'bridges_api_create_bridge',
                'http_method': 'POST',
                'servers': None,
            },
            params_map={
                'all': [
                    'bridge_request',
                ],
                'required': [],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'bridge_request':
                        (BridgeRequest,),
                },
                'attribute_map': {
                },
                'location_map': {
                    'bridge_request': 'body',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [
                    'application/json'
                ]
            },
            api_client=api_client
        )
        self.bridges_api_delete_bridge_endpoint = _Endpoint(
            settings={
                'response_type': None,
                'auth': [
                    'bearer'
                ],
                'endpoint_path': '/api/smartevents_mgmt/v2/bridges/{bridgeId}',
                'operation_id': 'bridges_api_delete_bridge',
                'http_method': 'DELETE',
                'servers': None,
            },
            params_map={
                'all': [
                    'bridge_id',
                ],
                'required': [
                    'bridge_id',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'bridge_id':
                        (str,),
                },
                'attribute_map': {
                    'bridge_id': 'bridgeId',
                },
                'location_map': {
                    'bridge_id': 'path',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client
        )
        self.bridges_api_get_bridge_endpoint = _Endpoint(
            settings={
                'response_type': (BridgeResponse,),
                'auth': [
                    'bearer'
                ],
                'endpoint_path': '/api/smartevents_mgmt/v2/bridges/{bridgeId}',
                'operation_id': 'bridges_api_get_bridge',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'bridge_id',
                ],
                'required': [
                    'bridge_id',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                    'bridge_id',
                ]
            },
            root_map={
                'validations': {
                    ('bridge_id',): {

                        'min_length': 1,
                    },
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'bridge_id':
                        (str,),
                },
                'attribute_map': {
                    'bridge_id': 'bridgeId',
                },
                'location_map': {
                    'bridge_id': 'path',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client
        )
        self.bridges_api_get_bridges_endpoint = _Endpoint(
            settings={
                'response_type': (BridgeListResponse,),
                'auth': [
                    'bearer'
                ],
                'endpoint_path': '/api/smartevents_mgmt/v2/bridges',
                'operation_id': 'bridges_api_get_bridges',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'name',
                    'page',
                    'size',
                    'status',
                ],
                'required': [],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                    'page',
                    'size',
                    'status',
                ]
            },
            root_map={
                'validations': {
                    ('page',): {

                        'inclusive_minimum': 0,
                    },
                    ('size',): {

                        'inclusive_maximum': 100,
                        'inclusive_minimum': 1,
                    },
                    ('status',): {

                    },
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'name':
                        (str,),
                    'page':
                        (int,),
                    'size':
                        (int,),
                    'status':
                        ([ManagedResourceStatus],),
                },
                'attribute_map': {
                    'name': 'name',
                    'page': 'page',
                    'size': 'size',
                    'status': 'status',
                },
                'location_map': {
                    'name': 'query',
                    'page': 'query',
                    'size': 'query',
                    'status': 'query',
                },
                'collection_format_map': {
                    'status': 'multi',
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client
        )

    def bridges_api_create_bridge(
        self,
        **kwargs
    ):
        """Create a Bridge instance  # noqa: E501

        Create a Bridge instance for the authenticated user.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.bridges_api_create_bridge(async_req=True)
        >>> result = thread.get()


        Keyword Args:
            bridge_request (BridgeRequest): [optional]
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _spec_property_naming (bool): True if the variable names in the input data
                are serialized names, as specified in the OpenAPI document.
                False if the variable names in the input data
                are pythonic names, e.g. snake case (default)
            _content_type (str/None): force body content-type.
                Default is None and content-type will be predicted by allowed
                content-types and body.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            _request_auths (list): set to override the auth_settings for an a single
                request; this effectively ignores the authentication
                in the spec for a single request.
                Default is None
            async_req (bool): execute request asynchronously

        Returns:
            BridgeResponse
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_spec_property_naming'] = kwargs.get(
            '_spec_property_naming', False
        )
        kwargs['_content_type'] = kwargs.get(
            '_content_type')
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['_request_auths'] = kwargs.get('_request_auths', None)
        return self.bridges_api_create_bridge_endpoint.call_with_http_info(**kwargs)

    def bridges_api_delete_bridge(
        self,
        bridge_id,
        **kwargs
    ):
        """Delete a Bridge instance  # noqa: E501

        Delete a Bridge instance of the authenticated user by ID.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.bridges_api_delete_bridge(bridge_id, async_req=True)
        >>> result = thread.get()

        Args:
            bridge_id (str):

        Keyword Args:
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _spec_property_naming (bool): True if the variable names in the input data
                are serialized names, as specified in the OpenAPI document.
                False if the variable names in the input data
                are pythonic names, e.g. snake case (default)
            _content_type (str/None): force body content-type.
                Default is None and content-type will be predicted by allowed
                content-types and body.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            _request_auths (list): set to override the auth_settings for an a single
                request; this effectively ignores the authentication
                in the spec for a single request.
                Default is None
            async_req (bool): execute request asynchronously

        Returns:
            None
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_spec_property_naming'] = kwargs.get(
            '_spec_property_naming', False
        )
        kwargs['_content_type'] = kwargs.get(
            '_content_type')
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['_request_auths'] = kwargs.get('_request_auths', None)
        kwargs['bridge_id'] = \
            bridge_id
        return self.bridges_api_delete_bridge_endpoint.call_with_http_info(**kwargs)

    def bridges_api_get_bridge(
        self,
        bridge_id,
        **kwargs
    ):
        """Get a Bridge instance  # noqa: E501

        Get a Bridge instance of the authenticated user by ID.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.bridges_api_get_bridge(bridge_id, async_req=True)
        >>> result = thread.get()

        Args:
            bridge_id (str):

        Keyword Args:
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _spec_property_naming (bool): True if the variable names in the input data
                are serialized names, as specified in the OpenAPI document.
                False if the variable names in the input data
                are pythonic names, e.g. snake case (default)
            _content_type (str/None): force body content-type.
                Default is None and content-type will be predicted by allowed
                content-types and body.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            _request_auths (list): set to override the auth_settings for an a single
                request; this effectively ignores the authentication
                in the spec for a single request.
                Default is None
            async_req (bool): execute request asynchronously

        Returns:
            BridgeResponse
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_spec_property_naming'] = kwargs.get(
            '_spec_property_naming', False
        )
        kwargs['_content_type'] = kwargs.get(
            '_content_type')
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['_request_auths'] = kwargs.get('_request_auths', None)
        kwargs['bridge_id'] = \
            bridge_id
        return self.bridges_api_get_bridge_endpoint.call_with_http_info(**kwargs)

    def bridges_api_get_bridges(
        self,
        **kwargs
    ):
        """Get the list of Bridge instances  # noqa: E501

        Get the list of Bridge instances for the authenticated user.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.bridges_api_get_bridges(async_req=True)
        >>> result = thread.get()


        Keyword Args:
            name (str): [optional]
            page (int): [optional] if omitted the server will use the default value of 0
            size (int): [optional] if omitted the server will use the default value of 100
            status ([ManagedResourceStatus]): [optional]
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _spec_property_naming (bool): True if the variable names in the input data
                are serialized names, as specified in the OpenAPI document.
                False if the variable names in the input data
                are pythonic names, e.g. snake case (default)
            _content_type (str/None): force body content-type.
                Default is None and content-type will be predicted by allowed
                content-types and body.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            _request_auths (list): set to override the auth_settings for an a single
                request; this effectively ignores the authentication
                in the spec for a single request.
                Default is None
            async_req (bool): execute request asynchronously

        Returns:
            BridgeListResponse
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_spec_property_naming'] = kwargs.get(
            '_spec_property_naming', False
        )
        kwargs['_content_type'] = kwargs.get(
            '_content_type')
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['_request_auths'] = kwargs.get('_request_auths', None)
        return self.bridges_api_get_bridges_endpoint.call_with_http_info(**kwargs)

