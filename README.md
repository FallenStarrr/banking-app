# banking-app

# account system
# Switch to the dev branch
# go run main.go to run - start up the project
Stack:
- DB: PostgresSQL
- Language:  Golang
- Router: Fiber
- Communication: RestAPI
- VCS - Git, GitHub
- Editor VS Code
- Request Sending - Postman

base_url : localhost:3000
- /account/15?field=type&value=credit - PUT update one account by id and query params
- /account/:id - GET one account by id
- /account/:id - DELETE delete one account by id
- /account -  POST create one account
- /send/account - PUT send money -
  Query Params:
  - from: str - id of the sender
  - to: str - id of the the person who get paid
  - amount: float64 - amount of money to be send
- /block/:id - PUT block the account
  
