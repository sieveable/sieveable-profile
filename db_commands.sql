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
