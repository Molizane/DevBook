INSERT IGNORE INTO
  usuarios (nome, nick, email, senha)
VALUES
  (
    "Benedito Angelo Molizane",
    "Angelo.Molizane",
    "molizane@gmail.com",
    "$2a$10$FcaN997de/rpd9yAciHpqO4LoWdTdUyRQ4ZYZULJL6Iv/zhq/aO7q" -- 123456
  ),
  (
    "Maria de Fátima Moreira Molizane",
    "Cocota",
    "fatimamolizane@gmail.com",
    "$2a$10$FcaN997de/rpd9yAciHpqO4LoWdTdUyRQ4ZYZULJL6Iv/zhq/aO7q"
  ),
  (
    "Amanda Moreira Molizane",
    "Nandinha",
    "amanda.molizane@gmail.com",
    "$2a$10$FcaN997de/rpd9yAciHpqO4LoWdTdUyRQ4ZYZULJL6Iv/zhq/aO7q"
  ),
  (
    "Bianca Moreira Molizane",
    "Bibi",
    "bianca.molizane@gmail.com",
    "$2a$10$FcaN997de/rpd9yAciHpqO4LoWdTdUyRQ4ZYZULJL6Iv/zhq/aO7q"
  );

INSERT IGNORE INTO
  publicacoes (titulo, conteudo, autor_id)
VALUES
  (
    "Primeira publicação de Angelo",
    "<b>Texto primeira publicação de Angelo.Molizane</b>",
    1
  ),
  (
    "Primeira publicação de Fátima",
    "Texto primeira publicação de Cocota",
    2
  ),
  (
    "Primeira publicação de Amanda",
    "Texto primeira publicação de Nandinha",
    3
  ),
  (
    "Primeira publicação de Bianca",
    "Texto primeira publicação de Bibi",
    4
  );

INSERT IGNORE INTO
  seguidores (seguidor_id, usuario_id)
VALUES
  (1, 2),
  (1, 3),
  (1, 4),
  (2, 1),
  (2, 3),
  (2, 4);
  