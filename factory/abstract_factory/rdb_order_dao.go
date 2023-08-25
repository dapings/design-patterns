package abstract_factory

import (
	"fmt"
)

// RDBOrderMainDao 关系型数据库的 OrderMainDao 实现
type RDBOrderMainDao struct{}

func (*RDBOrderMainDao) SaveOrderMain() {
	fmt.Print("rdb main save\n")
}

// RDBOrderDetailDao 关系型数据库的 OrderDetailDao 实现
type RDBOrderDetailDao struct{}

func (*RDBOrderDetailDao) SaveOrderDetail() {
	fmt.Print("rdb detail save\n")
}

// RDBOrderDaoFactory 抽象工厂实现
type RDBOrderDaoFactory struct{}

func (*RDBOrderDaoFactory) CreateOrderMainDao() OrderMainDao {
	return &RDBOrderMainDao{}
}

func (*RDBOrderDaoFactory) CreateOrderDetailDao() OrderDetailDao {
	return &RDBOrderDetailDao{}
}
