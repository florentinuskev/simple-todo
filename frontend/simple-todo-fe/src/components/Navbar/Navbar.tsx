import React from 'react'

const Navbar = ({ profile }: any) => {
    return (
        <div className='flex flex-row flex-1 items-center justify-center bg-blue-400 text-white p-3'>
            <h1 className='mr-auto ml-2 text-3xl font-bold'>Simple Todo</h1>
            <div className='flex flex-row ml-auto mr-2'>
                <p className='text-lg mx-2'>{profile.username}</p>
                <p className='text-lg'>Logout</p>
            </div>
        </div>
    )
}

export default Navbar