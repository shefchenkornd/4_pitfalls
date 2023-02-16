package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Пакет context в go позволяет нам передавать данные в нашу программу в каком-то «контексте».
// Context так же, как и таймаут, дедлайн или канал, сигнализирует прекращение работы и вызывает return.

// Лучшие практики
//
// 1. context.Background следует использовать только на самом высоком уровне, как корень всех производных контекстов.
// 2. context.TODО должен использоваться, когда вы не уверены, что использовать, или если текущая функция будет использовать context в будущем.
// 3. Отмены context рекомендуются, но эти функции могут занимать время, чтобы выполнить очистку и выход.
// 4. context.Value следует использовать как можно реже, и его нельзя применять для передачи необязательных параметров. Это делает API непонятным и может привести к ошибкам. Такие значения должны передаваться как аргументы.
// 5. Не храните context в структуре, передавайте их явно в функциях, предпочтительно в качестве первого аргумента.
// 6. Никогда не передавайте nil-context в качестве аргумента. Если сомневаетесь, используйте TODО.
// 7. Структура Context не имеет метода cancel, потому что только функция, которая порождает context, должна его отменять.

/*
Родит.контексты	->			context.Background								context.TODО
							/				\
				context.WithCancel  	context.WithValue
    			           /				 \
				context.WithTimeout			context.WithDeadline
 */

// Медленная функция
func sleepRandom(fromFunction string, ch chan int) {
	// Отложенная функция очистки
	defer func() { fmt.Println(fromFunction, "sleepRandom complete") }()

	// Выполним медленную задачу
	// В качестве примера,
	// «заснем» на рандомное время в мс
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	randomNumber := r.Intn(100)
	sleepTime := randomNumber + 100

	fmt.Println(fromFunction, "Starting sleep for", sleepTime, "ms")
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	fmt.Println(fromFunction, "Waking up, slept for ", sleepTime, "ms")

	// Напишем в канал, если он был передан
	if ch != nil {
		ch <- sleepTime
	}
}

// Функция, выполняющая медленную работу с использованием контекста
// Заметьте, что context - это первый аргумент
func sleepRandomContext(ctx context.Context, ch chan bool) {

	// Выполнение (прим. пер.: отложенное выполнение) действий по очистке
	// Созданных контекстов больше нет
	// Следовательно, отмена не требуется
	defer func() {
		fmt.Println("sleepRandomContext complete")
		ch <- true
	}()

	// Создаем канал
	sleepTimeChan := make(chan int)

	// Запускаем выполнение медленной задачи в горутине
	// Передаем канал для коммуникаций
	go sleepRandom("sleepRandomContext", sleepTimeChan)

	// Используем select для выхода по истечении времени жизни context'a
	select {
	case <-ctx.Done():
		// Если context отменен, выбирается этот случай.
		// Это случается, если заканчивается таймаут doWorkContext или
		// doWorkContext или main вызывает cancelFunction
		// Высвобождаем ресурсы, которые больше не нужны из-за прерывания работы
		// Посылаем сигнал всем горутинам, которые должны завершиться (используя каналы)
		// Обычно вы посылаете что-нибудь в канал,
		// ждете выхода из горутины, затем возвращаетесь
		// Или используете группы ожидания вместо каналов для синхронизации
		fmt.Println("sleepRandomContext: Time to return")

	case sleepTime := <-sleepTimeChan:
		// Этот вариант выбирается, когда работа завершается до отмены контекста
		fmt.Println("Slept for ", sleepTime, "ms")
	}
}

// Вспомогательная функция, которая в реальности может использоваться для разных целей
// Здесь она просто вызывает одну функцию
// В данном случае, она могла бы быть в main
func doWorkContext(ctx context.Context) {

	// От контекста с функцией отмены создаём производный контекст с тайм-аутом
	// Таймаут 150 мс
	// Все контексты, производные от этого, завершатся через 150 мс
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Duration(150)*time.Millisecond)

	// Функция отмены для освобождения ресурсов после завершения функции
	defer func() {
		fmt.Println("doWorkContext complete")
		cancel()
	}()

	// Создаем канал и вызываем функцию контекста
	// Можно также использовать группы ожидания для этого конкретного случая,
	// поскольку мы не используем возвращаемое значение, отправленное в канал
	ch := make(chan bool)
	go sleepRandomContext(ctxWithTimeout, ch)

	// Используем select для выхода при истечении контекста
	select {
	case <-ctx.Done():
		// Этот случай выбирается, когда переданный в качестве аргумента контекст уведомляет о завершении работы
		// В данном примере это произойдёт, когда в main будет вызвана cancel
		fmt.Println("doWorkContext: Time to return")
	case <-ch:
		// Этот вариант выбирается, когда работа завершается до отмены контекста
		fmt.Println("sleepRandomContext returned")
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// Отложенная отмена высвобождает все ресурсы
	// для этого и производных от него контекстов
	defer func() {
		fmt.Println("Main Defer: canceling context")
		cancel()
	}()

	// Отмена контекста после случайного тайм-аута
	// Если это происходит, все производные от него контексты должны завершиться
	go func() {
		sleepRandom("Main", nil)
		cancel()
		fmt.Println("Main Sleep complete. canceling context")
	}()

	// Выполнение работы
	doWorkContext(ctx)
}