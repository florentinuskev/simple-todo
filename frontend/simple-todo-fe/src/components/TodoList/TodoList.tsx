import React from 'react'

// Components
import TodoCard from '../TodoCard/TodoCard'

const TodoList = ({ todos }: any) => {
    return (
        <div>
            {todos.map((todo: any, i: any) => (
                <TodoCard key={i} todo={todo} />
            ))}
        </div>
    )
}

export default TodoList