package repository

const GetTodosQuery = `SELECT * FROM todos WHERE uid = $1;`
const GetTodoQuery = `SELECT * FROM todos WHERE id = $1;`
const NewTodoQuery = `INSERT INTO todos (id, uid, todo) VALUES ($1, $2, $3) RETURNING *;`
const UpdateTodoQuery = `UPDATE todos SET todo = COALESCE(NULLIF($1, ""), todo), updated_at = NOW() WHERE id=$2;`
const DeleteTodoQuery = `DELETE FROM todo WHERE id = $1;`
