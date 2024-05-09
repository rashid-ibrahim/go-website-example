DROP TABLE IF EXISTS work_history, job_detail_list, technologies;

CREATE TABLE work_history (
    id          INT AUTO_INCREMENT NOT NULL,
    place       VARCHAR(32),
    title       VARCHAR(16),
    start_date  VARCHAR(10),
    end_date    VARCHAR(10),
)

CREATE TABLE job_detail_list (
    id      INT AUTO_INCREMENT NOT NULL,
    entry   VARCHAR(32),

)

CREATE TABLE technologies (
    id          INT AUTO_INCREMENT NOT NULL,
    name        VARCHAR(16) NOT NULL,
    description VARCHAR(64),
)

CREATE TABLE job_detail_map (
    work_id INT NOT NULL,
    job_id  INT NOT NULL,
)

CREATE TABLE technologies_map (
    work_id         INT NOT NULL,
    technology_id   INT NOT NULL,
)