import React from "react";
import {Routes, Route, Link} from 'react-router-dom'
import SingIn from './components/SingIn'
import SingUp from './components/SingUp'
import Index from './components/Index'

export default function App () {
    return (
    <>
		<header>
			<h1>Dev City</h1>

			<nav>
				<Link to='/about'>Sobre</Link>
				<Link to='/singin'>Sing-in</Link>
				<Link to='/singup'>Sing-up</Link>
				<Link to='/'>Home</Link>
			</nav>
		</header>

		<main>
			<Routes>
				<Route path="/" element={<Index/>}></Route>
				<Route path="/singin" element={<SingIn/>}></Route>
				<Route path="/singup" element={<SingUp/>}></Route>
				{/* <Route path="/abouut" element={<About/>}></Route> */}
			</Routes>
		</main>
    </>
    )
}