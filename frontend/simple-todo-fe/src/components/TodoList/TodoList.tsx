import React from 'react'

// Components
import TodoCard from '../TodoCard/TodoCard'

const TodoList = ({ todos, setSelTodo, openEdit, openDelete }: any) => {
    return (
        <div>
            {todos.map((todo: any, i: any) => (
                <TodoCard key={i} todo={todo} setSelTodo={() => setSelTodo(todo)} openEdit={() => openEdit()} openDelete={() => openDelete()} />
            ))}
        </div>
    )
}

export default TodoList