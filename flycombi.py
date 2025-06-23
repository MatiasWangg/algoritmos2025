import sys
import utils as u
from grafo import Grafo

class SistemaVuelos:
    def __init__(self):
        pass
    
    def cargar_aeropuertos(self, aeropuertos):
        pass

    def cargar_vuelos(self, vuelos):
        pass



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
        comandos = u.procesar_entrada(linea.rstrip())

        if comandos[0] == "camino_mas":
         pass

        elif comandos[0] == "camino_escalas":
            pass
            
        elif comandos[0] == "centralidad":
           pass
            
        elif comandos[0] == "nueva_aerolinea":
            pass
                
        elif comandos[0] == "itinerario":
            pass
                
        elif comandos[0] == "exportar_kml":
            pass

main()
