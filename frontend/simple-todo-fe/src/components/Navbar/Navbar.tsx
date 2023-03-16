import React from 'react'
import { Menu } from '@headlessui/react'
import { useNavigate } from 'react-router-dom'

const Navbar = ({ profile }: any) => {

    const navigate = useNavigate();

    return (
        <div className='flex flex-row flex-1 items-center justify-center bg-blue-400 text-white p-3'>
            <h1 className='mr-auto ml-2 text-3xl font-bold'>Simple Todo</h1>
            <div className='flex flex-row ml-auto mr-2'>
                <p className='text-lg mx-2'>{profile.username}</p>
                <Menu as="div" className="relative inline-block text-left">
                    <Menu.Button>More</Menu.Button>
                    <Menu.Items className="absolute right-0 w-[100px] p-2 divide-gray-100 rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none text-black">
                        <Menu.Item as="div" className="hover:bg-gray-200" onClick={(e) => {
                            e.preventDefault()
                            localStorage.removeItem("token")
                            navigate("/login");
                        }}>
                            {({ active }) => (
                                <a
                                >
                                    Logout
                                </a>
                            )}
                        </Menu.Item>
                    </Menu.Items>
                </Menu>
            </div>
        </div>
    )
}

export default Navbar