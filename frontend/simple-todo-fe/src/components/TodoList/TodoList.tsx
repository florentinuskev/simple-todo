import React from 'react'

// Components
import TodoCard from '../TodoCard/TodoCard'

const TodoList = ({ todos, setSelTodo, openEdit, openDelete, updateDone }: any) => {
    return (
        <div>
            {todos.map((todo: any, i: any) => (
                <TodoCard key={i} todo={todo}
                    setSelTodo={() => setSelTodo(todo)}
                    openEdit={() => openEdit()}
                    openDelete={() => openDelete()}
                    updateDone={(done: any) => { updateDone(done) }}
                />
            ))}
        </div>
    )
}

export default TodoList