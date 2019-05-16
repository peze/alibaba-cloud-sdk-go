
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

// CloneDBInstanceForSecurity invokes the rds.CloneDBInstanceForSecurity API synchronously
// api document: https://help.aliyun.com/api/rds/clonedbinstanceforsecurity.html
func (client *Client) CloneDBInstanceForSecurity(request *CloneDBInstanceForSecurityRequest) (response *CloneDBInstanceForSecurityResponse, err error) {
response = CreateCloneDBInstanceForSecurityResponse()
err = client.DoAction(request, response)
return
}

// CloneDBInstanceForSecurityWithChan invokes the rds.CloneDBInstanceForSecurity API asynchronously
// api document: https://help.aliyun.com/api/rds/clonedbinstanceforsecurity.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CloneDBInstanceForSecurityWithChan(request *CloneDBInstanceForSecurityRequest) (<-chan *CloneDBInstanceForSecurityResponse, <-chan error) {
responseChan := make(chan *CloneDBInstanceForSecurityResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CloneDBInstanceForSecurity(request)
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

// CloneDBInstanceForSecurityWithCallback invokes the rds.CloneDBInstanceForSecurity API asynchronously
// api document: https://help.aliyun.com/api/rds/clonedbinstanceforsecurity.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CloneDBInstanceForSecurityWithCallback(request *CloneDBInstanceForSecurityRequest, callback func(response *CloneDBInstanceForSecurityResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CloneDBInstanceForSecurityResponse
var err error
defer close(result)
response, err = client.CloneDBInstanceForSecurity(request)
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

// CloneDBInstanceForSecurityRequest is the request struct for api CloneDBInstanceForSecurity
type CloneDBInstanceForSecurityRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    DBInstanceStorage     requests.Integer `position:"Query" name:"DBInstanceStorage"`
                    ClientToken     string `position:"Query" name:"ClientToken"`
                    TargetAliBid     string `position:"Query" name:"TargetAliBid"`
                    ResourceGroupId     string `position:"Query" name:"ResourceGroupId"`
                    DBInstanceId     string `position:"Query" name:"DBInstanceId"`
                    DBInstanceDescription     string `position:"Query" name:"DBInstanceDescription"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    BackupId     string `position:"Query" name:"BackupId"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    DBInstanceClass     string `position:"Query" name:"DBInstanceClass"`
                    TargetAliUid     string `position:"Query" name:"TargetAliUid"`
                    PayType     string `position:"Query" name:"PayType"`
}


// CloneDBInstanceForSecurityResponse is the response struct for api CloneDBInstanceForSecurity
type CloneDBInstanceForSecurityResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
            OrderId     string `json:"OrderId" xml:"OrderId"`
            ConnectionString     string `json:"ConnectionString" xml:"ConnectionString"`
            Port     string `json:"Port" xml:"Port"`
}

// CreateCloneDBInstanceForSecurityRequest creates a request to invoke CloneDBInstanceForSecurity API
func CreateCloneDBInstanceForSecurityRequest() (request *CloneDBInstanceForSecurityRequest) {
request = &CloneDBInstanceForSecurityRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "CloneDBInstanceForSecurity", "rds", "openAPI")
return
}

// CreateCloneDBInstanceForSecurityResponse creates a response to parse from CloneDBInstanceForSecurity response
func CreateCloneDBInstanceForSecurityResponse() (response *CloneDBInstanceForSecurityResponse) {
response = &CloneDBInstanceForSecurityResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

