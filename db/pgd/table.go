package pgd

var TableSql = `

 CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL
    );
    CREATE TABLE IF NOT EXISTS tasks (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        description TEXT NOT NULL,
        status VARCHAR(50) NOT NULL CHECK (status IN ('Pending', 'In Progress', 'Completed')),
        due_date DATE NOT NULL
    );
    CREATE TABLE IF NOT EXISTS map_users_with_tasks (
        user_id INT NOT NULL,
        task_id INT NOT NULL,
		assign_by INT NOT NULL,
        PRIMARY KEY (user_id, task_id),
        FOREIGN KEY (user_id) REFERENCES users(id),
        FOREIGN KEY (task_id) REFERENCES tasks(id)
    );
`
