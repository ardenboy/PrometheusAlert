package pvtz

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

// Record is a nested struct in pvtz response
type Record struct {
	RecordId int    `json:"RecordId" xml:"RecordId"`
	Rr       string `json:"Rr" xml:"Rr"`
	Type     string `json:"Type" xml:"Type"`
	Ttl      int    `json:"Ttl" xml:"Ttl"`
	Priority int    `json:"Priority" xml:"Priority"`
	Value    string `json:"Value" xml:"Value"`
	Status   string `json:"Status" xml:"Status"`
}
