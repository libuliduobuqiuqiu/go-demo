package main

import (
	"fmt"
	"godemo/pkg"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const (
	confPath = "/data/MyRepo/go-demo/configs/local_conf.json"
)

func main() {
	mysqlDSN := pkg.GenMysqlDSN(confPath)

	if mysqlDSN == "" {
		return
	}

	db, err := gorm.Open(mysql.Open(mysqlDSN))
	if err != nil {
		log.Fatal(err)
		return
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "../../internal/gostorage/gormgendemo/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery,
	})
	fmt.Println(g.OutPath)
	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateModel("book"),
		g.GenerateModelAs("model", "MyModel"),
		g.GenerateModel("history"),
		g.GenerateModel("device"),
		g.GenerateModel("users"),
	)
	// g.ApplyBasic(g.GenerateModel("book"), g.GenerateModel("history"))
	// g.ApplyBasic()

	g.Execute()
}
