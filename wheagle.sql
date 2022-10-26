/*
CREATE TABLE IF NOT EXISTS  `tablename`(
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
*/

/* APT */

/* End of APT */
CREATE TABLE IF NOT EXISTS  `threatactivity`(
  `name` varchar(255) NOT NULL,
  `activity_id` varchar(255) NOT NULL,
  `creatorid` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `validated` BOOLEAN,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS  `apt`(
  `aptname` varchar(255) NOT NULL,
  `code` int(255) NOT NULL,
  `id` varchar(255) NOT NULL,
  `description` TEXT,
  `active` BOOLEAN,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS  `virus`(
  `aptid` varchar(255) NOT NULL,
  `virusid` varchar(255) NOT NULL,
  `hash` varchar(255) NOT NULL,
  `virustype` varchar(255) NOT NULL,
  `filetype` varchar(255) NOT NULL,
  `communicationmode` varchar(255) NOT NULL,
  `ostype` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

/* End of APT */

/*  DATA */

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
  `address` INT UNSIGNED NOT NULL,
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
  `description` TEXT NOT NULL,
  `installed` BOOLEAN,
  `mothershipid`  varchar(255) NOT NULL,
  `minionip`  INT UNSIGNED NOT NULL,
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
  `ip_address`  INT UNSIGNED NOT NULL,
  `active`  BOOLEAN,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

/* End of data */

/* EDR  */

CREATE TABLE IF NOT EXISTS  `agent`(
  `minionid`  VARCHAR(255) NOT NULL,
  `agentid`  VARCHAR(255) NOT NULL,
  `name`  VARCHAR(255) NOT NULL,
  `agent_ip`  INT UNSIGNED NOT NULL,
  `ownerid`  VARCHAR(255) NOT NULL,
  `description` TEXT NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS  `assetfiles`(
  `assetid` VARCHAR(255) NOT NULL,
  `filename` VARCHAR(255) NOT NULL,
  `fileid` VARCHAR(255) NOT NULL,
  `filehash` VARCHAR(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS  `assets`(
  `name` VARCHAR(255) NOT NULL,
  `asset_id` VARCHAR(255) NOT NULL,
  `description` TEXT NOT NULL,
  `active` BOOLEAN,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS  `events`(
  `eventid`  varchar(255) NOT NULL,
  `userid`  varchar(255) NOT NULL,
  `description` TEXT NOT NULL,
  `evenrType` int(255) NOT NULL,
  `handled` BOOLEAN,
  `eventdata` TEXT NOT NULL,
  `handleerid`  varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS  `threats`(
  `name`  varchar(255) NOT NULL,
  `threat_ip` INT UNSIGNED NOT NULL,
  `virusid`  varchar(255) NOT NULL,
  `details`  varchar(255) NOT NULL,
  `active` BOOLEAN,
  `description` TEXT NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

/* End of EDR */

/* Workers */

CREATE TABLE IF NOT EXISTS  `plugins`(
  `owner` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `hash` varchar(255) NOT NULL,
  `plugin_type` INT(255) NOT NULL,
  `description` TEXT NOT NULL,
  `active` BOOLEAN,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS  `task`(
  `ownerid` varchar(255) NOT NULL,
  `taskid` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `type` BOOLEAN,
  `description` TEXT NOT NULL,
  `handled` BOOLEAN,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

/* End Of Workers */
