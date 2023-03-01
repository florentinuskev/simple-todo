package repository

const CreateUserQuery = `INSERT INTO users (id, username, password, created_at, updated_at) VALUES ($1, $2, $3, time.now(), time.now()) RETURNING *;`
const FindUserByIdQuery = `SELECT * FROM users WHERE id = $1;`
const FindUserByUsernameQuery = `SELECT * FROM users WHERE username = $2;`
