/* Caution : each query must be separated with empty comment */

CREATE TABLE IF NOT EXISTS mail
(
    _id     integer PRIMARY KEY AUTOINCREMENT not null,
    date    date,
    "to"    varchar(255)                      not null,
    subject varchar(255)                      not null,
    sent    boolean                           not null,
    error   text,
    viewed  boolean
);

--

CREATE TRIGGER set_viewed_param
    BEFORE INSERT
    ON "mail"
BEGIN
    INSERT INTO "mail" ("to", "subject", "sent", "error", "viewed", "date")
    VALUES (NEW."to",
            NEW."subject",
            NEW."sent",
            NEW."error",
            CASE WHEN NEW."error" IS NULL THEN 1 ELSE 0 END,
            DATE('now'));
    SELECT RAISE(IGNORE);
END;