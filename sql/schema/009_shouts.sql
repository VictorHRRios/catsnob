-- +goose Up

CREATE TABLE shouts(
	id uuid PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	user_id uuid NOT NULL,
	review_id uuid NOT NULL, 
	title VARCHAR(50) NOT NULL, 
	shout_text TEXT NOT NULL CHECK (char_length(shout_text) <= 500),
	foreign key (user_id) references users(id) on delete cascade,
	foreign key (review_id) references album_reviews(id) on delete cascade,
	UNIQUE(user_id, review_id)
);

-- +goose Down
drop table shouts;

