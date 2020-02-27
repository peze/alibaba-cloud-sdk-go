package r_kvstore

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ReleaseInstances invokes the r_kvstore.ReleaseInstances API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/releaseinstances.html
func (client *Client) ReleaseInstances(request *ReleaseInstancesRequest) (response *ReleaseInstancesResponse, err error) {
	response = CreateReleaseInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// ReleaseInstancesWithChan invokes the r_kvstore.ReleaseInstances API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/releaseinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseInstancesWithChan(request *ReleaseInstancesRequest) (<-chan *ReleaseInstancesResponse, <-chan error) {
	responseChan := make(chan *ReleaseInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseInstances(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ReleaseInstancesWithCallback invokes the r_kvstore.ReleaseInstances API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/releaseinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseInstancesWithCallback(request *ReleaseInstancesRequest, callback func(response *ReleaseInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseInstancesResponse
		var err error
		defer close(result)
		response, err = client.ReleaseInstances(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ReleaseInstancesRequest is the request struct for api ReleaseInstances
type ReleaseInstancesRequest struct {
	*requests.RpcRequest
	Async        requests.Boolean `position:"Query" name:"async"`
	InstanceName string           `position:"Query" name:"instanceName"`
	InstanceID   requests.Integer `position:"Query" name:"instanceID"`
	Channel      string           `position:"Query" name:"channel"`
	Force        requests.Boolean `position:"Query" name:"force"`
	AliUID       requests.Integer `position:"Query" name:"aliUID"`
	ResourceName string           `position:"Query" name:"resourceName"`
	Operator     string           `position:"Query" name:"operator"`
}

// ReleaseInstancesResponse is the response struct for api ReleaseInstances
type ReleaseInstancesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReleaseInstancesRequest creates a request to invoke ReleaseInstances API
func CreateReleaseInstancesRequest() (request *ReleaseInstancesRequest) {
	request = &ReleaseInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "ReleaseInstances", "", "")
	return
}

// CreateReleaseInstancesResponse creates a response to parse from ReleaseInstances response
func CreateReleaseInstancesResponse() (response *ReleaseInstancesResponse) {
	response = &ReleaseInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}