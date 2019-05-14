
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

// GetGatedLaunchPolicy invokes the acs.GetGatedLaunchPolicy API synchronously
// api document: https://help.aliyun.com/api/acs/getgatedlaunchpolicy.html
func (client *Client) GetGatedLaunchPolicy(request *GetGatedLaunchPolicyRequest) (response *GetGatedLaunchPolicyResponse, err error) {
response = CreateGetGatedLaunchPolicyResponse()
err = client.DoAction(request, response)
return
}

// GetGatedLaunchPolicyWithChan invokes the acs.GetGatedLaunchPolicy API asynchronously
// api document: https://help.aliyun.com/api/acs/getgatedlaunchpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetGatedLaunchPolicyWithChan(request *GetGatedLaunchPolicyRequest) (<-chan *GetGatedLaunchPolicyResponse, <-chan error) {
responseChan := make(chan *GetGatedLaunchPolicyResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.GetGatedLaunchPolicy(request)
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

// GetGatedLaunchPolicyWithCallback invokes the acs.GetGatedLaunchPolicy API asynchronously
// api document: https://help.aliyun.com/api/acs/getgatedlaunchpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetGatedLaunchPolicyWithCallback(request *GetGatedLaunchPolicyRequest, callback func(response *GetGatedLaunchPolicyResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *GetGatedLaunchPolicyResponse
var err error
defer close(result)
response, err = client.GetGatedLaunchPolicy(request)
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

// GetGatedLaunchPolicyRequest is the request struct for api GetGatedLaunchPolicy
type GetGatedLaunchPolicyRequest struct {
*requests.RoaRequest
                    ProductName     string `position:"Path" name:"ProductName"`
                    PolicyName     string `position:"Path" name:"PolicyName"`
                    Accept     string `position:"Header" name:"Accept"`
}


// GetGatedLaunchPolicyResponse is the response struct for api GetGatedLaunchPolicy
type GetGatedLaunchPolicyResponse struct {
*responses.BaseResponse
            ProductName     string `json:"productName" xml:"productName"`
            PolicyName     string `json:"policyName" xml:"policyName"`
            Percentage     int `json:"percentage" xml:"percentage"`
            Description     string `json:"description" xml:"description"`
                    Parameters ParametersInGetGatedLaunchPolicy `json:"Parameters" xml:"Parameters"`
}

// CreateGetGatedLaunchPolicyRequest creates a request to invoke GetGatedLaunchPolicy API
func CreateGetGatedLaunchPolicyRequest() (request *GetGatedLaunchPolicyRequest) {
request = &GetGatedLaunchPolicyRequest{
RoaRequest: &requests.RoaRequest{},
}
request.InitWithApiInfo("Acs", "2015-01-01", "GetGatedLaunchPolicy", "/GatedLaunchPolicy/[ProductName]/[PolicyName]", "", "")
request.Method = requests.GET
return
}

// CreateGetGatedLaunchPolicyResponse creates a response to parse from GetGatedLaunchPolicy response
func CreateGetGatedLaunchPolicyResponse() (response *GetGatedLaunchPolicyResponse) {
response = &GetGatedLaunchPolicyResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


