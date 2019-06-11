CREATE USER 'fred'@'localhost'
  IDENTIFIED BY 'flintstone';
GRANT ALL
  ON 'loa'.*
  TO 'fred'@'localhost'
  WITH GRANT OPTION;
