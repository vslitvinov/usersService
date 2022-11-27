CREATE TABLE accounts (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    firstname varchar(255) NOT NULL,
    lastname varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    phone varchar(255) NULL,
    username varchar(16) NOT NULL,
    password varchar(255) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_archive bool NOT NULL DEFAULT false,
    is_verified bool NOT NULL DEFAULT false,
    CONSTRAINT accounts_email_key UNIQUE (email),
    CONSTRAINT accounts_phone_key UNIQUE (phone),
    CONSTRAINT accounts_pkey PRIMARY KEY (id),
    CONSTRAINT accounts_username_key UNIQUE (username)
);