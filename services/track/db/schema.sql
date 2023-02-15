CREATE TABLE IF NOT EXISTS artists(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS tracks(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
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
    tag_id INTEGER REFERENCES tags(id),
    CONSTRAINT track_id_tag_id UNIQUE (track_id, tag_id)
);

CREATE TABLE IF NOT EXISTS albums(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    playcounts INTEGER,
    listeners INTEGER,
    artist_id INTEGER NOT NULL REFERENCES artists(id) ON DELETE CASCADE,
    CONSTRAINT album_artist_id UNIQUE (name, artist_id)
);

CREATE TABLE IF NOT EXISTS albums_tags(
    album_id INTEGER REFERENCES albums(id),
    tag_id INTEGER REFERENCES tags(id),
    CONSTRAINT album_id_tag_id UNIQUE (album_id, tag_id)
);

CREATE TABLE IF NOT EXISTS albums_tracks(
    album_id INTEGER REFERENCES albums(id),
    track_id INTEGER REFERENCES tracks(id),
    CONSTRAINT album_id_track_id UNIQUE (album_id, track_id)
);

DROP TABLE IF EXISTS albums_tracks;

DROP TABLE IF EXISTS albums_tags;

DROP TABLE IF EXISTS albums;

DROP TABLE IF EXISTS tracks_tags;

DROP TABLE IF EXISTS tags;

DROP TABLE IF EXISTS tracks;

DROP TABLE IF EXISTS artists;
