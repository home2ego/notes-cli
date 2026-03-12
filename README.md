# Notes Tool

Notes Tool is a command-line program written in Go that allows a user to manage single-line notes stored in a text file.

## Collaborators

[Artem Zhyrnyi](https://gitea.kood.tech/artemzhyrnyi)  
[Ruslan Aleev](https://gitea.kood.tech/ruslanaleev)  
[Sonia Cienfuegos Torres](https://gitea.kood.tech/soniacienfuegostorres)

## Key Features

* Supports CRUD operations to create, read, update and delete files
* Handles invalid inputs and errors
* Creates a file based on the given name and optionally timestamp to store data
* Each file has its own password

## Project Structure

```text
./notes
├── constants
│   └── colors.go
├── go.mod
├── notestool.go
├── README.md
└── utils
    ├── crud.go
    ├── formatname.go
    └── handlesecurity.go
```

## How to run

```bash
// Clone repository
git clone https://gitea.kood.tech/artemzhyrnyi/notes.git

// Compile
go build -o notestool

// Run the program
./notestool <filename>
```

## Usage

```text
$ ./notestool testtag
No password found. Set a new password for this file:

> 1234

Welcome to the notes tool!

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.

> 1

Notes:
001 - note one [2026-03-03 12:03]
002 - note two [2026-03-03 12:04]

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.

> 2

Enter the note text:
note three

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.

> 1

Notes:
001 - note one [2026-03-03 12:03]
002 - note two [2026-03-03 12:04]
003 - note three [2026-03-03 12:05]

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.

> 3

Enter the number of note to remove or 0 to cancel:

> 3

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.

> 1

Notes:
001 - note one [2026-03-03 12:03]
002 - note two [2026-03-03 12:04]

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.

> 4

Goodbye!
```
