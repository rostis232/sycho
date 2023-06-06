CREATE TABLE users (
    user_id SERIAL PRIMARY KEY NOT NULL,
    login VARCHAR(20) NOT NULL,
    pass VARCHAR(20) NOT NULL,
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(20) NOT NULL,
    email VARCHAR(30),
    phone VARCHAR(13),
    role INT DEFAULT 0 NOT NULL,
);

-- Admin account insertion
INSERT INTO users (login, pass, first_name, last_name, role)
VALUES ('admin', 'admin', 'Admin', 'Admin', 1);

CREATE TABLE beneficiaries (
    bnf_id SERIAL PRIMARY KEY NOT NULL,
    first_name VARCHAR(20) NOT NULL,
    middle_name VARCHAR(20),
    last_name VARCHAR(20) NOT NULL,
    phone VARCHAR(13),
    birthday VARCHAR(10),
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE SET NULL
);

CREATE TABLE activities (
    act_id SERIAL PRIMARY KEY NOT NULL,
    time VARCHAR(20) NOT NULL,
    bnf_id INT NOT NULL,
    user_id INT NOT NULL,
    description TEXT DEFAULT NULL,
    FOREIGN KEY (bnf_id) REFERENCES beneficiaries (bnf_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE SET NULL
);

CREATE TABLE sessions (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE
);