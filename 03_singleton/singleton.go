package singleton

import "sync"

//Singleton 是单例模式类
type Singleton struct{}

var singleton *Singleton

var once sync.Once	// 我尝试了去掉这个Once,外面引用singleton后地址也是相同的,
//TODO packaage是天然的单例模式吗? 感觉package是随着程序初始化创建的,是属于饿汉式单例

//GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}

/*

// init函数将在包初始化时执行，实例化单例
func init() {
	singleton = &Singleton{}
}
*/
