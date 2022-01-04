-- Tablas independientes

INSERT INTO persona_responsable VALUES ('12345789-k', 'El Pepe', 22, 'Calle Ete Sech', '42042069');
INSERT INTO persona_responsable VALUES ('11122334-6', 'Mox', 20, 'Calle La Calle', '1111111');
INSERT INTO persona_responsable VALUES ('10000000-5', 'Lolo', 25, 'No se', '1010101');
INSERT INTO persona_responsable VALUES ('11111111-1', 'Don Fede', 999, 'Al lado del metro san joaquin', '696969');
INSERT INTO persona_responsable VALUES ('88888888-1', 'John', 50, 'En alguna casa', '114201369');
INSERT INTO persona_responsable VALUES ('45834647-k', 'Doktor Miau', 35, 'En el planeta Tierra', '38364328');

INSERT INTO veterinario VALUES ('45834647-k', 'Doktor Miau', 40, 'Universidad de Gatitos', '1000100', 'Centro Vet.');
INSERT INTO veterinario VALUES ('23232322-1', 'Doktor Guagu', 33, 'Universidad de Perritos', '20001023', 'Movil');
INSERT INTO veterinario VALUES ('54835483-5', 'Doctor Animal', 50, 'Universidad de Veterinarios', '32915', 'Centro Vet.');

INSERT INTO alimentacion VALUES (1001, 'Dog Chau Cacho.', 4);
INSERT INTO alimentacion VALUES (1002, 'Dog Chau Adulto', 15);
INSERT INTO alimentacion VALUES (1003, 'Cat Chau Cacho.', 8);
INSERT INTO alimentacion VALUES (1004, 'Cat Chau Adulto', 18);
INSERT INTO alimentacion VALUES (1005, 'Comida Casera', 24);
INSERT INTO alimentacion VALUES (1006, 'UltraDiet Cat', 48);
INSERT INTO alimentacion VALUES (1007, 'Nadie lo usa', 40);

INSERT INTO vacuna VALUES (691, 'Primera');
INSERT INTO vacuna VALUES (692, 'Segunda');
INSERT INTO vacuna VALUES (693, 'Tercera');

INSERT INTO entrenamiento VALUES (501, 'Quieto');
INSERT INTO entrenamiento VALUES (502, 'Dar la pata');
INSERT INTO entrenamiento VALUES (503, 'No morder');
INSERT INTO entrenamiento VALUES (504, 'Responder');
INSERT INTO entrenamiento VALUES (505, 'Saltar cuerda');
INSERT INTO entrenamiento VALUES (506, 'Frisbi');
INSERT INTO entrenamiento VALUES (507, 'Flojear');
INSERT INTO entrenamiento VALUES (508, 'Rodar');
INSERT INTO entrenamiento VALUES (509, 'Sentarse');






-- Tablas dependientes

INSERT INTO mascota VALUES (111, 'Pepito', 10, 'perro', 'Pomerania', 'Blanco', 'Ninguno', 'Compañía', 'Rescate', '12345789-k');
INSERT INTO mascota VALUES (222, 'Coti', 84, 'perro', 'Dachshund', 'Café', 'Ninguno', 'Compañía', 'Adopcion', '11122334-6');
INSERT INTO mascota VALUES (333, 'Botas', 95, 'gato', 'Tuxedo', 'Negro', 'Patas blancas', 'Compañía', 'Adopcion', '11122334-6');
INSERT INTO mascota VALUES (444, 'Garfil', 60, 'gato', 'Tabby', 'Naranjo', 'Atrigrado', 'Apoyo', 'Compra', '88888888-1');
INSERT INTO mascota VALUES (555, 'Feddi', 24, 'perro', 'Husky siberiano', 'Gris', 'Pecho blanco', 'Apoyo', 'Compra', '11111111-1');
INSERT INTO mascota VALUES (666, 'Lala', 7, 'gato', 'Angora', 'Blanco', 'Ninguno', 'Compañía', 'Compra', '10000000-5');
INSERT INTO mascota VALUES (777, 'Oddi', 30, 'perro', 'No se', 'Amarillo', 'Ninguno', 'Apoyo', 'Adopcion', '88888888-1');
INSERT INTO mascota VALUES (888, 'Boni', 150, 'perro', 'Yorkshire', 'Blanco', 'Ninguno', 'Compañía', 'Adopcion', '11122334-6');
INSERT INTO mascota VALUES (999, 'Miau', 34, 'gato', 'Minx', 'Blanco', 'Manchas cafes y negras', 'Compañía', 'Rescate', '45834647-k');
INSERT INTO mascota VALUES (1000, 'Pepito', 56, 'perro', 'Poodle', 'Blanco', 'Ninguno', 'Compañía', 'Compra', '11111111-1');

INSERT INTO atencion_veterinaria VALUES ('45834647-k', 333, '09-10-2021', 'Control de salud. Correcto.');
INSERT INTO atencion_veterinaria VALUES ('45834647-k', 444, '05-10-2021', 'Obesidad. Cambio de dieta.');
INSERT INTO atencion_veterinaria VALUES ('23232322-1', 111, '25-05-2021', 'Vacunacion, primera.');
INSERT INTO atencion_veterinaria VALUES ('23232322-1', 222, '29-09-2021', 'Cirugia.');
INSERT INTO atencion_veterinaria VALUES ('23232322-1', 555, '10-03-2020', 'Control de salud. Correcto.');
INSERT INTO atencion_veterinaria VALUES ('23232322-1', 111, '25-06-2021', 'Vacunacion, segunda.');
INSERT INTO atencion_veterinaria VALUES ('45834647-k', 666, '02-04-2021', 'Vacunacion, primera.');
INSERT INTO atencion_veterinaria VALUES ('23232322-1', 777, '21-03-2020', 'Control de salud. Correcto.');
INSERT INTO atencion_veterinaria VALUES ('23232322-1', 888, '20-04-2018', 'Control de salud. Correcto.');
INSERT INTO atencion_veterinaria VALUES ('45834647-k', 999, '04-02-2017', 'Vacunacion, primera.');
INSERT INTO atencion_veterinaria VALUES ('45834647-k', 999, '04-03-2017', 'Vacunacion, segunda.');
INSERT INTO atencion_veterinaria VALUES ('45834647-k', 999, '04-03-2018', 'Vacunacion, tercera.');
INSERT INTO atencion_veterinaria VALUES ('45834647-k', 999, '06-09-2020', 'Control de salud. Correcto.');
INSERT INTO atencion_veterinaria VALUES ('54835483-5', 444, '05-02-2018', 'Cambio alimentacion.');
INSERT INTO atencion_veterinaria VALUES ('54835483-5', 222, '04-02-2019', 'Cambio alimentacion.');
INSERT INTO atencion_veterinaria VALUES ('54835483-5', 555, '04-05-2018', 'Control de salud. Correcto.');
INSERT INTO atencion_veterinaria VALUES ('54835483-5', 1000, '09-09-2019', 'Control de salud. Correcto.');

INSERT INTO adiestramiento VALUES (503, 222, '15-03-2017');
INSERT INTO adiestramiento VALUES (504, 222, '28-08-2017');
INSERT INTO adiestramiento VALUES (507, 333, '01-01-2019');
INSERT INTO adiestramiento VALUES (506, 555, '04-10-2018');
INSERT INTO adiestramiento VALUES (505, 555, '16-09-2020');
INSERT INTO adiestramiento VALUES (503, 555, '23-03-2019');
INSERT INTO adiestramiento VALUES (504, 555, '04-05-2018');

INSERT INTO vacunacion VALUES (111, 691, '25-05-2021');
INSERT INTO vacunacion VALUES (111, 692, '25-06-2021');
INSERT INTO vacunacion VALUES (666, 691, '02-04-2021');
INSERT INTO vacunacion VALUES (999, 691, '04-02-2017');
INSERT INTO vacunacion VALUES (999, 692, '04-03-2017');
INSERT INTO vacunacion VALUES (999, 693, '04-03-2018');

INSERT INTO dieta VALUES (1004, 333, '23-05-2019', 1000);
INSERT INTO dieta VALUES (1006, 444, '05-10-2021', 500);
INSERT INTO dieta VALUES (1002, 222, '04-02-2019', 300);
INSERT INTO dieta VALUES (1005, 222, '04-02-2019', 500);
INSERT INTO dieta VALUES (1001, 111, '02-09-2021', 300);
INSERT INTO dieta VALUES (1002, 555, '04-05-2018', 800);
INSERT INTO dieta VALUES (1003, 666, '02-08-2021', 200);
INSERT INTO dieta VALUES (1002, 777, '12-07-2020', 800);
INSERT INTO dieta VALUES (1002, 888, '04-08-2015', 600);
INSERT INTO dieta VALUES (1004, 999, '25-07-2019', 600);
INSERT INTO dieta VALUES (1002, 1000, '09-09-2019', 500);

INSERT INTO historial_conflicto VALUES (9991, 444, 'Garfil araño a John porque no le dio lasaña', '13-10-2021');
INSERT INTO historial_conflicto VALUES (9992, 333, 'Botas araño a Garfil porque se comio su comida', '29-04-2020');
INSERT INTO historial_conflicto VALUES (9993, 444, 'Garfil ataco a Oddie porque no lo dejaba comer', '14-11-2019');
INSERT INTO historial_conflicto VALUES (9994, 777, 'Oddie ataco a Garfil porque no deja de molestarlo', '24-12-2019');
