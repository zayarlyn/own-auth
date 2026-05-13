CREATE TABLE user_pool (
    -- Identity
    id      UUID  PRIMARY KEY DEFAULT gen_random_uuid(),
    name    TEXT  NOT NULL,
 
    -- Token lifetimes
    token_expiry_seconds    INT  NOT NULL DEFAULT 900,   -- 15 min
    refresh_expiry_days     INT  NOT NULL DEFAULT 30,
 
    -- Password rules
    -- { min_length, require_uppercase, require_lowercase,
    --   require_numbers, require_symbols }
    password_policy  JSONB NOT NULL DEFAULT '{
        "min_length": 8,
        "require_uppercase": false,
        "require_lowercase": false,
        "require_numbers":   false,
        "require_symbols":   false
    }',
 
    allowed_flows       TEXT[] NOT NULL DEFAULT ARRAY['email:password'],
    callback_urls       TEXT[] NOT NULL DEFAULT ARRAY[]::TEXT[],
 
    -- Timestamps
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);
 

CREATE TABLE user (
    id       UUID  PRIMARY KEY DEFAULT gen_random_uuid(),
    pool_id  UUID  NOT NULL REFERENCES user_pools(id) ON DELETE CASCADE,
 
    email          TEXT        NOT NULL,
    username       TEXT        NOT NULL,
    password_hash  TEXT        NOT NULL,
 
    status      TEXT NOT NULL DEFAULT 'UNCONFIRMED',
    attributes  JSONB NOT NULL DEFAULT '{}',

    last_login_at     TIMESTAMPTZ DEFAULT NULL,
    email_verified_at TIMESTAMPTZ DEFAULT NULL,
 
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at    TIMESTAMPTZ

);
