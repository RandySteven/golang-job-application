CREATE DATABASE job_application_db;

DROP DATABASE job_application_db;

SELECT * FROM users;
SELECT * FROM auths;
SELECT * FROM jobs;

INSERT INTO users (name, date_of_birth)
VALUES
    ('Randy Steven', '2001-04-11'),
    ('Juju Man', '2001-05-01');

INSERT INTO auths (email, password, user_id)
VALUES
    ('randysteven12@gmail.com', 'test_1234', 1),
    ('jujuman@gmail.com', 'test_1234', 2);

INSERT INTO jobs (name, quota, status, job_poster_id, expiry_date)
VALUES
    ('Software Engineer', 12, 'Open', 1, '2024-12-01'),
    ('Software Development Engineer in Test', 21, 'Open', 1, '2024-11-04'),
    ('Data Sciencetist', 11, 'Open', 1, '2024-10-10');