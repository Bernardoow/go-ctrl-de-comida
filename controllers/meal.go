package controllers

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "app/models"
    "app/forms"
    "fmt"
)

type MealController struct {
    beego.Controller
}

func (this *MealController) List() {
    o := orm.NewOrm()
    meet := new(models.Meet)
    var meet_list []models.Meet
    qs := o.QueryTable(meet)
    qs.All(&meet_list)

    var new_meal_list []models.Meet

    for _, meal := range meet_list {
        o.LoadRelated(&meal, "Foods")
        new_meal_list = append(new_meal_list, meal)
    }
    this.Data["Meal_List"] = new_meal_list
    this.TplName = "meat/_list.tpl"
}

func (this *MealController) New() {
    if this.Ctx.Input.Is("POST") {
        fmt.Println(this.Input())
        form := forms.MeatForm{}
        if err := this.ParseForm(&form); err == nil {
            o := orm.NewOrm()
            // transaction
            tran_err := o.Begin()
            if tran_err != nil{
                this.Abort("500")
            }

            meet := new(models.Meet)
            meet.LactosePresence = form.LactosePresence
            _, err := o.Insert(meet)
            if err == nil {
                m2m := o.QueryM2M(meet, "Foods")
                _, err := m2m.Add(form.Foods)
                if err == nil {
                    err = o.Commit()
                    this.Redirect(this.URLFor("MealController.List"), 302)
                }else{
                    _ = o.Rollback()
                    this.Abort("500")
                }
            }else{
                _ = o.Rollback()
                this.Abort("500")
            }
        }
    }
    o := orm.NewOrm()
    food := new(models.Food)
    var food_list []models.Food
    qs := o.QueryTable(food)
    qs.All(&food_list)

    this.Data["Food_List"] = food_list
    this.TplName = "meat/_create.tpl"
}
