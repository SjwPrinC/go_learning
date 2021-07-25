package dao

import "database/sql"

type CustomerDao struct{}

type Customer struct {
	Name    string
	Age     int16
	Address string
}

//在dao层如果未找到对应记录不应该出现error，应该吃掉异常显示空值表示未找到，（如果是集合查找应该是返回长度为0的空数组）
//在service等业务层，应该根据业务需要包装相应的业务异常，如客户未找到
func (cd *CustomerDao) FindCustomer(id int32) (*Customer, error) {

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var customer *Customer
	err1 := db.QueryRow("select name, age, address from customers where status = 'Active' and id = ?", id).Scan(&customer)
	if err1 == sql.ErrNoRows {
		return nil, nil
	}

	return customer, err1
}
