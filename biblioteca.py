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