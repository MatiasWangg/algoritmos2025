import random
class Grafo:

    def __init__(self, es_dirigido = False, vertices_init = []):
        self.dirigido = es_dirigido
        self.vertices = {}
        for vertice in vertices_init:
            self.agregar_vertice(vertice)
    
    def vertice_existe(self,v):
        return v in self.vertices
    
    def agregar_vertice(self, v):
        if v not in self.vertices:
            self.vertices[v] = {}  

    def borrar_vertice(self, v):
        if not self.vertice_existe(v):
            raise ValueError(f"No se encuentra el vertice {v} en el grafo")
        self.vertices.pop(v)
        for dato in  self.vertices.values():
            if v in dato:
                dato.pop(v)

    def agregar_arista(self, v, w, peso = 1):
        if not self.vertice_existe(v):
            raise ValueError(f"No se encuentra el vertice {v} en el grafo")
        elif not self.vertice_existe(w):
            raise ValueError(f"No se encuentra el vertice {w} en el grafo")

        self.vertices[v][w] = peso
        if not self.dirigido:
            self.vertices[w][v] = peso
    
    def borrar_arista(self, v, w):
        if not self.vertice_existe(v):
            raise ValueError(f"No se encuentra el vertice {v} en el grafo")
        elif not self.vertice_existe(w):
            raise ValueError(f"No se encuentra el vertice {w} en el grafo")

        self.vertices[v].pop(w)
        if not self.dirigido:
            self.vertices[w].pop(v)
    
    def estan_unidos(self, v, w):
        if not self.vertice_existe(v) or not self.vertice_existe(w):
            return False
        if self.dirigido:
            return w in self.vertices[v]
        else:
            return w in self.vertices[v] or v in self.vertices[w]
    
    def peso_arista(self, v, w):
        if self. estan_unidos(v, w):
            return self.vertices[v][w]
        return None
    
    def obtener_vertices(self):
        resultado = []
        for vertice in self.vertices:
            resultado.append(vertice)
        return resultado
    
    def adyacentes(self, v):
        adyacentes = []
        if self.vertice_existe(v):
            for w in self.vertices[v]:
                adyacentes.append(w)
        return adyacentes

    def vertice_aleatorio(self):
        vertices = self.obtener_vertices()
        return random.choice(vertices)
