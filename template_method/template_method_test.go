package template_method

func ExampleHTTPDownloader() {
	downloader := NewHTTPDownloader()
	downloader.Download("https://example.com/adb.zip")

	// Output:
	// prepare downloading
	// download https://example.com/adb.zip via http
	// http save
	// finish downloading
}

func ExampleFTPDownloader() {
	downloader := NewFTPDownloader()
	downloader.Download("ftp://example.com/adb.zip")

	// Output:
	// prepare downloading
	// download ftp://example.com/adb.zip via ftp
	// default save
	// finish downloading
}

func ExampleConcrete() {
	a := &ConcreteClassA{}
	a.TemplateMethod()

	b := &ConcreteClassB{}
	b.TemplateMethod()

	// Output:
	// 具体类A的模板方法
	// 具体类A的操作1
	// 具体类A的操作2
	// 具体类B的模板方法
	// 具体类B的操作1
	// 具体类B的操作2
}
