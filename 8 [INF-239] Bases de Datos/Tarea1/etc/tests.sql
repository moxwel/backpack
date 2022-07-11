-- TESTS VISTA_NOMBRES

SELECT nombre FROM mascota GROUP BY nombre;






-- TESTS VISTA_CONSUMO

SELECT id_alimentacion, count(*) as "CANTIDAD_DIETAS" FROM dieta
    GROUP BY id_alimentacion;

SELECT id_alimentacion, tipo_alimentacion FROM alimentacion;



SELECT a.id_alimentacion, a.tipo_alimentacion, d.cantidad_dietas FROM alimentacion a
    LEFT JOIN (
        SELECT id_alimentacion, count(*) as "CANTIDAD_DIETAS" FROM dieta
            GROUP BY id_alimentacion
    ) d ON
        a.id_alimentacion = d.id_alimentacion;
        
SELECT a.id_alimentacion, a.tipo_alimentacion, d.cantidad_dietas FROM alimentacion a
    LEFT JOIN (
        SELECT id_alimentacion, count(*) as "CANTIDAD_DIETAS" FROM dieta
            GROUP BY id_alimentacion
    ) d ON
        a.id_alimentacion = d.id_alimentacion
    WHERE
        d.cantidad_dietas >= 5;
        
        
        
SELECT al.id_alimentacion, al.tipo_alimentacion, count(*) AS "CANTIDAD_DIETAS" FROM alimentacion al
    JOIN dieta di ON
        al.id_alimentacion = di.id_alimentacion
    GROUP BY
        al.id_alimentacion, al.tipo_alimentacion;


SELECT id_alimentacion, tipo_alimentacion
FROM (
    SELECT al.id_alimentacion, al.tipo_alimentacion, count(*) AS "CANTIDAD_DIETAS" FROM alimentacion al
        JOIN dieta di ON
            al.id_alimentacion = di.id_alimentacion
        GROUP BY
            al.id_alimentacion, al.tipo_alimentacion
) WHERE
    cantidad_dietas >= 5;





-- TESTS VISTA_COINCIDENCIA

SELECT rut_veterinario FROM veterinario;

SELECT rut_responsable FROM persona_responsable;

SELECT p.rut_responsable FROM persona_responsable p
    JOIN veterinario v ON
        p.rut_responsable = v.rut_veterinario;








SELECT m.id_mascota, m.nombre, m.rut_responsable FROM mascota m; 
        
SELECT m.id_mascota, m.nombre, m.rut_responsable FROM mascota m
    WHERE m.rut_responsable IN (
        SELECT p.rut_responsable FROM persona_responsable p
            JOIN veterinario v ON
                p.rut_responsable = v.rut_veterinario
    );
    
SELECT m.id_mascota FROM mascota m
    WHERE m.rut_responsable IN (
        SELECT p.rut_responsable FROM persona_responsable p
            JOIN veterinario v ON
                p.rut_responsable = v.rut_veterinario
    );
    
    
    
    
        
        
        

SELECT a.rut_veterinario, a.id_mascota, count(*) as "N_ATENCIONES" FROM atencion_veterinaria a
    WHERE a.id_mascota IN (
        SELECT m.id_mascota FROM mascota m
            WHERE m.rut_responsable IN (
                SELECT p.rut_responsable FROM persona_responsable p
                    JOIN veterinario v ON
                        p.rut_responsable = v.rut_veterinario
            )
    )
    GROUP BY a.rut_veterinario, a.id_mascota;


SELECT x.id_mascota, y.nombre FROM (
    SELECT a.rut_veterinario, a.id_mascota, count(*) as "N_ATENCIONES" FROM atencion_veterinaria a
        WHERE a.id_mascota IN (
            SELECT m.id_mascota FROM mascota m
                WHERE m.rut_responsable IN (
                    SELECT p.rut_responsable FROM persona_responsable p
                        JOIN veterinario v ON
                            p.rut_responsable = v.rut_veterinario
                )
        )
        GROUP BY a.rut_veterinario, a.id_mascota
) x
    JOIN mascota y ON
        x.id_mascota = y.id_mascota;

        
        
        
        
        
        
        
-- TESTS VISTA_RESPONSABLE
        
SELECT rut_responsable, nombre FROM persona_responsable;

SELECT rut_responsable, count(*) FROM mascota
    GROUP BY rut_responsable;
    
SELECT p.rut_responsable, p.nombre FROM persona_responsable p
    LEFT JOIN mascota m ON
        p.rut_responsable = m.rut_responsable;
        
SELECT p.rut_responsable, count(*) as "N_MASCOTAS" FROM persona_responsable p
    LEFT JOIN mascota m ON
        p.rut_responsable = m.rut_responsable
    GROUP BY
        p.rut_responsable;
        
    


        
        
-- TESTS VISTA_HISTORIAL

SELECT * FROM historial_conflicto;

SELECT id_mascota, count(*) as "N_CONFLICTO" FROM historial_conflicto
    GROUP BY id_mascota
    ORDER BY n_conflicto DESC;
    
SELECT m.id_mascota, m.nombre, c.n_conflicto FROM mascota m
    JOIN (
        SELECT id_mascota, count(*) as "N_CONFLICTO" FROM historial_conflicto
            GROUP BY id_mascota
            ORDER BY n_conflicto DESC
    ) c ON
        m.id_mascota = c.id_mascota
    ORDER BY c.n_conflicto DESC
    FETCH FIRST 10 ROWS ONLY;
    
    
    
    
    
    
-- TESTS VISTA_RECOMENDACIONES

SELECT
    ms.id_mascota,
    ms.nombre,
    ms.edad,
    ad.id_entrenamiento,
    ad.fecha,
    en.habilidad,
    di.id_alimentacion,
    di.fecha,
    al.tipo_alimentacion,
    al.edad_recomendada,
    av.fecha
    FROM
        mascota ms
    LEFT JOIN adiestramiento ad ON
        ms.id_mascota = ad.id_mascota
    JOIN entrenamiento en ON
        ad.id_entrenamiento = en.id_entrenamiento
    JOIN dieta di ON
        di.id_mascota = ms.id_mascota
    JOIN alimentacion al ON
        di.id_alimentacion = al.id_alimentacion
    JOIN atencion_veterinaria av ON
        av.id_mascota = ms.id_mascota
    WHERE
        ms.edad >= al.edad_recomendada
        AND
        ad.id_entrenamiento IN (501,502,503,504)
        AND
        ((av.fecha IN ad.fecha) OR (av.fecha IN di.fecha));
        
SELECT DISTINCT
    ms.id_mascota,
    ms.nombre
    FROM
        mascota ms
    LEFT JOIN adiestramiento ad ON
        ms.id_mascota = ad.id_mascota
    JOIN entrenamiento en ON
        ad.id_entrenamiento = en.id_entrenamiento
    JOIN dieta di ON
        di.id_mascota = ms.id_mascota
    JOIN alimentacion al ON
        di.id_alimentacion = al.id_alimentacion
    JOIN atencion_veterinaria av ON
        av.id_mascota = ms.id_mascota
    WHERE
        ms.edad >= al.edad_recomendada
        AND
        ad.id_entrenamiento IN (501,502,503,504)
        AND
        ((av.fecha IN ad.fecha) OR (av.fecha IN di.fecha));
        
        
        



-- Tests procedure_entrenamiento
        
SELECT ms.id_mascota, ms.nombre, ad.id_entrenamiento FROM mascota ms
    LEFT JOIN adiestramiento ad ON
        ad.id_mascota = ms.id_mascota;
        

--Mascotas entrenadas con habildiad
SELECT ms.id_mascota FROM mascota ms
    LEFT JOIN adiestramiento ad ON
        ad.id_mascota = ms.id_mascota
    WHERE
        ad.id_entrenamiento = 503;
        
        
SELECT m.id_mascota FROM mascota m
    WHERE m.id_mascota NOT IN (
        SELECT ms.id_mascota FROM mascota ms
            LEFT JOIN adiestramiento ad ON
                ad.id_mascota = ms.id_mascota
            WHERE
                ad.id_entrenamiento = 505
    );
        
        
        
        
        

SELECT id_mascota FROM mascota;


CREATE OR REPLACE PROCEDURE
    procedure_entrenamiento (param_habilidad IN number)
AS BEGIN

    FOR variable_mascota IN (
        SELECT m.id_mascota FROM mascota m
            WHERE m.id_mascota NOT IN (
                SELECT ms.id_mascota FROM mascota ms
                    LEFT JOIN adiestramiento ad ON
                        ad.id_mascota = ms.id_mascota
                    WHERE
                        ad.id_entrenamiento = param_habilidad
            )
    ) LOOP
    
        INSERT INTO adiestramiento VALUES (param_habilidad, variable_mascota.id_mascota, CURRENT_DATE);
        
    
    END LOOP;

END;
/

EXEC procedure_entrenamiento (501);
