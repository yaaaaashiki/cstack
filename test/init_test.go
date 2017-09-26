package test

import (
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const dsn = "root:@/cstack_test?parseTime=true&collation=utf8_general_ci&interpolateParams=true"

func startUpDatabase() *gorm.DB {
	gorm, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Databas open error. Error:", err.Error())
	}
	return gorm
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Errorf(err.Error())
	}
}
