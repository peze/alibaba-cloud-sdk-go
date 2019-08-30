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

// CopyImage invokes the ecs.CopyImage API synchronously
// api document: https://help.aliyun.com/api/ecs/copyimage.html
func (client *Client) CopyImage(request *CopyImageRequest) (response *CopyImageResponse, err error) {
	response = CreateCopyImageResponse()
	err = client.DoAction(request, response)
	return
}

// CopyImageWithChan invokes the ecs.CopyImage API asynchronously
// api document: https://help.aliyun.com/api/ecs/copyimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CopyImageWithChan(request *CopyImageRequest) (<-chan *CopyImageResponse, <-chan error) {
	responseChan := make(chan *CopyImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CopyImage(request)
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

// CopyImageWithCallback invokes the ecs.CopyImage API asynchronously
// api document: https://help.aliyun.com/api/ecs/copyimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CopyImageWithCallback(request *CopyImageRequest, callback func(response *CopyImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CopyImageResponse
		var err error
		defer close(result)
		response, err = client.CopyImage(request)
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

// CopyImageRequest is the request struct for api CopyImage
type CopyImageRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ImageId                string           `position:"Query" name:"ImageId"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	DestinationImageName   string           `position:"Query" name:"DestinationImageName"`
	DestinationRegionId    string           `position:"Query" name:"DestinationRegionId"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	Encrypted              requests.Boolean `position:"Query" name:"Encrypted"`
	Tag                    *[]CopyImageTag  `position:"Query" name:"Tag"  type:"Repeated"`
	KMSKeyId               string           `position:"Query" name:"KMSKeyId"`
	DestinationDescription string           `position:"Query" name:"DestinationDescription"`
}

// CopyImageTag is a repeated param struct in CopyImageRequest
type CopyImageTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CopyImageResponse is the response struct for api CopyImage
type CopyImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ImageId   string `json:"ImageId" xml:"ImageId"`
}

// CreateCopyImageRequest creates a request to invoke CopyImage API
func CreateCopyImageRequest() (request *CopyImageRequest) {
	request = &CopyImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CopyImage", "", "")
	return
}

// CreateCopyImageResponse creates a response to parse from CopyImage response
func CreateCopyImageResponse() (response *CopyImageResponse) {
	response = &CopyImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
