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

// CountRedisaInstances invokes the r_kvstore.CountRedisaInstances API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/countredisainstances.html
func (client *Client) CountRedisaInstances(request *CountRedisaInstancesRequest) (response *CountRedisaInstancesResponse, err error) {
	response = CreateCountRedisaInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// CountRedisaInstancesWithChan invokes the r_kvstore.CountRedisaInstances API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/countredisainstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CountRedisaInstancesWithChan(request *CountRedisaInstancesRequest) (<-chan *CountRedisaInstancesResponse, <-chan error) {
	responseChan := make(chan *CountRedisaInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CountRedisaInstances(request)
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

// CountRedisaInstancesWithCallback invokes the r_kvstore.CountRedisaInstances API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/countredisainstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CountRedisaInstancesWithCallback(request *CountRedisaInstancesRequest, callback func(response *CountRedisaInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CountRedisaInstancesResponse
		var err error
		defer close(result)
		response, err = client.CountRedisaInstances(request)
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

// CountRedisaInstancesRequest is the request struct for api CountRedisaInstances
type CountRedisaInstancesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CountRedisaInstancesResponse is the response struct for api CountRedisaInstances
type CountRedisaInstancesResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	InstancesCount int    `json:"InstancesCount" xml:"InstancesCount"`
}

// CreateCountRedisaInstancesRequest creates a request to invoke CountRedisaInstances API
func CreateCountRedisaInstancesRequest() (request *CountRedisaInstancesRequest) {
	request = &CountRedisaInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "CountRedisaInstances", "", "")
	return
}

// CreateCountRedisaInstancesResponse creates a response to parse from CountRedisaInstances response
func CreateCountRedisaInstancesResponse() (response *CountRedisaInstancesResponse) {
	response = &CountRedisaInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}