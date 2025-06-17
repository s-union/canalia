-- Drop tables in reverse order of creation (to handle foreign key constraints)
DROP TABLE IF EXISTS qa_answers;
DROP TABLE IF EXISTS contact_emails;
DROP TABLE IF EXISTS event_tag_relations;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS places;
DROP TABLE IF EXISTS event_genres;
DROP TABLE IF EXISTS file_news;
DROP TABLE IF EXISTS text_news;
DROP TABLE IF EXISTS answers;
DROP TABLE IF EXISTS answered_forms;
DROP TABLE IF EXISTS question_options;
DROP TABLE IF EXISTS questions;
DROP TABLE IF EXISTS form_tags;
DROP TABLE IF EXISTS event_tags;
DROP TABLE IF EXISTS forms;
DROP TABLE IF EXISTS user_permissions;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS leaders;
DROP TABLE IF EXISTS circles;
DROP TABLE IF EXISTS users;

-- Drop enum types
DROP TYPE IF EXISTS place_type;
DROP TYPE IF EXISTS category_type;
DROP TYPE IF EXISTS campus_type;