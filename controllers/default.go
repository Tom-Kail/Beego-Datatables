package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Time      time.Time `json:"time"`
}
type RspData struct {
	Draw            int32  `json:"draw"`
	RecordsTotal    int32  `json:"recordsTotal"`
	RecordsFiltered int32  `json:"recordsFiltered"`
	Data            []User `json:"data"`
}

func (this *MainController) TestAjax() {
	//	Data post from front end with map[string][ []string ] format
	//  Example of get 'draw' option for the map
	//    draw := strconv.Atoi(this.Input()["draw"][0])
	//  'draw' -- record request times. Get from datatables api and must sent back in integer type

	fmt.Println("Ajax data from DataTables: ", this.Input())
	fmt.Println("columns[0][data]: ", this.Input()["columns[0][data]"][0])
	fmt.Println("columns[1][searchable]:", this.Input()["columns[1][searchable]"][0])
	draw, _ := strconv.Atoi(this.Input()["draw"][0])
	//Send back []User
	users := make([]User, 0)
	//Test timer refresh
	TimeSecond := time.Now()
	fmt.Println("Now:", TimeSecond)
	users = append(users, User{"<button onclick=\"alert('dd');\">yes</button>", "cool", TimeSecond})
	users = append(users, User{"cool", "guy", TimeSecond})
	users = append(users, User{"ga", "te", TimeSecond})
	//Send back this object
	rst := RspData{
		Draw:            int32(draw),
		RecordsTotal:    3,
		RecordsFiltered: 3,
		Data:            users,
	}
	//Convert the object to json format
	this.Data["json"] = &rst
	this.ServeJson()
}
func (this *MainController) Get() {
	this.TplNames = "home.html"
}
