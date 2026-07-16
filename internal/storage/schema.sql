PRAGMA user_version = 2;

CREATE TABLE translations (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    abbreviation TEXT NOT NULL,
    language_tag TEXT NOT NULL,
    language_name TEXT NOT NULL,
    edition TEXT NOT NULL,
    canon TEXT NOT NULL,
    text_edition TEXT NOT NULL,
    source_publisher TEXT NOT NULL,
    source_homepage TEXT NOT NULL,
    source_archive_url TEXT NOT NULL,
    source_archive_sha256 TEXT NOT NULL CHECK (length(source_archive_sha256) = 64),
    source_retrieved_at TEXT NOT NULL,
    rights_status TEXT NOT NULL,
    rights_notice_url TEXT NOT NULL,
    trademark_notice TEXT NOT NULL,
    text_policy TEXT NOT NULL
) STRICT;

CREATE TABLE books (
    translation_id TEXT NOT NULL,
    id TEXT NOT NULL,
    source_code TEXT NOT NULL,
    position INTEGER NOT NULL CHECK (position > 0),
    name TEXT NOT NULL,
    PRIMARY KEY (translation_id, id),
    UNIQUE (translation_id, source_code),
    UNIQUE (translation_id, position),
    FOREIGN KEY (translation_id) REFERENCES translations (id) ON DELETE CASCADE
) STRICT;

CREATE TABLE verses (
    id INTEGER PRIMARY KEY,
    translation_id TEXT NOT NULL,
    book_id TEXT NOT NULL,
    chapter INTEGER NOT NULL CHECK (chapter > 0),
    verse INTEGER NOT NULL CHECK (verse > 0),
    text TEXT NOT NULL,
    UNIQUE (translation_id, book_id, chapter, verse),
    FOREIGN KEY (translation_id, book_id)
        REFERENCES books (translation_id, id) ON DELETE CASCADE
) STRICT;

CREATE VIRTUAL TABLE verses_fts USING fts5(
    text,
    content = 'verses',
    content_rowid = 'id',
    tokenize = 'unicode61 remove_diacritics 2'
);
