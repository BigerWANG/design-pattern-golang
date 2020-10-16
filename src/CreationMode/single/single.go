package main

import (
	"sync"
	"fmt"
)

var lock = &sync.Mutex{}

type single struct {

}

var signleInstance *single

func getInstance() *single{
	if signleInstance == nil { // 第一层nil检查是要确保 signleInstance 实例在最开始为空. 并且防止在每次调用getInstance时都去消耗资源执行锁定操作
		lock.Lock() // 加上互斥锁
		defer lock.Unlock()
		if signleInstance == nil {

			fmt.Println("Createing signle instance now.")
			signleInstance = &single{}

		} else {
			fmt.Println("Signle instance is already created.")
		}
	}else {
		fmt.Println("Signle instance is already created.")

	}
	return signleInstance
}

func main() {
	for i:=0; i<30; i++{
		go getInstance()
	}

	fmt.Scanln()
}


