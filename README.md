# Sat Store

this is an online store Back-End API Written with go-lang 

i wrote it as a training 

- [Sat Store](#sat-store)
  - [Used technologies](#used-technologies)
  - [TODO List](#todo-list)
  - [Routes](#routes)
    - [User's Routes](#users-routes)
      - [Register](#register)
        - [Request](#request)
        - [Response](#response)
      - [Login](#login)
        - [Request](#request-1)
        - [Response](#response-1)
      - [Get current user](#get-current-user)
        - [Request](#request-2)
        - [Response](#response-2)

## Used technologies 

| package |          import          |                using                 |
| :------ | :----------------------: | :----------------------------------: |
| gin     | github.com/gin-gonic/gin | web framework written in Go (Golang) |

## TODO List 

- [x] select file structure 
- [x] select database type ( SQLite )
- [x] create models ( users - products - categories - cart )
  - [x] users 
  - [x] categories
  - [x] products
  - [x] cart
- [ ] create routs
  - [x] setup routs gropes
  - [x] user routs
    - [x] login
    - [x] register
    - [x] get current user
  - [ ] product routs
  - [ ] category routs
  - [ ] cart routs
- [x] setup auth system ( JWT )
- [ ] request validation
- [ ] error handling
  - [ ] UNIQUE constraint failed: users.email

## Routes 

### User's Routes 

#### Register
`POST:domain:8080/api/v1/user/register`

##### Request 
  ``` 
  body : {
    name,
    email,
    username,
    password,
  }
  ```

##### Response 
  ```
  body : {
    name,
    username,
    token,
  }
  ```

#### Login
`POST:domain:8080/api/v1/user/login`

##### Request 
  ``` 
  body : {
    username,
    password,
  }
  ```

##### Response 
  ```
  body : {
    name,
    username,
    token,
  }
  ```

#### Get current user 
`GET:domain:8080/api/v1/user/me`

##### Request 
  ``` 
  head : {
    token,
  }
  ```

##### Response 
  ```
  body : {
    name,
    email,
    username,
    token,
  }
  ```