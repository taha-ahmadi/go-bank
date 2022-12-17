# A Bank
A bank to create and manage account (owner, balance, currency) and Record all balances changes,
Money transfer transaction (Perform money transfer between 2 account consistently within a transaction)

### Database diagram - UML
![UML](documents/assets/uml.png)

TODO:

- [ ] make DockerFile & docker-compose
- [x] use PostgresSQL
- [ ] Implement CRUD transaction services
- [ ] Fix DB transaction lock and handle deadlock issue
- [ ] use gRPC
- [x] Setup GitHub action
- [ ] Implement user login with JWT access token
- [ ] use k8s
