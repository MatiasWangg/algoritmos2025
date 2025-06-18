from collections import deque

class Cola:
    def __init__(self):
        self.items = deque()

    def estaVacia(self):
        return len(self.items) == 0
    
    def encolar(self, item):
        self.items.append(item)

    def desencolar(self):
        return self.items.popleft()