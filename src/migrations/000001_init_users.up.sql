CREATE TYPE user_role AS ENUM ('member', 'admin');
CREATE TYPE player_position AS ENUM ('goalkeeper', 'field');

CREATE TABLE users (
    id                 BIGSERIAL PRIMARY KEY,
    name               VARCHAR(100)  NOT NULL,
    email              VARCHAR(255)  NOT NULL,
    password_hash      VARCHAR(255)  NOT NULL,
    role               user_role     NOT NULL DEFAULT 'member',
    bio                TEXT,
    avatar_url         VARCHAR(500),
    preferred_position player_position,
    is_public          BOOLEAN       NOT NULL DEFAULT TRUE,
    created_at         TIMESTAMPTZ   NOT NULL DEFAULT now(),
    updated_at         TIMESTAMPTZ   NOT NULL DEFAULT now()
);

-- E-mail is the login identifier: unique regardless of case.
CREATE UNIQUE INDEX users_email_lower_uq ON users (lower(email));
