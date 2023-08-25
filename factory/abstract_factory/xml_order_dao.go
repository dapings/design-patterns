package abstract_factory

import (
	"fmt"
)

// XMLOrderMainDao xml存储订单主记录
type XMLOrderMainDao struct{}

func (*XMLOrderMainDao) SaveOrderMain() {
	fmt.Print("xml main save\n")
}

// XMLOrderDetailDao xml存储订单详细记录
type XMLOrderDetailDao struct{}

func (*XMLOrderDetailDao) SaveOrderDetail() {
	fmt.Print("xml detail save\n")
}

// XMLOrderDaoFactory  xml 抽象工厂实现
type XMLOrderDaoFactory struct{}

func (*XMLOrderDaoFactory) CreateOrderMainDao() OrderMainDao {
	return &XMLOrderMainDao{}
}

func (*XMLOrderDaoFactory) CreateOrderDetailDao() OrderDetailDao {
	return &XMLOrderDetailDao{}
}
