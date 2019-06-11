CREATE USER 'Fred'@'localhost'
  IDENTIFIED BY 'Flintstone';
GRANT ALL
  ON loa.*
  TO 'Fred'@'localhost'
  WITH GRANT OPTION;
