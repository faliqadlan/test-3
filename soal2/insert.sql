CREATE TABLE users (
	user_id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by varchar(100),
    username varchar(100),
    name VARCHAR(200));

INSERT INTO users(created_by, username, name)
	VALUES 
		('','andi123','andi'),
        ('1','bd','budi');