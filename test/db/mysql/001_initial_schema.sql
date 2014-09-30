CREATE TABLE IF NOT EXISTS call_records 
(
    imsi VARCHAR(15),
    timestamp TIMESTAMP,
    duration INT NOT NULL,
    region VARCHAR(20) NOT NULL,
    calling_number VARCHAR(15) NOT NULL,
    called_number VARCHAR(15) NOT NULL,
    PRIMARY KEY(imsi, timestamp)
);