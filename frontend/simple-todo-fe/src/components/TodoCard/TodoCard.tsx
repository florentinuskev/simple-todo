import React from 'react'

// React Icons
import { HiTrash, HiPencilSquare } from "react-icons/hi2";

const TodoCard = ({ todo, setSelTodo, openEdit, openDelete, updateDone }: any) => {
    return (
        <div className={`flex flex-row justify-between mx-2 my-2 p-10 shadow-lg items-center ${todo.is_done ? 'bg-gray-300' : ''}`}>
            <p className='text-3xl'>{todo.todo}</p>
            <div className='flex flex-row'>
                <input type="checkbox" name="" id="" className='mx-5' defaultChecked={todo.is_done} onChange={(e) => {
                    setSelTodo()
                    updateDone({ isDone: e.target.checked, id: todo.id })
                }} />
                <HiPencilSquare size={20} className='mr-5' onClick={() => {
                    setSelTodo()
                    openEdit()
                }} />
                <HiTrash size={20} onClick={() => {
                    setSelTodo()
                    openDelete()
                }} />
            </div>
        </div>
    )
}

export default TodoCard