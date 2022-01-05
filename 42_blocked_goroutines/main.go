package main


// В выступлении «Go Concurrency Patterns» на конференции Google I/O в 2012-м Роб Пайк рассказал о нескольких
// фундаментальных concurrency-шаблонах. Один из них — извлечение первого результата.
// Для каждой копии (replica) поиска функция запускает отдельную горутину. Каждая из горутин отправляет свои поисковые
// результаты в канал результатов. Возвращается первое значение из канала.
//
// А что с результатами от других горутин? И что насчёт них самих?
//
// В функции First() канал результатов не буферизован. Это значит, что возвращается только первая горутина.
// Все остальные застревают в попытке отправить свои результаты. Получается, что если у вас более одной копии (replica), то при каждом вызове происходит утечка ресурсов.

// First
// Решение №1
// Чтобы этого избежать, все горутины должны завершиться (exit). Одно из возможных решений: использовать достаточно
// большой буферизованный канал результатов, способный вместить все результаты.
func First(query string, replicas ...Search) Result {
	c := make(chan Result,len(replicas))
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}


// First2
// Решение №2
// использовать выражение select со сценарием (case) default и буферизованный канал на одно значение.
// Сценарий default позволяет быть уверенным, что горутина не застряла, даже если канал результатов не может принимать сообщения.
func First2(query string, replicas ...Search) Result {
	c := make(chan Result,1)
	searchReplica := func(i int) {
		select {
		case c <- replicas[i](query):
		default:
		}
	}
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}


// First3
// Решение №3
// Также можно использовать специальный канал отмены (special cancellation channel) для прерывания рабочих горутин.
func First3(query string, replicas ...Search) Result {
	c := make(chan Result)
	done := make(chan struct{})
	defer close(done)
	searchReplica := func(i int) {
		select {
		case c <- replicas[i](query):
		case <- done:
		}
	}
	for i := range replicas {
		go searchReplica(i)
	}

	return <-c
}