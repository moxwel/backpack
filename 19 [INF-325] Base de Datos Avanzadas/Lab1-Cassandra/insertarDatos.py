from cassandra.cluster import Cluster
import pandas as pd

# Conectar a Cassandra
cluster = Cluster(['127.0.0.1'], port=7001)
session = cluster.connect('universia_keyspace')

# Cargar el archivo Excel
df = pd.read_excel('postulaciones.xlsx', sheet_name=0)

def insertar_datos():
    # Insertar los datos
    for index, row in df.iterrows():
        if row['PACE'] == None:
            pace = ""
        else: 
            pace = row['PACE']

        query = """
            INSERT INTO postulantes_por_carrera (
                periodo, carrera, puntaje_psu, region, sexo, cedula, preferencia, matriculado,
                facultad, puntaje, grupo_depen, latitud, longitud, puntaje_nem, pace, gratuidad
            ) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)
            """
        session.execute(query, (
            int(row['PERIODO']), row['CARRERA'], int(row['PSU_PROMLM']), row['REGION'], row['SEXO'], int(row['CEDULA']),
            int(row['PREFERENCIA']), row['MATRICULADO'], row['FACULTAD'], int(row['PUNTAJE']), row['GRUPO_DEPEN'],
            float(row['LATITUD']), float(row['LONGITUD']), int(row['PTJE_NEM']),str(pace) , row['GRATUIDAD ']
        ))

    print("Datos insertados con éxito")


def query1():

    query="""
        SELECT carrera, periodo, cedula FROM postulantes_por_carrera
        WHERE carrera='MEDICINA'
        """
    
    resp = session.execute(query)

    for fila in resp:
        print(fila)

query1()