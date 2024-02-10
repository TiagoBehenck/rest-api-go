## My first API using GO

In this document, we will guide you through the process of installing Go and running a simple API using Go. We will cover everything from installing Go to running a sample application.

### Installing Go

To get started, you need to install Go on your system. Follow these steps:

#### Linux

1. Open a terminal.
2. Execute the following command to download the Go installer:

   ```bash
   wget https://golang.org/dl/go1.XX.X.linux-amd64.tar.gz
   ```

   Replace `XX.X` with the latest version of Go available on the [official Go website](https://golang.org/dl/).

3. Unpack the downloaded file:

   ```bash
   sudo tar -C /usr/local -xzf go1.XX.X.linux-amd64.tar.gz
   ```

4. Add the Go bin directory to your PATH by adding the following line to your `~/.profile` or `~/.bashrc` file:

   ```bash
   export PATH=$PATH:/usr/local/go/bin
   ```

5. Update your shell environment:

   ```bash
   source ~/.profile
   ```

#### macOS

1. Download the Go installer from [https://golang.org/dl/](https://golang.org/dl/).
2. Open the installer and follow the on-screen instructions.

#### Windows

1. Download the Go installer from [https://golang.org/dl/](https://golang.org/dl/).
2. Open the installer and follow the on-screen instructions.

### Running the API

Now that you have installed Go, let's create and run a simple API as an example.

1. Clone this repo using the command:

   ```bash
    git clone https://github.com/TiagoBehenck/rest-api-go
   ```

2. In the terminal, run the following command to install the dependencies:
   
   ```bash
   go mod tidy
   ```

3. Not just run the following command to start the server:

   ```bash
   go run main.go
   ```

4. Open a web browser and access `http://localhost:9000`. You should see the "null" message, because doesn't has todo yet

### Create the postgres database using docker

To create a new container using the image of postgres, follow the command: 

1. First, create the container. The name of container is api-todo, the deafult port is 5432 and the password is 1234:
  ```bash 
  docker run -d --name api-todo -p 5432:5432 -e POSTGRES_PASSWORD=1234 postgres:13.5
  ```

2. Now, open the container to create the database, table and user:
  ```bash
  docker exec -it api-todo psql -U postgres
  ```

3. Into the postgres, create the database:
  ```bash
  create database api_todo;
  ```

4. Create the user:
  ```bash
  create user user_todo;
  ```

5. Create the password to user

  ```bash
  alter user user_todo with encrypted password '1122';
  ```

6. Add permissions

  ```bash
  grant all privileges on database api_todos to user_todo;
  grant all privileges on all sequences in schema public to user_todo;
  grant all privileges on all tables in sequence public to user_todo;
  ```

7. Now create the table
   1. First, connect to database:
   ```bash
   \c api_todo
   ```
   2. Just create the table with correct columns:
    ```bash
   create table todos (id serial primary key, title varchar, description text, done bool default FALSE);
   ```
   3. To see the new table, run the command: 
   ```bash
   \dt
   ```

If all steps were successful, your API with the database connection should be working correctly :)