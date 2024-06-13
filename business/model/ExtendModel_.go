package model

type ExtendModel struct {
	ID         int64 `json:"id" gorm:"column:id;primary_key;type:bigint;not null;;comment:主键"`
	DelFlag    int8  `json:"-" gorm:"column:del_flag;type:int;not null;default:0;comment:删除标志 0 有效 1删除"`
	AddTime    *Time `json:"addTime" gorm:"autoCreateTime:milli;column:add_time;type:datetime;not null;default:current_timestamp;comment:新增时间"`
	UpdateTime *Time `json:"updateTime" gorm:"autoUpdateTime:milli;column:update_time;type:datetime;not null;default:current_timestamp;comment:更新时间"`
}
