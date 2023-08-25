package abstract_factory

func getMainAndDetail(f OrderDaoFactory) {
	f.CreateOrderMainDao().SaveOrderMain()
	f.CreateOrderDetailDao().SaveOrderDetail()
}

func ExampleRDBOrderDaoFactory() {
	var f = &RDBOrderDaoFactory{}
	getMainAndDetail(f)
	// Output:
	// rdb main save
	// rdb detail save
}

func ExampleXMLOrderDaoFactory() {
	var f = &XMLOrderDaoFactory{}
	getMainAndDetail(f)
	// Output:
	// xml main save
	// xml detail save
}
