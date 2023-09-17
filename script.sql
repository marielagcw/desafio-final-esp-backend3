CREATE DATABASE `db_desafio_final`;

USE `db_desafio_final`;

CREATE TABLE `odontologo` (
  `id` int NOT NULL AUTO_INCREMENT,
  `apellido` varchar(45) DEFAULT NULL,
  `nombre` varchar(45) DEFAULT NULL,
  `matricula` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
SELECT * FROM db_desafio_final.odontologo;

CREATE TABLE `paciente` (
  `id` int NOT NULL AUTO_INCREMENT,
  `apellido` varchar(45) DEFAULT NULL,
  `nombre` varchar(45) DEFAULT NULL,
  `dni` int DEFAULT NULL,
  `domicilio` varchar(45) DEFAULT NULL,
  `fecha_alta` date DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;


CREATE TABLE `turno` (
   `id` int NOT NULL AUTO_INCREMENT,
  `fecha` date DEFAULT NULL,
  `hora` time DEFAULT NULL,
  `descripcion` varchar(45) DEFAULT NULL,
  `odontologo_id` int NOT NULL,
  `paciente_id` int NOT NULL,
  PRIMARY KEY (`id`,`odontologo_id`,`paciente_id`),
  KEY `fk_turno_odontologo_idx` (`odontologo_id`),
  KEY `fk_turno_paciente1_idx` (`paciente_id`),
  CONSTRAINT `fk_turno_odontologo` FOREIGN KEY (`odontologo_id`) REFERENCES `odontologo` (`id`),
  CONSTRAINT `fk_turno_paciente1` FOREIGN KEY (`paciente_id`) REFERENCES `paciente` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

INSERT INTO `db_desafio_final`.`odontologo`
(`id`,
`apellido`,
`nombre`,
`matricula`)
VALUES
(1,
"Simpson",
"Homero",
"ABC123");

INSERT INTO `db_desafio_final`.`odontologo`
(`id`,
`apellido`,
`nombre`,
`matricula`)
VALUES
(2,
"Simpson",
"Lisa",
"ABC456");

INSERT INTO `db_desafio_final`.`paciente`
(`id`,
`apellido`,
`nombre`,
`dni`,
`domicilio`,
`fecha_alta`)
VALUES
(1,
"Simpson",
"Bart",
123456789,
"Springfield",
"2023-09-15");

INSERT INTO `db_desafio_final`.`paciente`
(`id`,
`apellido`,
`nombre`,
`dni`,
`domicilio`,
`fecha_alta`)
VALUES
(2,
"Simpson",
"Maggie",
987654321,
"Springfield",
"2023-09-15");

INSERT INTO `db_desafio_final`.`turno`
(`id`,
`fecha`,
`hora`,
`descripcion`,
`odontologo_id`,
`paciente_id`)
VALUES
(1,
"2023-09-15",
"10:00:00",
"Dolor",
1,
1);

INSERT INTO `db_desafio_final`.`turno`
(`id`,
`fecha`,
`hora`,
`descripcion`,
`odontologo_id`,
`paciente_id`)
VALUES
(2,
"2023-09-15",
"10:00:00",
"Dolor",
1,
2);

INSERT INTO `db_desafio_final`.`turno`
(`id`,
`fecha`,
`hora`,
`descripcion`,
`odontologo_id`,
`paciente_id`)
VALUES
(3,
"2023-09-16",
"10:00:00",
"Dolor",
2,
1);