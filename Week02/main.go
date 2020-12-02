package main

import (
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/tools/go/ssa/interp/testdata/src/os"
)

type User struct {
	id   int
	name string
	age  int
}

func main() {
	err := biz()
	fmt.Printf("service: %+v\n",err)
	os.Exit(1)
}

func biz() error {
	ErrAge:=errors.New("你不行啊")
	u := User{
		id: 1,
	}
	user, err := dao(u)
	if err != nil{
		//err直接返回，不继续wrap，继续会造成双倍堆栈信息
		return err
	}
	if user.age <= 18{
		return errors.Wrap(ErrAge,"hhhhh")
	}
	return nil
}

func dao(u User)(user User,err error) {
	//db操作wrap err把根因返回
	return User{age: 18}, nil
}