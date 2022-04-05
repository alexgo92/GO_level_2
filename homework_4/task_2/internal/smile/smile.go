package smile

import (
	"context"
	"fmt"
	"time"
)

func Run(ctx context.Context) {
	// создали тикет с интервалом в одну секунду
	// секунда выбрана исходя из ограничения по времени в задании
	ticker := time.NewTicker(time.Second)

	go func() {
		for {
			select {
			// каждую секунду ticker.C отправляет тик
			case <-ticker.C:
				fmt.Println("Smile")
			// при срабатывании cancelFunc выполнится данный кейс
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()
}
