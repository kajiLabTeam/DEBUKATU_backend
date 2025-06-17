CREATE TABLE User_Data (
    user_id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    password VARCHAR(30),
    age INT,
    height FLOAT,
    deleted BOOLEAN DEFAULT FALSE
)CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE Model_Data (
    model_id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    model_weight FLOAT,
    length_of_days INT,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES User_Data(user_id)
);

CREATE TABLE Weight_Data (
    weight_id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    model_id INT,
    user_id INT,
    current_weight FLOAT,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (model_id) REFERENCES Model_Data(model_id)
    FOREIGN KEY (user_id) REFERENCES User_Data(user_id)
);

CREATE TABLE Calorie_Data (
    calorie_id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    weight_id INT,
    model_id INT,
    must_calorie INT,
    current_calorie INT,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (weight_id) REFERENCES Weight_Data(weight_id)
    FOREIGN KEY (model_id) REFERENCES Model_Data(model_id)
);
