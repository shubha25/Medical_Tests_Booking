-- dbname = medical_tests

CREATE TABLE users (
    id int NOT NULL AUTO_INCREMENT primary key,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20),
    address TEXT,
    role VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE tests (
    id int NOT NULL AUTO_INCREMENT primary key,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price float NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP    
);



CREATE TABLE bookings_tracker (
    id int NOT NULL AUTO_INCREMENT primary key,
    user_id int REFERENCES users(id) ,
    test_id int REFERENCES tests(id) ,
    appointment_date DATE NOT NULL,
    appointment_time TIME NOT NULL,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


CREATE TABLE admin(
    id int NOT NULL AUTO_INCREMENT primary key,
    Username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
)


insert into admin(Username,password) values ('Ram','password111' ),('Sham','password222' ),('Browny','password333' ), ('John','password444' ),