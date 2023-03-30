-- name: RegisterUser :execresult
INSERT INTO users (username, password, name, email, mobile, creation_date, gender)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetUserPassword :one
SELECT password
FROM users
WHERE id = ? LIMIT 1;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = ? LIMIT 1;

-- name: GetUserByName :one
SELECT *
FROM users
WHERE username = ? LIMIT 1;

-- name: GetUserByPhone :one
SELECT *
FROM users
WHERE mobile = ? LIMIT 1;

-- name: GetUserPhoneByID :one
SELECT mobile
FROM users
WHERE id = ? LIMIT 1;

-- name: UpdateUserLoginDate :exec
UPDATE users
SET login_date          = sqlc.arg(login_date),
    login_failure_count = sqlc.arg(login_failure_count)
where id = ?
   or username = ?
   or mobile = ?;

-- name: UpdateUserPassword :exec
UPDATE users
SET password = sqlc.arg(password)
where id = ?
   or username = ?;

-- name: UpdateUserInfo :exec
UPDATE users
SET avatar   = sqlc.arg(avatar),
    username = sqlc.arg(username),
    name     = sqlc.arg(name),
    gender   = sqlc.arg(gender),
    birth    = sqlc.arg(birth)
where id = ?;

-- name: UpdateUserUpdateDate :exec
UPDATE users
SET last_updated_date = sqlc.arg(last_updated_date)
where id = ?
   or username = ?
   or mobile = ?;