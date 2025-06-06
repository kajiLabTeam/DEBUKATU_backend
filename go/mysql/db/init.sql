-- ユーザーデータ
CREATE TABLE User_Data (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    deleted BOOLEAN DEFAULT FALSE
);

-- モデル（目標）データ：ユーザーの目標体重・期間
CREATE TABLE Model_Data (
    model_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    model_weight FLOAT,
    length_of_month INT,
    created_date DATETIME,
    FOREIGN KEY (user_id) REFERENCES User_Data(user_id)
);

-- 現在の体重データ（Model_Data に紐づく）
CREATE TABLE Weight_Data (
    weight_id INT AUTO_INCREMENT PRIMARY KEY,
    model_id INT,
    current_weight FLOAT,
    created_date DATETIME,
    FOREIGN KEY (model_id) REFERENCES Model_Data(model_id)
);

-- カロリーデータ（Weight_Data に紐づく）
CREATE TABLE Calorie_Data (
    calorie_id INT AUTO_INCREMENT PRIMARY KEY,
    weight_id INT,
    must_calorie INT,
    current_calorie INT,
    created_date DATETIME,
    FOREIGN KEY (weight_id) REFERENCES Weight_Data(weight_id)
);
