
package ams

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

// AppInfo is a nested struct in ams response
type AppInfo struct {
            Name     string `json:"Name" xml:"Name"`
            AppSecret     string `json:"AppSecret" xml:"AppSecret"`
            ProductId     int `json:"ProductId" xml:"ProductId"`
            Status     int `json:"Status" xml:"Status"`
            ModifyTime     string `json:"ModifyTime" xml:"ModifyTime"`
            TaobaoUserId     int `json:"TaobaoUserId" xml:"TaobaoUserId"`
            AppKey     int `json:"AppKey" xml:"AppKey"`
            BizName     string `json:"BizName" xml:"BizName"`
            CreateTime     string `json:"CreateTime" xml:"CreateTime"`
            PackageName     string `json:"PackageName" xml:"PackageName"`
            CallbackUrl     string `json:"CallbackUrl" xml:"CallbackUrl"`
            AppId     int `json:"AppId" xml:"AppId"`
            CertDevelopAvail     bool `json:"CertDevelopAvail" xml:"CertDevelopAvail"`
            CertDevelopExpiration     string `json:"CertDevelopExpiration" xml:"CertDevelopExpiration"`
            LogoUrl     string `json:"LogoUrl" xml:"LogoUrl"`
            BundleId     string `json:"BundleId" xml:"BundleId"`
            CertProductAvail     bool `json:"CertProductAvail" xml:"CertProductAvail"`
            CertProductExpiration     string `json:"CertProductExpiration" xml:"CertProductExpiration"`
            UserId     int `json:"UserId" xml:"UserId"`
            Description     string `json:"Description" xml:"Description"`
            Type     int `json:"Type" xml:"Type"`
}