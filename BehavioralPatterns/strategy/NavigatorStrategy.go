package main

import "fmt"

/*
更简单的策略模式的实现
 */

// 根据客户端选择不同的导航策略


// 路径策略接口
type RouteStrategy interface {

	buildRoute()
}

type CarsStrategy struct {

}

func(s *CarsStrategy)buildRoute(){
	fmt.Println("this a car navigator strategy")
}


type PublicTransportStrategy struct {

}

func(s *PublicTransportStrategy)buildRoute(){
	fmt.Println("this a PublicTransport strategy")
}


type WalkingStrategy struct {

}

func(s *WalkingStrategy)buildRoute(){
	fmt.Println("this a WalkingStrategy strategy")
}


type Navigator struct {

	RouteStrategy RouteStrategy  // 将接口作为成员变量
}


func(n *Navigator)setRouteStrategy(r RouteStrategy){
	n.RouteStrategy = r  // 由客户端调用此方法将导航策略设置到 Navigator 中
}


func routestrategyExecutor(evict string){
	var n Navigator
	switch evict {
	case "car":
		n.setRouteStrategy(&CarsStrategy{})
	case "public":
		n.setRouteStrategy(&PublicTransportStrategy{})
	case "walking":
		n.setRouteStrategy(&WalkingStrategy{})
	}

	n.RouteStrategy.buildRoute()

}

func main() {
	//routestrategyExecutor("car")
	//routestrategyExecutor("walking")
	//routestrategyExecutor("public")
	var nav Navigator

	nav.setRouteStrategy(&CarsStrategy{})
	nav.RouteStrategy.buildRoute()
}