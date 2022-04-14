ALTER TABLE users
	ADD creator varchar(100);

SET SQL_SAFE_UPDATES = 0;
UPDATE users
SET creator = 'andi 123'
WHERE username = 'bd';