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

// DescribeReservedInstancePrice invokes the ecs.DescribeReservedInstancePrice API synchronously
// api document: https://help.aliyun.com/api/ecs/describereservedinstanceprice.html
func (client *Client) DescribeReservedInstancePrice(request *DescribeReservedInstancePriceRequest) (response *DescribeReservedInstancePriceResponse, err error) {
	response = CreateDescribeReservedInstancePriceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeReservedInstancePriceWithChan invokes the ecs.DescribeReservedInstancePrice API asynchronously
// api document: https://help.aliyun.com/api/ecs/describereservedinstanceprice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeReservedInstancePriceWithChan(request *DescribeReservedInstancePriceRequest) (<-chan *DescribeReservedInstancePriceResponse, <-chan error) {
	responseChan := make(chan *DescribeReservedInstancePriceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeReservedInstancePrice(request)
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

// DescribeReservedInstancePriceWithCallback invokes the ecs.DescribeReservedInstancePrice API asynchronously
// api document: https://help.aliyun.com/api/ecs/describereservedinstanceprice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeReservedInstancePriceWithCallback(request *DescribeReservedInstancePriceRequest, callback func(response *DescribeReservedInstancePriceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeReservedInstancePriceResponse
		var err error
		defer close(result)
		response, err = client.DescribeReservedInstancePrice(request)
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

// DescribeReservedInstancePriceRequest is the request struct for api DescribeReservedInstancePrice
type DescribeReservedInstancePriceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer                    `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer                    `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer                    `position:"Query" name:"PageSize"`
	Tag                  *[]DescribeReservedInstancePriceTag `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceOwnerAccount string                              `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                              `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer                    `position:"Query" name:"OwnerId"`
	ReservedInstanceId   *[]string                           `position:"Query" name:"ReservedInstanceId"  type:"Repeated"`
}

// DescribeReservedInstancePriceTag is a repeated param struct in DescribeReservedInstancePriceRequest
type DescribeReservedInstancePriceTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// DescribeReservedInstancePriceResponse is the response struct for api DescribeReservedInstancePrice
type DescribeReservedInstancePriceResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	PriceInfo PriceInfo `json:"PriceInfo" xml:"PriceInfo"`
}

// CreateDescribeReservedInstancePriceRequest creates a request to invoke DescribeReservedInstancePrice API
func CreateDescribeReservedInstancePriceRequest() (request *DescribeReservedInstancePriceRequest) {
	request = &DescribeReservedInstancePriceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2016-03-14", "DescribeReservedInstancePrice", "", "")
	return
}

// CreateDescribeReservedInstancePriceResponse creates a response to parse from DescribeReservedInstancePrice response
func CreateDescribeReservedInstancePriceResponse() (response *DescribeReservedInstancePriceResponse) {
	response = &DescribeReservedInstancePriceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
