BEGIN;

DROP TYPE IF EXISTS enum;

DROP TABLE IF EXISTS users,
    whitelist,
    discord_tokens,
    bans;

COMMIT;