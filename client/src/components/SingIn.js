import React, { useEffect, useState } from "react";
import Axios from 'axios'
const baseUrl = 'http://localhost:8080'

export default function SingIn () {
    useEffect(() => {
        document.getElementsByTagName('header')[0].classList.add('center')
        document.getElementsByTagName('main')[0].classList.add('center')
    })

    const [form, setForm] = useState({
        name: '',
        email: '',
        password: ''
    })

    const getValues = (e) => {
        if (typeof(e.target.value) === 'string' && e.target.value !== null) {
            setForm({...form, [e.target.name]: e.target.value})
        }

        e.target.classList.add('on')
    }

    const sendValues = () => {
        setForm({
            name: form.name.toString(),
            email: form.email.toString(), 
            password: form.password.toString()
        })

        Axios.post(`${baseUrl}/user`, {
            name: form.name,
            email: form.email,
            password: form.password,
            bio: 'teste'
        }).then(res => {console.log(res.data)})
    }

    const mudarForm = () => {
        let url = window.location.href

        let urlTmp = url.split('/')
        url = urlTmp[3]

        console.log(url)

        if (url !== 'signin') {
            window.location.href = `${urlTmp[0]}/${urlTmp[1]}/${urlTmp[2]}/singup`
            console.log(window.location.href)
        }
    }

    return (
        <div className="form">
            <div className="signin">
                <h2>Sing in</h2>

                <div className="groupInputs">
                    <div className="groupInput">
                        <label htmlFor='name'>Nome de usuário:</label>

                        <input type='text' placeholder="Digite seu nome de usuário" name='name' id="name" value={form.name} onChange={(e) => {getValues(e)}} />
                    </div>

                    <div className="groupInput">
                        <label htmlFor='email'>E-mail:</label>

                        <input type='text' placeholder="Digite seu e-mail" name='email' id="email" value={form.email} onChange={(e) => {getValues(e)}} />
                    </div>

                    <div className="groupInput">
                        <label htmlFor='password'>Senha</label>

                        <input type='text' placeholder="Digite sua senha" name='password' id="password" value={form.password} onChange={(e) => {getValues(e)}} />
                    </div>
                </div>

                <button onClick={() => {sendValues()}}>Sign in</button>
            </div>
            <button className="mudarForm" onClick={() => {mudarForm()}}>Sign up</button>
        </div>
    )
}