package model

import "gorm.io/plugin/soft_delete"

type ExtendModel struct {
	ID         int64                 `json:"id" gorm:"column:id;primary_key;type:bigint;not null;;comment:主键"`
	DelFlag    soft_delete.DeletedAt `gorm:"softDelete:flag"`
	AddTime    *Time                 `json:"addTime" gorm:"autoCreateTime:milli;column:add_time;type:datetime;not null;default:current_timestamp;comment:新增时间"`
	UpdateTime *Time                 `json:"updateTime" gorm:"autoUpdateTime:milli;column:update_time;type:datetime;not null;default:current_timestamp;comment:更新时间"`
}
