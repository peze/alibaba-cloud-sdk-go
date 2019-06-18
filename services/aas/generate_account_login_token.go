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

// GenerateAccountLoginToken invokes the aas.GenerateAccountLoginToken API synchronously
// api document: https://help.aliyun.com/api/aas/generateaccountlogintoken.html
func (client *Client) GenerateAccountLoginToken(request *GenerateAccountLoginTokenRequest) (response *GenerateAccountLoginTokenResponse, err error) {
	response = CreateGenerateAccountLoginTokenResponse()
	err = client.DoAction(request, response)
	return
}

// GenerateAccountLoginTokenWithChan invokes the aas.GenerateAccountLoginToken API asynchronously
// api document: https://help.aliyun.com/api/aas/generateaccountlogintoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GenerateAccountLoginTokenWithChan(request *GenerateAccountLoginTokenRequest) (<-chan *GenerateAccountLoginTokenResponse, <-chan error) {
	responseChan := make(chan *GenerateAccountLoginTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GenerateAccountLoginToken(request)
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

// GenerateAccountLoginTokenWithCallback invokes the aas.GenerateAccountLoginToken API asynchronously
// api document: https://help.aliyun.com/api/aas/generateaccountlogintoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GenerateAccountLoginTokenWithCallback(request *GenerateAccountLoginTokenRequest, callback func(response *GenerateAccountLoginTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GenerateAccountLoginTokenResponse
		var err error
		defer close(result)
		response, err = client.GenerateAccountLoginToken(request)
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

// GenerateAccountLoginTokenRequest is the request struct for api GenerateAccountLoginToken
type GenerateAccountLoginTokenRequest struct {
	*requests.RpcRequest
	TargetPk string `position:"Query" name:"TargetPk"`
}

// GenerateAccountLoginTokenResponse is the response struct for api GenerateAccountLoginToken
type GenerateAccountLoginTokenResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	LoginToken LoginToken `json:"LoginToken" xml:"LoginToken"`
}

// CreateGenerateAccountLoginTokenRequest creates a request to invoke GenerateAccountLoginToken API
func CreateGenerateAccountLoginTokenRequest() (request *GenerateAccountLoginTokenRequest) {
	request = &GenerateAccountLoginTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Aas", "2015-07-01", "GenerateAccountLoginToken", "aas", "openAPI")
	return
}

// CreateGenerateAccountLoginTokenResponse creates a response to parse from GenerateAccountLoginToken response
func CreateGenerateAccountLoginTokenResponse() (response *GenerateAccountLoginTokenResponse) {
	response = &GenerateAccountLoginTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
