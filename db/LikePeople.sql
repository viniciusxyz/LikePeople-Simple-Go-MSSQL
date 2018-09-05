USE master;

IF EXISTS(SELECT * FROM SYS.DATABASES WHERE NAME = 'LIKEPEOPLE')
	DROP DATABASE LIKEPEOPLE 

CREATE DATABASE LIKEPEOPLE

USE LIKEPEOPLE

CREATE TABLE PESSOAS 
(
    ID_PESSOA       INT IDENTITY(1,1),
    NOME            VARCHAR(255),
    SOBRENOME       VARCHAR(255),
    APELIDO         VARCHAR(50),
    LIKES           INT,
    DESLIKES        INT,
    CONSTRAINT PRIMARY_PESSOAS PRIMARY KEY (ID_PESSOA)
)

CREATE TABLE USUARIOS
(
    ID_USUARIO       INT IDENTITY(1,1),
    LOGIN_USUARIO    VARCHAR(20),
    SENHA_USUARIO    VARCHAR(255),
    FK_PESSOAS       INT
    CONSTRAINT FK_USUARIOS_PESSOAS FOREIGN KEY (ID_USUARIO) REFERENCES PESSOAS(ID_PESSOA)
)

INSERT INTO PESSOAS (NOME,SOBRENOME,APELIDO,LIKES,DESLIKES) VALUES(
    'Vinicius',
    'Santos',
    'Vini',
    0,
    0
),
(
    'Israel',
    'Leite',
    'Rael',
    1,
    0
),
(
    'Denis',
    'Takano',
    'Japa da federal',
    1,
    0
)

select * from PESSOAS
