# Введение

В этом репозитории собраны и подробно описаны:
* :bug: частые ошибки на Golang
* :rocket: лайфхаки на Golang
* :technologist: примеры кода на Golang

Всё это необходимо для того, чтобы человеку, начинающему изучать Golang, было проще понять этот ЯП, его идеологию и
избежать "подводных" камней.

# Оглавление

1. Что такое slice?
2. Подводные камни в Golang, на которые часто натыкаются
3. Буферизированный ввод-вывод
4. Приведение типов
5. Channels
    1. close
    2. deadlock
    3. 
    4. 
    5. 
    6. 
6. Interface
    1. проверка Interface
    2. вложение (embedding)
    3. полиморфизм на примере имплементации метода interface{} и структур
    4. interface io.Reader & io.Writer
    5. сравнение interface{} с nil
    6. тип interface{} в переменной
7. Конкатенация строк 4 способа
8. defer и его принцип работы LIFO (англ. Last In, First Out – «последним пришёл — первым ушёл»)
9. Функция
    1. Неопределенное количество параметров в функции [ставится многоточие: numbers ...int]
    2. Функция как параметр другой функции
10. Goroutine: 
    1. reading from a channel
    2. select
11. sync.Mutex
    1. Mutex
    2. Cond
    3. waitGroup
    4. Once
12. net/http packages
13. Организация кода в Golang
14. os package
15. Стандартные потоки вывода в Golang
16. panic()
    1. простой пример с recovery()
    2. panic() в goroutine
17. Утечка горутины в заблокированном канале при отсутствии default в select{}
18. Паттерн «Fan-in» | асинхронный merge каналов
19. Generics
    


19. Указатели как параметры функции
20. Struct and nested structs
21. Именованные типы. Именованный тип основывается на уже существующем типе.
22. tests
    1. unit-test
    2. http test
23. sync.WaitGroup
24. Прохождение онлайн-заданий на одном из сайтов
25. переключатель типов (type switch)
26. Многомерные slices и array
27. Неизменяемость строк (смотрите пример с unsafe.Pointer())
28. Десериализация (unmarshalling) JSON-чисел в интерфейсные значения. По умолчанию Golang обращается с числовыми
    значениями в JSON как с числами float64
29. Сравнение struct, array, slice и map. Можно использовать оператор эквивалентности == для сравнения переменных
    структур, если каждое поле структуры можно сравнить с помощью этого оператора.
30. Перехват panic (переместил, можно удалять)
31. Обновление и привязка значений полей в slice, array и map в выражениях for range
32. «Скрытые данные» в slices
33. Patterns: «Цепочка обязанностей» - это поведенческий паттерн проектирования, который позволяет передавать задачи на
    обработку последовательно по цепочке
34. «Повреждение» данных в slices или Three-index slices sl[:i:j], где j - это ёмкость исходного массива
35. «Устаревшие» slices
36. Методы и объявления типов. Если вам нужны методы из исходного типа, вы можете задать новый тип структуры, встроив
    исходный в качестве анонимного поля.
37. Как выбраться из кодовых блоков for switch и for select с помощью метки
38. Итерационные переменные и замыкания в выражениях for
39. Вычисление аргумента блока defer (Deferred Function Call Argument Evaluation)
40. Вызов блока defer в цикле FOR{}
41. Блокированные горутины и утечки ресурсов
42. Обновление полей значений в хеш-таблице
43. GOMAXPROCS, согласованность (concurrency) и параллелизм
44. Изменение порядка операций чтения и записи в горутинах
45. context
46. Channel Axioms by Dave Cheney
    1. Sending to a nil channel blocks forever
    2. Receiving from a nil channel blocks forever
    3. A send to a closed channel panics
    4. Receiving from a closed channel returns the zero value immediately
    5. Can't close nil channel will panics
47. Curious Channels by Dave Cheney
    1. A closed channel never blocks
    2. A nil channel always blocks; a channel can be assigned a nil value
49. Strings, bytes, runes and characters in Golang:
    1. strings
    2. utf-8 and string literals
    3. code points and rune
    4. for range loop
49. Decode (unmarshal) JSON to struct
50. Decode (unmarshal) JSON to map
51. Marshal struct to JSON file
52. Reset slice (two ways)
53. Simple HTTP Server
54. How To Use Dates and Times in Golang
55. Copy map
55. Concurrency map
57. Reverse string
58. Summation
59. Find the next perfect square root | sqrt(n)
60. Write a function that returns both the minimum and maximum number of the given list/array.
61. Take 2 strings s1 and s2 including only letters from a to z. Return a new sorted string, the longest possible, containing distinct letters - each taken only once - coming from s1 or s2.
62. Your task is to sort a given string.
63. Such that each lowercase letter becomes uppercase and each uppercase letter becomes lowercase. For example:
64. Max consecutive strings.
68. Проблема с циклом for{} и pointer'ами
67. Build a pyramid-shaped tower, as an array/list of strings, given a positive integer number of floors. A tower block is represented with "*" character.
70. Список всех типов данных в Golang и их размер в bytes с помощью функции unsafe.Sizeof()
71. Изменить букву в строке через unsafe.Pointer() 
72. В языке Go нет наследования — есть только композиция (встраивание)
73. Написать простенькую реализацию cache используя map[string]string
74. Тестовое задание