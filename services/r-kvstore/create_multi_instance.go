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

// CreateMultiInstance invokes the r_kvstore.CreateMultiInstance API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/createmultiinstance.html
func (client *Client) CreateMultiInstance(request *CreateMultiInstanceRequest) (response *CreateMultiInstanceResponse, err error) {
	response = CreateCreateMultiInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMultiInstanceWithChan invokes the r_kvstore.CreateMultiInstance API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/createmultiinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMultiInstanceWithChan(request *CreateMultiInstanceRequest) (<-chan *CreateMultiInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateMultiInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMultiInstance(request)
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

// CreateMultiInstanceWithCallback invokes the r_kvstore.CreateMultiInstance API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/createmultiinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMultiInstanceWithCallback(request *CreateMultiInstanceRequest, callback func(response *CreateMultiInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMultiInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateMultiInstance(request)
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

// CreateMultiInstanceRequest is the request struct for api CreateMultiInstance
type CreateMultiInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	KVStoreInstances     string           `position:"Query" name:"KVStoreInstances"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Token                string           `position:"Query" name:"Token"`
}

// CreateMultiInstanceResponse is the response struct for api CreateMultiInstance
type CreateMultiInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

// CreateCreateMultiInstanceRequest creates a request to invoke CreateMultiInstance API
func CreateCreateMultiInstanceRequest() (request *CreateMultiInstanceRequest) {
	request = &CreateMultiInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "CreateMultiInstance", "", "")
	return
}

// CreateCreateMultiInstanceResponse creates a response to parse from CreateMultiInstance response
func CreateCreateMultiInstanceResponse() (response *CreateMultiInstanceResponse) {
	response = &CreateMultiInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}