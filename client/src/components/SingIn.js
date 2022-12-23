import React, { useState } from "react";
import Form from '../styles/form'
import Axios from 'axios'
const baseUrl = 'http://localhost:8080'

export default function SingIn () {
    const [form, setForm] = useState({
        name: 'Seu nome aqui',
        email: 'Seu email aqui',
        password: 'Sua senha aqui',
        bio: 'Fale um pouco sobre você',
    })

    const getValues = (e) => {
        if (typeof(e.target.value) === 'string' && e.target.value !== null) {
            setForm({...form, [e.target.name]: e.target.value})
        }

        e.target.classList.add('on')
    }

    const sendValues = () => {
        Axios.get(`${baseUrl}/user`, {
            name: form.name,
            email: form.email,
            password: form.password,
            bio: form.bio
        }).then(res => {console.log(res.data)})
    }

    return (
    <Form>
        <h2>Sing in</h2>

        <div>
            <label>
                Nome:

                <input type='text' name='name' value={form.name} onChange={(e) => {getValues(e)}} />
            </label>

            <label>
                E-mail:

                <input type='text' name='email' value={form.email} onChange={(e) => {getValues(e)}} />
            </label>

            <label>
                Senha:

                <input type='text' name='password' value={form.password} onChange={(e) => {getValues(e)}} />
            </label>

            <label>
                Bio:

                <input type='text' name='bio' value={form.bio} onChange={(e) => {getValues(e)}} />
            </label>
        </div>

        <button onClick={() => {sendValues()}}>Login</button>
    </Form>
    )
}