package Iterator


/*
golang 之 迭代器模式
 */


type user struct {
	name string
	age  int
}

 // 先实现一个Iterator的接口，包含 get_next 和 has_next方法
type Iterator interface {
	getNext() *user
	hasNext() bool
}


// 实现一个collection create方法返回Iterator的实现
type Collection interface {
	Create() Iterator
}



// 实现collection接口

type UserCollection struct {
	users []*user

}

func(u *UserCollection) Create() Iterator{
	return &UserIter{
		users: u.users,
	}
}



type UserIter struct {
	index int
	users []*user

}


func(u *UserIter) getNext()*user{
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}


func(u *UserIter) hasNext()bool{
	if u.index < len(u.users) {
		return true
	}
	return false
}