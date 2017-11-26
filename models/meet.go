package models

import (
    "github.com/astaxie/beego/orm"
    "time"
)


type Meet struct {
    Id          int `orm:pk;`
    Foods []*Food `orm:"rel(m2m)"`
    Created time.Time `orm:"auto_now_add;type(datetime)"`
    LactosePresence bool `orm:type(bool)`
}


func init() {
    orm.RegisterModel(new(Meet))
}
