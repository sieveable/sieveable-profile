DROP DATABASE IF EXISTS test_apps;
CREATE DATABASE test_apps;
USE test_apps;
CREATE TABLE IF NOT EXISTS app (id VARCHAR(200) NOT NULL PRIMARY KEY,
    package_name VARCHAR(120) NOT NULL,
    version_code INT NOT NULL,
    version_name VARCHAR(80) NOT NULL,
    downloads INT NOT NULL DEFAULT 0,
    ratings FLOAT NOT NULL DEFAULT 0.0,
    release_date DATE NOT NULL);
CREATE TABLE IF NOT EXISTS category (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(120) NOT NULL UNIQUE,
    type ENUM('ui','code','manifest','listing'),
    description VARCHAR(300) NULL);
CREATE TABLE IF NOT EXISTS feature (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(120) NOT NULL UNIQUE,
    description VARCHAR(300) NULL,
    sieveable_query VARCHAR(400) NOT NULL,
    category_id INT NOT NULL,
    INDEX cat_index (category_id),
    FOREIGN KEY (category_id)
    REFERENCES category(id)
    ON UPDATE CASCADE ON DELETE CASCADE);
CREATE TABLE IF NOT EXISTS app_feature (
    app_id VARCHAR(200) NOT NULL,
    feature_id INT NOT NULL,
    PRIMARY KEY(app_id, feature_id),
    UNIQUE INDEX(app_id, feature_id),
    FOREIGN KEY (app_id) REFERENCES app(id)
    ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (feature_id) REFERENCES feature(id)
    ON UPDATE CASCADE ON DELETE CASCADE);
# insert data
INSERT INTO category(name, type, description) VALUES
    ('material-design', 'ui', 'Apps that implement Material Design');
SET @cat_id = LAST_INSERT_ID();
INSERT INTO app(id, package_name, version_code, version_name, downloads, 
     ratings, release_date) VALUES('com.example.app-8', 'com.example.app',
     8, '1.2', 100, 4.2, '2016-01-16');
SET @app_1_id = LAST_INSERT_ID();
INSERT INTO app(id, package_name, version_code, version_name, downloads, 
     ratings, release_date) VALUES('com.example.app-9', 'com.example.app',
     9, '1.3', 100, 4.2, '2016-03-22');
SET @app_2_id = LAST_INSERT_ID();
INSERT INTO feature(name, description, sieveable_query, category_id) VALUES(
     'first_feature_name', 'feature_description', 'MATCH app...', @cat_id);
SET @feature_1_id = LAST_INSERT_ID();
INSERT INTO feature(name, description, sieveable_query, category_id) VALUES(
     'second_feature_name', 'feature_description', 'MATCH app...', @cat_id);
SET @feature_2_id = LAST_INSERT_ID();
INSERT INTO app_feature VALUES('com.example.app-8', @feature_1_id);
INSERT INTO app_feature VALUES('com.example.app-8', @feature_2_id);
INSERT INTO app_feature VALUES('com.example.app-9', @feature_2_id);
