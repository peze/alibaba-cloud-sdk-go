
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

// ItemsItem is a nested struct in rds response
type ItemsItem struct {
            ConflictGtid     string `json:"ConflictGtid" xml:"ConflictGtid"`
            DestinationDetail     string `json:"DestinationDetail" xml:"DestinationDetail"`
            TotalConsume     int `json:"TotalConsume" xml:"TotalConsume"`
            Role     string `json:"Role" xml:"Role"`
            InstanceNetworkType     string `json:"InstanceNetworkType" xml:"InstanceNetworkType"`
            Template     string `json:"Template" xml:"Template"`
            RecoveryMode     string `json:"RecoveryMode" xml:"RecoveryMode"`
            AbnormalType     string `json:"AbnormalType" xml:"AbnormalType"`
            TotalScanRows     int `json:"TotalScanRows" xml:"TotalScanRows"`
            DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
            Select     string `json:"select" xml:"select"`
            DatabaseName     string `json:"DatabaseName" xml:"DatabaseName"`
            SecurityIPList     string `json:"SecurityIPList" xml:"SecurityIPList"`
            TotalUpdateRows     int `json:"TotalUpdateRows" xml:"TotalUpdateRows"`
            Schema     string `json:"Schema" xml:"Schema"`
            STATUS     string `json:"STATUS" xml:"STATUS"`
            ReadWriteType     string `json:"ReadWriteType" xml:"ReadWriteType"`
            AvgConsume     float64 `json:"AvgConsume" xml:"AvgConsume"`
            FinishTime     string `json:"FinishTime" xml:"FinishTime"`
            HasInternetIP     bool `json:"HasInternetIP" xml:"HasInternetIP"`
            OccurTime     string `json:"OccurTime" xml:"OccurTime"`
            Progress     string `json:"Progress" xml:"Progress"`
            Insert     int `json:"insert" xml:"insert"`
            SqlType     string `json:"SqlType" xml:"SqlType"`
            SourceInstanceId     string `json:"SourceInstanceId" xml:"SourceInstanceId"`
            TOTALROWS     int `json:"TOTAL_ROWS" xml:"TOTAL_ROWS"`
            TotalCounts     int `json:"TotalCounts" xml:"TotalCounts"`
            InconsistentType     string `json:"InconsistentType" xml:"InconsistentType"`
            ConfictKey     string `json:"ConfictKey" xml:"ConfictKey"`
            KeyType     string `json:"KeyType" xml:"KeyType"`
            Status     string `json:"Status" xml:"Status"`
            ReplicaId     string `json:"ReplicaId" xml:"ReplicaId"`
            Update     int `json:"update" xml:"update"`
            PROCESSROWS     int `json:"PROCESS_ROWS" xml:"PROCESS_ROWS"`
            ConfictReason     string `json:"ConfictReason" xml:"ConfictReason"`
            InternetIP     string `json:"InternetIP" xml:"InternetIP"`
            SourceDetail     string `json:"SourceDetail" xml:"SourceDetail"`
            InconsistentFields     string `json:"InconsistentFields" xml:"InconsistentFields"`
            Delete     int `json:"delete" xml:"delete"`
            InstanceIdA     string `json:"InstanceIdA" xml:"InstanceIdA"`
            Key     string `json:"Key" xml:"Key"`
            RegionId     string `json:"RegionId" xml:"RegionId"`
            FAILMSG     string `json:"FAIL_MSG" xml:"FAIL_MSG"`
            LinkExpiredTime     string `json:"LinkExpiredTime" xml:"LinkExpiredTime"`
            AvgUpdateRows     float64 `json:"AvgUpdateRows" xml:"AvgUpdateRows"`
            UPDATETIME     string `json:"UPDATE_TIME" xml:"UPDATE_TIME"`
            DownloadLink     string `json:"DownloadLink" xml:"DownloadLink"`
            DetailInfo     string `json:"DetailInfo" xml:"DetailInfo"`
            TemplateHash     string `json:"TemplateHash" xml:"TemplateHash"`
            AvgScanRows     float64 `json:"AvgScanRows" xml:"AvgScanRows"`
            DestinationInstanceId     string `json:"DestinationInstanceId" xml:"DestinationInstanceId"`
            InstanceIdB     string `json:"InstanceIdB" xml:"InstanceIdB"`
            CREATETIME     string `json:"CREATE_TIME" xml:"CREATE_TIME"`
            CurrentStep     string `json:"CurrentStep" xml:"CurrentStep"`
            JOBNAME     string `json:"JOB_NAME" xml:"JOB_NAME"`
            Show     int `json:"show" xml:"show"`
                ReadDBInstanceNames []    string  `json:"ReadDBInstanceNames" xml:"ReadDBInstanceNames"`
                ReadDelayTimes []    string  `json:"ReadDelayTimes" xml:"ReadDelayTimes"`
                    InconsistentFileds []InconsistentFiledsItem  `json:"InconsistentFileds" xml:"InconsistentFileds"`
                    Accounts []AccountsItem  `json:"Accounts" xml:"Accounts"`
                    ReadonlyInstanceDelay ReadonlyInstanceDelay `json:"ReadonlyInstanceDelay" xml:"ReadonlyInstanceDelay"`
}
