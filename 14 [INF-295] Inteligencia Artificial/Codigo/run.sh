#!/bin/bash

# Verificar que el archivo 'main' existe
if [ ! -f main ]; then
    echo "No se encuentra el archivo 'main'"
    exit 1
fi

# Verificar que el primer argumento es un directorio
if [ ! -d $1 ]; then
    echo "El primer argumento debe ser un directorio"
    exit 1
fi

# Copiar archivo 'results_template.tsv' a la carpeta 'output'
cp results_template.tsv output

# Renombrar archivo 'results_template.tsv' a 'results.tsv'
mv output/results_template.tsv output/results.tsv

for file in $1/*
do
    # Verificar que el archivo no es un directorio
    if [ ! -d $file ]; then
        # Ejecutar el programa 'main' con el archivo como argumento
        ./main "$file"
    else
        echo "El archivo $file es un directorio."
    fi
done

ruta=$1
ultimo_dir=$(basename $ruta)

# Crear carpeta con el nombre del directorio
mkdir $ultimo_dir

# Mover todos los archivos de output a la carpeta creada
mv ./output/* $ultimo_dir
