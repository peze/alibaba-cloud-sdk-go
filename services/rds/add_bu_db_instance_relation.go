
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

// AddBuDBInstanceRelation invokes the rds.AddBuDBInstanceRelation API synchronously
// api document: https://help.aliyun.com/api/rds/addbudbinstancerelation.html
func (client *Client) AddBuDBInstanceRelation(request *AddBuDBInstanceRelationRequest) (response *AddBuDBInstanceRelationResponse, err error) {
response = CreateAddBuDBInstanceRelationResponse()
err = client.DoAction(request, response)
return
}

// AddBuDBInstanceRelationWithChan invokes the rds.AddBuDBInstanceRelation API asynchronously
// api document: https://help.aliyun.com/api/rds/addbudbinstancerelation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddBuDBInstanceRelationWithChan(request *AddBuDBInstanceRelationRequest) (<-chan *AddBuDBInstanceRelationResponse, <-chan error) {
responseChan := make(chan *AddBuDBInstanceRelationResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.AddBuDBInstanceRelation(request)
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

// AddBuDBInstanceRelationWithCallback invokes the rds.AddBuDBInstanceRelation API asynchronously
// api document: https://help.aliyun.com/api/rds/addbudbinstancerelation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddBuDBInstanceRelationWithCallback(request *AddBuDBInstanceRelationRequest, callback func(response *AddBuDBInstanceRelationResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *AddBuDBInstanceRelationResponse
var err error
defer close(result)
response, err = client.AddBuDBInstanceRelation(request)
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

// AddBuDBInstanceRelationRequest is the request struct for api AddBuDBInstanceRelation
type AddBuDBInstanceRelationRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    BusinessUnit     string `position:"Query" name:"BusinessUnit"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
}


// AddBuDBInstanceRelationResponse is the response struct for api AddBuDBInstanceRelation
type AddBuDBInstanceRelationResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            BusinessUnit     string `json:"BusinessUnit" xml:"BusinessUnit"`
            DBInstanceName     string `json:"DBInstanceName" xml:"DBInstanceName"`
}

// CreateAddBuDBInstanceRelationRequest creates a request to invoke AddBuDBInstanceRelation API
func CreateAddBuDBInstanceRelationRequest() (request *AddBuDBInstanceRelationRequest) {
request = &AddBuDBInstanceRelationRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "AddBuDBInstanceRelation", "rds", "openAPI")
return
}

// CreateAddBuDBInstanceRelationResponse creates a response to parse from AddBuDBInstanceRelation response
func CreateAddBuDBInstanceRelationResponse() (response *AddBuDBInstanceRelationResponse) {
response = &AddBuDBInstanceRelationResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

