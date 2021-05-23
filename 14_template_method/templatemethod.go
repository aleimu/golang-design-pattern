package templatemethod

import "fmt"

//Downloader 接口
type Downloader interface {
	Download(uri string)
}

// 模板中不变的部分
type template struct {
	implement
	uri string
}

// 模板中需要继承者改变的部分
type implement interface {
	download()
	save()
}

// newTemplate 创建模板
func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

//Download 模板方法
func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Print("finish downloading\n")
}

func (t *template) save() {
	fmt.Print("default save\n")
}

//HTTPDownloader 继承者
type HTTPDownloader struct {
	*template // 因为Golang不提供继承机制，需要使用匿名组合模拟实现继承
}

//NewHTTPDownloader 子类的实现
func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

// 子类对可变部分的重写
func (d *HTTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

// 子类对可变部分的重写
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
