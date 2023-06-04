CREATE TABLE projects (
    prj_id SERIAL PRIMARY KEY NOT NULL,
    short_title VARCHAR(20) NOT NULL,
    full_title VARCHAR(250),
    code VARCHAR(10)
);

CREATE TABLE organisations (
    org_id SERIAL PRIMARY KEY NOT NULL,
    title VARCHAR(30) NOT NULL,
    code VARCHAR(5)
);

-- Main organisation insertion
INSERT INTO organisations (title, code)
VALUES ('Карітас України', 'CU');

CREATE TABLE users (
    user_id SERIAL PRIMARY KEY NOT NULL,
    login VARCHAR(20) NOT NULL,
    pass VARCHAR(20) NOT NULL,
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(20) NOT NULL,
    email VARCHAR(30),
    phone VARCHAR(13),
    org_id INT NOT NULL,
    role INT DEFAULT 0 NOT NULL,
    FOREIGN KEY (org_id) REFERENCES organisations (org_id) ON DELETE RESTRICT
);

-- Admin account insertion
INSERT INTO users (login, pass, first_name, last_name, org_id, role)
VALUES ('admin', 'admin', 'Admin', 'Admin', 1, 1);

CREATE TABLE beneficiaries (
    bnf_id SERIAL PRIMARY KEY NOT NULL,
    first_name VARCHAR(20) NOT NULL,
    middle_name VARCHAR(20),
    last_name VARCHAR(20) NOT NULL,
    phone VARCHAR(13),
    birthday VARCHAR(10),
    prj_id INT,
    org_id INT,
    user_id INT,
    done BOOLEAN DEFAULT FALSE NOT NULL,
    FOREIGN KEY (prj_id) REFERENCES projects (prj_id) ON DELETE SET NULL,
    FOREIGN KEY (org_id) REFERENCES organisations (org_id) ON DELETE SET NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE SET NULL
);

CREATE TABLE reqs (
    req_id SERIAL PRIMARY KEY NOT NULL,
    title VARCHAR(100) NOT NULL
);

CREATE TABLE activities (
    act_id SERIAL PRIMARY KEY NOT NULL,
    time VARCHAR(20) NOT NULL,
    req INT,
    bnf_id INT NOT NULL,
    user_id INT,
    FOREIGN KEY (req) REFERENCES reqs (req_id) ON DELETE SET NULL,
    FOREIGN KEY (bnf_id) REFERENCES beneficiaries (bnf_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE SET NULL
);

CREATE TABLE prj_org (
    po_id SERIAL PRIMARY KEY NOT NULL,
    prj_id INT NOT NULL,
    org_id INT NOT NULL,
    FOREIGN KEY (prj_id) REFERENCES projects (prj_id) ON DELETE CASCADE,
    FOREIGN KEY (org_id) REFERENCES organisations (org_id) ON DELETE CASCADE
);

CREATE TABLE prj_usr(
    pu_id SERIAL PRIMARY KEY NOT NULL,
    prj_id INT NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (prj_id) REFERENCES projects (prj_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE
);