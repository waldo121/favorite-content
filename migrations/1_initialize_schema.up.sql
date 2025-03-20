PRAGMA foreign_keys = ON;
CREATE TABLE IF NOT EXISTS Items (
    id INTEGER PRIMARY KEY AUTOINCREMENT, 
    video_id TEXT NOT NULL UNIQUE,
    platform TEXT NOT NULL
);
CREATE INDEX idx_items ON Items (id);