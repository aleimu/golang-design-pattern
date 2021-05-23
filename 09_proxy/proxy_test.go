package proxy

import "testing"

func TestProxy(t *testing.T) {
	var sub Subject //抽象主题角色
	sub = &Proxy{}  //代理主题角色

	res := sub.Do()

	if res != "pre:real:after" {
		t.Fail()
	}
}

/*
以上涉及的是静态代理模式
动态代理模式涉及到反射 ,如下面代码的实现,用到了reflect
https://gitee.com/sqxwww/aop/blob/master/aop/aop.go#

Python如果实现所谓的动态代理功能

class Proxy(object):
    def __init__(self, target):
        self.target = target

    def __getattribute__(self, name):
        target = object.__getattribute__(self, "target")
        attr = object.__getattribute__(target, name)

        def newAttr(*args, **kwargs):  # 包装
            print "before print"
            res = attr(*args, **kwargs)
            print "after print"
            return res
        return newAttr

class RealHello(object):
    def prints(self, s):
        print 'hello', s

if __name__ == '__main__':
    t = RealHello()
    p = Proxy(t)
    p.prints("world")

*/