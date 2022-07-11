-- Borrar tablas

DROP TABLE dieta;
DROP TABLE vacunacion;
DROP TABLE adiestramiento;
DROP TABLE atencion_veterinaria;
DROP TABLE historial_conflicto;
DROP TABLE mascota;

DROP TABLE alimentacion;
DROP TABLE vacuna;
DROP TABLE entrenamiento;
DROP TABLE veterinario;
DROP TABLE persona_responsable;



-- Borrar Vistas

DROP VIEW vista_consumo;
DROP VIEW vista_historial;
DROP VIEW vista_nombres;
DROP VIEW vista_responsable;
DROP VIEW vista_coincidencia;
DROP VIEW vista_recomendaciones;



-- Borrar Procedimientos

DROP PROCEDURE procedure_entrenamiento;
DROP PROCEDURE procedure_encontrar;



-- Borrar Disparadores
DROP TRIGGER trigger_veterinario;
DROP TRIGGER trigger_alimentacion;