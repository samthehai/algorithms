class PriorityQueueNode(object):
    def __init__(self, key, obj) -> None:
        super().__init__()
        self.key = key
        self.obj = obj

    def __repr__(self):
        return str(self.obj) + ': ' + str(self.key)

class PriorityQueue(object):
    def __init__(self) -> None:
        super().__init__()
        self.array = []

    def __len__(self):
        return len(self.array)

    def extract_min(self):
        return None

    def insert(self, key, obj):
        return None

    # def
