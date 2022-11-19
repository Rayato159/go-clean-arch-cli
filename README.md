<h1>📦 Golang Clean Architecture CLI</h1>
<h3>🧑🏼‍💻 Established by</h3>
<ul>
    <li><strong>Ruangyot Nanchiang</strong></li>
</ul>

<h3>🚀 Version</h3>
<ul>
    <li><strong>1.0.0</strong></li>
</ul>

<h3>📃 Introduction</h3>
<ul>
    <li>This project is to create a file and a folder for the clean architecture of Uncle Bob in automatic (CLI) with Golang </li>
</ul>

<h3>Quick Start</h3>
<ul>
<li>

<strong>Clone the project</strong>

```
git clone https://github.com/Rayato159/go-clean-arch-cli
```

</li>

<li>

<strong>Build main.go</strong>

```
cd ./go-clean-arch-cli/src
go build main.go
```

</li>

<li>

<strong>Usage</strong>

You need to move a `main.exe` to the destination that you want to establish the project before you execute it

```
main [command] [project-name|module-name]
```

</li>

</li>

<li>

<strong>Command List</strong>

| Command | Short Term | Prefix | Description                 |
| ------- | ---------- | ------ | --------------------------- |
| init    | i          | none   | To initialize a new project |
| module  | m          | -      | To generate a new module    |

</li>

<li>

<strong>Example</strong>

Initialize a new project

```
main i my-project
```

Generate a new module **(Need to init a project first)**

```
main -m users
```

</li>
</ul>

<h3>🔩 Structure of Creation</h3>
<ul>

```zsh
📂 app/
├─ 📄 main.go
📂 assets/
├─ 📂 logs/
├─ 📂 screenshots/
📂 configs/
├─ 📄 configs.go
📂 internals/
├─ 📂 servers/
│  ├─ 📄 server.go
├─ 📂 entities/
│  ├─ 📄 books.go
├─ 📂 books/
│  ├─ 📂 controllers/
│  │  ├─ 📄 book_controller.go
│  ├─ 📂 usecases/
│  │  ├─ 📄 book_usecase.go
│  ├─ 📂 repositories/
│  │  ├─ 📄 book_repository.go
📂 package/
├─ 📂 databases/
├─ 📂 middlewares/
├─ 📂 utils/
📂 tests/
📄 .env.dev
📄 .env.prod
📄 .env.test
```

</ul>
