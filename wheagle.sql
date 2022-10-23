/*
CREATE TABLE IF NOT EXISTS  `tablename`(
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

*/

CREATE TABLE IF NOT EXISTS  `apikey`(
  `apikey`  varchar(255) NOT NULL,
  `ownerid` varchar(255) NOT NULL,
  `active`  BOOLEAN,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE IF NOT EXISTS  `mothership`(
  `ownerid`  varchar(255) NOT NULL,
  `name`   varchar(255) NOT NULL,
  `mothershipid`   varchar(255) NOT NULL,
  `address` INT UNSIGNED,
  `port` INT(255) NOT NULL,
  `description`  varchar(255) NOT NULL,
  `active` BOOLEAN,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE IF NOT EXISTS  `minion`(
  `minionid`  varchar(255) NOT NULL,
  `name`  varchar(255) NOT NULL,
  `uname` varchar(255) NOT NULL,
  `userid`  varchar(255) NOT NULL,
  `groupid` varchar(255) NOT NULL,
  `homedir` varchar(255) NOT NULL,
  `miniontype`  varchar(255) NOT NULL,
  `ostype`  varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `installed` BOOLEAN,
  `mothershipid`  varchar(255) NOT NULL,
  `minionip`  INT UNSIGNED,
  `ownerid` varchar(255) NOT NULL,
  `lastseen`  varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS  `user`(
  `userid` varchar(255) NOT NULL,
  `ownerid` varchar(255) NOT NULL,
  `username`  varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password`  varchar(255) NOT NULL,
  `active`  varchar(255) NOT NULL,
  `anonymous` BOOLEAN,
  `verified` BOOLEAN,
  `admin` BOOLEAN,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS  `temp`(
  `userid`  VARCHAR(255) NOT NULL,
  `sessionid` varchar(255) NOT NULL,
  `ip_address`  INT UNSIGNED,
  `active`  BOOLEAN,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;


//* EDR  */

CREATE TABLE IF NOT EXISTS  `events`(
  `description` TEXT NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
