package baas

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

// CreateChaincode invokes the baas.CreateChaincode API synchronously
// api document: https://help.aliyun.com/api/baas/createchaincode.html
func (client *Client) CreateChaincode(request *CreateChaincodeRequest) (response *CreateChaincodeResponse, err error) {
	response = CreateCreateChaincodeResponse()
	err = client.DoAction(request, response)
	return
}

// CreateChaincodeWithChan invokes the baas.CreateChaincode API asynchronously
// api document: https://help.aliyun.com/api/baas/createchaincode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateChaincodeWithChan(request *CreateChaincodeRequest) (<-chan *CreateChaincodeResponse, <-chan error) {
	responseChan := make(chan *CreateChaincodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateChaincode(request)
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

// CreateChaincodeWithCallback invokes the baas.CreateChaincode API asynchronously
// api document: https://help.aliyun.com/api/baas/createchaincode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateChaincodeWithCallback(request *CreateChaincodeRequest, callback func(response *CreateChaincodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateChaincodeResponse
		var err error
		defer close(result)
		response, err = client.CreateChaincode(request)
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

// CreateChaincodeRequest is the request struct for api CreateChaincode
type CreateChaincodeRequest struct {
	*requests.RpcRequest
	OrganizationId string `position:"Body" name:"OrganizationId"`
	OssBucket      string `position:"Body" name:"OssBucket"`
	OssUrl         string `position:"Body" name:"OssUrl"`
	EndorsePolicy  string `position:"Body" name:"EndorsePolicy"`
	Location       string `position:"Body" name:"Location"`
	ChannelId      string `position:"Body" name:"ChannelId"`
	ConsortiumId   string `position:"Body" name:"ConsortiumId"`
}

// CreateChaincodeResponse is the response struct for api CreateChaincode
type CreateChaincodeResponse struct {
	*responses.BaseResponse
	RequestId string                  `json:"RequestId" xml:"RequestId"`
	Success   bool                    `json:"Success" xml:"Success"`
	ErrorCode int                     `json:"ErrorCode" xml:"ErrorCode"`
	Result    ResultInCreateChaincode `json:"Result" xml:"Result"`
}

// CreateCreateChaincodeRequest creates a request to invoke CreateChaincode API
func CreateCreateChaincodeRequest() (request *CreateChaincodeRequest) {
	request = &CreateChaincodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-07-31", "CreateChaincode", "", "")
	return
}

// CreateCreateChaincodeResponse creates a response to parse from CreateChaincode response
func CreateCreateChaincodeResponse() (response *CreateChaincodeResponse) {
	response = &CreateChaincodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
