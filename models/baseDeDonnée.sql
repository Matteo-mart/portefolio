-- mariadb -u matteo -p
-- matteo

CREATE DATABASE IF NOT EXISTS portefolio;
USE portefolio;

--contacts
CREATE TABLE IF NOT EXISTS contacts (     
    id INT AUTO_INCREMENT PRIMARY KEY,     
    telephone VARCHAR(20),     
    email VARCHAR(100) NOT NULL,     
    linkedin VARCHAR(255),    
    github VARCHAR(255),     
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP 
) ENGINE=InnoDB;

-- corbeille
CREATE TABLE IF NOT EXISTS corbeille (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    project_id BIGINT UNSIGNED NOT NULL,
    titre VARCHAR(255),
    date_suppression TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    date_creation DATE,
    description TEXT,
    technologie VARCHAR(255),
    explication TEXT,
    probleme TEXT,
    solution TEXT,
    url_source VARCHAR(255)
) ENGINE=InnoDB;

-- corbeille_image
CREATE TABLE IF NOT EXISTS corbeille_image (
    id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    project_id bigint(20) unsigned NOT NULL,
    url varchar(255) NOT NULL,
    mime_type varchar(100),
    PRIMARY KEY (id)
);

-- project
CREATE TABLE IF NOT EXISTS project (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY, 
    titre VARCHAR(255) NOT NULL,
    date_creation DATE,
    description TEXT,
    technologie VARCHAR(255),
    explication TEXT,
    probleme TEXT,
    solution TEXT,
    url_source VARCHAR(255)
) ENGINE=InnoDB;

-- project_image
CREATE TABLE IF NOT EXISTS project_image (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    project_id BIGINT UNSIGNED NOT NULL, 
    url VARCHAR(255) NOT NULL,
    CONSTRAINT fk_project 
        FOREIGN KEY (project_id) 
        REFERENCES project(id) 
        ON DELETE CASCADE
) ENGINE=InnoDB;

-- Table des technologies 
CREATE TABLE IF NOT EXISTS corbeille_technologies (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    tech_id BIGINT UNSIGNED NOT NULL,
    nom VARCHAR(100) NOT NULL,
    icone VARCHAR(255),
    url_source VARCHAR(255),
    date_suppression DATETIME DEFAULT NOW()
) ENGINE=InnoDB;
