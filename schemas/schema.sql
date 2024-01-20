CREATE TABLE List (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at DATE NOT NULL
);

CREATE TABLE Task (
    id SERIAL PRIMARY KEY,
    parent INTEGER REFERENCES List(id) on DELETE CASCADE,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    completed BOOLEAN NOT NULL,
    created_at DATE NOT NULL
)

CREATE TABLE Subtask (
    id SERIAL PRIMARY KEY,
    parent INTEGER REFERENCES Task(id) on DELETE CASCADE,
    name TEXT NOT NULL,
    completed BOOLEAN NOT NULL,
    created_at DATE NOT NULL
)
