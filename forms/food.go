package forms

import (
    "github.com/astaxie/beego/validation"
    "strings"
)

type FoodForm struct {
    Id    int         `form:"-"`
    Name  string `form:"name,text,foood"`
}

func (f FoodForm) IsValidate() (bool, map[string]string) {
    valid := validation.Validation{}
    valid.Required(f.Name, "name")
    valid.MaxSize(f.Name, 50, "nameMax")
    valid.MinSize(f.Name, 2, "nameMin")

    erros := map[string]string{"name":""}

    if valid.HasErrors() {
        for _, err := range valid.Errors {
            if strings.HasPrefix(err.Key, "name"){
                erros["name"] += err.Message + "; "
            }

        }
    }
    return !valid.HasErrors(), erros
}