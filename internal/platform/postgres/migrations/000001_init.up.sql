-- CorpTodo schema: Telegram users, tasks, and persisted FSM sessions (applied at startup).

CREATE TABLE IF NOT EXISTS app_users (
    id         BIGSERIAL PRIMARY KEY,
    telegram_id BIGINT NOT NULL UNIQUE,
    username    TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS todos (
    id          BIGSERIAL PRIMARY KEY,
    user_id     BIGINT NOT NULL REFERENCES app_users (id) ON DELETE CASCADE,
    title       TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    status      TEXT NOT NULL DEFAULT 'open'
        CHECK (status IN ('open', 'done', 'archived')),
    priority    TEXT NOT NULL DEFAULT 'normal'
        CHECK (priority IN ('low', 'normal', 'high')),
    due_at      TIMESTAMPTZ,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS todos_user_updated
    ON todos (user_id, updated_at DESC);

CREATE INDEX IF NOT EXISTS todos_user_status
    ON todos (user_id, status);

CREATE TABLE IF NOT EXISTS bot_sessions (
    app_user_id BIGINT PRIMARY KEY
        REFERENCES app_users (id) ON DELETE CASCADE,
    state       TEXT NOT NULL,
    payload     JSONB NOT NULL DEFAULT '{}'::jsonb,
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);
