package controllers

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "app/models"
    "app/forms"
)

type FoodController struct {
    beego.Controller
}

func (c *FoodController) List() {
    o := orm.NewOrm()
    food := new(models.Food)
    var food_list []models.Food
    qs := o.QueryTable(food)
    qs.All(&food_list)

    c.Data["Food_List"] = food_list
    c.TplName = "food/_list.tpl"
}

func (this *FoodController) New() {
    if this.Ctx.Input.Is("POST") {
        form := forms.FoodForm{}
        if err := this.ParseForm(&form); err == nil {
            if is_ok, errors := form.IsValidate(); is_ok {
                o := orm.NewOrm()
                food := new(models.Food)
                food.Name = form.Name
                _, err := o.Insert(food)
                if err == nil {
                    this.Redirect(this.URLFor("FoodController.List"), 302)
                }else{
                    this.Abort("500")
                }
            }else{
                this.Data["errors"] = errors
            }
        }else
        {
            this.Abort("500")
        }
    }else{
        this.Data["errors"] = map[string]string{"name":""}
    }
    this.TplName = "food/_create.tpl"
}



