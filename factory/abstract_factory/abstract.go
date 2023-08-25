package abstract_factory

// OrderDaoFactory 订单信息记录，抽象模式工厂接口
type OrderDaoFactory interface {
	CreateOrderMainDao() OrderMainDao
	CreateOrderDetailDao() OrderDetailDao
}

// OrderMainDao 主订单信息纪录
type OrderMainDao interface {
	SaveOrderMain()
}

// OrderDetailDao 订单详细信息记录
type OrderDetailDao interface {
	SaveOrderDetail()
}
