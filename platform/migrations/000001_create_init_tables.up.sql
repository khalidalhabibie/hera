BEGIN;

-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SET TIMEZONE="Asia/Jakarta";

    -- Create Author table
    CREATE TABLE authors (
        id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
        updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
        username VARCHAR (255) NOT NULL UNIQUE,
        password_hash VARCHAR (255) NOT NULL,
        active BOOLEAN NOT NULL
    );

    -- Add indexes
    CREATE INDEX id_authors ON authors (id);

    CREATE INDEX username_users ON authors (username);

    COMMIT