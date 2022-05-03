package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `gorm: "type: varchar(20);not null"`
	Telephone string `gorm: "type: varchar(11); not null; unique"`
	Password string `gorm: "size: 255; not null"`
	Desc string `gorm: "size: 255; not null"`
}



type Host struct {
	gorm.Model
	IP string `gorm: "type: varchar(20);not null"`
	Hostname string `gorm: "type: varchar(20);not null"`
	HostBelongCluster string `gorm: "type: varchar(20);not null"`
	HostCpus uint8 `gorm: "type: int;not null, column: f_host_cpus"`
	HostMemSize uint64 `gorm: "type: int;not null, column:f_host_mem_size"`
	HostAddress      string  `gorm:"index:addr"`
}

