package template_method

import (
	"fmt"
)

// 如实例代码中，通用步骤在父类中实现（准备、下载、保存、收尾），而下载和保存的具体实现留到子类中，并且提供保存方法的默认实现。

type implement interface {
	download()
	save()
}

type Downloader interface {
	Download(uri string)
}

// 父类
type template struct {
	implement
	uri string
}

func newTemplate(impl implement) *template {
	return &template{implement: impl}
}

// Download 在父类中实现（准备、下载、保存、收尾）
func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Println("prepare downloading")
	t.implement.download()
	t.implement.save()
	fmt.Println("finish downloading")
}

// 保存方法的默认实现
func (t *template) save() {
	fmt.Println("default save")
}

type HTTPDownloader struct {
	*template
}

func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *HTTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

func (*HTTPDownloader) save() {
	fmt.Printf("http save\n")
}

type FTPDownloader struct {
	*template
}

func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *FTPDownloader) download() {
	fmt.Printf("download %s via ftp\n", d.uri)
}
