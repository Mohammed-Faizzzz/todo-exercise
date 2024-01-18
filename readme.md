# Todo List Application

Welcome to the Todo List Application! This is a simple tool to help you keep track of your tasks and to-dos. In this version (main branch), we do not have the "completed" field, making it easy to create and manage your tasks.

## Get Started with Todoexercise Blockchain

If you're interested in blockchain development, you can also explore our Todoexercise blockchain, which is built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

### Get started

To set up and run the Todoexercise blockchain, follow these steps:

1. Install Ignite CLI:
   ```
   npm install -g @ignetwork/ignite-cli
   ```
2. Start the Blockchain
   ```
   ignite chain serve
   ```
  `serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

## How to Use

To manage your tasks using the command line interface (CLI), follow these commands:

1. **Add a Task:**
   - To add a task, use the following command:
     ```
     todo-exercised tx todo create-task "[title]" "[description]" --from [user] --chain-id [chain-id]
     ```
   - Replace `[title]`, `[description]`, `[user]`, `[chain-id]` with the relevant information.

2. **View Tasks:**
   - to view ALL tasks:
     ```
     todo-exercised query todo list-task --chain-id [chain-id]
     ```
   - to view a specific task:
     ```
     todo-exercised query todo show-task [task-id] --chain-id [chain-id]
     ```

3. **Edit a Task:**
   - To edit a task, use the following command (Current Version):
     ```
     todo-exercised tx todo update-task [task-id] "[new-title]" "[new-description]" --from [user] --chain-id [chain-id]
     ```
   - To edit a task, use the following command (New Version *Coming Soon*):
     ```
     todo-exercised tx todo update-task [task-id] "[new-title]" "[new-description]" [completed-status] --from [user] --chain-id [chain-id]
     ```

4. **Delete a Task:**
   - To delete a task, use the following command:
     ```
     todo-exercised tx todo delete-task [task-id] --from [user] --chain-id [chain-id]
     ```

These CLI commands allow you to add, edit, and delete tasks in your Todo List Application. Enjoy using the CLI interface to manage your tasks efficiently.


### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project. 


For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/todo-exercise@latest! | sudo bash
```
`username/todo-exercise` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)


Enjoy using the Todo List Application, and if you're interested in blockchain development, don't hesitate to explore our todo blockchain project for a deeper dive into the world of decentralized applications!
