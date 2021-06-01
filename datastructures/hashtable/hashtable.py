from typing import Any


class Item(object):
    def __init__(self, key, value) -> None:
        super().__init__()
        self.key = key
        self.value = value


class HashTable(object):
    def __init__(self, size) -> None:
        super().__init__()
        self.size = size
        self.table = [[] for _ in range(self.size)]

    def _hash_function(self, key) -> int:
        return key % self.size

    def set(self, key, value) -> None:
        hash_index = self._hash_function(key)
        for item in self.table[hash_index]:
            if item.key == key:
                item.value = value
                return

        self.table[hash_index].append(Item(key, value))

    def get(self, key) -> Any:
        hash_index = self._hash_function(key)
        for item in self.table[hash_index]:
            if item.key == key:
                return item.value

        raise KeyError('Key not found')

    def remove(self, key) -> None:
        hash_index = self._hash_function(key)
        for index,item in enumerate(self.table[hash_index]):
            if item.key == key:
                del self.table[hash_index][index]
                return

        raise KeyError('Key not found')
