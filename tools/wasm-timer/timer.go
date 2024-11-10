package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func playSound() {
	// 音を鳴らす方法。macOS/Linuxの場合、ベル音を出す。
	// Windowsの場合、別の方法で音を鳴らす処理が必要
	cmd := exec.Command("cmd", "/C", "echo ^G") // Windowsの場合のベル音
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error playing sound:", err)
	}
}

func startTimer(duration time.Duration) {
	ticker := time.NewTicker(1 * time.Second) // 1秒ごとに通知
	defer ticker.Stop()

	timerEnd := time.Now().Add(duration)

	for {
		select {
		case <-ticker.C:
			remaining := time.Until(timerEnd)
			if remaining <= 0 {
				fmt.Println("タイマー終了！")
				playSound() // 終了時に音を鳴らす
				return
			}
			fmt.Printf("残り時間: %v\n", remaining.Round(time.Second))
		}
	}
}

func main() {
	// タイマーのデフォルトは10分（600秒）
	defaultDuration := 10 * time.Minute

	// 引数があればそれを使用して、無ければデフォルト（10分）
	if len(os.Args) > 1 {
		// コマンドライン引数で指定された時間（分）
		var minutes int
		_, err := fmt.Sscanf(os.Args[1], "%d", &minutes)
		if err != nil {
			fmt.Println("引数の形式が無効です。数字で時間を指定してください。")
			return
		}
		defaultDuration = time.Duration(minutes) * time.Minute
	}

	// タイマー開始
	fmt.Printf("タイマーを%vに設定しました...\n", defaultDuration)
	startTimer(defaultDuration)
}
