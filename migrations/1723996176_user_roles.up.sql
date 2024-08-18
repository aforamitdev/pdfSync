CREATE TABLE IF NOT EXISTS user_roles(id INT PRIMARY KEY, system_role TEXT);
CREATE TABLE IF NOT EXISTS user_roles_mapping(
    id INT PRIMARY KEY,
    user_id INT,
    role_id INT,
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(role_id) REFERENCES user_roles(id)
)