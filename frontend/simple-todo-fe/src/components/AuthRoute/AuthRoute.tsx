import React, { useState, useRef, useEffect } from 'react'
import { useNavigate } from 'react-router-dom';

import axios from 'axios';

const AuthRoute = ({ children }: any) => {

    const [isLogged, setLogged] = useState(false);
    const [profile, setProfile] = useState({});

    const navigate = useNavigate();

    const fetchData = async () => {
        try {
            const res = await axios.get("http://127.0.0.1:3000/api/v1/auth/profile", {
                headers: {
                    Authorization: localStorage.getItem("token")
                }
            })

            if (res.status == 200) {
                setLogged(true);
                setProfile(res.data?.user);
            } else {
                navigate("/login")
            }
        } catch (err) {
            navigate("/login")
        }
    }

    const dataFetcherRef = useRef(false);
    useEffect(() => {
        if (dataFetcherRef.current) return;
        dataFetcherRef.current = true
        fetchData();
    }, []);


    return (
        <div>
            {children}
        </div>
    )
}

export default AuthRoute