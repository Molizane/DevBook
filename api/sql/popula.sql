INSERT INTO usuarios (nome, nick, email, senha)
VALUES (
         "João da Silva",
         "Joao.Silva",
         "joao.silva@gmail.com",
           "$2a$10$FSiEYryjCHcCmHhbUyTheuQWuP9iqOwbWHnbR6IM0H3fWsOVFcLY2" -- 123456
       ),
       (
         "Maria de Souza",
         "Maria.Souza",
         "maria.souza@gmail.com",
           "$2a$10$FSiEYryjCHcCmHhbUyTheuQWuP9iqOwbWHnbR6IM0H3fWsOVFcLY2"
       ),
       (
         "Manoela Cruz",
         "Manoela.Cruz",
         "manoela.cruz@gmail.com",
           "$2a$10$FSiEYryjCHcCmHhbUyTheuQWuP9iqOwbWHnbR6IM0H3fWsOVFcLY2"
       ),
       (
         "Rute Santos",
         "Rute.Santos",
         "rute.santos@gmail.com",
           "$2a$10$FSiEYryjCHcCmHhbUyTheuQWuP9iqOwbWHnbR6IM0H3fWsOVFcLY2"
       );

INSERT IGNORE INTO publicacoes (titulo, conteudo, autor_id)
VALUES ("Primeira publicação de João da Silva", "<b>Texto primeira publicação de João da Silva</b>", 1),
       ("Primeira publicação de Maria de Souza", "Texto primeira publicação de Maria de Souza", 2),
       ("Primeira publicação de Manoela Cruz", "Texto primeira publicação de Manoela Cruz", 3),
       ("Primeira publicação de Rute Santos", "Texto primeira publicação de Rute Santos", 4);

INSERT IGNORE INTO seguidores (usuario_id, seguidor_id)
VALUES (2, 1),
       (3, 1),
       (4, 1),
       (3, 2),
       (4, 2);
