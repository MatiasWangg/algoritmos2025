import heapq

class Heap:
    def __init__(self, is_min_heap = True):
        self.heap = []
        self.contador = 0
        if is_min_heap:
            self.multiplicador = 1
        else:
            self.multiplicador = -1
    
    def encolar(self, elemento, prioridad):
        heapq.heappush(self.heap, (self.multiplicador * prioridad, self.contador, elemento))
        self.contador += 1
    
    def desencolar(self):
        return heapq.heappop(self.heap)[2]
    
    def estaVacia(self):
        return len(self.heap) == 0
    
    def heapify(self, diccionario):
        self.heap = []
        for clave, valor in diccionario.items():
            elemento = (self.multiplicador * valor, self.contador, clave)
            self.heap.append(elemento)
        self.contador = len(diccionario)
        heapq.heapify(self.heap)