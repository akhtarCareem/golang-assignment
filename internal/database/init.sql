CREATE DATABASE careem;

\c careem;

CREATE TABLE users (
                       user_id SERIAL PRIMARY KEY,
                       name TEXT NOT NULL
);

CREATE TABLE rides (
                       ride_id SERIAL PRIMARY KEY,
                       source TEXT NOT NULL,
                       destination TEXT NOT NULL,
                       distance INTEGER NOT NULL,
                       cost INTEGER NOT NULL
);

CREATE TABLE bookings (
                          booking_id SERIAL PRIMARY KEY,
                          user_id INTEGER NOT NULL,
                          ride_id INTEGER NOT NULL,
                          time TIMESTAMP NOT NULL,
                          FOREIGN KEY(user_id) REFERENCES users(user_id),
                          FOREIGN KEY(ride_id) REFERENCES rides(ride_id)
);