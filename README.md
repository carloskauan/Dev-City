# Dev City

uma comunidade simples mas poderosa para devs interagir!

## Conteúdo

- [Screenshot](#screenshot)
<!-- - [Checklist](#checklist) -->
- [Como criei o projeto](#como-criei-o-projeto)
    - [Tecnologias](#tecnologias)
    - [Minha maior dificuldade](#minha-maior-dificuldade)
- [O que aprendi](#o-que-aprendi)
- [Links](#links)
- [Créditos](#créditos)
- [Autor(es)](#autores)

## Screenshot

<!-- ![#](imagem do projeto) -->

<!-- 
    ## Checklist do que adicionar no projeto

    - [x] Escrever o README.md
    - [ ] colocar uma screenshot no readme
    - [ ] adicionar a opção de dark mode no projeto
    - [x] colocar o projeto no ar 
-->

## Como criei o projeto
<!-- coloque aqui os passo (claro que você pode colocar as coisas que você achar mais relevantes) que você fez para criar o projeto -->

### Tecnologias
<!-- liste algumas tecnologias que você usou no projeto, exemplo -->

- HTML
- CSS
- JAVASCRIPT
- Nodejs
- Golang
- Gorm
- Gin

### Minha maior dificuldade
<!-- coloque aqui sua maior dificuldade e como você fez para solucionar ela ou peça ajudar para o leitor em relação a sua dificuldade, também mencione o artigo ou usuário que te ajudou a resolver  -->

## O que aprendi
<!-- coloque aqui o que você aprendeu nesse projeto -->

## Links
<!-- coloque links sobre o projeto, como um protótipo no ar -->

[link que o usuário vai querer acessar](http://teste.com)

## Créditos
<!-- coloque aqui os conteúdos ou usuário que ajudaram a criar o projeto -->

## Autor(es)
<!-- coloque links relacionados as suas redes sociais a as pessoas que participaram no projeto -->

- carloskauan -[Github](https://github.com/carloskauan)
- Victor Eduardo Art -[Redes Sociais](https://linktr.ee/victor_eduardo_art)

# Api
## Users Endpoints
### Estrutura
~~~json
{
    "name":"string",
    "email":"string",
    "password":"string",
    "bio":"string"
}
~~~
> Atente-se aos tipos de dados de cada campo
### /user
#### Post
~~~go
"/user"
~~~
Este metodo registra os dados do usuario no banco e retorna esses dados
> Requer o JSON com os dados

Esta etapa contem um validador de email e caso seja passado um email invalido a operação de cadastro e cancelada e e temos o seguinte retorno
~~~json
{
    "error":"invalid email"
}
~~~

#### Get
~~~go
"/user"
~~~
Este metodo retorna todos os usarios registrados no banco

### /user/"email"
#### Get
~~~
get("/user/"email")
~~~
Esta rota retorna retorna os dados do usuario do email informado

### /user/"email"/"senha"
#### Get
Este metodo deve ser chamado quando uma autenticação de usuario for necessario. O metodo retorna um status com um booleano

Exemplo:
~~~
get("/user/jão12345@gmail.com/12345678")
~~~
Caso as credenciais estjam certas e retornado
~~~json
{
    "status": true
}
~~~
Caso o usuario estejam errado ou não exista e retornado
~~~json
{
    "error":"User not found",
    "status": false
}
~~~
E em caso de senha erra temos 
~~~json
{
    "error":"Password don't macth",
    "status":false
}
~~~

### /user/verify/"email"
#### Post
O metodo requer um email na qual sera gerado um pin de 6 numeros e enviado para o email informado e tbm sera salvo o pin no servidor . O pin sera guardado criptografado juntamente com o email.

Esse metodo retorna dois estados. Um erro ou um aviso de sucesso de envio
