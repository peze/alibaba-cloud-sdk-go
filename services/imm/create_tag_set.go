package imm

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

// CreateTagSet invokes the imm.CreateTagSet API synchronously
// api document: https://help.aliyun.com/api/imm/createtagset.html
func (client *Client) CreateTagSet(request *CreateTagSetRequest) (response *CreateTagSetResponse, err error) {
	response = CreateCreateTagSetResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTagSetWithChan invokes the imm.CreateTagSet API asynchronously
// api document: https://help.aliyun.com/api/imm/createtagset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTagSetWithChan(request *CreateTagSetRequest) (<-chan *CreateTagSetResponse, <-chan error) {
	responseChan := make(chan *CreateTagSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTagSet(request)
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

// CreateTagSetWithCallback invokes the imm.CreateTagSet API asynchronously
// api document: https://help.aliyun.com/api/imm/createtagset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTagSetWithCallback(request *CreateTagSetRequest, callback func(response *CreateTagSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTagSetResponse
		var err error
		defer close(result)
		response, err = client.CreateTagSet(request)
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

// CreateTagSetRequest is the request struct for api CreateTagSet
type CreateTagSetRequest struct {
	*requests.RpcRequest
	Project string `position:"Query" name:"Project"`
}

// CreateTagSetResponse is the response struct for api CreateTagSet
type CreateTagSetResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	SetId      string `json:"SetId" xml:"SetId"`
	Status     string `json:"Status" xml:"Status"`
	Photos     int    `json:"Photos" xml:"Photos"`
	CreateTime string `json:"CreateTime" xml:"CreateTime"`
	ModifyTime string `json:"ModifyTime" xml:"ModifyTime"`
}

// CreateCreateTagSetRequest creates a request to invoke CreateTagSet API
func CreateCreateTagSetRequest() (request *CreateTagSetRequest) {
	request = &CreateTagSetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "CreateTagSet", "imm", "openAPI")
	return
}

// CreateCreateTagSetResponse creates a response to parse from CreateTagSet response
func CreateCreateTagSetResponse() (response *CreateTagSetResponse) {
	response = &CreateTagSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
