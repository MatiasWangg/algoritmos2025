import sys
import csv
import utils as u
import biblioteca as b
from grafo import Grafo


class SistemaVuelos:
    def __init__(self):
        self.conexiones = {} #Clave:(aeropuerto, aeropuerto)  Valor: [datos]
        self.aeropuertos_en_ciudad = {}  #Clave:ciudad  Valores: [aeropuertos]
        self.ciudad_de_aeropuerto = {} #Clave:aeropuerto  Valor: [datos]
        self.grafo_precio = Grafo(es_dirigido=False) #camino_mas barato y nueva_aerolinea
        self.grafo_tiempo = Grafo(es_dirigido=False) #camino_mas rapido e itinerario
        self.grafo_frecuencia = Grafo(es_dirigido=False) #Para centralidad

    def cargar_aeropuertos(self, aeropuertos):
        for informacion in aeropuertos:
            ciudad = informacion[0]
            codigo_aeropuerto = informacion[1]
            self.ciudad_de_aeropuerto[codigo_aeropuerto] = informacion

            if ciudad not in self.aeropuertos_en_ciudad:
                self.aeropuertos_en_ciudad[ciudad] = [codigo_aeropuerto]
            else:
                self.aeropuertos_en_ciudad[ciudad].append(codigo_aeropuerto)
            self.grafo_precio.agregar_vertice(codigo_aeropuerto)
            self.grafo_tiempo.agregar_vertice(codigo_aeropuerto)
            self.grafo_frecuencia.agregar_vertice(codigo_aeropuerto)

    def cargar_vuelos(self, vuelos):
        vuelos_totales = 0
        for informacion in vuelos:
            vuelos_totales += float(informacion[4])
        
        for informacion in vuelos:
            aeropuerto1 = informacion[0]
            aeropuerto2 = informacion[1]

            tiempo = informacion[2]
            precio = informacion[3]
            cantidad_vuelos = informacion[4]
            freq = 100 * float(cantidad_vuelos)/vuelos_totales

            self.grafo_precio.agregar_arista(aeropuerto1, aeropuerto2, precio)
            self.grafo_tiempo.agregar_arista(aeropuerto1, aeropuerto2, tiempo)
            self.grafo_frecuencia.agregar_arista(aeropuerto1, aeropuerto2, freq)

            self.conexiones[aeropuerto1 + "-" + aeropuerto2] = informacion

    def caminoMas(self, grafo, desde, hasta):
        padres,dist = b.camino_minimo_dijkstra(grafo, desde,hasta)
        act = hasta
        camino = []
        while act != desde:
            camino.append(act)
            act = padres[act]
        camino.append(desde)
        return camino, dist[hasta]
    
    def caminoEscalas(self, grafo, desde, hasta):
        min_escalas = float("inf")
        mejor_camino = []

        for aeropuerto_origen in self.aeropuertos_en_ciudad.get(desde, []):
            padres, distancias = b.camino_minimo_bfs(grafo, aeropuerto_origen)
            for aeropuerto_destino in self.aeropuertos_en_ciudad.get(hasta, []):
                if distancias[aeropuerto_destino] < min_escalas:
                    min_escalas = distancias[aeropuerto_destino]
                    camino = []
                    actual = aeropuerto_destino
                    while actual is not None:
                        camino.append(actual)
                        actual = padres[actual]
                    mejor_camino = camino

        mejor_camino.reverse()
        return mejor_camino
    
    def centralidad(self, grafo, k):
        centralidades = b.centralidad(grafo)
        centralidades_topK = u.topK(centralidades, k)
        return centralidades_topK
    
    def nueva_aerolinea(self, grafo, archivo):
        arbol = b.mst_prim(grafo)
        vuelos_escribir = []
        for v in arbol.obtener_vertices():
            for w in arbol.adyacentes(v):
                nodo1 = v
                nodo2 = w
                if nodo1 > nodo2:
                    nodo1, nodo2 = nodo2, nodo1
                clave = (nodo1, nodo2)
                datos = self.conexiones.get("-".join(clave))
                if datos:
                    vuelos_escribir.append(datos)

        with open(archivo, "w", newline='', encoding="utf-8") as salida:
            writer = csv.writer(salida)
            for vuelo in vuelos_escribir:
                writer.writerow(vuelo)
        

def main():
    argumentos = sys.argv
    #Funciones para recibir la informacion de los argumentos   
    if len(argumentos) != 3:
        sys.exit()
    aeropuertos, vuelos = argumentos[1], argumentos[2]
    informacion_aeropuerto = u.procesar_informacion(aeropuertos)
    informacion_vuelos = u.procesar_informacion(vuelos)

    #Creo el TDA y cargo la informacion recibida
    sistema = SistemaVuelos()
    sistema.cargar_aeropuertos(informacion_aeropuerto)
    sistema.cargar_vuelos(informacion_vuelos)
    
    #Llamada a funcion que procese la entrada y se llame a respectivo funcion 
    for linea in sys.stdin:
        comandos = u.procesar_entrada(linea.rstrip(" "))

        if comandos[0] == "camino_mas":
            if len(comandos) != 4:
                print("Error al utilizar el comando 'camino_mas'")
                continue
            min_costo = float("inf")
            mejor_camino = []
            for aeropuerto_origen in sistema.aeropuertos_en_ciudad.get(comandos[2], []):
                for aeropuerto_destino in sistema.aeropuertos_en_ciudad.get(comandos[3], []):
                    if comandos[1] == "barato":
                        camino_actual, costo_actual = sistema.caminoMas(sistema.grafo_precio ,aeropuerto_origen, aeropuerto_destino)
                    elif comandos[1] == "rapido":
                        camino_actual, costo_actual = sistema.caminoMas(sistema.grafo_tiempo, aeropuerto_origen, aeropuerto_destino)
                    if costo_actual < min_costo:
                        min_costo = costo_actual
                        mejor_camino = camino_actual
            if mejor_camino:
                mejor_camino.reverse()
                print(" -> ".join(mejor_camino))
            else:
                print("No existe camino")

        elif comandos[0] == "camino_escalas":
            if len(comandos) != 3:
                print("Error al utilizar el comando 'camino_escalas'")
                continue
            origen = comandos[1]
            destino = comandos[2]
            camino = sistema.caminoEscalas(sistema.grafo_precio, origen, destino)
            if camino:
                print(" -> ".join(camino))
            else:
                print("No existe camino")
                
        elif comandos[0] == "centralidad":
            if len(comandos) != 2:
                print("Error al utilizar el comando 'centralidad'")
                continue
            k = int(comandos[1])
            mas_importantes = sistema.centralidad(sistema.grafo_frecuencia, k)
            if mas_importantes:
                print(", ".join(mas_importantes))
            else:
                print("No existe mas importantes")
            
        elif comandos[0] == "nueva_aerolinea":
            if len(comandos) != 2:
                print("Error al utilizar el comando 'nueva_aerolinea'")
                continue
            archivo = comandos[1]
            sistema.nueva_aerolinea(sistema.grafo_precio, archivo)
            print("OK")

        elif comandos[0] == "itinerario":
            pass
                
        elif comandos[0] == "exportar_kml":
            pass
        else:
            print("ERROR:No se reconoce el comando ingresado")
main()
