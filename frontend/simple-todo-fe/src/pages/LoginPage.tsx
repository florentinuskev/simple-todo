import axios from 'axios';
import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom';

const LoginPage = () => {

    const navigate = useNavigate();

    const loginFunc = async (e: any) => {
        e.preventDefault();
        try {
            const res = await axios.post("http://127.0.0.1:3000/api/v1/auth/login", {
                username: username,
                password: password
            })

            if (res.status == 200) {
                localStorage.setItem("token", res.data.token)
                navigate("/")
            }
        } catch (err) {
            console.log(err);
        }
    }

    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    return (
        <div className='flex justify-center items-center min-h-[100vh]'>
            <div className='flex flex-col shadow-lg rounded p-10 items-center w-[30%]'>
                <h1 className='text-3xl mb-10'>Auth</h1>
                <input type="text" className='p-2 w-[100%] border rounded my-1' placeholder='Username' onChange={(e) => { setUsername(e.target.value) }} />
                <input type="text" className='p-2 w-[100%] border rounded my-1' placeholder='Password' onChange={(e) => { setPassword(e.target.value) }} />

                <button className='border rounded p-2 w-[100%] my-3'><p className='font-bold' onClick={(e) => { loginFunc(e) }}>Login</p></button>
            </div>
        </div>
    )
}

export default LoginPage