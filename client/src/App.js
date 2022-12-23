import React from "react";
import {Routes, Route, Link} from 'react-router-dom'
import Index from './components/Index'
import SingIn from './components/SingIn'
import SingUp from './components/SingUp'
import styled from 'styled-components'

const Header = styled.header `
	display: flex;
	justify-content: space-between;
	padding: 2rem 3rem;
	align-items: center;
	background-color: #111;

	h1 {
		font-size: 3rem;
	}

	a {
		color: #fff;
		font-size: 1.8rem;
		margin: 0 1rem;

		&:hover {
			cursor: pointer;
			color: #ddf;
		}
	}
`

const Main = styled.main `
	margin-top: 1rem;
`

export default function App () {
    return (
    <>
		<Header>
			<h1>Dev City</h1>

			<nav>
				<Link to='/'>Home</Link>
				<Link to='/singin'>Sing-in</Link>
				<Link to='/singup'>Sing-up</Link>
			</nav>
		</Header>

		<Main>
			<Routes>
				<Route path="/" element={<Index/>}></Route>
				<Route path="/singin" element={<SingIn/>}></Route>
				<Route path="/singup" element={<SingUp/>}></Route>
			</Routes>
		</Main>
    </>
    )
}