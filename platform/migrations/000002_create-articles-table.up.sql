

    
 -- Create Author table
CREATE TABLE articles (
        id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
        author_id UUID NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
        updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
        start_published_at TIMESTAMP NOT NULL,
        end_published_at TIMESTAMP NOT NULL,
        title VARCHAR (255) NOT NULL,
        body VARCHAR (255) NOT NULL,
        is_published BOOLEAN NOT NULL,    
        category VARCHAR (255) NOT NULL
);

    -- Add indexes
CREATE INDEX id_articles ON articles (id);

CREATE INDEX username_author_id ON articles (author_id);

