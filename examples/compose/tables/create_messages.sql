CREATE TABLE IF NOT EXISTS messages (
  page BIGINT(20) UNSIGNED,
  time_created_ns BIGINT(20) UNSIGNED,
  message VARCHAR(10000),
  PRIMARY KEY (page, time_created_ns)
) ENGINE=InnoDB;
