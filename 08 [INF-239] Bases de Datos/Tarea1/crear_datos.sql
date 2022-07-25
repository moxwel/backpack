----- TABLAS INDEPENDIENTES -----

CREATE TABLE persona_responsable (
    rut_responsable         varchar2(10),
    nombre                  varchar2(30),
    edad                    number(3),
    direccion_residencial   varchar2(50),
    telefono                varchar2(15),
    
    PRIMARY KEY (rut_responsable)
);


CREATE TABLE veterinario (
    rut_veterinario         varchar2(10),
    nombre                  varchar2(30),
    edad                    number(3),
    universidad             varchar2(30),
    telefono                varchar2(15),
    tipo_consulta           varchar2(15),
    
    PRIMARY KEY (rut_veterinario)
);


CREATE TABLE alimentacion (
    id_alimentacion         number(6),
    tipo_alimentacion       varchar2(15),
    edad_recomendada        number(3),
    
    PRIMARY KEY (id_alimentacion)
);

CREATE TABLE vacuna (
    id_vacuna               number(6),
    nombre_vacuna           varchar2(15),
    
    PRIMARY KEY (id_vacuna)
);

CREATE TABLE entrenamiento (
    id_entrenamiento        number(4),
    habilidad               varchar2(20),
    
    PRIMARY KEY (id_entrenamiento)
);

---------------------------------





--------TABLAS DEPENDIENTES------

CREATE TABLE mascota (
    id_mascota              number(10),
    nombre                  varchar2(30),
    edad                    number(3),
    especie                 varchar2(6),
    raza                    varchar2(15),
    color                   varchar2(10),
    patron                  varchar2(30),
    razon_tenencia          varchar2(15),
    modo_obtencion          varchar2(10),
    rut_responsable         varchar2(10),
    
    PRIMARY KEY (id_mascota),
    FOREIGN KEY (rut_responsable)
        REFERENCES persona_responsable (rut_responsable)
);

CREATE TABLE dieta (
    id_alimentacion         number(6),
    id_mascota              number(10),
    fecha                   date,
    gramos                  number(4),
    
    PRIMARY KEY (id_alimentacion, id_mascota, fecha),
    FOREIGN KEY (id_alimentacion)
        REFERENCES alimentacion (id_alimentacion),
    FOREIGN KEY (id_mascota)
        REFERENCES mascota (id_mascota)
);

CREATE TABLE atencion_veterinaria (
    rut_veterinario         varchar2(10),
    id_mascota              number(10),
    fecha                   date,
    observaciones           varchar2(500),
    
    PRIMARY KEY (rut_veterinario, id_mascota, fecha),
    FOREIGN KEY (rut_veterinario)
        REFERENCES veterinario (rut_veterinario),
    FOREIGN KEY (id_mascota)
        REFERENCES mascota (id_mascota)
);

CREATE TABLE vacunacion (
    id_mascota              number(10),
    id_vacuna               number(6),
    fecha                   date,
    
    PRIMARY KEY (id_mascota, id_vacuna, fecha),
    FOREIGN KEY (id_mascota)
        REFERENCES mascota (id_mascota),
    FOREIGN KEY (id_vacuna)
        REFERENCES vacuna (id_vacuna)
);

CREATE TABLE historial_conflicto (
    id_conflicto            number(8),
    id_mascota              number(10),
    descripcion             varchar2(500),
    fecha                   date,
    
    PRIMARY KEY (id_conflicto),
    FOREIGN KEY (id_mascota)
        REFERENCES mascota (id_mascota)
);

CREATE TABLE adiestramiento (
    id_entrenamiento        number(4),
    id_mascota              number(10),
    fecha                   date,
    
    PRIMARY KEY (id_entrenamiento, id_mascota),
    FOREIGN KEY (id_entrenamiento)
        REFERENCES entrenamiento (id_entrenamiento),
    FOREIGN KEY (id_mascota)
        REFERENCES mascota (id_mascota)
);

---------------------------------





----- VISTAS -----

CREATE OR REPLACE VIEW vista_nombres AS (
    SELECT nombre FROM mascota GROUP BY nombre
);

CREATE OR REPLACE VIEW vista_consumo AS (
    SELECT id_alimentacion, tipo_alimentacion
    FROM (
        SELECT al.id_alimentacion, al.tipo_alimentacion, count(*) AS "CANTIDAD_DIETAS" FROM alimentacion al
            JOIN dieta di ON
                al.id_alimentacion = di.id_alimentacion
            GROUP BY
                al.id_alimentacion, al.tipo_alimentacion
    ) WHERE
        cantidad_dietas >= 5
) ORDER BY id_alimentacion ASC;

CREATE OR REPLACE VIEW vista_responsable AS (
    SELECT p.rut_responsable, count(*) as "N_MASCOTAS" FROM persona_responsable p
        LEFT JOIN mascota m ON
            p.rut_responsable = m.rut_responsable
        GROUP BY
            p.rut_responsable
);

CREATE OR REPLACE VIEW vista_historial AS (
    SELECT m.id_mascota, m.nombre, c.n_conflicto FROM mascota m
        JOIN (
            SELECT id_mascota, count(*) as "N_CONFLICTO" FROM historial_conflicto
                GROUP BY id_mascota
                ORDER BY n_conflicto DESC
        ) c ON
            m.id_mascota = c.id_mascota
) ORDER BY c.n_conflicto DESC
FETCH FIRST 10 ROWS ONLY;

CREATE OR REPLACE VIEW vista_coincidencia AS (
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
            x.id_mascota = y.id_mascota
);

CREATE OR REPLACE VIEW vista_recomendaciones AS (
    SELECT DISTINCT
        ms.id_mascota,
        ms.nombre
        FROM
            mascota ms
        JOIN adiestramiento ad ON
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
            ((av.fecha IN ad.fecha) AND (av.fecha IN di.fecha))
);

------------------





-- PROCEDIMIENTOS --

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

--------------------





-- DISPARADORES --

CREATE OR REPLACE TRIGGER trigger_alimentacion
    AFTER INSERT ON mascota
    FOR EACH ROW
BEGIN
    INSERT INTO dieta VALUES
    (
        (
        SELECT al.id_alimentacion FROM alimentacion al
        WHERE al.edad_recomendada < :new.edad
        ORDER BY al.edad_recomendada DESC
        FETCH FIRST 1 ROW ONLY
        ),
        :new.id_mascota, CURRENT_DATE, 500
    );
END;
/

--------------------