-- COmando en CMD: psql -U tu_usuario -d postgres -f Tablas.sql

-- Crear la base de datos
CREATE DATABASE BasedatosChat;

-- Conectarse a la base (solo funciona desde psql CLI)
\c BasedatosChat;


--Copiar desde aqui si es por medio de Query en pgAdmin
-- Crear tablas en orden

CREATE TABLE Usuarios (
    ID SERIAL PRIMARY KEY,
    Nombre VARCHAR(100),
    Email VARCHAR(150) UNIQUE,
    Contrasena VARCHAR(50),
    Rol INT
);

CREATE TABLE Chats (
    ID SERIAL PRIMARY KEY
);

CREATE TABLE Grupos (
    ID SERIAL PRIMARY KEY,
    Nombre VARCHAR(30),
    ID_Chat INT UNIQUE,
    FOREIGN KEY (ID_Chat) REFERENCES Chats(ID)
);

CREATE TABLE Mensajes (
    ID SERIAL PRIMARY KEY,
    ID_Usuario INT,
    ID_Chat INT,
    Contenido TEXT NOT NULL,
    FechaEnvio TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (ID_Usuario) REFERENCES Usuarios(ID),
    FOREIGN KEY (ID_Chat) REFERENCES Chats(ID)
);

CREATE TABLE ChatUsuarios (
    ID_Usuario INT,
    ID_Chat INT,
    PRIMARY KEY (ID_Usuario, ID_Chat),
    FOREIGN KEY (ID_Usuario) REFERENCES Usuarios(ID),
    FOREIGN KEY (ID_Chat) REFERENCES Chats(ID)
);

CREATE TABLE GrupoMiembros (
    ID_Usuario INT,
    ID_Grupo INT,
    Admin BOOLEAN,
    PRIMARY KEY (ID_Usuario, ID_Grupo),
    FOREIGN KEY (ID_Usuario) REFERENCES Usuarios(ID),
    FOREIGN KEY (ID_Grupo) REFERENCES Grupos(ID)
);

CREATE TABLE Contactos (
    ID_Usuario INT,
    ID_Contacto INT,
    PRIMARY KEY (ID_Usuario, ID_Contacto),
    FOREIGN KEY (ID_Usuario) REFERENCES Usuarios(ID),
    FOREIGN KEY (ID_Contacto) REFERENCES Usuarios(ID)
);

-- Poblar tabla Usuario
INSERT INTO Usuarios (Nombre, Email, Contrasena, Rol) VALUES
('Omar Sáez', 'omar@example.com', 'clave123', 1),
('Lucía Torres', 'lucia@example.com', 'segura456', 2),
('Diego Ramírez', 'diego@example.com', 'admin789', 3);
