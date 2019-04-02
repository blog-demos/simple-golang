/*
@Time :     2018/9/6 15:10 
@Author :   Hickey Q
@File :     gotest_test.go
@Software:  GoLand
*/

package gotest

import "testing"

func Test_Division_1(t *testing.T) {
    if i, e := Division(6, 2); i != 3 || e != nil{
        t.Error("除数函数测试未通过")
    } else {
        t.Log("第一个测试通过")
    }
}

func Test_Division_2(t *testing.T) {
    t.Error("未通过")
}

func Test_Division_3(t *testing.T) {
    if i, e := Division(6, 0); i != 3 || e != nil{
        t.Error("除数函数测试未通过")
    } else {
        t.Log("第一个测试通过")
    }
}