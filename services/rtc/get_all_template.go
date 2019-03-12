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

package rtc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetAllTemplate invokes the rtc.GetAllTemplate API synchronously
// api document: https://help.aliyun.com/api/rtc/getalltemplate.html
func (client *Client) GetAllTemplate(request *GetAllTemplateRequest) (response *GetAllTemplateResponse, err error) {
	response = CreateGetAllTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// GetAllTemplateWithChan invokes the rtc.GetAllTemplate API asynchronously
// api document: https://help.aliyun.com/api/rtc/getalltemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAllTemplateWithChan(request *GetAllTemplateRequest) (<-chan *GetAllTemplateResponse, <-chan error) {
	responseChan := make(chan *GetAllTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAllTemplate(request)
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

// GetAllTemplateWithCallback invokes the rtc.GetAllTemplate API asynchronously
// api document: https://help.aliyun.com/api/rtc/getalltemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAllTemplateWithCallback(request *GetAllTemplateRequest, callback func(response *GetAllTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAllTemplateResponse
		var err error
		defer close(result)
		response, err = client.GetAllTemplate(request)
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

// GetAllTemplateRequest is the request struct for api GetAllTemplate
type GetAllTemplateRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	AppId   string           `position:"Query" name:"AppId"`
}

// GetAllTemplateResponse is the response struct for api GetAllTemplate
type GetAllTemplateResponse struct {
	*responses.BaseResponse
	RequestId   string   `json:"RequestId" xml:"RequestId"`
	TemplateIds []string `json:"TemplateIds" xml:"TemplateIds"`
}

// CreateGetAllTemplateRequest creates a request to invoke GetAllTemplate API
func CreateGetAllTemplateRequest() (request *GetAllTemplateRequest) {
	request = &GetAllTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "GetAllTemplate", "rtc", "openAPI")
	return
}

// CreateGetAllTemplateResponse creates a response to parse from GetAllTemplate response
func CreateGetAllTemplateResponse() (response *GetAllTemplateResponse) {
	response = &GetAllTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}