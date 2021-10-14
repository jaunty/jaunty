BEGIN;

CREATE TYPE whitelist_status AS ENUM (
    'pending',
    'approved',
    'rejected'
);

CREATE TABLE users (
    id bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    
    sf text NOT NULL UNIQUE,

    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE whitelist (
    id bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,

    sf text NOT NULL,
    uuid text NOT NULL UNIQUE,
    whitelist_status whitelist_status NOT NULL DEFAULT 'pending',

    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW(),

    UNIQUE(sf, uuid)
);

CREATE TABLE discord_tokens (
    id bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,

    sf text NOT NULL UNIQUE,
    access_token text NOT NULL,
    token_type text NOT NULL,
    expiry timestamp NOT NULL,
    refresh_token text NOT NULL,

    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW()
);

COMMIT;