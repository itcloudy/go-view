// Copyright 2018 itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package all_fileds

import "time"

//角色
type Role struct {
	ID             int       `json:"id" comment:"主键ID"`
	CreatedAt      time.Time `json:"created_at,omitempty" comment:"记录创建时间"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"  comment:"记录更新时间"`
	Name           string    `json:"name" comment:"角色名称" binding:"required" `
	Code           string    `json:"code" comment:"角色编码"`
	InheritIds     []int     `json:"inherit_ids,omitempty"  comment:"所继承的角色ID"`
	InheritStrings string    `json:"inherit_strings"  comment:"所继承角色ID逗号分隔"`
	Inherits       []*Role   `json:"inherits,omitempty" comment:"继承的角色"`
	IsActive       bool      `json:"is_active" comment:"有效"`
}

func (mh *Role) TableName() string {
	return "role"
}
