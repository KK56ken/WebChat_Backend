CREATE DATABASE IF NOT EXISTS chatDB;
use chatDB;
CREATE TABLE IF NOT EXISTS chatDB.users(
  `userid` int(11) NOT NULL AUTO_INCREMENT UNIQUE,
  `email` varchar(255) NOT NULL PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `token` varchar(255) NOT NULL,
  `avatorurl` varchar(255)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS chatDB.messages(
  `messageid` int(11) NOT NULL AUTO_INCREMENT UNIQUE,
  `sendUserId` int(11) NOT NULL,
  `receiveUserId` int(11) NOT NULL,
  `message` varchar(255) NOT NULL,
  `time` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS chatDB.friends(
  `id` int(11) NOT NULL AUTO_INCREMENT UNIQUE,
  `userId` int(11) NOT NULL,
  `friendId` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;