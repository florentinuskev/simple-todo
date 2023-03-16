import React, { useState, useRef, useEffect } from 'react'

import axios from 'axios'

// Components
import Navbar from '../components/Navbar/Navbar'
import TodoList from '../components/TodoList/TodoList'
import Fab from '../components/FAB/Fab'
import Modal from '../components/Modal/Modal'

const MainPage = () => {

    const [profile, setProfile] = useState({});
    const [todos, setTodos] = useState([]);

    // Modal State
    const [newModalOpen, setNewModalOpen] = useState(false);
    const [editModalOpen, setEditModalOpen] = useState(false);
    const [delModalOpen, setDelModalOpen] = useState(false);

    // Another state
    const [activity, setActivity] = useState("");
    const [selectedTodo, setSelTodo] = useState({
        id: "",
        uid: "",
        todo: "",
        is_done: false,
    });

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

    const createTodo = async (e: any) => {
        e.preventDefault()
        try {
            const res = await axios.post("http://127.0.0.1:3000/api/v1/todos/", {
                todo: activity
            }, {
                headers: {
                    Authorization: localStorage.getItem("token")
                },
            })

            if (res.status == 201) {
                fetchTodos();
                setActivity("");
            }

        } catch (err) {
            console.log(err);
        }
    }

    const editTodo = async (e: any) => {
        e.preventDefault();
        try {
            const res = await axios.patch("http://127.0.0.1:3000/api/v1/todos/" + selectedTodo.id, {
                todo: activity
            }, {
                headers: {
                    Authorization: localStorage.getItem("token")
                },
            })

            if (res.status == 200) {
                fetchTodos();
                setEditModalOpen(false);
            }
        } catch (err) {
            console.log(err);
        }
    }

    const updateTodoDone = async (done: any) => {
        try {
            console.log("HELLO DONE", done);
            const res = await axios.patch("http://127.0.0.1:3000/api/v1/todos/" + done.id, {
                is_done: done.isDone
            }, {
                headers: {
                    Authorization: localStorage.getItem("token")
                },
            })

            if (res.status == 200) {
                fetchTodos();
            }
        } catch (err) {
            console.log(err);
        }
    }

    const deleteTodo = async (e: any) => {
        e.preventDefault();
        try {
            const res = await axios.delete("http://127.0.0.1:3000/api/v1/todos/" + selectedTodo.id, {
                headers: {
                    Authorization: localStorage.getItem("token")
                },
            })

            if (res.status == 200) {
                fetchTodos();
                setDelModalOpen(false);
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
            <TodoList todos={todos}
                setSelTodo={(todo: any) => setSelTodo(todo)}
                openEdit={() => setEditModalOpen(true)}
                openDelete={() => setDelModalOpen(true)}
                updateDone={(done: any) => { updateTodoDone(done) }}
            />

            <Fab onClick={() => setNewModalOpen(true)} />
            <Modal title={"Create New Todo"} desc={"Enter the activity you want to add to database."} isOpen={newModalOpen} setIsOpen={(open: any) => setNewModalOpen(open)}>
                <h1>Activity:</h1>
                <input type="text" className='border rounded w-[100%] p-3 my-3' onChange={e => setActivity(e.target.value)} />

                <div className='flex flex-row justify-between'>
                    <button className='bg-green-200 rounded p-3' onClick={(e) => createTodo(e)}>Create</button>
                    <button className='bg-red-200 rounded p-3' onClick={() => setNewModalOpen(false)}>Cancel</button>
                </div>
            </Modal>
            <Modal title={"Edit Todo"} desc={"Enter the activity you want to put for current todo."} isOpen={editModalOpen} setIsOpen={(open: any) => setEditModalOpen(open)}>
                <h1>Activity:</h1>
                <input type="text" className='border rounded w-[100%] p-3 my-3' onChange={e => setActivity(e.target.value)} defaultValue={selectedTodo.todo} />

                <div className='flex flex-row justify-between'>
                    <button className='bg-green-200 rounded p-3' onClick={(e) => editTodo(e)}>Edit</button>
                    <button className='bg-red-200 rounded p-3' onClick={() => setEditModalOpen(false)}>Cancel</button>
                </div>
            </Modal>
            <Modal title={"Delete Todo Confirmation"} desc={"Confirm to delete the todo, are you sure?"} isOpen={delModalOpen} setIsOpen={(open: any) => { setDelModalOpen(open) }}>
                <div className='flex flex-row justify-between'>
                    <button className='bg-green-200 rounded p-3' onClick={(e) => deleteTodo(e)}>Delete</button>
                    <button className='bg-red-200 rounded p-3' onClick={() => setEditModalOpen(false)}>Cancel</button>
                </div>
            </Modal>
        </div>
    )
}

export default MainPage