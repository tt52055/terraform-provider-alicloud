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

// RdsInstanceResourceSetting is a nested struct in rds response
type RdsInstanceResourceSetting struct {
	StartDate          string `json:"StartDate" xml:"StartDate"`
	EndDate            string `json:"EndDate" xml:"EndDate"`
	ResourceNiche      string `json:"ResourceNiche" xml:"ResourceNiche"`
	NoticeBarContent   string `json:"NoticeBarContent" xml:"NoticeBarContent"`
	PoppedUpButtonText string `json:"PoppedUpButtonText" xml:"PoppedUpButtonText"`
	PoppedUpButtonType string `json:"PoppedUpButtonType" xml:"PoppedUpButtonType"`
	PoppedUpButtonUrl  string `json:"PoppedUpButtonUrl" xml:"PoppedUpButtonUrl"`
	PoppedUpContent    string `json:"PoppedUpContent" xml:"PoppedUpContent"`
}
