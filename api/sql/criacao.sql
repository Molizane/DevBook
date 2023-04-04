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
        id int auto_increment,
        nome varchar(50) not null,
        nick varchar(50) not null unique,
        email varchar(50) not null unique,
        senha varchar(100) not null,
        criadoEm timestamp DEFAULT current_timestamp() not null,
        CONSTRAINT pk_usuarios PRIMARY KEY (id)
) ENGINE = INNODB;

CREATE TABLE seguidores
(
        usuario_id int not null,
        seguidor_id int not null,
        desde timestamp DEFAULT current_timestamp() not null,
        bloqueado tinyint DEFAULT 0 not null,
        CONSTRAINT fk_seguidores_seguido FOREIGN KEY (usuario_id) REFERENCES usuarios (id) ON DELETE CASCADE,
        CONSTRAINT fk_seguidores_seguidor FOREIGN KEY (seguidor_id) REFERENCES usuarios (id) ON DELETE CASCADE,
        CONSTRAINT pk_seguidores PRIMARY KEY (usuario_id, seguidor_id)
) ENGINE = INNODB;

CREATE UNIQUE INDEX idx_seguidores_usuario_seguidor
ON seguidores (seguidor_id, usuario_id);

CREATE TABLE publicacoes
(
        id int auto_increment not null,
        titulo varchar(100) not null unique,
        conteudo varchar(300) not null,
        autor_id int not null,
        curtidas int default 0 not null,
        criadaEm timestamp DEFAULT current_timestamp() not null,
        CONSTRAINT fk_publicacoes_autor FOREIGN KEY (autor_id) REFERENCES usuarios (id) ON DELETE CASCADE,
        CONSTRAINT pk_publicacoes PRIMARY KEY (id)
) ENGINE = INNODB;

CREATE TABLE curtidas
(
        publicacao_id int not null,
        usuario_id int not null,
        criadaEm timestamp DEFAULT current_timestamp() not null,
        CONSTRAINT fk_curtidas_publicacoes FOREIGN KEY (publicacao_id) REFERENCES publicacoes (id) ON DELETE CASCADE,
        CONSTRAINT fk_curtidas_usuarios FOREIGN KEY (usuario_id) REFERENCES usuarios (id) ON DELETE CASCADE,
        CONSTRAINT pk_curtidas PRIMARY KEY (publicacao_id, usuario_id)
) ENGINE = INNODB;

CREATE UNIQUE INDEX idx_curtidas_usuario_publicacao
ON curtidas (usuario_id, publicacao_id);

CREATE TABLE descurtidas
(
        publicacao_id int not null,
        usuario_id int not null,
        criadaEm timestamp DEFAULT current_timestamp() not null,
        CONSTRAINT fk_descurtidas_publicacoes FOREIGN KEY (publicacao_id) REFERENCES publicacoes (id) ON DELETE CASCADE,
        CONSTRAINT fk_descurtidas_usuarios FOREIGN KEY (usuario_id) REFERENCES usuarios (id) ON DELETE CASCADE,
        CONSTRAINT pk_descurtidas PRIMARY KEY (publicacao_id, usuario_id)
) ENGINE = INNODB;

CREATE UNIQUE INDEX idx_descurtidas_usuario_publicacao
ON descurtidas (usuario_id, publicacao_id);
