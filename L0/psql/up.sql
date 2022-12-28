DROP TABLE IF EXISTS msg;
CREATE TABLE msg (
                       id VARCHAR(32) PRIMARY KEY,
                       body TEXT NOT NULL,
                       created_at TIMESTAMP WITH TIME ZONE NOT NULL
);