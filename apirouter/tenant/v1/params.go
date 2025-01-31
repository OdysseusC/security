/*
Copyright 2021 The tKeel Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

type TenantCreteIn struct {
	Title  string        `json:"title"`
	Remark string        `json:"remark"`
	Admin  *UserCreateIn `json:"admin"`
}
type TenantCreateOut struct {
	ID     int          `json:"tenant_id"`
	Title  string       `json:"title"`
	Remark string       `json:"remark"`
	Admin  UserCreateIn `json:"admin"`
}

type UserCreateIn struct {
	TenantID int    `json:"tenant_id" `
	UserName string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
}

type Tenant struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Remark string `json:"remark"`
}

type User struct {
	ID         string `json:"user_id"`
	ExternalID string `json:"external_id"`
	TenantID   int    `json:"tenant_id"`
	UserName   string `json:"username"`
	NickName   string `json:"nick_name"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
}
