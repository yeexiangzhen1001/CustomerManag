package service

import (
	"../model"
	"fmt"
)
//该CustomerService，完成对Customer的增删改查
type CustomerService struct {
	customers []model.Customer
	//表示当前切片已经含有多少个客户
	//该字段后面，还可以作为新客户的id+1
	customerNum int
}

//编写一个方法，可以返回*CustomerSevice
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "11111111111", "1111111111@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}
//返回客户切片
func (this *CustomerService) List() []model.Customer{
	return this.customers
}

//添加客户到customers切片
func (this *CustomerService) Add (customer model.Customer) bool{
	//确定一个分配id的规则，就是添加的顺序
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers,customer)
	return true
}
//根据id删除客户（从切片中删除）
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)

	if index == -1 {
		fmt.Println("查无此人！")
		return false
	}
	//从切片中删除一个元素
	this.customers = append(this.customers[:index],this.customers[index+1:]...)
	return true
}
//根据id查找客户在切片中对应的下标。如果没有该客户则返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1
	//遍历this.customers 切片
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			//找到
			index = i
		}
	}
return index
}

func (this *CustomerService) GetinfoById(id int) model.Customer {
	i := id - 1
	return this.customers[i]
}
//根据id修改客户信息
func (this *CustomerService) Update(id int, customer model.Customer) bool {
	for i := 0; i < len(this.customers); i++ {
		if  this.customers[i].Id == id {

			this.customers[i].Name = customer.Name
			this.customers[i].Gender = customer.Gender
			this.customers[i].Age = customer.Age
			this.customers[i].Phone = customer.Phone
			this.customers[i].Email = customer.Email
		}
	}
	return true
}