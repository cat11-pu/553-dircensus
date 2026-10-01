-- ledger table
CREATE TABLE ledger (
	id INTEGER PRIMARY KEY,
	amount INTEGER NOT NULL
);

--no space after dashes
INSERT INTO ledger VALUES (1, 42);
