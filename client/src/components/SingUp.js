import React, { useEffect, useState } from "react";
import Axios from 'axios'
const baseUrl = 'http://localhost:8080'

export default function SingUp () {
    useEffect(() => {
        document.getElementsByTagName('header')[0].classList.add('center')
        document.getElementsByTagName('main')[0].classList.add('center')
    })

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
        setForm({
            email: form.email.toString(), 
            password: form.password.toString()
        })

        Axios.get(`${baseUrl}/user/${form.email}/${form.password}`).then(res => {console.log(res.data)})
    }

    const mudarForm = () => {
        let url = window.location.href

        let urlTmp = url.split('/')
        url = urlTmp[3]

        console.log(url)

        if (url !== 'signup') {
            window.location.href = `${urlTmp[0]}/${urlTmp[1]}/${urlTmp[2]}/singin`
            console.log(window.location.href)
        }
    }

    return (
        <div className="form">
            <div className="signup">
                <h2>Sing up</h2>

                <div className="groupInputs">
                    <div className="groupInput">
                        <label htmlFor='name'>E-mail:</label>

                        <input type='text' placeholder="Digite seu e-mail" name='email' id="email" value={form.email} onChange={(e) => {getValues(e)}} />
                    </div>

                    <div className="groupInput">
                        <label htmlFor='password'>Senha</label>

                        <input type='text' placeholder="Digite sua senha" name='password' id="password" value={form.password} onChange={(e) => {getValues(e)}} />
                    </div>
                </div>

                <button onClick={() => {getUser()}}>Sign up</button>
            </div>
            
            <button className="mudarForm" onClick={() => {mudarForm()}}>Sign in</button>
        </div>
    )
}