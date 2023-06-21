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
	playerMap := map[int]string{
		0: "玩家一",
		1: "玩家二",
		2: "玩家三",
		3: "玩家四",
		4: "玩家五",
		5: "玩家六",
	}
OuterLoop:
	for {
		for i := 0; i < numOfChambers; i++ {
			chamberWithBullet := rand.Intn(numOfChambers) // 随机选取一个弹药膛放置子弹
			player := playerMap[i]
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

//时间复杂度O(1)
