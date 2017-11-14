ALTER TABLE users DROP deleted_at;

/* Change id to user_id as primary key */
ALTER TABLE anime_stats DROP deleted_at;
ALTER TABLE anime_stats DROP created_at;
ALTER TABLE anime_stats DROP CONSTRAINT anime_stats_pkey;
ALTER TABLE anime_stats ADD COLUMN user_id INTEGER;
UPDATE anime_stats a SET user_id=u.id FROM users u WHERE a.username=u.username;
/* Clean up duplicated rows */
UPDATE users SET fetched=false FROM (SELECT user_id, count(*) FROM anime_stats group by user_id having count(*)>1) sq WHERE users.id=sq.user_id;
DELETE FROM anime_stats WHERE user_id IN (SELECT user_id FROM (SELECT user_id, count(*) FROM anime_stats group by user_id having count(*)>1) AS t2);
/* Transform id into user_id */
ALTER TABLE anime_stats ALTER COLUMN user_id SET NOT NULL;
ALTER TABLE anime_stats ADD PRIMARY KEY(user_id);
ALTER TABLE anime_stats DROP id;
ALTER TABLE anime_stats DROP username;

/* Change id to user_id as primary key */
ALTER TABLE manga_stats DROP deleted_at;
ALTER TABLE manga_stats DROP created_at;
ALTER TABLE manga_stats DROP CONSTRAINT manga_stats_pkey;
ALTER TABLE manga_stats ADD COLUMN user_id INTEGER;
UPDATE manga_stats a SET user_id=u.id FROM users u WHERE a.username=u.username;
/* Clean up duplicated rows */
UPDATE users SET fetched=false FROM (SELECT user_id, count(*) FROM manga_stats group by user_id having count(*)>1) sq WHERE users.id=sq.user_id;
DELETE FROM manga_stats WHERE user_id IN (SELECT user_id FROM (SELECT user_id, count(*) FROM manga_stats group by user_id having count(*)>1) AS t2);
/* Transform id into user_id */
ALTER TABLE manga_stats ALTER COLUMN user_id SET NOT NULL;
ALTER TABLE manga_stats ADD PRIMARY KEY(user_id);
ALTER TABLE manga_stats DROP id;
ALTER TABLE manga_stats DROP username;
