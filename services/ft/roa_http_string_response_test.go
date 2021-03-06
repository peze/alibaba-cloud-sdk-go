package ft

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

// RoaHttpStringResponseTest invokes the ft.RoaHttpStringResponseTest API synchronously
// api document: https://help.aliyun.com/api/ft/roahttpstringresponsetest.html
func (client *Client) RoaHttpStringResponseTest(request *RoaHttpStringResponseTestRequest) (response *RoaHttpStringResponseTestResponse, err error) {
	response = CreateRoaHttpStringResponseTestResponse()
	err = client.DoAction(request, response)
	return
}

// RoaHttpStringResponseTestWithChan invokes the ft.RoaHttpStringResponseTest API asynchronously
// api document: https://help.aliyun.com/api/ft/roahttpstringresponsetest.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RoaHttpStringResponseTestWithChan(request *RoaHttpStringResponseTestRequest) (<-chan *RoaHttpStringResponseTestResponse, <-chan error) {
	responseChan := make(chan *RoaHttpStringResponseTestResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RoaHttpStringResponseTest(request)
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

// RoaHttpStringResponseTestWithCallback invokes the ft.RoaHttpStringResponseTest API asynchronously
// api document: https://help.aliyun.com/api/ft/roahttpstringresponsetest.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RoaHttpStringResponseTestWithCallback(request *RoaHttpStringResponseTestRequest, callback func(response *RoaHttpStringResponseTestResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RoaHttpStringResponseTestResponse
		var err error
		defer close(result)
		response, err = client.RoaHttpStringResponseTest(request)
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

// RoaHttpStringResponseTestRequest is the request struct for api RoaHttpStringResponseTest
type RoaHttpStringResponseTestRequest struct {
	*requests.RoaRequest
	QueryParam string `position:"Query" name:"QueryParam"`
}

// RoaHttpStringResponseTestResponse is the response struct for api RoaHttpStringResponseTest
type RoaHttpStringResponseTestResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Params    Params `json:"Params" xml:"Params"`
}

// CreateRoaHttpStringResponseTestRequest creates a request to invoke RoaHttpStringResponseTest API
func CreateRoaHttpStringResponseTestRequest() (request *RoaHttpStringResponseTestRequest) {
	request = &RoaHttpStringResponseTestRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Ft", "2019-08-02", "RoaHttpStringResponseTest", "/web/getData", "ft", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRoaHttpStringResponseTestResponse creates a response to parse from RoaHttpStringResponseTest response
func CreateRoaHttpStringResponseTestResponse() (response *RoaHttpStringResponseTestResponse) {
	response = &RoaHttpStringResponseTestResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
