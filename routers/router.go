package routers

import (
	"app/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MealController{},"get:List")
    beego.Router("/api/list",&controllers.FoodController{},"get:List")
    beego.Router("/food/new",&controllers.FoodController{},"get:New")
    beego.Router("/food/new",&controllers.FoodController{},"post:New")
    beego.Router("/meal/new",&controllers.MealController{},"get:New")
    beego.Router("/meal/new",&controllers.MealController{},"post:New")
    beego.Router("/meels",&controllers.MealController{},"get:List")
}
