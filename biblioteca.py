from grafo import Grafo
from cola import Cola
from heap import Heap
"""
Funciones:
Camino Minimo: Dijkstra y BFS
MST: Prim
Centralidad
Orden Topológico
"""
def grados_entrada(grafo):
    gr_entrada = {}
    for v in grafo.obtener_vertices():
        gr_entrada[v] = 0

    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            gr_entrada[w] = gr_entrada[w] + 1
    
    return gr_entrada

def orden_topologico(grafo):
    grados = grados_entrada(grafo)
    cola = Cola()
    for v in grafo.obtener_vertices():
        if grados[v] == 0:
            cola.encolar(v)

    resultado = []
    while not cola.estaVacia():
        v = cola.desencolar()
        resultado.append(v)
        for w in grafo.adyacentes(v):
            grados[w] -= 1
            if grados[w] == 0:
                cola.encolar(w)
    return resultado

def camino_minimo_dijkstra(grafo, origen, destino=None):
    distancia = {}
    padre = {}
    for v in grafo.obtener_vertices():
        distancia[v] = float('inf')
    distancia[origen] = 0
    padre[origen] = None
    heap = Heap()
    heap.encolar(origen, 0)
    while not heap.estaVacia():
        v = heap.desencolar()
        if destino != None and v == destino:
            return padre, distancia
        
        for w in grafo.adyacentes(v):
            nueva_distancia = distancia[v] + float(grafo.peso_arista(v, w))
            if nueva_distancia < distancia[w]:
                distancia[w] = nueva_distancia
                padre[w] = v
                heap.encolar(w, nueva_distancia)
    return padre, distancia

def camino_minimo_bfs(grafo, origen):
    distancia = {}
    padre = {}
    visitados = set()
    for v in grafo.obtener_vertices():
        distancia[v] = float('inf')
    distancia[origen] = 0
    padre[origen] = None
    visitados.add(origen)
    cola = Cola()
    cola.encolar(origen)

    while not cola.estaVacia():
        v = cola.desencolar()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                distancia[w] = distancia[v] + 1
                padre[w] = v
                visitados.add(w)
                cola.encolar(w)
    return padre, distancia


def mst_prim(grafo):
    v = grafo.vertice_aleatorio()
    visitados = set()
    visitados.add(v)
    heap = Heap()
    for w in grafo.adyacentes(v):
        heap.encolar((v, w), grafo.peso_arista(v, w))
    arbol = Grafo(es_dirigido= False, vertices_init= grafo.obtener_vertices())
    while not heap.estaVacia():
        (v, w) = heap.desencolar()
        if w in visitados:
            continue
        arbol.agregar_arista(v, w, grafo.peso_arista(v, w))
        visitados.add(w)
        for x in grafo.adyacentes(w):
            if x not in visitados:
                heap.encolar((w, x), grafo.peso_arista(w, x))
    return arbol


def ordenar(grafo, distancias):
    vertices = grafo.obtener_vertices()
    vertices_ordenados = sorted(vertices, key=lambda v: distancias[v])
    return vertices_ordenados

def centralidad(grafo):
    centralidad = {}
    for v in grafo.obtener_vertices():
        centralidad[v] = 0
    for v in grafo.obtener_vertices():
        padre, distancias = camino_minimo_dijkstra(grafo, v)
        centralidadAux = {}
        for w in grafo.obtener_vertices():
            centralidadAux[w] = 0
        vertices_ordenados = ordenar(grafo, distancias)
        for w in vertices_ordenados:
            if padre[w] is None:
                continue
            centralidadAux[padre[w]] += 1 + centralidadAux[w]
        for w in grafo.obtener_vertices():
            if w == v:
                continue
            centralidad[w] += centralidadAux[w]
    return centralidad
