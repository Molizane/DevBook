CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

CREATE USER 'golang' @'localhost' IDENTIFIED BY 'golang';

GRANT ALL PRIVILEGES ON devbook.* TO 'golang' @'localhost';

DROP TABLE IF EXISTS descurtidas;
DROP TABLE IF EXISTS curtidas;
DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios
(
        id int auto_increment primary key,
        nome varchar(50) not null,
        nick varchar(50) not null unique,
        email varchar(50) not null unique,
        senha varchar(100) not null,
        criadoEm timestamp DEFAULT current_timestamp() not null
) ENGINE = INNODB;

CREATE TABLE seguidores
(
        seguidor_id int not null,
        usuario_id int not null,
        desde timestamp DEFAULT current_timestamp() not null,
        bloqueado bit(1) DEFAULT 0 not null,
        FOREIGN KEY (usuario_id) REFERENCES usuarios (id) ON DELETE CASCADE,
        FOREIGN KEY (seguidor_id) REFERENCES usuarios (id) ON DELETE CASCADE,
        PRIMARY KEY (seguidor_id, usuario_id)
) ENGINE = INNODB;

CREATE UNIQUE INDEX idx_seguidores_usuario_seguidor
ON seguidores (usuario_id, seguidor_id);

CREATE TABLE publicacoes
(
        id int auto_increment not null primary key,
        titulo varchar(100) not null unique,
        conteudo varchar(300) not null,
        autor_id int not null,
        curtidas int default 0 not null,
        criadaEm timestamp DEFAULT current_timestamp() not null,
        FOREIGN KEY (autor_id) REFERENCES usuarios (id) ON DELETE CASCADE
) ENGINE = INNODB;

CREATE TABLE curtidas
(
        publicacao_id int not null,
        usuario_id int not null,
        criadaEm timestamp DEFAULT current_timestamp() not null,
        FOREIGN KEY (publicacao_id) REFERENCES publicacoes (id) ON DELETE CASCADE,
        FOREIGN KEY (usuario_id) REFERENCES usuarios (id) ON DELETE CASCADE,
        PRIMARY KEY (publicacao_id, usuario_id)
) ENGINE = INNODB;

CREATE UNIQUE INDEX idx_curtidas_usuario_publicacao
ON curtidas (usuario_id, publicacao_id);

CREATE TABLE descurtidas
(
        publicacao_id int not null,
        usuario_id int not null,
        criadaEm timestamp DEFAULT current_timestamp() not null,
        FOREIGN KEY (publicacao_id) REFERENCES publicacoes (id) ON DELETE CASCADE,
        FOREIGN KEY (usuario_id) REFERENCES usuarios (id) ON DELETE CASCADE,
        PRIMARY KEY (publicacao_id, usuario_id)
) ENGINE = INNODB;

CREATE UNIQUE INDEX idx_descurtidas_usuario_publicacao
ON descurtidas (usuario_id, publicacao_id);
