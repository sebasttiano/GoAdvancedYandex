package retrier_test

import (
	"context"
	retrier "doc.com"
	"fmt"
	"time"
)

func Example() {
	op := func(_ context.Context) error {
		return fmt.Errorf("что-то пошло не так")
	}

	// Определяем контекст с ограничением по времени.
	opCtx, opCancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer opCancel()

	// Выполняем операцию op, переопределяя стандартные значения min и max.
	retrier.Do(opCtx, op, retrier.WithMinMaxDelay(50*time.Millisecond, 1*time.Second))
}
