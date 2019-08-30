package ecs

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

// DescribeLaunchTemplates invokes the ecs.DescribeLaunchTemplates API synchronously
// api document: https://help.aliyun.com/api/ecs/describelaunchtemplates.html
func (client *Client) DescribeLaunchTemplates(request *DescribeLaunchTemplatesRequest) (response *DescribeLaunchTemplatesResponse, err error) {
	response = CreateDescribeLaunchTemplatesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLaunchTemplatesWithChan invokes the ecs.DescribeLaunchTemplates API asynchronously
// api document: https://help.aliyun.com/api/ecs/describelaunchtemplates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLaunchTemplatesWithChan(request *DescribeLaunchTemplatesRequest) (<-chan *DescribeLaunchTemplatesResponse, <-chan error) {
	responseChan := make(chan *DescribeLaunchTemplatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLaunchTemplates(request)
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

// DescribeLaunchTemplatesWithCallback invokes the ecs.DescribeLaunchTemplates API asynchronously
// api document: https://help.aliyun.com/api/ecs/describelaunchtemplates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLaunchTemplatesWithCallback(request *DescribeLaunchTemplatesRequest, callback func(response *DescribeLaunchTemplatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLaunchTemplatesResponse
		var err error
		defer close(result)
		response, err = client.DescribeLaunchTemplates(request)
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

// DescribeLaunchTemplatesRequest is the request struct for api DescribeLaunchTemplates
type DescribeLaunchTemplatesRequest struct {
	*requests.RpcRequest
	LaunchTemplateName      *[]string                             `position:"Query" name:"LaunchTemplateName"  type:"Repeated"`
	ResourceOwnerId         requests.Integer                      `position:"Query" name:"ResourceOwnerId"`
	PageNumber              requests.Integer                      `position:"Query" name:"PageNumber"`
	PageSize                requests.Integer                      `position:"Query" name:"PageSize"`
	TemplateTag             *[]DescribeLaunchTemplatesTemplateTag `position:"Query" name:"TemplateTag"  type:"Repeated"`
	LaunchTemplateId        *[]string                             `position:"Query" name:"LaunchTemplateId"  type:"Repeated"`
	ResourceOwnerAccount    string                                `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string                                `position:"Query" name:"OwnerAccount"`
	TemplateResourceGroupId string                                `position:"Query" name:"TemplateResourceGroupId"`
	OwnerId                 requests.Integer                      `position:"Query" name:"OwnerId"`
}

// DescribeLaunchTemplatesTemplateTag is a repeated param struct in DescribeLaunchTemplatesRequest
type DescribeLaunchTemplatesTemplateTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// DescribeLaunchTemplatesResponse is the response struct for api DescribeLaunchTemplates
type DescribeLaunchTemplatesResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	TotalCount         int                `json:"TotalCount" xml:"TotalCount"`
	PageNumber         int                `json:"PageNumber" xml:"PageNumber"`
	PageSize           int                `json:"PageSize" xml:"PageSize"`
	LaunchTemplateSets LaunchTemplateSets `json:"LaunchTemplateSets" xml:"LaunchTemplateSets"`
}

// CreateDescribeLaunchTemplatesRequest creates a request to invoke DescribeLaunchTemplates API
func CreateDescribeLaunchTemplatesRequest() (request *DescribeLaunchTemplatesRequest) {
	request = &DescribeLaunchTemplatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeLaunchTemplates", "", "")
	return
}

// CreateDescribeLaunchTemplatesResponse creates a response to parse from DescribeLaunchTemplates response
func CreateDescribeLaunchTemplatesResponse() (response *DescribeLaunchTemplatesResponse) {
	response = &DescribeLaunchTemplatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
