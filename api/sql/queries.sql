-- Seguidor, Seguindo
SELECT seguidor.id IdSeguidor,
       seguidor.Nome Seguidor,
       seguindo.id IdSeguindo,
       seguindo.Nome Seguindo,
       s.bloqueado Bloqueado
FROM seguidores s
INNER JOIN usuarios seguindo
ON seguindo.id = s.usuario_id
INNER JOIN usuarios seguidor
ON seguidor.id = s.seguidor_id
ORDER BY seguidor.Nome, seguindo.Nome;

-- Curtidas / Descurtidas
SELECT 'Curtidas' AS Engajamento, p.titulo AS Publicacao, a.nome AS Autor,
       u.nome AS Usuario, c.criadaEm AS Data
FROM curtidas c
INNER JOIN publicacoes p
ON p.id = c.publicacao_id
INNER JOIN usuarios u
ON u.id = c.usuario_id
INNER JOIN usuarios a
ON a.id = p.autor_id
UNION
SELECT 'Descurtidas' AS Engajamento, p.titulo AS Publicacao, a.nome AS Autor,
       u.nome AS Usuario, d.criadaEm AS Data
FROM descurtidas d
INNER JOIN publicacoes p
ON p.id = d.publicacao_id
INNER JOIN usuarios u
ON u.id = d.usuario_id
INNER JOIN usuarios a
ON a.id = p.autor_id;
