package emr

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

// ConfigHistory is a nested struct in emr response
type ConfigHistory struct {
	NewValue       string `json:"NewValue" xml:"NewValue"`
	ConfigFileName string `json:"ConfigFileName" xml:"ConfigFileName"`
	Comment        string `json:"Comment" xml:"Comment"`
	Applied        bool   `json:"Applied" xml:"Applied"`
	OldValue       string `json:"OldValue" xml:"OldValue"`
	Author         string `json:"Author" xml:"Author"`
	CreateTime     int64  `json:"CreateTime" xml:"CreateTime"`
	ConfigVersion  string `json:"ConfigVersion" xml:"ConfigVersion"`
	ServiceName    string `json:"ServiceName" xml:"ServiceName"`
	ConfigItemName string `json:"ConfigItemName" xml:"ConfigItemName"`
}