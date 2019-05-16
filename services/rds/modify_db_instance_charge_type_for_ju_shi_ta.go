
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

// ModifyDBInstanceChargeTypeForJuShiTa invokes the rds.ModifyDBInstanceChargeTypeForJuShiTa API synchronously
// api document: https://help.aliyun.com/api/rds/modifydbinstancechargetypeforjushita.html
func (client *Client) ModifyDBInstanceChargeTypeForJuShiTa(request *ModifyDBInstanceChargeTypeForJuShiTaRequest) (response *ModifyDBInstanceChargeTypeForJuShiTaResponse, err error) {
response = CreateModifyDBInstanceChargeTypeForJuShiTaResponse()
err = client.DoAction(request, response)
return
}

// ModifyDBInstanceChargeTypeForJuShiTaWithChan invokes the rds.ModifyDBInstanceChargeTypeForJuShiTa API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydbinstancechargetypeforjushita.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceChargeTypeForJuShiTaWithChan(request *ModifyDBInstanceChargeTypeForJuShiTaRequest) (<-chan *ModifyDBInstanceChargeTypeForJuShiTaResponse, <-chan error) {
responseChan := make(chan *ModifyDBInstanceChargeTypeForJuShiTaResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ModifyDBInstanceChargeTypeForJuShiTa(request)
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

// ModifyDBInstanceChargeTypeForJuShiTaWithCallback invokes the rds.ModifyDBInstanceChargeTypeForJuShiTa API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydbinstancechargetypeforjushita.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceChargeTypeForJuShiTaWithCallback(request *ModifyDBInstanceChargeTypeForJuShiTaRequest, callback func(response *ModifyDBInstanceChargeTypeForJuShiTaResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ModifyDBInstanceChargeTypeForJuShiTaResponse
var err error
defer close(result)
response, err = client.ModifyDBInstanceChargeTypeForJuShiTa(request)
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

// ModifyDBInstanceChargeTypeForJuShiTaRequest is the request struct for api ModifyDBInstanceChargeTypeForJuShiTa
type ModifyDBInstanceChargeTypeForJuShiTaRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    JuShiTaInstanceUid     requests.Integer `position:"Query" name:"JuShiTaInstanceUid"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    ChargeType     string `position:"Query" name:"ChargeType"`
}


// ModifyDBInstanceChargeTypeForJuShiTaResponse is the response struct for api ModifyDBInstanceChargeTypeForJuShiTa
type ModifyDBInstanceChargeTypeForJuShiTaResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBInstanceChargeTypeForJuShiTaRequest creates a request to invoke ModifyDBInstanceChargeTypeForJuShiTa API
func CreateModifyDBInstanceChargeTypeForJuShiTaRequest() (request *ModifyDBInstanceChargeTypeForJuShiTaRequest) {
request = &ModifyDBInstanceChargeTypeForJuShiTaRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceChargeTypeForJuShiTa", "rds", "openAPI")
return
}

// CreateModifyDBInstanceChargeTypeForJuShiTaResponse creates a response to parse from ModifyDBInstanceChargeTypeForJuShiTa response
func CreateModifyDBInstanceChargeTypeForJuShiTaResponse() (response *ModifyDBInstanceChargeTypeForJuShiTaResponse) {
response = &ModifyDBInstanceChargeTypeForJuShiTaResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

