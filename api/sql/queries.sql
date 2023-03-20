SELECT 'Curtidas' AS Engajamento, c.*
FROM curtidas c
UNION
SELECT 'Descurtidas' AS Engajamento, d.*
FROM descurtidas d;

-- Seguidores
SELECT DISTINCT u.id, u.nome, u.nick, u.email, u.criadoEm
FROM usuarios u
INNER JOIN seguidores s
ON s.seguidor_id = u.id
WHERE s.usuario_id = 1;

-- Seguindo
SELECT DISTINCT u.id, u.nome, u.nick, u.email, u.criadoEm
FROM usuarios u
INNER JOIN seguidores s
ON s.usuario_id = u.id
WHERE s.seguidor_id = 1;

-- Seguido, seguindo
SELECT seguidor.id IdSeguidor, seguidor.Nome Seguidor, seguido.id IdSeguido, seguido.Nome Seguido
FROM seguidores s
INNER JOIN usuarios seguido
ON seguido.id = s.usuario_id
INNER JOIN usuarios seguidor
ON seguidor.id = s.seguidor_id
ORDER BY seguidor.Nome, seguido.Nome;
