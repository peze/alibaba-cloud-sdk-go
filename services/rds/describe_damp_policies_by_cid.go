
package rds

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

// DescribeDampPoliciesByCid invokes the rds.DescribeDampPoliciesByCid API synchronously
// api document: https://help.aliyun.com/api/rds/describedamppoliciesbycid.html
func (client *Client) DescribeDampPoliciesByCid(request *DescribeDampPoliciesByCidRequest) (response *DescribeDampPoliciesByCidResponse, err error) {
response = CreateDescribeDampPoliciesByCidResponse()
err = client.DoAction(request, response)
return
}

// DescribeDampPoliciesByCidWithChan invokes the rds.DescribeDampPoliciesByCid API asynchronously
// api document: https://help.aliyun.com/api/rds/describedamppoliciesbycid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDampPoliciesByCidWithChan(request *DescribeDampPoliciesByCidRequest) (<-chan *DescribeDampPoliciesByCidResponse, <-chan error) {
responseChan := make(chan *DescribeDampPoliciesByCidResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeDampPoliciesByCid(request)
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

// DescribeDampPoliciesByCidWithCallback invokes the rds.DescribeDampPoliciesByCid API asynchronously
// api document: https://help.aliyun.com/api/rds/describedamppoliciesbycid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDampPoliciesByCidWithCallback(request *DescribeDampPoliciesByCidRequest, callback func(response *DescribeDampPoliciesByCidResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeDampPoliciesByCidResponse
var err error
defer close(result)
response, err = client.DescribeDampPoliciesByCid(request)
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

// DescribeDampPoliciesByCidRequest is the request struct for api DescribeDampPoliciesByCid
type DescribeDampPoliciesByCidRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    SecurityToken     string `position:"Query" name:"SecurityToken"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}


// DescribeDampPoliciesByCidResponse is the response struct for api DescribeDampPoliciesByCid
type DescribeDampPoliciesByCidResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
                    Policies Policies `json:"Policies" xml:"Policies"`
}

// CreateDescribeDampPoliciesByCidRequest creates a request to invoke DescribeDampPoliciesByCid API
func CreateDescribeDampPoliciesByCidRequest() (request *DescribeDampPoliciesByCidRequest) {
request = &DescribeDampPoliciesByCidRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDampPoliciesByCid", "rds", "openAPI")
return
}

// CreateDescribeDampPoliciesByCidResponse creates a response to parse from DescribeDampPoliciesByCid response
func CreateDescribeDampPoliciesByCidResponse() (response *DescribeDampPoliciesByCidResponse) {
response = &DescribeDampPoliciesByCidResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


