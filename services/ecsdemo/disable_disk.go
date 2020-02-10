package ecsdemo

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

// DisableDisk invokes the ecsdemo.DisableDisk API synchronously
// api document: https://help.aliyun.com/api/ecsdemo/disabledisk.html
func (client *Client) DisableDisk(request *DisableDiskRequest) (response *DisableDiskResponse, err error) {
	response = CreateDisableDiskResponse()
	err = client.DoAction(request, response)
	return
}

// DisableDiskWithChan invokes the ecsdemo.DisableDisk API asynchronously
// api document: https://help.aliyun.com/api/ecsdemo/disabledisk.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableDiskWithChan(request *DisableDiskRequest) (<-chan *DisableDiskResponse, <-chan error) {
	responseChan := make(chan *DisableDiskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableDisk(request)
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

// DisableDiskWithCallback invokes the ecsdemo.DisableDisk API asynchronously
// api document: https://help.aliyun.com/api/ecsdemo/disabledisk.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableDiskWithCallback(request *DisableDiskRequest, callback func(response *DisableDiskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableDiskResponse
		var err error
		defer close(result)
		response, err = client.DisableDisk(request)
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

// DisableDiskRequest is the request struct for api DisableDisk
type DisableDiskRequest struct {
	*requests.RpcRequest
	TypeCode string `position:"Query" name:"TypeCode"`
}

// DisableDiskResponse is the response struct for api DisableDisk
type DisableDiskResponse struct {
	*responses.BaseResponse
	Success bool `json:"Success" xml:"Success"`
	Data    Data `json:"Data" xml:"Data"`
}

// CreateDisableDiskRequest creates a request to invoke DisableDisk API
func CreateDisableDiskRequest() (request *DisableDiskRequest) {
	request = &DisableDiskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EcsDemo", "2019-06-20", "DisableDisk", "", "")
	return
}

// CreateDisableDiskResponse creates a response to parse from DisableDisk response
func CreateDisableDiskResponse() (response *DisableDiskResponse) {
	response = &DisableDiskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
