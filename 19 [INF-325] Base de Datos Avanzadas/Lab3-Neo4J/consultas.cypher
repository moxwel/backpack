// 001 ¿Cuantos kg de fruta se produjeron en la ultima semana?
MATCH (f:Formulario)
WHERE f.fecha <= 20241022 AND f.fecha >= 20241022 - 7
RETURN sum(f.fruta_kg)

// 002 ¿Cuantos aspersores estan presentando fallos en la ultima semana?
MATCH (f:Formulario)
WHERE f.asp_estado <> "operativo" AND f.fecha <= 20241022 AND f.fecha >=20241022-7
MATCH (a)-[:REGISTRADO_EN]->(f)
RETURN count(*)

// 003 ¿Cual es el sector con mas plagas en la ultima semana?
MATCH (s:Sector)<-[:ASOCIADO_A]-(f:Formulario)<-[:REGISTRADO_EN]-(a:Arbol)
WITH a, s, f
ORDER BY f.fecha DESC
WITH a, s, collect(f)[0] AS last_form
WHERE last_form.plaga_tipo <> "" AND last_form.fecha <= 20241022 AND last_form.fecha >= 20241022-7
RETURN s.id AS Sector, count(*) AS Cantidad_Plagas ORDER BY Cantidad_Plagas DESC

// 004 ¿Cuantos arboles tiene cada sector?
MATCH (a:Arbol)-[:PERTENECE_A]->(s:Sector)
RETURN s.id AS Sector, count(*) AS Cantidad_Arboles ORDER BY Cantidad_Arboles DESC

// 005 ¿Cuantos arboles sin repetir visito cada encargado en la ultima semana?
MATCH (e:Encargado)-[r:VISITA]->(a:Arbol)
WHERE r.fecha <= 20241022 AND r.fecha >= 20241022-7
RETURN e.nombre as Encargado, count(DISTINCT a) as Arboles_Visitados

// 006 ¿Cual es el % de cumplimiento de cada encargado la ultima semana?
MATCH (e:Encargado)-[:RESPONSABLE_DE]->(s:Sector)<-[:PERTENECE_A]-(a:Arbol)
WITH e, s, count(a) as total_arboles
MATCH (e)-[r:VISITA]->(arbol_visitado:Arbol)
WHERE r.fecha <= 20241022 AND r.fecha >= 20241022-7
WITH e, s, total_arboles, count(DISTINCT arbol_visitado) as arboles_visitados
RETURN 
    e.nombre as Encargado,
    s.id as Sector,
    total_arboles as Total_Arboles_Sector,
    arboles_visitados as Arboles_Visitados,
    round(100.0 * arboles_visitados / total_arboles) as Porcentaje_Cumplimiento
ORDER BY Porcentaje_Cumplimiento DESC;

// 007 ¿Que arboles le fata por visitar a cada encargado la semana pasada?
MATCH (e:Encargado)-[:RESPONSABLE_DE]->(s:Sector)<-[:PERTENECE_A]-(a:Arbol)
OPTIONAL MATCH (e)-[v:VISITA]->(a)
WHERE v.fecha <= 20241022 AND v.fecha >= 20241022-7
WITH e, s, a, v
WHERE v IS NULL
RETURN 
    e.nombre as Encargado,
    s.id as Sector,
    collect({
        id: a.id,
        tipo: a.tipo
    }) as Arboles_No_Visitados