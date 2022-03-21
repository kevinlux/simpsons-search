package main

const SqlSearch = `
SELECT
title, ts_headline(content, q)
FROM (
SELECT
title, content, ts_rank(tsv, q) as rank, q
FROM
simpsons_episodes, plainto_tsquery($1) q
WHERE
tsv @@ q
ORDER BY
rank DESC
LIMIT
10
) AS subq
ORDER BY
rank DESC;
`
const sqlGetEpisode = `
SELECT title, content
FROM simpsons_episodes
WHERE title=$1
`
