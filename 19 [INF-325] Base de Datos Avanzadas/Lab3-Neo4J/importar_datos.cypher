// Nodos y relaciones de Encargados y Sectores (RESPONSABLE DE)
LOAD CSV WITH HEADERS FROM 'file:///encargados_sectores.csv' AS csv_encargado
MERGE (e:Encargado {id: toInteger(csv_encargado.encargado_id)})
SET e.nombre = csv_encargado.encargado_nombre
MERGE (s:Sector {id: toInteger(csv_encargado.sector_id)})
MERGE (e)-[:RESPONSABLE_DE]->(s);

// Nodos de Arboles
LOAD CSV WITH HEADERS FROM 'file:///arboles.csv' AS csv_arboles
MERGE (a:Arbol {id: toInteger(csv_arboles.arbol_id)})
SET a.tipo = csv_arboles.arbol_tipo,
    a.nombreCientifico = csv_arboles.arbol_nombCi,
    a.longitud = toInteger(csv_arboles.arbol_lon),
    a.latitud = toInteger(csv_arboles.arbol_lat);

// Relaciones entre Arboles y Sectores (PERTENECE_A)
LOAD CSV WITH HEADERS FROM 'file:///arboles.csv' AS csv_arboles
MATCH (a:Arbol {id: toInteger(csv_arboles.arbol_id)})
MATCH (s:Sector {id: toInteger(csv_arboles.sector_id)})
MERGE (a)-[:PERTENECE_A]->(s);

// Nodos de Formularios
LOAD CSV WITH HEADERS FROM 'file:///formularios.csv' AS csv_formulario
MERGE (f:Formulario {id: toInteger(csv_formulario.form_id)})
SET f.fecha = toInteger(csv_formulario.fecha),
    f.fruta_num = toInteger(csv_formulario.fruta_num),
    f.fruta_kg = toInteger(csv_formulario.fruta_kg),
    f.fruta_foto = csv_formulario.fruta_foto,
    f.plaga_tipo = csv_formulario.plaga_tipo,
    f.plaga_grado = csv_formulario.plaga_grado,
    f.plaga_pobl = csv_formulario.plaga_pobl,
    f.plaga_foto = csv_formulario.plaga_foto,
    f.plaga_obs = csv_formulario.plaga_obs,
    f.enf_tipo = csv_formulario.enf_tipo,
    f.enf_grado = csv_formulario.enf_grado,
    f.enf_foto = csv_formulario.enf_foto,
    f.enf_obs = csv_formulario.enf_obs,
    f.asp_estado = csv_formulario.asp_estado,
    f.asp_foto = csv_formulario.asp_foto,
    f.asp_obs = csv_formulario.asp_obs;

// Relaciones Encargado y Formulario (REGISTRA)
LOAD CSV WITH HEADERS FROM 'file:///formularios.csv' AS csv_formulario
MATCH (e:Encargado {id: toInteger(csv_formulario.encargado_id)})
MATCH (f:Formulario {id: toInteger(csv_formulario.form_id)})
MERGE (e)-[:REGISTRA {fecha: toInteger(csv_formulario.fecha)}]->(f);

// Relaciones Arbol y Formulario (REGISTRADO_EN)
LOAD CSV WITH HEADERS FROM 'file:///formularios.csv' AS csv_formulario
MATCH (a:Arbol {id: toInteger(csv_formulario.arbol_id)})
MATCH (f:Formulario {id: toInteger(csv_formulario.form_id)})
MERGE (a)-[:REGISTRADO_EN]->(f);

// Relacion Encargado y Arbol (VISITA)
LOAD CSV WITH HEADERS FROM 'file:///formularios.csv' AS csv_formulario
MATCH (e:Encargado {id: toInteger(csv_formulario.encargado_id)})
MATCH (a:Arbol {id: toInteger(csv_formulario.arbol_id)})
MERGE (e)-[:VISITA {fecha: toInteger(csv_formulario.fecha)}]->(a);

// Relacion Formulario y Sector (ASOCIADO_A)
LOAD CSV WITH HEADERS FROM 'file:///formularios.csv' AS csv_formulario
MATCH (f:Formulario {id: toInteger(csv_formulario.form_id)})
MATCH (a:Arbol {id: toInteger(csv_formulario.arbol_id)})
MATCH (a)-[:PERTENECE_A]->(s:Sector)
MERGE (f)-[:ASOCIADO_A]->(s);
