package main

import (
	_ "app/routers"
    "github.com/astaxie/beego/orm"
    _ "github.com/lib/pq"
    _ "app/models"
    "os"
)

func init() {
    orm.RegisterDriver("postgres", orm.DRPostgres)
    // orm.RegisterDataBase("default", "postgres", "user=postgres password=postgres host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
    orm.RegisterDataBase("default", "postgres", os.Getenv("DATABASE_URL"))
}

func db() {
    orm.RunCommand()
}

