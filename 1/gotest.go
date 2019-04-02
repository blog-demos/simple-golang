/*
@Time :     2018/9/6 15:08 
@Author :   Hickey Q
@File :     gotest.go
@Software:  GoLand
*/

package gotest

import "errors"

func Division(a, b float64) (float64, error){
    if b == 0{
        return 0, errors.New("除数不能为0")
    }

    return a / b, nil
}
