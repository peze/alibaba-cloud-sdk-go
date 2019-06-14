package acs

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

// DeleteApiInDaily invokes the acs.DeleteApiInDaily API synchronously
// api document: https://help.aliyun.com/api/acs/deleteapiindaily.html
func (client *Client) DeleteApiInDaily(request *DeleteApiInDailyRequest) (response *DeleteApiInDailyResponse, err error) {
	response = CreateDeleteApiInDailyResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteApiInDailyWithChan invokes the acs.DeleteApiInDaily API asynchronously
// api document: https://help.aliyun.com/api/acs/deleteapiindaily.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteApiInDailyWithChan(request *DeleteApiInDailyRequest) (<-chan *DeleteApiInDailyResponse, <-chan error) {
	responseChan := make(chan *DeleteApiInDailyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteApiInDaily(request)
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

// DeleteApiInDailyWithCallback invokes the acs.DeleteApiInDaily API asynchronously
// api document: https://help.aliyun.com/api/acs/deleteapiindaily.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteApiInDailyWithCallback(request *DeleteApiInDailyRequest, callback func(response *DeleteApiInDailyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteApiInDailyResponse
		var err error
		defer close(result)
		response, err = client.DeleteApiInDaily(request)
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

// DeleteApiInDailyRequest is the request struct for api DeleteApiInDaily
type DeleteApiInDailyRequest struct {
	*requests.RoaRequest
	Environment string `position:"Query" name:"Environment"`
	Name        string `position:"Query" name:"Name"`
	ProductName string `position:"Query" name:"ProductName"`
	VersionName string `position:"Query" name:"VersionName"`
}

// DeleteApiInDailyResponse is the response struct for api DeleteApiInDaily
type DeleteApiInDailyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteApiInDailyRequest creates a request to invoke DeleteApiInDaily API
func CreateDeleteApiInDailyRequest() (request *DeleteApiInDailyRequest) {
	request = &DeleteApiInDailyRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Acs", "2015-01-01", "DeleteApiInDaily", "/deleteApiInDaily", "dsafsd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteApiInDailyResponse creates a response to parse from DeleteApiInDaily response
func CreateDeleteApiInDailyResponse() (response *DeleteApiInDailyResponse) {
	response = &DeleteApiInDailyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
