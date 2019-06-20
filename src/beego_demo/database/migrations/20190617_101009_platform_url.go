package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type PlatformUrl_20190617_101009 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &PlatformUrl_20190617_101009{}
	m.Created = "20190617_101009"

	migration.Register("PlatformUrl_20190617_101009", m)
}

// Run the migrations
func (m *PlatformUrl_20190617_101009) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE platform_url(`id` int(11) NOT NULL AUTO_INCREMENT,`platform_id` varchar(128) NOT NULL,`url` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *PlatformUrl_20190617_101009) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `platform_url`")
}
