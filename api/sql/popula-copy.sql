INSERT INTO usuarios (nome, nick, email, senha)
VALUES (
           "Benedito Angelo Molizane",
           "Molizane",
           "molizane@gmail.com",
           "$2a$10$FSiEYryjCHcCmHhbUyTheuQWuP9iqOwbWHnbR6IM0H3fWsOVFcLY2" -- 123456
       ),
       (
           "Maria de Fátima Moreira Molizane",
           "Cocota",
           "fatimamolizane@gmail.com",
           "$2a$10$FSiEYryjCHcCmHhbUyTheuQWuP9iqOwbWHnbR6IM0H3fWsOVFcLY2"
       ),
       (
           "Amanda Moreira Molizane",
           "Nandinha",
           "amanda.molizane@gmail.com",
           "$2a$10$FSiEYryjCHcCmHhbUyTheuQWuP9iqOwbWHnbR6IM0H3fWsOVFcLY2"
       ),
       (
           "Bianca Moreira Molizane",
           "Bibi",
           "bianca.molizane@gmail.com",
           "$2a$10$FSiEYryjCHcCmHhbUyTheuQWuP9iqOwbWHnbR6IM0H3fWsOVFcLY2"
       );

INSERT IGNORE INTO publicacoes (titulo, conteudo, autor_id)
VALUES ("Primeira publicação de Angelo", "Texto primeira publicação de Angelo", 1),
       ("Primeira publicação de Fátima", "Texto primeira publicação de Fátima", 2),
       ("Primeira publicação de Amanda", "Texto primeira publicação de Amanda", 3),
       ("Primeira publicação de Bianca", "Texto primeira publicação de Bianca", 4);

INSERT IGNORE INTO seguidores (usuario_id, seguidor_id)
VALUES (2, 1),
       (3, 1),
       (4, 1),
       (3, 2),
       (4, 2);
