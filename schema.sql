drop table if exists consumeroffset;

CREATE TABLE consumeroffset (
    consumer_id serial PRIMARY KEY,
    consumer_offset int
);
