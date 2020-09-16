package main

import (
	"../model"
	"../service"
	"fmt"
)


type customerView struct {
	//定义必要字段
	key string//接收用户输入
	loop bool //表示是否循环显示菜单
	//增加一个字段 customerService
	customerService *service.CustomerService
}

//显示主菜单
func (this *customerView) mainMenu() {
	for {
		fmt.Println("---------------客户信息管理系统---------------")
		fmt.Println("1 添 加 客 户")
		fmt.Println("2 修 改 客 户")
		fmt.Println("3 删 除 客 户")
		fmt.Println("4 客 户 列 表")
		fmt.Println("5 退 出 系 统")
		fmt.Print("请选择（1-5）：")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1" :
			this.add()
		case "2" :
			this.update()
		case "3" :
			this.delete()
		case "4" :
			this.list()
		case "5" :
			this.exit()
		default:
			fmt.Println("输入有误，请重新输入！")

		}
		if !this.loop {
			break
		}

	}
	fmt.Println("您以退出客户信息管理系统！")
}
//显示所有客户的信息
func  (this *customerView) list() {
	//首先获取当前所有客服的信息（存在切片中）
	customer := this.customerService.List()
	//显示
	fmt.Println("-------------------客户列表-------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱\t")
	for i := 0; i<len(customer); i++ {
		fmt.Println(customer[i].GetInfo())
	}
	fmt.Println("-----------------客户列表完成-----------------\n\n")
}
//得到用户的输入，信息构建新的客户，并完成添加
func (this *customerView) add() {
	fmt.Println("-------------------添加信息-------------------")
	fmt.Println("姓名：")
	name := " "
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := " "
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := " "
	fmt.Scanln(&phone)
	fmt.Println("邮件：")
	email := " "
	fmt.Scanln(&email)
	//构建一个新的customer实例
	//注意： id是唯一的，需要系统分配
	customer := model.NewCustomer2(name,gender,age,phone,email)
	//调用
	if this.customerService.Add(customer) {
		fmt.Println("添加完成！")
	}else {
		fmt.Println("添加失败！")
	}


}

//得到用户输入的ID，删除该ID对应的用户
func (this *customerView) delete() {
	fmt.Println("-------------------删除客户-------------------")
	fmt.Println("请输入待删除的客户id（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return//放弃删除
	}

	fmt.Println("确认是否删除（Y/N):")
	choice := " "
	fmt.Scanln(&choice)
	if choice == "Y"|| choice == "y" {

		//调用customerService delete的方法
		if this.customerService.Delete(id) {
			fmt.Println("删除成功！")
		}else {
			fmt.Println("删除失败，输入ID号不存在~")
		}
	}

}
//退出
func (this *customerView) exit() {
	fmt.Println("确认是否退出（Y/N）：")
	for {
		fmt.Scanln(&this.key)
		if this.key == "y" || this.key == "Y" || this.key == "N" || this.key == "n" {

			break
		}
		fmt.Println("您的输入有误,确认是否退出（Y/N）：")
	}
	if this.key == "y" || this.key == "Y" {
		this.loop = false
	}
}

//修改客户信息
func (this *customerView) update() {
	fmt.Print("请输入要修改的id：")
	id := 0
	fmt.Scanln(&id)
	if this.customerService.FindById(id) != -1 {
		customer := this.customerService.GetinfoById(id)
		fmt.Printf("姓名（%v）：", customer.Name)
		name := ""
		fmt.Scanln(&name)
		fmt.Printf("性别（%v）：", customer.Gender)
		gender := ""
		fmt.Scanln(&gender)
		fmt.Printf("年龄（%v）：", customer.Age)
		age := 0
		fmt.Scanln(&age)
		fmt.Printf("电话（%v）：", customer.Phone)
		phone := ""
		fmt.Scanln(&phone)
		fmt.Printf("邮箱（%v）：", customer.Email)
		email := ""
		fmt.Scanln(&email)
		customer2 := model.NewCustomer2(name, gender, age, phone, email)
		this.customerService.Update(id, customer2)
	} else {
		fmt.Println("输入id不存在，请重新输入~")
	}
}

func main() {
	//在main函数中，创建一个customerView，并运行显示主菜单
	customerView := customerView{
		key : " ",
		loop : true,
	}
	//完成对customerService的初始化
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}
