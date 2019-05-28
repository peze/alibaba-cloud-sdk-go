
package r_kvstore

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

// DescribeCertification invokes the r_kvstore.DescribeCertification API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/describecertification.html
func (client *Client) DescribeCertification(request *DescribeCertificationRequest) (response *DescribeCertificationResponse, err error) {
response = CreateDescribeCertificationResponse()
err = client.DoAction(request, response)
return
}

// DescribeCertificationWithChan invokes the r_kvstore.DescribeCertification API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describecertification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCertificationWithChan(request *DescribeCertificationRequest) (<-chan *DescribeCertificationResponse, <-chan error) {
responseChan := make(chan *DescribeCertificationResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeCertification(request)
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

// DescribeCertificationWithCallback invokes the r_kvstore.DescribeCertification API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describecertification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCertificationWithCallback(request *DescribeCertificationRequest, callback func(response *DescribeCertificationResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeCertificationResponse
var err error
defer close(result)
response, err = client.DescribeCertification(request)
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

// DescribeCertificationRequest is the request struct for api DescribeCertification
type DescribeCertificationRequest struct {
*requests.RpcRequest
                    ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
                    SecurityToken     string `position:"Query" name:"SecurityToken"`
                    ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
                    OwnerAccount     string `position:"Query" name:"OwnerAccount"`
                    OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
                    InstanceId     string `position:"Query" name:"InstanceId"`
                    Parameters     string `position:"Query" name:"Parameters"`
}


// DescribeCertificationResponse is the response struct for api DescribeCertification
type DescribeCertificationResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            NoCertification     bool `json:"NoCertification" xml:"NoCertification"`
}

// CreateDescribeCertificationRequest creates a request to invoke DescribeCertification API
func CreateDescribeCertificationRequest() (request *DescribeCertificationRequest) {
request = &DescribeCertificationRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeCertification", "kvstore", "openAPI")
return
}

// CreateDescribeCertificationResponse creates a response to parse from DescribeCertification response
func CreateDescribeCertificationResponse() (response *DescribeCertificationResponse) {
response = &DescribeCertificationResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


