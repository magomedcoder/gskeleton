CREATE TABLE user_logs
(
    user_id    Int64,
    log        String,
    created_at DateTime default now()
) ENGINE = MergeTree()
      ORDER BY (user_id, created_at);
