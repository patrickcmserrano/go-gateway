CREATE TABLE transactions (
    id TEXT NOT NULL,
    account TEXT NOT NULL,
    amount REAL NOT NULL,
    status TEXT NOT NULL,
    error_message TEXT NOT NULL,
    created_at TEXT NOT NULL, 
    updated_at TEXT NOT NULL,
    PRIMARY KEY (id)
);