# Введение

В этом репозитории собраны и подробно описаны:
* :bug: частые ошибки на Golang
* :rocket: лайфхаки на Golang
* :technologist: примеры кода на Golang

Всё это необходимо для того, чтобы человеку, начинающему изучать Golang было проще понять этот ЯП, его идеологию,
избежать "подводных" камней.

# Оглавление

1. 4 подводных камня на Go, на которые часто натыкаются
2. Вопрос по Golang на миллион долларов
3. Буферизированный ввод-вывод
4. Приведение типов
5. Каналы
6. Проверка на имплементацию структуры repo интерфейса Repo
7. Конкатенация строк 4 способа
8. defer и его принцип работы LIFO (англ. Last In, First Out – «последним пришёл — первым ушёл»)
9. Неопределенное количество параметров в функции [ставится многоточие: numbers ...int]
10. Goroutine: чтение из канала
11. Goroutine select
12. Полиморфизм на примере имплементации метода интерфейса и структур
13. Интерфейс io.Reader & io.Writer
14. sync.Mutex
15. net/http packages
16. Организация кода в Go
17. os package
18. Стандартные потоки вывода в Go
19. panic()
20. Указатели как параметры функции
21. Struct and nested structs
22. Именованные типы. Именованный тип основывается на уже существующем типе.
23. tests
    1. unit-test
    2. http test
24. sync.WaitGroup
25. Прохождение онлайн-заданий на одном из сайтов
26. переключатель типов (type switch)
27. Многомерные slices и array
28. Неизменяемость строк
29. Десериализация (unmarshalling) JSON-чисел в интерфейсные значения. По умолчанию Golang обращается с числовыми
    значениями в JSON как с числами float64
30. Сравнение struct, array, slice и map. Можно использовать оператор эквивалентности == для сравнения переменных
    структур, если каждое поле структуры можно сравнить с помощью этого оператора.
31. Перехват panic
32. Обновление и привязка значений полей в slice, array и map в выражениях for range
33. «Скрытые данные» в slices
34. Patterns: «Цепочка обязанностей» - это поведенческий паттерн проектирования, который позволяет передавать задачи на
    обработку последовательно по цепочке
35. «Повреждение» данных в slices или Three-index slices sl[:i:j], где j - это ёмкость исходного массива
36. «Устаревшие» slices
37. Методы и объявления типов. Если вам нужны методы из исходного типа, вы можете задать новый тип структуры, встроив
    исходный в качестве анонимного поля.
38. Как выбраться из кодовых блоков for switch и for select с помощью метки
39. Итерационные переменные и замыкания в выражениях for
40. Вычисление аргумента блока defer (Deferred Function Call Argument Evaluation)
41. Вызов блока defer в цикле FOR{}
42. Блокированные горутины и утечки ресурсов
43. Обновление полей значений в хеш-таблице
44. GOMAXPROCS, согласованность (concurrency) и параллелизм
45. Изменение порядка операций чтения и записи в горутинах
46. Разбираемся с пакетом Context в Golang
47. Channel Axioms by Dave Cheney
    1. A send to a nil channel blocks forever
    2. A receive from a nil channel blocks forever
    3. A send to a closed channel panics
    4. A receive from a closed channel returns the zero value immediately
    5. Can't close nil channel will panics
48. Curious Channels by Dave Cheney
    1. A closed channel never blocks
    2. A nil channel always blocks; каналу можно присвоить nil значение
49. Strings, bytes, runes and characters in Go:
    1. strings
    2. utf-8 and string literals
    3. code points and rune
    4. for range loop
50. Decode (unmarshal) JSON to struct
51. Decode (unmarshal) JSON to map
52. Marshal struct to JSON file
53. Reset slice (two ways)
54. Simple HTTP Server
55. How To Use Dates and Times in Go