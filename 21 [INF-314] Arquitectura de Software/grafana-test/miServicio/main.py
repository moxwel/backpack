from fastapi import FastAPI
import logging
import os

logging.basicConfig(level=logging.INFO, format='[LOG] %(levelname)s:\t%(message)s')

env_nombre = os.getenv("NOMBRE_SERVICIO", "Sin Nombre")

app = FastAPI()

@app.get("/")
def read_root():
    logging.info("Se llamó a la ruta raíz")
    return {"message": "Hola Mundo", "serviceName": env_nombre}

@app.get("/info")
def logging_info():
    logging.info("Este es un mensaje de info")
    return {"message": "Log info emitido"}

@app.get("/warning")
def logging_warning():
    logging.warning("Este es un mensaje de warning")
    return {"message": "Log warning emitido"}

@app.get("/error")
def logging_error():
    logging.error("Este es un mensaje de error")
    return {"message": "Log error emitido"}

@app.get("/debug")
def logging_debug():
    logging.debug("Este es un mensaje de debug")
    return {"message": "Log debug emitido"}
