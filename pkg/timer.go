package pkg

import (
	"fmt"
	"time"
)

func RunPomodoro(focusDuration, breakDuration, longBreakDuration, cycles int) {
	for i := 1; i <= cycles; i++ {
		// 集中時間
		fmt.Printf("🍅 Focus %d: %d minutes\n", i, focusDuration)
		countdown(focusDuration)

		// セッションが最後のサイクルでない場合、休憩
		if i < cycles {
			fmt.Printf("💤 Short Break: %d minutes\n", breakDuration)
			countdown(breakDuration)
		} else {
			// 最後のサイクルが終わったら長い休憩
			fmt.Printf("🌴 Long Break: %d minutes\n", longBreakDuration)
			countdown(longBreakDuration)
		}

		fmt.Println("🍅 Session Complete!")
	}
}

// タイマーのカウントダウン
func countdown(minutes int) {
	duration := time.Duration(minutes) * time.Minute
	ticker := time.NewTicker(time.Second)
	done := time.After(duration)

	for {
		select {
		case <-done:
			ticker.Stop()
			fmt.Println("\n⏰ Time's up!")
			return
		case <-ticker.C:
			// 時間経過ごとに何か表示する場合、ここに処理を追加
		}
	}
}
