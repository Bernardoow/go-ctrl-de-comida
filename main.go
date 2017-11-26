package main

import (
	_ "app/routers"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/lib/pq"
    _ "app/models"
    "os"
    "strconv"
    "fmt"
)

func init() {
    orm.RegisterDriver("postgres", orm.DRPostgres)
    // orm.RegisterDataBase("default", "postgres", "user=postgres password=postgres host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
    orm.RegisterDataBase("default", "postgres", os.Getenv("DATABASE_URL"))

}

func main() {

    port, err := strconv.Atoi(os.Getenv("PORT"));

    if err != nil {
        port = 8000
    }

    fmt.Println(port)

    beego.BConfig.Listen.HTTPPort = port
    beego.Run()
}

