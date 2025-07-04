#!/usr/bin/python3
import sys
import csv
import utils as u
import biblioteca as b
from grafo import Grafo


class SistemaVuelos:
    def __init__(self):
        self.conexiones = {} 
        self.aeropuertos_en_ciudad = {}  
        self.ciudad_de_aeropuerto = {} 
        self.grafo_precio = Grafo(es_dirigido=False) 
        self.grafo_tiempo = Grafo(es_dirigido=False) 
        self.grafo_frecuencia = Grafo(es_dirigido=False)
        self.ultima_ruta = [] 

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

            tiempo = float(informacion[2])
            precio = float(informacion[3])
            cantidad_vuelos = informacion[4]
            freq = 100 * float(cantidad_vuelos) / vuelos_totales
            if freq != 0:
                peso_freq = 1 / freq
            else:
                peso_freq = float('inf')
                
            self.grafo_frecuencia.agregar_arista(aeropuerto1, aeropuerto2, peso_freq)
            self.grafo_precio.agregar_arista(aeropuerto1, aeropuerto2, precio)
            self.grafo_tiempo.agregar_arista(aeropuerto1, aeropuerto2, tiempo)
            

            self.conexiones[aeropuerto1 + "-" + aeropuerto2] = informacion

    def caminoMas(self, grafo, desde, hasta):
        padres,dist = b.camino_minimo_dijkstra(grafo, desde,hasta)
        act = hasta
        camino = []
        while act != desde:
            camino.append(act)
            act = padres[act]
        camino.append(desde)
        camino.reverse()
        self.ultima_ruta = [(camino[i], camino[i+1]) for i in range(len(camino)-1)]
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
        self.ultima_ruta = [(mejor_camino[i], mejor_camino[i+1]) for i in range(len(mejor_camino)-1)]

        return mejor_camino
    
    def centralidad(self, grafo, k):
        centralidades = b.centralidad(grafo)
        centralidades_topK = u.topK(centralidades, k)
        return centralidades_topK
    
    def nueva_aerolinea(self, grafo, archivo):
        arbol = b.mst_prim(grafo)
        vuelos_escribir = []
        visitados = set() 

        for v in arbol.obtener_vertices():
            for w in arbol.adyacentes(v):
                nodo1, nodo2 = sorted([v, w])
                clave1 = f"{nodo1}-{nodo2}"
                clave2 = f"{nodo2}-{nodo1}"
                datos = self.conexiones.get(clave1) or self.conexiones.get(clave2)
                
                if datos and (clave1 not in visitados):
                    vuelos_escribir.append(datos)
                    visitados.add(clave1)

        with open(archivo, "w", newline='', encoding="utf-8") as salida:
            writer = csv.writer(salida)
            for vuelo in vuelos_escribir:
                writer.writerow(vuelo)

    
    def itinerario(self, grafo, archivo_itinerario):
        informacion = u.procesar_informacion(archivo_itinerario)
        nuevo_grafo = Grafo(es_dirigido=True)
        for ciudad in informacion[0]:
            nuevo_grafo.agregar_vertice(ciudad)
        for orden in informacion[1:]:
            nuevo_grafo.agregar_arista(orden[0], orden[1])
        nuevo_orden = b.orden_topologico(nuevo_grafo)
        print(", ".join(nuevo_orden))

        self.ultima_ruta = []  

        for i in range(len(nuevo_orden) - 1):
            ciudad_origen = nuevo_orden[i]
            ciudad_destino = nuevo_orden[i + 1]

            min_costo = float("inf")
            mejor_camino = []

            for aeropuerto_origen in self.aeropuertos_en_ciudad.get(ciudad_origen, []):
                for aeropuerto_destino in self.aeropuertos_en_ciudad.get(ciudad_destino, []):
                    camino_actual, costo_actual = self.caminoMas(self.grafo_tiempo, aeropuerto_origen, aeropuerto_destino)
                    if costo_actual < min_costo:
                        min_costo = costo_actual
                        mejor_camino = camino_actual

            if mejor_camino:
                mejor_camino.reverse()
                print(" -> ".join(mejor_camino))
                
                self.ultima_ruta.extend([(mejor_camino[i], mejor_camino[i+1]) for i in range(len(mejor_camino)-1)])
            else:
                print(f"No existe camino entre {ciudad_origen} y {ciudad_destino}")

    def exportar_kml(self, archivo):
        if not self.ultima_ruta:
            print("No hay ruta para exportar")
            return

        try:
            with open(archivo, "w", encoding="utf-8") as f:
                f.write('<?xml version="1.0" encoding="UTF-8"?>\n')
                f.write('<kml xmlns="http://earth.google.com/kml/2.1">\n')
                f.write('  <Document>\n')
                f.write('    <name>Ruta exportada</name>\n')
                f.write('    <description>Ruta generada por FlyCombi</description>\n')

                aeropuertos_visitados = []
                for origen, destino in self.ultima_ruta:
                    if origen not in aeropuertos_visitados:
                        aeropuertos_visitados.append(origen)
                    if destino not in aeropuertos_visitados:
                        aeropuertos_visitados.append(destino)

                for aeropuerto in aeropuertos_visitados:
                    info = self.ciudad_de_aeropuerto.get(aeropuerto)
                    if info:
                        codigo_iata = info[1]
                        lat, lon = info[2], info[3]
                        f.write('    <Placemark>\n')
                        f.write(f'      <name>{codigo_iata}</name>\n') 
                        f.write('      <Point>\n')
                        f.write(f'        <coordinates>{lon},{lat}</coordinates>\n')
                        f.write('      </Point>\n')
                        f.write('    </Placemark>\n')

                for origen, destino in self.ultima_ruta:
                    info_origen = self.ciudad_de_aeropuerto.get(origen)
                    info_destino = self.ciudad_de_aeropuerto.get(destino)
                    if not info_origen or not info_destino:
                        continue

                    lat_o, lon_o = info_origen[2], info_origen[3]
                    lat_d, lon_d = info_destino[2], info_destino[3]

                    f.write('    <Placemark>\n')
                    f.write('      <LineString>\n')
                    f.write(f'        <coordinates>{lon_o},{lat_o} {lon_d},{lat_d}</coordinates>\n')
                    f.write('      </LineString>\n')
                    f.write('    </Placemark>\n')

                f.write('  </Document>\n')
                f.write('</kml>\n')

            print("OK")

        except Exception as e:
            print(f"Error al exportar KML: {e}")




def main():
    argumentos = sys.argv
       
    if len(argumentos) != 3:
        sys.exit()
    aeropuertos, vuelos = argumentos[1], argumentos[2]
    informacion_aeropuerto = u.procesar_informacion(aeropuertos)
    informacion_vuelos = u.procesar_informacion(vuelos)

    
    sistema = SistemaVuelos()
    sistema.cargar_aeropuertos(informacion_aeropuerto)
    sistema.cargar_vuelos(informacion_vuelos)
    
     
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
            if len(comandos) != 2:
                print("Error al utilizar el comando 'itinerario'")
                continue
            archivo = comandos[1]
            sistema.itinerario(sistema.grafo_tiempo, archivo)
                
        elif comandos[0] == "exportar_kml":
            if len(comandos) != 2:
                print("Error al utilizar el comando 'exportar_kml'")
                continue
            archivo = comandos[1]
            sistema.exportar_kml(archivo)

        else:
            print("ERROR:No se reconoce el comando ingresado")
main()
