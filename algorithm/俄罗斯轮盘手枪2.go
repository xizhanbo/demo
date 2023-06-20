// 将俄罗斯轮盘手枪demo规则改为每开一枪之后进行一次弹药膛的转动
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func playGame() {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子

	numOfChambers := 6 // 设置弹药膛数量

	fmt.Println("弹药膛中有一颗子弹，游戏即将开始！")

	currentChamber := 0 // 当前弹药膛的位置
OuterLoop:
	for {
		for i := 0; i < numOfChambers; i++ {
			chamberWithBullet := rand.Intn(numOfChambers) // 随机选取一个弹药膛放置子弹
			var player string

			switch i {
			case 0:
				player = "玩家一"
			case 1:
				player = "玩家二"
			case 2:
				player = "玩家三"
			case 3:
				player = "玩家四"
			case 4:
				player = "玩家五"
			case 5:
				player = "玩家六"
			}

			if currentChamber == chamberWithBullet {
				fmt.Println(player, "中了子弹！他被淘汰了！")
				break OuterLoop
			} else {
				fmt.Println(player, "运气不错！没有中子弹，他依然活着。")
			}

			currentChamber = (currentChamber + 1) % numOfChambers // 将当前弹药膛位置加 1，并且取模操作保证不会超过弹药膛数量
		}
	}

}

func main() {
	playGame()
}
