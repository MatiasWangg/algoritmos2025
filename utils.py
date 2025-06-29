import csv
from heap import Heap

def procesar_informacion(archivo):
    resultado = []
    with open(archivo, 'r', encoding='utf-8') as archivo_abierto:
        archivo_lectura = csv.reader(archivo_abierto)
        for linea in archivo_lectura:
            resultado.append(linea)
    return resultado

def procesar_entrada(entrada):
    entrada = entrada.strip()
    if not entrada:
        return []

    partes = entrada.split(',')
    primeros = partes[0].split(' ', 1)
    comando = primeros[0]

    comando_final = [comando]

    if len(primeros) > 1:
        comando_final.append(primeros[1])

  
    for parametro in partes[1:]:
        comando_final.append(parametro.strip())

    return comando_final

def topK(diccionario, k):
    resultado = []
    heap = Heap(is_min_heap=False)
    heap.heapify(diccionario)
    for _ in range(k):
        if heap.estaVacia():
            break
        resultado.append(heap.desencolar())
    return resultado