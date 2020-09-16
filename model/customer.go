package model

import "fmt"

//声明一个customer结构体，表示一个客户信息

type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string

}
//返回一个Customer的实例

func NewCustomer (id int, name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Id: id,
		Name: name,
		Gender: gender,
		Age: age,
		Phone: phone,
		Email: email,
	}
}

//第二种创建customer实例的办法，不带id
func NewCustomer2 (name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Name: name,
		Gender: gender,
		Age: age,
		Phone: phone,
		Email: email,
	}
}
//返回用户信息,格式化后的字符串
func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v  \t%v  \t%v  \t%v  \t%v  \t%v  \t",this.Id,this.Name,this.Gender,
		this.Age,this.Phone,this.Email)
	return info
}
