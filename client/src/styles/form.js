import styled from 'styled-components'

const Form = styled.div `
    background-color: #005;
    width: 25vw;
    margin-top: 5vh;
    margin: auto;
    border-radius: 1rem;
    border: 1px solid #fff;
    display: flex;
    flex-direction: column;
    padding: 1rem;

    h2 {
        text-align: center;
        margin-bottom: 1rem;
    }

    label {
        display: flex;
        flex-direction: column;
        margin: 1rem 0;

        input.on {
            width: 100%;
        }

        input {
            transition: 1s all;
            margin-top: 1rem;
            width: 55%;
            padding: 1rem;
            border-radius: 1rem;
            color: #fff;
            border: none;
            background-color: #002;
        }

        &:hover {
            input {
                width: 100%;
            }
        }
    }

    button {
        background-color: #558;
        color: #fff;
        padding: 0.8rem;
        border-radius: 1rem;
        margin-top: 1rem;
        font-size: 1.4rem;
        transition: 1s all;

        &:hover {
            opacity: 80%;
            cursor: pointer;
        }
    }
`

export default Form