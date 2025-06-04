CREATE TABLE User_Data (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    created_date DATETIME,
    deleted_date DATETIME
);

CREATE TABLE Weight_Data (
    user_id INT,
    current_weight FLOAT,
    model_weight FLOAT,
    length_of_month INT,
    calorie INT,
    FOREIGN KEY (user_id) REFERENCES User_Data(user_id)
);

CREATE TABLE Calorie_Data (
    user_id INT,
    model_calorie INT,
    current_calorie INT,
    created_date DATETIME,
    FOREIGN KEY (user_id) REFERENCES User_Data(user_id)
);
