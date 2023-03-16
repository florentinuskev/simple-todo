import React from 'react'

// React Icons
import { HiTrash, HiPencilSquare } from "react-icons/hi2";

const TodoCard = ({ todo }: any) => {
    return (
        <div className='flex flex-row justify-between mx-2 p-10 shadow-lg items-center'>
            <p className='text-3xl'>{todo.todo}</p>
            <div className='flex flex-row'>
                <input type="checkbox" name="" id="" className='mx-5' />
                <HiPencilSquare size={20} className='mr-5' />
                <HiTrash size={20} />
            </div>
        </div>
    )
}

export default TodoCard