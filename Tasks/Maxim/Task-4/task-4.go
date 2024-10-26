/* Task 4. Написать lru кэш

Кеш имеет фиксированный максимальный размер.
Если после вставки нового элемента превышается максимально допустимый размер кеша, из него нужно удалить самый старый элемент.

	Пример:
cache = new LRU(capacity=2)

cache.put(A, 1); // cache is {A=1)
cache.put(B, 2): // cache is {A=1, B=2}
cache.get(A); // return 1
cache.put(C, 3); // {A=1, C=3}, element B=2 was evicted
cache.get(B); // returns -1 (not found)
cache.put(D, 4); // (D=4, C=3}
cache.get(A); // return -1 (not found)
cache.get(C); // return 3
cache.get (D); // return 4

A D C

*/

package main
