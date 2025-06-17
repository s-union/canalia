-- Create enum types
CREATE TYPE campus_type AS ENUM ('野田', '葛飾', '神楽坂', '３キャンパス合同');
CREATE TYPE category_type AS ENUM ('運動系', '演技系', '音楽系', '文化系');
CREATE TYPE place_type AS ENUM ('屋外', '屋内', '特殊');

-- Users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    contact_email VARCHAR(255),
    university_id VARCHAR(50) UNIQUE,
    is_verified BOOLEAN NOT NULL DEFAULT FALSE,
    phone_number VARCHAR(20),
    family_name VARCHAR(100) NOT NULL,
    given_name VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Circles table
CREATE TABLE circles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    description TEXT,
    contact_email VARCHAR(255),
    website_url VARCHAR(500),
    twitter_url VARCHAR(500),
    instagram_url VARCHAR(500),
    youtube_url VARCHAR(500),
    campus campus_type NOT NULL,
    category category_type NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Leaders table
CREATE TABLE leaders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    circle_id INTEGER NOT NULL REFERENCES circles(id) ON DELETE CASCADE,
    priority INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    UNIQUE(user_id, circle_id)
);

-- Permissions table
CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT
);

-- User permissions table
CREATE TABLE user_permissions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    permission_id INTEGER NOT NULL REFERENCES permissions(id) ON DELETE CASCADE,
    UNIQUE(user_id, permission_id)
);

-- Forms table
CREATE TABLE forms (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    summary TEXT,
    is_published BOOLEAN NOT NULL DEFAULT FALSE,
    start_time TIMESTAMP WITH TIME ZONE NOT NULL,
    end_time TIMESTAMP WITH TIME ZONE NOT NULL,
    committee_start_time TIMESTAMP WITH TIME ZONE NOT NULL,
    committee_end_time TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Event tags table (referenced by forms)
CREATE TABLE event_tags (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT
);

-- Form tags table
CREATE TABLE form_tags (
    id SERIAL PRIMARY KEY,
    form_id INTEGER NOT NULL REFERENCES forms(id) ON DELETE CASCADE,
    event_tag_id INTEGER NOT NULL REFERENCES event_tags(id) ON DELETE CASCADE,
    UNIQUE(form_id, event_tag_id)
);

-- Questions table
CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    form_id INTEGER NOT NULL REFERENCES forms(id) ON DELETE CASCADE,
    "order" INTEGER NOT NULL,
    title VARCHAR(200) NOT NULL,
    description TEXT,
    type VARCHAR(50) NOT NULL,
    is_required BOOLEAN NOT NULL DEFAULT FALSE,
    validation_rules JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    UNIQUE(form_id, "order")
);

-- Question options table
CREATE TABLE question_options (
    id SERIAL PRIMARY KEY,
    question_id INTEGER NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
    "order" INTEGER NOT NULL,
    value VARCHAR(500) NOT NULL,
    label VARCHAR(500) NOT NULL,
    UNIQUE(question_id, "order")
);

-- Answered forms table
CREATE TABLE answered_forms (
    id SERIAL PRIMARY KEY,
    form_id INTEGER NOT NULL REFERENCES forms(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    circle_id INTEGER NOT NULL REFERENCES circles(id) ON DELETE CASCADE,
    status VARCHAR(50) NOT NULL DEFAULT 'draft',
    submitted_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    UNIQUE(form_id, user_id, circle_id)
);

-- Answers table
CREATE TABLE answers (
    id SERIAL PRIMARY KEY,
    answered_form_id INTEGER NOT NULL REFERENCES answered_forms(id) ON DELETE CASCADE,
    question_id INTEGER NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
    answer JSONB NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    UNIQUE(answered_form_id, question_id)
);

-- Text news table
CREATE TABLE text_news (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    summary TEXT,
    contents TEXT NOT NULL,
    is_published BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- File news table
CREATE TABLE file_news (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    summary TEXT,
    file_path VARCHAR(500) NOT NULL,
    is_published BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Event genres table
CREATE TABLE event_genres (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    start_application TIMESTAMP WITH TIME ZONE NOT NULL,
    end_application TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Places table
CREATE TABLE places (
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    type place_type NOT NULL,
    description TEXT
);

-- Events table
CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    event_genre_id INTEGER NOT NULL REFERENCES event_genres(id) ON DELETE CASCADE,
    circle_id INTEGER NOT NULL REFERENCES circles(id) ON DELETE CASCADE,
    place_id INTEGER NOT NULL REFERENCES places(id) ON DELETE CASCADE,
    name VARCHAR(200) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Event tag relations table
CREATE TABLE event_tag_relations (
    id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    event_tag_id INTEGER NOT NULL REFERENCES event_tags(id) ON DELETE CASCADE,
    UNIQUE(event_id, event_tag_id)
);

-- Contact emails table
CREATE TABLE contact_emails (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- QA answers table
CREATE TABLE qa_answers (
    id SERIAL PRIMARY KEY,
    contact_email_id INTEGER NOT NULL REFERENCES contact_emails(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    contents TEXT NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Create indexes for better performance
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_university_id ON users(university_id);
CREATE INDEX idx_leaders_user_id ON leaders(user_id);
CREATE INDEX idx_leaders_circle_id ON leaders(circle_id);
CREATE INDEX idx_user_permissions_user_id ON user_permissions(user_id);
CREATE INDEX idx_questions_form_id ON questions(form_id);
CREATE INDEX idx_question_options_question_id ON question_options(question_id);
CREATE INDEX idx_answered_forms_form_id ON answered_forms(form_id);
CREATE INDEX idx_answered_forms_user_id ON answered_forms(user_id);
CREATE INDEX idx_answered_forms_circle_id ON answered_forms(circle_id);
CREATE INDEX idx_answers_answered_form_id ON answers(answered_form_id);
CREATE INDEX idx_answers_question_id ON answers(question_id);
CREATE INDEX idx_events_circle_id ON events(circle_id);
CREATE INDEX idx_events_event_genre_id ON events(event_genre_id);
CREATE INDEX idx_events_place_id ON events(place_id);
CREATE INDEX idx_event_tag_relations_event_id ON event_tag_relations(event_id);
CREATE INDEX idx_qa_answers_user_id ON qa_answers(user_id);
CREATE INDEX idx_qa_answers_contact_email_id ON qa_answers(contact_email_id);