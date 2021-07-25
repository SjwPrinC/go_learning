package service

import (
	"mydb/dao"

	"github.com/pkg/errors"
)

type CustomerService struct {
	customerDao *dao.CustomerDao
}

func (cs *CustomerService) FindCustomer(customerId int32) (*dao.Customer, error) {
	customer, err := cs.customerDao.FindCustomer(customerId)

	if err != nil {
		return nil, errors.Wrap(err, "an error occured when find customer")
	}

	if customer == nil {
		return nil, errors.New("no valid customer found")
	}

	return customer, nil
}
