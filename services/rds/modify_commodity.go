
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

// ModifyCommodity invokes the rds.ModifyCommodity API synchronously
// api document: https://help.aliyun.com/api/rds/modifycommodity.html
func (client *Client) ModifyCommodity(request *ModifyCommodityRequest) (response *ModifyCommodityResponse, err error) {
response = CreateModifyCommodityResponse()
err = client.DoAction(request, response)
return
}

// ModifyCommodityWithChan invokes the rds.ModifyCommodity API asynchronously
// api document: https://help.aliyun.com/api/rds/modifycommodity.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyCommodityWithChan(request *ModifyCommodityRequest) (<-chan *ModifyCommodityResponse, <-chan error) {
responseChan := make(chan *ModifyCommodityResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ModifyCommodity(request)
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

// ModifyCommodityWithCallback invokes the rds.ModifyCommodity API asynchronously
// api document: https://help.aliyun.com/api/rds/modifycommodity.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyCommodityWithCallback(request *ModifyCommodityRequest, callback func(response *ModifyCommodityResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ModifyCommodityResponse
var err error
defer close(result)
response, err = client.ModifyCommodity(request)
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

// ModifyCommodityRequest is the request struct for api ModifyCommodity
type ModifyCommodityRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
}


// ModifyCommodityResponse is the response struct for api ModifyCommodity
type ModifyCommodityResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            Commodity     string `json:"Commodity" xml:"Commodity"`
}

// CreateModifyCommodityRequest creates a request to invoke ModifyCommodity API
func CreateModifyCommodityRequest() (request *ModifyCommodityRequest) {
request = &ModifyCommodityRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "ModifyCommodity", "rds", "openAPI")
return
}

// CreateModifyCommodityResponse creates a response to parse from ModifyCommodity response
func CreateModifyCommodityResponse() (response *ModifyCommodityResponse) {
response = &ModifyCommodityResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

