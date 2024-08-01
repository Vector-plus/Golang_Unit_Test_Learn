package designpatterns

import "errors"

// 集合祖先，定义集合中的各种方法
type Collection interface {
	getIterator() Iterater
}

// 迭代器
type Iterater interface {
	hasNext() bool
	getNext() (any, error)
}

type User struct {
	name string
	age  int64
}

// 实体数组集合
type List struct {
	index    int
	userList []any
}

type ListCollection struct {
	userList []any
}

func (uc *ListCollection) getIterator() Iterater {
	return &List{
		index:    0,
		userList: uc.userList,
	}
}

// 实现迭代器方法
func (list *List) hasNext() bool {
	return list.index < len(list.userList)
}

func (list *List) getNext() (user any, err error) {
	if list.index >= len(list.userList) {
		// return nil,errors.New("数组越界")
		err = errors.New("数组越界")
		return
	}

	user = list.userList[list.index]
	list.index++
	return
}
