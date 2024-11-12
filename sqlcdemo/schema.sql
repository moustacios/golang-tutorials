-- ~/Workspace/sqlcdemo/schema.sql

CREATE TABLE comments (
  id              BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  email           VARCHAR(255)    NOT NULL,
  comment_text    TEXT            NOT NULL,
  bot_probability TINYINT         DEFAULT 0
);
