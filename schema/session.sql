
CREATE TABLE session (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    account_id varchar(255) NOT NULL,
    provider varchar(255) NOT NULL,
    user_agent varchar(255) NOT NULL,
    ip varchar(255) NOT NULL,
    ttl integer NOT NULL,
    expires_at bigint NOT NULL,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT session_pkey PRIMARY KEY (id)
);