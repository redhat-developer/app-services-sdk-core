# rhoas_registry_instance_sdk.UsersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_current_user_info**](UsersApi.md#get_current_user_info) | **GET** /users/me | Get current user


# **get_current_user_info**
> UserInfo get_current_user_info()

Get current user

Returns information about the currently authenticated user.

### Example


```python
import time
import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.api import users_api
from rhoas_registry_instance_sdk.model.user_info import UserInfo
from rhoas_registry_instance_sdk.model.error import Error
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_registry_instance_sdk.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with rhoas_registry_instance_sdk.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = users_api.UsersApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Get current user
        api_response = api_instance.get_current_user_info()
        pprint(api_response)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling UsersApi->get_current_user_info: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**UserInfo**](UserInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Response when the endpoint is successfully invoked. |  -  |
**500** | Common response for all operations that can fail with an unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

