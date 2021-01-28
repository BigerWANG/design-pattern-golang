package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/gpmgo/gopm/modules/log"
)

// 讲状态操作抽象成接口
type state interface {
	addItem(int) 		   error
	requestItem() 		   error
	insertMoney(money int) error
	dispenseItem() 		   error
}


// 自动售卖器
// 处于4种不同的状态
// hasItem 		  	有商品
// noItem  		  	无商品
// itemRequested  	商品已请求
// hasMoney 		收到纸币

// 定义4种状态, 每种状态都实现了state接口定义的四种操作

type vendingMachine struct {
	hasItem 		state
	itemRequestd	state
	hasMoney		state
	noItem 			state

	currentState    state // 售卖机当前状态
	itemCount 		int
	itemPrice 		int

	}


func newVendingMachine(itemCount, itemPrice int) *vendingMachine{

	v := &vendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}

	hasItemState := &hasItemState{
		v,
	}

	itemRequestedState := &itemRequestedState{
		v,
	}
	hasMoneyState := &hasMoneyState{
		v,
	}
	noItemState := &noItemState{
		v,
	}

	v.SetState(hasItemState)

	v.hasItem = hasItemState
	v.itemRequestd = itemRequestedState

	v.hasMoney = hasMoneyState
	v.noItem = noItemState

	return v

}


// 对外暴露设置状态方法
func (v *vendingMachine)SetState(s state){
	v.currentState = s
}

func (v *vendingMachine) incrementItemCount(count int){
	fmt.Printf("add %d items\n", count)
	v.itemCount = v.itemCount + count
}

// 查询每个状态的方法
func (v *vendingMachine) requestItem() error{
	return v.currentState.requestItem()
}

func (v *vendingMachine)addItem(count int) error{
	return v.currentState.addItem(count)
}

func (v *vendingMachine)insertMoney(money int) error  {

	fmt.Printf("投币 %d 元", money)
	return v.currentState.insertMoney(money)
}

func (v *vendingMachine)dispenseItem()error{
	return v.currentState.dispenseItem()
}


// 处于4种不同的具体实现
// hasItem 		  	有商品
// noItem  		  	无商品
// itemRequested  	商品已请求
// hasMoney 		收到纸币


// hasItem 		  	有商品
type hasItemState struct {
	v *vendingMachine
}


func (i *hasItemState) requestItem() error{
	if i.v.itemCount == 0{
		i.v.SetState(i.v.noItem)
		return errors.New("No item present")
	}
	fmt.Println("Item requestd")
	i.v.SetState(i.v.itemRequestd)
	return nil
}


func (i *hasItemState)addItem(count int) error{
	fmt.Printf("%d items added\n", count)
	i.v.incrementItemCount(count)
	i.v.SetState(i.v.hasItem)
	return nil
}

func (i *hasItemState)insertMoney(money int) error {
	return errors.New("Please select item first")
}


func (i *hasItemState)dispenseItem()error{
	return errors.New("Please select item first")

}


// noItem  		  	无商品
type noItemState struct {
	v *vendingMachine
}


func (i *noItemState) requestItem() error{
	return errors.New("Item out of stock")
}


func (i *noItemState)addItem(count int) error {
	i.v.incrementItemCount(count)
	i.v.SetState(i.v.hasItem)
	return nil
}


func (i *noItemState)insertMoney(money int) error {
	return errors.New("Item out of stock")
}


func (i *noItemState)dispenseItem()error{
	return errors.New("Item out of stock")

}


// itemRequested  	商品已请求
type itemRequestedState struct {
	v *vendingMachine
}

func (i *itemRequestedState) requestItem() error{
	return errors.New("Item out of stock")
}


func (i *itemRequestedState)addItem(count int) error{
	i.v.incrementItemCount(count)
	i.v.SetState(i.v.hasItem)
	return nil
}


func (i *itemRequestedState)insertMoney(money int) error {
	return errors.New("Item out of stock")
}


func (i *itemRequestedState)dispenseItem()error{
	return errors.New("Item out of stock")

}



// hasMoney 		收到纸币
type hasMoneyState struct {
	v *vendingMachine
}

func (i *hasMoneyState) requestItem() error{
	return errors.New("Item out of stock")
}


func (i *hasMoneyState)addItem(count int) error{
	i.v.incrementItemCount(count)
	i.v.SetState(i.v.hasItem)
	return nil
}


func (i *hasMoneyState)insertMoney(money int) error {
	return errors.New("Item out of stock")
}


func (i *hasMoneyState)dispenseItem()error{
	return errors.New("Item out of stock")

}



func main() {
	vendingMachine := newVendingMachine(1, 10)

	err := vendingMachine.requestItem()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println()

	err = vendingMachine.addItem(2)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println()

	err = vendingMachine.requestItem()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatal(err.Error())
	}
}