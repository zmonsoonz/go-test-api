CREATE TABLE IF NOT EXISTS users (
    user_id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS libraries (
    library_id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    total_tracks INTEGER NOT NULL,
    
    CONSTRAINT fk_libraries_user
        FOREIGN KEY (user_id)
        REFERENCES users(user_id)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tracks (
    track_id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    library_id INTEGER NOT NULL,
    file_path VARCHAR(255) NOT NULL,
    duration INTEGER NOT NULL,
    uploaded_at TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_tracks_library
        FOREIGN KEY (library_id)
        REFERENCES libraries(library_id)
        ON DELETE CASCADE
);