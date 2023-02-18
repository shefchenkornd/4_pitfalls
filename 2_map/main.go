package main

import "fmt"

/*
Q: зачем аллоцировать память в map?
R: m := make(map[int]int, 1000) | при такой записи эвакуация данных не произойдёт и наш код будет работать намного быстрее

Q: почему нельзя взять указатель на элемент map'ы?
R: потому что может произойти эвакуация данных из bucket'a

Q: почему порядок обхода случайный?
R: при итерации по map'e авторы языка используют fastrand() или fastrand64() [генераторы случайных чисел] func mapiterinit(t *maptype, h *hmap, it *hiter


Внутри map есть buckets
bucket = hash(key)
функция hash(key) должна обладать след.св-вами:
 * равномерность
 * быстрота
 * детерминированность
 * криптоустойчивость

Как обойтись без дженериков в map'е?
 * все операции внутри map выполняются с помощью unsafe.Pointer
 * Мета-информация о типе хранится в type descriptor
 * type descriptor предоставляет операции: hash, equal, copy
 * в качестве структуру используем maptype, который хранится в: [/usr/local/go/src/runtime/type.go]
 * в maptype описан type descriptor


path: [/usr/local/go/src/runtime/type.go]
type maptype struct {
	typ    _type
	key    *_type
	elem   *_type
	bucket *_type // internal type representing a hash bucket
	// function for hashing keys (ptr to key, seed) -> hash
	hasher     func(unsafe.Pointer, uintptr) uintptr
	keysize    uint8  // size of key slot
	elemsize   uint8  // size of elem slot
	bucketsize uint16 // size of bucket
	flags      uint32
}

// A header for a Go map.
path: [/usr/local/go/src/runtime/map.go]
type hmap struct {
	// Note: the format of the hmap is also encoded in cmd/compile/internal/reflectdata/reflect.go.
	// Make sure this stays in sync with the compiler's definition.
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	extra *mapextra // optional fields
}

path: [/usr/local/go/src/runtime/map.go]
v := m["k"] 	-> func mapaccess1(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer
v, ok := m["k"] -> func mapaccess2(t *maptype, h *hmap, key unsafe.Pointer) (unsafe.Pointer, bool)
m["k"] = 101 	-> func mapassign(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer
delete(m, "k")	-> func mapdelete(t *maptype, h *hmap, key unsafe.Pointer)

*/
func main() {
	m := make(map[int]int, 1000)

	// нельзя взять указатель на элемент map
	el := &m[123]
	fmt.Println(el) // invalid operation: cannot take address of m[123] (map index expression of type int)
}
