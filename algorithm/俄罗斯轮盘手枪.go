package main

import (
	"math/rand"
	"time"
)

/*
俄罗斯轮盘手枪游戏的算法十分简单。游戏开始，弹药膛中仅有一颗子弹，然后将弹药膛旋转，使得玩家不知道下一颗子弹在哪个位置。

接下来，每个玩家按照顺序依次拿起手枪，瞄准自己的头或者面部，然后扣动扳机。如果这个玩家没有听到子弹发射的声音，并且没有被击中，那么游戏继续，手枪交给下一个玩家。

否则，如果选择的这个弹药膛恰好有子弹，那么这个玩家就会被打中而死亡或受重伤。接下来，游戏也会立即结束。

这个游戏很危险，因为每个玩家都无法知道下一颗子弹放在哪个位置。因此，玩家承担着巨大的风险，玩这个游戏是非常不明智的。
*/
import (
	"fmt"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子

	numOfChambers := 6                            // 设置弹药膛数量
	chamberWithBullet := rand.Intn(numOfChambers) // 随机选取一个弹药膛放置子弹

	for i := 0; i < numOfChambers; i++ {
		if i == chamberWithBullet {
			fmt.Println("请玩家", i+1, "击发扳机！（有子弹！）")
			break
		} else {
			fmt.Println("请玩家", i+1, "击发扳机！（无子弹。）")
		}
	}
}

//时间复杂度O(1)
