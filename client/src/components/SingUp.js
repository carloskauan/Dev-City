import React, { useState } from "react";
import Form from '../styles/form'
import Axios from 'axios'
const baseUrl = 'http://localhost:8080'

export default function SingUp () {
    const [form, setForm] = useState({
        email: 'Seu email aqui',
        password: 'Sua senha aqui'
    })

    const getValues = (e) => {
        if (typeof(e.target.value) === 'string' && e.target.value !== null) {
            setForm({...form, [e.target.name]: e.target.value})
        }

        e.target.classList.add('on')
    }

    const getUser = () => {
        Axios.get(`${baseUrl}/user/${form.email}/${form.password}`).then(res => {console.log(res.data)})
    }

    return (
    <Form>
        <h2>Sing up</h2>

        <div>
            <label>
                E-mail:

                <input type='text' name='email' value={form.email} onChange={(e) => {getValues(e)}} />
            </label>

            <label>
                Senha:

                <input type='password' name='password' value={form.password} onChange={(e) => {getValues(e)}} />
            </label>
        </div>

        <button onClick={() => {getUser()}}>Criar conta</button>
    </Form>
    )
}