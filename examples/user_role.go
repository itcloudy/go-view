// Copyright 2018 itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package all_fileds

import "time"

//用户角色
type UserRole struct {
	ID        int       `json:"id" comment:"主键ID"`
	CreatedAt time.Time `json:"created_at,omitempty" comment:"记录创建时间"`
	UpdatedAt time.Time `json:"updated_at,omitempty" comment:"记录更新时间"`
	Role      *Role     `json:"role" comment:"角色"`
	RoleID    int       `json:"role_id" comment:"角色ID"`
	User      *User     `json:"user" comment:"用户"`
	UserID    int       `json:"user_id" comment:"用户ID"`
}

func (mh *UserRole) TableName() string {
	return "user_role"
}
