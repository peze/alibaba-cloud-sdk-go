
package acs

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

// PutGatedLaunchPolicy invokes the acs.PutGatedLaunchPolicy API synchronously
// api document: https://help.aliyun.com/api/acs/putgatedlaunchpolicy.html
func (client *Client) PutGatedLaunchPolicy(request *PutGatedLaunchPolicyRequest) (response *PutGatedLaunchPolicyResponse, err error) {
response = CreatePutGatedLaunchPolicyResponse()
err = client.DoAction(request, response)
return
}

// PutGatedLaunchPolicyWithChan invokes the acs.PutGatedLaunchPolicy API asynchronously
// api document: https://help.aliyun.com/api/acs/putgatedlaunchpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutGatedLaunchPolicyWithChan(request *PutGatedLaunchPolicyRequest) (<-chan *PutGatedLaunchPolicyResponse, <-chan error) {
responseChan := make(chan *PutGatedLaunchPolicyResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.PutGatedLaunchPolicy(request)
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

// PutGatedLaunchPolicyWithCallback invokes the acs.PutGatedLaunchPolicy API asynchronously
// api document: https://help.aliyun.com/api/acs/putgatedlaunchpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutGatedLaunchPolicyWithCallback(request *PutGatedLaunchPolicyRequest, callback func(response *PutGatedLaunchPolicyResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *PutGatedLaunchPolicyResponse
var err error
defer close(result)
response, err = client.PutGatedLaunchPolicy(request)
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

// PutGatedLaunchPolicyRequest is the request struct for api PutGatedLaunchPolicy
type PutGatedLaunchPolicyRequest struct {
*requests.RoaRequest
                    ContentLength     requests.Integer `position:"Header" name:"Content-Length"`
                    BodyContent     string `position:"Body" name:"BodyContent"`
                    ContentMD5     string `position:"Header" name:"Content-MD5"`
                    ContentType     string `position:"Header" name:"Content-Type"`
                    Accept     string `position:"Header" name:"Accept"`
}


// PutGatedLaunchPolicyResponse is the response struct for api PutGatedLaunchPolicy
type PutGatedLaunchPolicyResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreatePutGatedLaunchPolicyRequest creates a request to invoke PutGatedLaunchPolicy API
func CreatePutGatedLaunchPolicyRequest() (request *PutGatedLaunchPolicyRequest) {
request = &PutGatedLaunchPolicyRequest{
RoaRequest: &requests.RoaRequest{},
}
request.InitWithApiInfo("Acs", "2015-01-01", "PutGatedLaunchPolicy", "/GatedLaunchPolicy", "", "")
request.Method = requests.PUT
return
}

// CreatePutGatedLaunchPolicyResponse creates a response to parse from PutGatedLaunchPolicy response
func CreatePutGatedLaunchPolicyResponse() (response *PutGatedLaunchPolicyResponse) {
response = &PutGatedLaunchPolicyResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

