CREATE DATABASE sansappdb;
USE sansappdb;

CREATE TABLE cuenta (
    fecha_registro              DATE            NOT NULL DEFAULT current_timestamp(),
    id_cuenta                   INT(11)         NOT NULL AUTO_INCREMENT,

    rol_usm                     INT(10)         NOT NULL,
    correo                      VARCHAR(255)    NOT NULL,
    contrasenna                 VARCHAR(255)    NOT NULL,
    nombre                      VARCHAR(50)     NOT NULL,
    fecha_nacimiento            DATE            NOT NULL,

    telefono                    INT(9)          NOT NULL DEFAULT -1,
    num_compras                 INT(11)         NOT NULL DEFAULT 0,
    num_ventas                  INT(11)         NOT NULL DEFAULT 0,

    PRIMARY KEY (id_cuenta),
    UNIQUE (rol_usm),
    UNIQUE (correo)
);

CREATE TABLE etiqueta (
    nombre                      VARCHAR(20)     NOT NULL,

    PRIMARY KEY (nombre)
);

CREATE TABLE producto (
    id_producto                 INT(11)         NOT NULL AUTO_INCREMENT,
    vendedor                    INT(11)         NOT NULL,

    nombre                      VARCHAR(20)     NOT NULL,
    descripcion                 TEXT            NOT NULL,
    etiqueta                    VARCHAR(20)     NOT NULL,
    fecha_publicacion           DATE            NOT NULL DEFAULT current_timestamp(),

    precio                      INT(7)          NOT NULL,
    unidades                    INT(11)         NOT NULL,

    prom_calificacion           DECIMAL(2,1)    NOT NULL DEFAULT 0.0,
    ventas                      INT(11)         NOT NULL DEFAULT 0,

    PRIMARY KEY (id_producto),
    FOREIGN KEY (vendedor) REFERENCES cuenta (id_cuenta) ON DELETE CASCADE
);

CREATE TABLE comentario (
    id_comentario               INT(11)         NOT NULL AUTO_INCREMENT,
    autor                       INT(11)         NOT NULL,
    producto                    INT(11)         NOT NULL,

    titulo                      VARCHAR(20)     NOT NULL,
    descripcion                 TEXT            NOT NULL,
    fecha_comentario            DATE            NOT NULL DEFAULT current_timestamp(),

    PRIMARY KEY (id_comentario),
    FOREIGN KEY (autor) REFERENCES cuenta (id_cuenta) ON DELETE CASCADE,
    FOREIGN KEY (producto) REFERENCES producto (id_producto) ON DELETE CASCADE
);

CREATE TABLE calificacion (
    autor                       INT(11)         NOT NULL,
    producto                    INT(11)         NOT NULL,

    estrellas                   INT(1)          NOT NULL,

    PRIMARY KEY (autor, producto),
    FOREIGN KEY (autor) REFERENCES cuenta (id_cuenta) ON DELETE CASCADE,
    FOREIGN KEY (producto) REFERENCES producto (id_producto) ON DELETE CASCADE
);

CREATE TABLE productos_carrito (
    comprador                   INT(11)         NOT NULL,
    producto                    INT(11)         NOT NULL,

    cantidad                    INT(11)         NOT NULL,
    precio_total                INT(11)         NOT NULL DEFAULT 0,

    PRIMARY KEY (comprador, producto),
    FOREIGN KEY (comprador) REFERENCES cuenta (id_cuenta) ON DELETE CASCADE,
    FOREIGN KEY (producto) REFERENCES producto (id_producto) ON DELETE CASCADE
);

CREATE TABLE etiquetado (
    producto                    INT(11)         NOT NULL,
    etiqueta                    VARCHAR(20)     NOT NULL,

    PRIMARY KEY (producto, etiqueta),
    FOREIGN KEY (producto) REFERENCES producto (id_producto),
    FOREIGN KEY (etiqueta) REFERENCES etiqueta (nombre)
);

CREATE TABLE transaccion (
    id_transacc                 INT(11)         NOT NULL AUTO_INCREMENT,
    fecha_transacc              DATE            NOT NULL DEFAULT current_timestamp(),

    vendedor                    INT(11)         NOT NULL,
    producto                    INT(11)         NOT NULL,

    comprador                   INT(11)         NOT NULL,
    cantidad                    INT(11)         NOT NULL,
    precio_total                INT(11)         NOT NULL,

    PRIMARY KEY (id_transacc),
    FOREIGN KEY (vendedor) REFERENCES cuenta (id_cuenta) ON DELETE CASCADE,
    FOREIGN KEY (comprador) REFERENCES cuenta (id_cuenta) ON DELETE CASCADE,
    FOREIGN KEY (producto) REFERENCES producto (id_producto) ON DELETE CASCADE
);





CREATE OR REPLACE VIEW top5_ventas AS (
    SELECT id_cuenta, nombre, num_ventas
    FROM cuenta
    ORDER BY num_ventas DESC
    LIMIT 5
);

CREATE OR REPLACE VIEW top5_prod_calif AS (
    SELECT id_producto, nombre, prom_calificacion
    FROM producto
    ORDER BY prom_calificacion DESC
    LIMIT 5
);

CREATE OR REPLACE VIEW top5_prod_venta AS (
    SELECT id_producto, nombre, ventas
    FROM producto
    ORDER BY ventas DESC
    LIMIT 5
);





DELIMITER |
CREATE OR REPLACE TRIGGER calc_promedio_insert
    AFTER INSERT ON calificacion
    FOR EACH ROW
BEGIN
    SET @prom = (SELECT AVG(estrellas) FROM calificacion WHERE producto=NEW.producto);
    UPDATE producto SET prom_calificacion=@prom WHERE id_producto=NEW.producto;
END;
|
DELIMITER ;

DELIMITER |
CREATE OR REPLACE TRIGGER calc_promedio_update
    AFTER UPDATE ON calificacion
    FOR EACH ROW
BEGIN
    SET @prom = (SELECT AVG(estrellas) FROM calificacion WHERE producto=NEW.producto);
    UPDATE producto SET prom_calificacion=@prom WHERE id_producto=NEW.producto;
END;
|
DELIMITER ;

DELIMITER |
CREATE OR REPLACE TRIGGER calc_promedio_delete
    AFTER DELETE ON calificacion
    FOR EACH ROW
BEGIN
    SET @prom = (SELECT AVG(estrellas) FROM calificacion WHERE producto=OLD.producto);
    UPDATE producto SET prom_calificacion=@prom WHERE id_producto=OLD.producto;
END;
|
DELIMITER ;

DELIMITER |
CREATE OR REPLACE TRIGGER añadir_etiqueta_insert
    AFTER INSERT ON producto
    FOR EACH ROW
BEGIN
    INSERT IGNORE INTO etiqueta SET nombre=NEW.etiqueta;
END;
|
DELIMITER ;

DELIMITER |
CREATE OR REPLACE TRIGGER añadir_etiqueta_update
    AFTER UPDATE ON producto
    FOR EACH ROW
BEGIN
    INSERT IGNORE INTO etiqueta SET nombre=NEW.etiqueta;
END;
|
DELIMITER ;

DELIMITER |
CREATE OR REPLACE TRIGGER calcular_precio_insert
    BEFORE INSERT ON productos_carrito
    FOR EACH ROW
BEGIN
    SET @precio = (SELECT precio FROM producto WHERE id_producto=NEW.producto);
    SET NEW.precio_total = @precio * NEW.cantidad;
END;
|
DELIMITER ;

DELIMITER |
CREATE OR REPLACE TRIGGER registrar_ventas
    AFTER INSERT ON transaccion
    FOR EACH ROW
BEGIN
    UPDATE producto SET unidades=unidades-NEW.cantidad, ventas=ventas+NEW.cantidad WHERE id_producto=NEW.producto;
    UPDATE cuenta SET num_compras=num_compras+1 WHERE id_cuenta=NEW.comprador;
    UPDATE cuenta SET num_ventas=num_ventas+1 WHERE id_cuenta=NEW.vendedor;
END;
|
DELIMITER ;
