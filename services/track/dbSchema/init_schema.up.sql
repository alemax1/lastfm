CREATE TABLE IF NOT EXISTS artists(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS tracks(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    listeners INTEGER,
    playcounts INTEGER,
    artist_id INTEGER NOT NULL REFERENCES artists(id) ON DELETE CASCADE,
    CONSTRAINT name_artist_id UNIQUE (name, artist_id)
);

CREATE TABLE IF NOT EXISTS tags(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS tracks_tags(
    track_id INTEGER REFERENCES tracks(id),
    tag_id INTEGER REFERENCES tags(id)
);

CREATE TABLE IF NOT EXISTS albums(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    artist_id INTEGER NOT NULL REFERENCES artists(id) ON DELETE CASCADE,
    CONSTRAINT album_artist_id UNIQUE (name, artist_id)
);

CREATE TABLE IF NOT EXISTS albums_tracks(
    album_id INTEGER REFERENCES albums(id),
    track_id INTEGER REFERENCES tracks(id)
);