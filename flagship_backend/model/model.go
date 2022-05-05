package model

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)



type TestGorm struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"` //设置字段的大小为255个字节
	MemberNumber *string `gorm:"unique;not null"` // 设置 memberNumber 字段唯一且不为空
	Num          int     `gorm:"AUTO_INCREMENT"` // 设置 Num字段自增
	Address      string  `gorm:"index:addr"` // 给Address 创建一个名字是  `addr`的索引
	IgnoreMe     int     `gorm:"-"` //忽略这个字段
	Field1     string     `gorm:"type:varchar(20);DEFAULT:null"` //默认值
	Field2     string     `gorm:"type:varchar(20);Column:f_field2"` //指定列的名称
	Field3     string     `gorm:"type:varchar(20);UNIQUE"` //设置唯一约束
	Field5     string     `gorm:"type:varchar(20);UNIQUE_INDEX"` //设置唯一索引
	Field4     string     `gorm:"size:20"` //设置size, 默认255
}

// 给TestGorm 定义表名 指定表
func (TestGorm) TableName() string {
	return "t_test_gorm"
}

// 给所有表设置表前缀



type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11); not null; unique"`
	Password  string `gorm:"size:255;not null"`
	Desc      string `gorm:"size:255;not null"`
}

type Host struct {
	gorm.Model
	IP                string `gorm:"type: varchar(20);not null"`
	Hostname          string `gorm:"type: varchar(20);not null"`
	HostBelongCluster string `gorm:"type: varchar(20);not null"`
	HostCpus          uint8  `gorm:"type: int;not null, column: f_host_cpus"`
	HostMemSize       uint64 `gorm:"type: int;not null, column:f_host_mem_size"`
	HostAddress       string `gorm:"index:addr"`
	HostLoginUser     string `gorm:"type: int;not null; column:f_host_login_user"`
	HostLoginPasswd   string `gorm:"type: int;not null; column:f_host_login_password"`
}

