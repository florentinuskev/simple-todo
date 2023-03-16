import React, { useState, useRef, useEffect } from 'react'

import axios from 'axios'

// Components
import Navbar from '../components/Navbar/Navbar'
import TodoList from '../components/TodoList/TodoList'
import Fab from '../components/FAB/FAB'

const MainPage = () => {

    const [profile, setProfile] = useState({});
    const [todos, setTodos] = useState([]);

    const fetchProfile = async () => {
        try {
            const res = await axios.get("http://127.0.0.1:3000/api/v1/auth/profile", {
                headers: {
                    Authorization: localStorage.getItem("token")
                }
            })

            if (res.status == 200) {
                setProfile(res.data?.user);
            }
        } catch (err) {
            console.log(err);
        }
    }

    const fetchTodos = async () => {
        try {
            const res = await axios.get("http://127.0.0.1:3000/api/v1/todos/", {
                headers: {
                    Authorization: localStorage.getItem("token")
                }
            })

            if (res.status == 200) {
                setTodos(res.data.todos)
            }
        } catch (err) {
            console.log(err);
        }
    }

    const dataFetcherRef = useRef(false);
    useEffect(() => {
        if (dataFetcherRef.current) return;
        dataFetcherRef.current = true
        fetchProfile();
        fetchTodos();
    }, []);

    return (
        <div>
            <Navbar profile={profile} />
            <TodoList todos={todos} />
            <Fab />
        </div>
    )
}

export default MainPage