/*
Navicat MySQL Data Transfer

Source Server         : king
Source Server Version : 60011
Source Host           : 127.0.0.1:3306
Source Database       : yinji

Target Server Type    : MYSQL
Target Server Version : 60011
File Encoding         : 65001

Date: 2019-12-08 23:30:00
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for audio
-- ----------------------------
DROP TABLE IF EXISTS `audio`;
CREATE TABLE `audio` (
  `id` bigint(16) NOT NULL,
  `name` varchar(16) NOT NULL,
  `code` varchar(20) NOT NULL,
  `create_time` datetime NOT NULL,
  `modify_time` datetime NOT NULL,
  `url` varchar(50) NOT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  `image` varchar(50) DEFAULT NULL,
  `time_length` int(1) DEFAULT NULL,
  `descript` text COMMENT '将其暂时设定为text模型，因为无法确定到底字符多少',
  PRIMARY KEY (`id`),
  KEY `fk_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for audio_history
-- ----------------------------
DROP TABLE IF EXISTS `audio_history`;
CREATE TABLE `audio_history` (
  `id` bigint(20) NOT NULL COMMENT '对应的id',
  `user_id` bigint(20) NOT NULL COMMENT '目标用户的id',
  `audio_id` bigint(20) NOT NULL COMMENT '目标音频的id',
  `browse_all_count` bigint(15) unsigned zerofill NOT NULL COMMENT '该用户收看的次数',
  `create_time` datetime NOT NULL COMMENT '创造时间 ， 现在，我们可以理解为 第一次收看的时间',
  `modify_time` datetime NOT NULL COMMENT '我们可以把这里理解为第一次修改的时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Table structure for file
-- ----------------------------
DROP TABLE IF EXISTS `file`;
CREATE TABLE `file` (
  `id` bigint(20) NOT NULL COMMENT 'file 文件 主要代码的 作用，便是让程序，能找到将来分散在各地的数据资源',
  `code` char(20) NOT NULL,
  `create_time` datetime NOT NULL,
  `modify_time` datetime NOT NULL,
  `path` char(50) NOT NULL,
  `server` char(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Table structure for image
-- ----------------------------
DROP TABLE IF EXISTS `image`;
CREATE TABLE `image` (
  `id` bigint(16) NOT NULL COMMENT '目标的id',
  `targetId` bigint(20) NOT NULL,
  `name` varchar(30) DEFAULT NULL,
  `image_type` char(6) DEFAULT NULL,
  `origin_file` bigint(16) DEFAULT NULL,
  `compress_30_file` bigint(16) DEFAULT NULL,
  `compress_60_file` bigint(16) DEFAULT NULL,
  `compress_80_file` bigint(16) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_origin_file` (`origin_file`),
  KEY `fk_compress_30_file` (`compress_30_file`),
  CONSTRAINT `fk_compress_30_file` FOREIGN KEY (`compress_30_file`) REFERENCES `file` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `fk_origin_file` FOREIGN KEY (`origin_file`) REFERENCES `file` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Table structure for login
-- ----------------------------
DROP TABLE IF EXISTS `login`;
CREATE TABLE `login` (
  `id` bigint(16) NOT NULL,
  `password` varchar(16) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(16) NOT NULL,
  `name` varchar(16) NOT NULL,
  `code` varchar(20) NOT NULL,
  `create_time` datetime NOT NULL,
  `modify_time` datetime NOT NULL,
  `image` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
