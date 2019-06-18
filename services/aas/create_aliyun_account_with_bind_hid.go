package aas

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

// CreateAliyunAccountWithBindHid invokes the aas.CreateAliyunAccountWithBindHid API synchronously
// api document: https://help.aliyun.com/api/aas/createaliyunaccountwithbindhid.html
func (client *Client) CreateAliyunAccountWithBindHid(request *CreateAliyunAccountWithBindHidRequest) (response *CreateAliyunAccountWithBindHidResponse, err error) {
	response = CreateCreateAliyunAccountWithBindHidResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAliyunAccountWithBindHidWithChan invokes the aas.CreateAliyunAccountWithBindHid API asynchronously
// api document: https://help.aliyun.com/api/aas/createaliyunaccountwithbindhid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAliyunAccountWithBindHidWithChan(request *CreateAliyunAccountWithBindHidRequest) (<-chan *CreateAliyunAccountWithBindHidResponse, <-chan error) {
	responseChan := make(chan *CreateAliyunAccountWithBindHidResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAliyunAccountWithBindHid(request)
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

// CreateAliyunAccountWithBindHidWithCallback invokes the aas.CreateAliyunAccountWithBindHid API asynchronously
// api document: https://help.aliyun.com/api/aas/createaliyunaccountwithbindhid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAliyunAccountWithBindHidWithCallback(request *CreateAliyunAccountWithBindHidRequest, callback func(response *CreateAliyunAccountWithBindHidResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAliyunAccountWithBindHidResponse
		var err error
		defer close(result)
		response, err = client.CreateAliyunAccountWithBindHid(request)
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

// CreateAliyunAccountWithBindHidRequest is the request struct for api CreateAliyunAccountWithBindHid
type CreateAliyunAccountWithBindHidRequest struct {
	*requests.RpcRequest
	InnerAccountHid string `position:"Query" name:"InnerAccountHid"`
}

// CreateAliyunAccountWithBindHidResponse is the response struct for api CreateAliyunAccountWithBindHid
type CreateAliyunAccountWithBindHidResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	PK        string `json:"PK" xml:"PK"`
}

// CreateCreateAliyunAccountWithBindHidRequest creates a request to invoke CreateAliyunAccountWithBindHid API
func CreateCreateAliyunAccountWithBindHidRequest() (request *CreateAliyunAccountWithBindHidRequest) {
	request = &CreateAliyunAccountWithBindHidRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Aas", "2015-07-01", "CreateAliyunAccountWithBindHid", "aas", "openAPI")
	return
}

// CreateCreateAliyunAccountWithBindHidResponse creates a response to parse from CreateAliyunAccountWithBindHid response
func CreateCreateAliyunAccountWithBindHidResponse() (response *CreateAliyunAccountWithBindHidResponse) {
	response = &CreateAliyunAccountWithBindHidResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
