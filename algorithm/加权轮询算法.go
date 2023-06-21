package main

import (
	"context"
	"fmt"
	"time"
)

type node struct {
	name         string
	weight       int
	curentWeight int
}

func (n *node) Print() {
	fmt.Println(n.name, ":", n.curentWeight)
}

type nodeList struct {
	list map[string]*node
}

func (nl *nodeList) add(n *node) {
	nl.list[n.name] = n
}

// O(n)
func (nl *nodeList) selectMaxNode() bool {
	maxNode := new(node)
	total := 0
	for _, v := range nl.list {
		v.Print()
		if v.curentWeight > maxNode.curentWeight {
			maxNode = v
		}
		total += v.weight
	}

	for _, v := range nl.list {
		if v.name == maxNode.name {
			v.curentWeight -= total
		}
		v.curentWeight += v.weight

	}
	fmt.Println(maxNode.name)
	return false
}
func (nl *nodeList) Println(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop")
			return
		default:
			time.Sleep(1 * time.Second)
			nl.selectMaxNode()
		}
	}

}

func main() {
	nodeList := &nodeList{
		list: make(map[string]*node),
	}
	arr := map[string]int{
		"a": 4,
		"b": 2,
		"c": 1,
	}

	for k, v := range arr {
		nodeList.add(&node{
			name:         k,
			weight:       v,
			curentWeight: v,
		})
	}

	nodeList.Println(context.Background())
	// context.WithTimeout
	// ctx, _ := context.WithTimeout(context.Background(), time.Second*10000)
	// nodeList.Println(ctx)

	//context.WithCancel  手动控制时间
	// ctx, cancel := context.WithCancel(context.Background())
	// go nodeList.Println(ctx)
	// time.Sleep(5 * time.Second)
	// cancel()
	// fmt.Println("wait main go stop")
	// time.Sleep(20 * time.Second)
	// fmt.Println("main go stop")

	// 创建一个子节点的context,3秒后自动取消
	// ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	// go nodeList.Println(ctx)
	// fmt.Println("现在开始等待5秒,time=", time.Now().Unix())
	// time.Sleep(5 * time.Second)
	// fmt.Println("等待5秒结束,准备调用cancel()函数，发现两个子协程已经结束了，time=", time.Now().Unix())
	// cancel()

	//注：WithTimeout和WithDeadline知识时间参数不同

}
