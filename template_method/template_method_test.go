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
