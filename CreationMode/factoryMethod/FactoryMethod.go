package main

import (
	"fmt"
)

type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// 实现iGun接口的gun对象，相当于所有具体gun实现的组合类
type gun struct {
	name string
	power int
}

func (g *gun) setName(name string){
	g.name = name
}

func (g *gun) getName() string{
	return g.name
}


func (g *gun) setPower(power int){
	g.power = power
}

func (g *gun) getPower() int{
	return g.power
}


// 具体产品
type Ak47 struct {
	gun  // 组合形式代替继承
	soundlike string
}

type Musket struct {
	gun
	soundlike string

} 


// 返回具体产品实例的方法
func newMusket() iGun {
	return &Musket{
		gun{
			"Musket",
			20,
		},
		"biubiubiu",

	}
}


func newAk47() iGun {

	return &Ak47{  // iGun的具体实现
		gun{
			"ak47",
			0,
		},
		"tututu",
	}
}


// 工厂方法，根据不同的参数返回不同的gun实例
func getGun(gunType string) (iGun, error){
		if gunType == "ak47"{
			return newAk47(), nil
		}
		if gunType == "musket"{
			return newMusket(), nil
		}
		return nil, fmt.Errorf("error gun type passed")
}


func main() {
	if ak47, err:= getGun("ak47"); err!=nil{
		panic(err)
	}else {
		fmt.Println(ak47)
		fmt.Println(ak47.getName())
		fmt.Println(ak47.getPower())
	}

	if musket, err := getGun("musket"); err!=nil{
		panic(err)
	}else {
		fmt.Println(musket)
		fmt.Println(musket.getName())
		fmt.Println(musket.getPower())
	}

	//var ak Ak47
	//ak = Ak47{gun: gun{"ak47", 100}, soundlike:"tututu"}
	//
	//ak.setPower(110)
	//fmt.Println(ak.getPower())
	//fmt.Println(ak.soundlike)
}

