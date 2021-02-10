package main

import "fmt"

/*

命令模式
 */


type command interface {
	execute()
}

type button struct {
	command command
}

type device interface {
	on()
	off()
}

func (b *button) press()  {
	b.command.execute()
}

type onCommand struct {

	device device
}



func (o *onCommand)deviceOn()  {
	o.device.on()
}


type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}


// 具体接收者
type tv struct {
	isRunning bool
}


func (t *tv)on(){
	t.isRunning = true
	fmt.Println("turning tv on")
}




func (t *tv)off(){
	t.isRunning = false
	fmt.Println("turning tv off")
}



type reader interface {
	read()
}

type wirter interface {
	wirte()
}

type RW struct {


}

func (rw *RW)read()  {


}


func (rw *RW)wirter()  {


}




func main()  {

}





