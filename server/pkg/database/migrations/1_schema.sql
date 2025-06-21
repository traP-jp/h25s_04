-- +goose Up

CREATE TABLE eateries (
    id CHAR(36) NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    PRIMARY KEY (id)
);

CREATE TABLE reviews (
    id CHAR(36) NOT NULL,
    eatery_id CHAR(36) NOT NULL,
    user_id VARCHAR(32) NOT NULL,
    content TEXT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (eatery_id) REFERENCES eateries(id)
);

CREATE TABLE images (
    id CHAR(36) NOT NULL,
    review_id CHAR(36) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (review_id) REFERENCES reviews(id)
);

-- +goose Down

DROP TABLE images;
DROP TABLE reviews;
DROP TABLE eateries;

