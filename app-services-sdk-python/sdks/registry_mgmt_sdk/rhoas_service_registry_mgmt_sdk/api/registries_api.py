"""
    Service Registry Management API

    Service Registry Management API is a REST API for managing Service Registry instances. Service Registry is a datastore for event schemas and API designs, which is based on the open source Apicurio Registry project.  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Contact: rhosak-eval-support@redhat.com
    Generated by: https://openapi-generator.tech
"""


import re  # noqa: F401
import sys  # noqa: F401

from rhoas_service_registry_mgmt_sdk.api_client import ApiClient, Endpoint as _Endpoint
from rhoas_service_registry_mgmt_sdk.model_utils import (  # noqa: F401
    check_allowed_values,
    check_validations,
    date,
    datetime,
    file_type,
    none_type,
    validate_and_convert_types
)
from rhoas_service_registry_mgmt_sdk.model.empty import Empty
from rhoas_service_registry_mgmt_sdk.model.error import Error
from rhoas_service_registry_mgmt_sdk.model.registry import Registry
from rhoas_service_registry_mgmt_sdk.model.registry_create import RegistryCreate
from rhoas_service_registry_mgmt_sdk.model.registry_list import RegistryList


class RegistriesApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client
        self.create_registry_endpoint = _Endpoint(
            settings={
                'response_type': (Registry,),
                'auth': [
                    'Bearer'
                ],
                'endpoint_path': '/api/serviceregistry_mgmt/v1/registries',
                'operation_id': 'create_registry',
                'http_method': 'POST',
                'servers': None,
            },
            params_map={
                'all': [
                    'registry_create',
                ],
                'required': [
                    'registry_create',
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
                    'registry_create':
                        (RegistryCreate,),
                },
                'attribute_map': {
                },
                'location_map': {
                    'registry_create': 'body',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json',
                    '*/*'
                ],
                'content_type': [
                    'application/json'
                ]
            },
            api_client=api_client
        )
        self.delete_registry_endpoint = _Endpoint(
            settings={
                'response_type': None,
                'auth': [
                    'Bearer'
                ],
                'endpoint_path': '/api/serviceregistry_mgmt/v1/registries/{id}',
                'operation_id': 'delete_registry',
                'http_method': 'DELETE',
                'servers': None,
            },
            params_map={
                'all': [
                    'id',
                ],
                'required': [
                    'id',
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
                    'id':
                        (str,),
                },
                'attribute_map': {
                    'id': 'id',
                },
                'location_map': {
                    'id': 'path',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    '*/*',
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client
        )
        self.get_registries_endpoint = _Endpoint(
            settings={
                'response_type': (RegistryList,),
                'auth': [
                    'Bearer'
                ],
                'endpoint_path': '/api/serviceregistry_mgmt/v1/registries',
                'operation_id': 'get_registries',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'page',
                    'size',
                    'order_by',
                    'search',
                ],
                'required': [],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                    'page',
                    'size',
                ]
            },
            root_map={
                'validations': {
                    ('page',): {

                        'inclusive_minimum': 0,
                    },
                    ('size',): {

                        'inclusive_maximum': 500,
                        'inclusive_minimum': 1,
                    },
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'page':
                        (int,),
                    'size':
                        (int,),
                    'order_by':
                        (str,),
                    'search':
                        (str,),
                },
                'attribute_map': {
                    'page': 'page',
                    'size': 'size',
                    'order_by': 'orderBy',
                    'search': 'search',
                },
                'location_map': {
                    'page': 'query',
                    'size': 'query',
                    'order_by': 'query',
                    'search': 'query',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json',
                    '*/*'
                ],
                'content_type': [],
            },
            api_client=api_client
        )
        self.get_registry_endpoint = _Endpoint(
            settings={
                'response_type': (Registry,),
                'auth': [
                    'Bearer'
                ],
                'endpoint_path': '/api/serviceregistry_mgmt/v1/registries/{id}',
                'operation_id': 'get_registry',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'id',
                ],
                'required': [
                    'id',
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
                    'id':
                        (str,),
                },
                'attribute_map': {
                    'id': 'id',
                },
                'location_map': {
                    'id': 'path',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json',
                    '*/*'
                ],
                'content_type': [],
            },
            api_client=api_client
        )

    def create_registry(
        self,
        registry_create,
        **kwargs
    ):
        """create_registry  # noqa: E501

        Create a new Registry instance  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.create_registry(registry_create, async_req=True)
        >>> result = thread.get()

        Args:
            registry_create (RegistryCreate): A new `Registry` instance to be created.

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
            Registry
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
        kwargs['registry_create'] = \
            registry_create
        return self.create_registry_endpoint.call_with_http_info(**kwargs)

    def delete_registry(
        self,
        id,
        **kwargs
    ):
        """Delete a Registry instance  # noqa: E501

        Deletes an existing `Registry` instance and all of the data that it stores. Important: Users should export the registry data before deleting the instance, e.g., using the Service Registry web console, core REST API, or `rhoas` CLI.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.delete_registry(id, async_req=True)
        >>> result = thread.get()

        Args:
            id (str): The id of the object you wish to interact with.

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
        kwargs['id'] = \
            id
        return self.delete_registry_endpoint.call_with_http_info(**kwargs)

    def get_registries(
        self,
        **kwargs
    ):
        """get_registries  # noqa: E501

        Get the list of all Registry instances  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.get_registries(async_req=True)
        >>> result = thread.get()


        Keyword Args:
            page (int): Page index.. [optional]
            size (int): Number of items in each page.. [optional]
            order_by (str): Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement. Each query can be ordered by any of the request fields. For example, to retrieve all Registry instances ordered by their name:  ```sql name asc ```  Or to retrieve all Registry instances ordered by their name _and_ created date:  ```sql name asc, created_at asc ```  If the parameter isn't provided, or if the value is empty,  the results are ordered by name.. [optional]
            search (str): Search criteria.  The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement. Allowed fields in the search are: `name`, `status`. Allowed comparators are `=` or `LIKE`. Allowed joins are `AND` and `OR`, however there is a limit of max 10 joins in the search query.  Examples:  To retrieve a request with name equal `my-registry`, the value should be:  ``` name = my-registry  ```  To retrieve a request with its name starting with `my`, the value should be:  ``` name like my%25 ```  If the parameter isn't provided, or if the value is empty, all the Registry instances that the user has permission to see are returned.  Note: If the query is invalid, an error is returned. . [optional]
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
            RegistryList
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
        return self.get_registries_endpoint.call_with_http_info(**kwargs)

    def get_registry(
        self,
        id,
        **kwargs
    ):
        """Get a Registry instance  # noqa: E501

        Gets the details of a single instance of a `Registry`.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.get_registry(id, async_req=True)
        >>> result = thread.get()

        Args:
            id (str): The id of the object you wish to interact with.

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
            Registry
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
        kwargs['id'] = \
            id
        return self.get_registry_endpoint.call_with_http_info(**kwargs)

