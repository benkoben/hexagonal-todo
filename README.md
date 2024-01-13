# hexagonal-todo

A SaaS todo app which can be used from the terminal with cloud syncing.

# Architecture

I want to learn more about hexagonal architectures for this project. Which is an architecture that defines how an application is designed with a separation between core logic and external dependencies. Core logic is the business logic that the application is designed for. Whereas external dependencies might be databases, caches and services such as middle ware.
## Terminology:

### Core

Here is where the logic for handling todos are defined. 

I problably want to use hexagonal architecture combined with a microservices architecture. This way I can keep the todo logic contained and let the todo-core call external services for stuff like authentication.

> Lets start with the todo logic and then see how we want to do with middleware

### Ports (interfaces)

Ports define how the application communicates with external systems or services, kind of like a contact. For example there can be a port for connecting to a database or all other webservices. 

Ports belong to the core, because the core defines which actions are required to achieve the business logic goals.
### Adapters

Adapters are the ones who implement the contracts or interfaces defined by ports. Adapter are responsible for making sure the application can interact to databases, or other things. They handle the technical details.

### Driver Actors

Driver actors are the initiators of communication with the core. They reach out to the core to request specific services. Examples of driver actors can be HTTP request or command line interfaces (CLI).

### Driven Actors

Driven actors are the ones that triggered by the core. If the core needs something from external services, it sends a request to the adapter, instructing it to perform a particular action. For example, if the core needs to store data in a Postgres database, it triggers communication with the Postgres client to execute an INSERT query. In this scenario, the core initiates the communication.

# Infrastructure

The application should be container based so that I can deploy it to Kubernetes.

For persistence i want to use postgresql

# Database

```
Table List {
  id integer [primary key]
  name varchar
  created_at timestamp
}

Table Task {
  id integer [primary key]
  listid integer
  name varchar
  description text
  completed bool
  created_at timestamp
}

Table Subtask {
  id integer
  name varchar
  created_at timestamp
  completed bool
}

Ref: Task.id > List.id // many-to-ne

Ref: Subtask.id > Task.id

```
![[images/database.png]]

# Core libraries

Handling HTTP requests andresponses - Gin
Postgres SQL with pgx driver
slog for logging
User interface - Bubble tea 
