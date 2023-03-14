package repository

const CreateUserQuery = `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING *;`
const FindUserByIdQuery = `SELECT * FROM users WHERE id = $1;`
const FindUserByUsernameQuery = `SELECT * FROM users WHERE username = $1;`
