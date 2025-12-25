CREATE TABLE IF NOT EXISTS tasks (
    id INT PRIMARY KEY NOT NULL,
    headline VARCHAR(255) NOT NULL,
    description TEXT ,
    done BOOL DEFAULT FALSE ,
    created_at TIMESTAMP DEFAULT NOW(),
    completed_at TIMESTAMP
    );
