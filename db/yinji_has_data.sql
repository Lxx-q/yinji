/*
Navicat MySQL Data Transfer

Source Server         : king
Source Server Version : 60011
Source Host           : 127.0.0.1:3306
Source Database       : yinji

Target Server Type    : MYSQL
Target Server Version : 60011
File Encoding         : 65001

Date: 2019-12-08 23:30:19
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
-- Records of audio
-- ----------------------------
INSERT INTO `audio` VALUES ('2', 'South Music', 'm_2', '2019-11-04 20:59:53', '2019-11-05 20:59:56', 'resources/audio/2.mp3', '2', 'resources/image/audio/2.png', '150', null);
INSERT INTO `audio` VALUES ('3', '来不及勇敢', 'm_3', '2019-11-04 21:02:20', '2019-11-04 21:02:23', 'resources/audio/3.mp3', '3', 'resources/image/audio/3.jpg', '125', null);
INSERT INTO `audio` VALUES ('4', '大鱼', 'm_4', '2019-11-04 21:02:48', '2019-11-05 21:02:52', 'resources/audio/4.mp3', '3', 'resources/image/audio/4.png', '240', null);
INSERT INTO `audio` VALUES ('1574951701701', '大闹天宫', '1574951701701', '2019-11-28 22:35:01', '2019-11-28 22:35:01', 'resources/audio/1574951701701.mp3', '2', 'resources/image/audio/1574951701701.jpg', '12', null);
INSERT INTO `audio` VALUES ('1574951746746', '', '1574951746746', '2019-11-28 22:35:46', '2019-12-01 22:26:41', 'resources/audio/1574951746746.mp3', '2', 'resources/image/audio/1574951746746.jpg', '30', null);
INSERT INTO `audio` VALUES ('1574952541541', '飘移，止不住的飘移~', '1574952541541', '2019-11-28 22:49:01', '2019-12-01 20:49:13', 'resources/audio/1574952541541.mp3', '2', 'resources/image/audio/1574952541541.png', '30', null);
INSERT INTO `audio` VALUES ('1574954699699', '回梦游仙', '1574954699699', '2019-11-28 23:24:59', '2019-12-05 21:52:27', 'resources/audio/1574954699699.mp3', '2', 'resources/image/audio/1574954699699.JPG', '30', null);
INSERT INTO `audio` VALUES ('1575114448448', '兰贵人', '1575114448448', '2019-11-30 19:47:28', '2019-12-01 20:50:25', 'resources/audio/1575114448448.mp3', '2', 'resources/image/audio/1575114448448.png', '30', null);
INSERT INTO `audio` VALUES ('1575128897897', '落入凡尘', '1575128897897', '2019-11-30 23:48:17', '2019-11-30 23:48:17', 'resources/audio/1575128897897.mp3', '2', 'resources/image/audio/1575128897897.jpg', '12', null);
INSERT INTO `audio` VALUES ('1575129126126', 'London works', '1575129126126', '2019-11-30 23:52:06', '2019-12-05 21:35:12', 'resources/audio/1575129126126.mp3', '2', 'resources/image/audio/1575129126126.jpg', '30', null);
INSERT INTO `audio` VALUES ('1575172401401', '鱼养肥了，该下锅了', '1575172401401', '2019-12-01 11:53:21', '2019-12-01 20:28:51', 'resources/audio/1575172401401.mp3', '2', 'resources/image/audio/1575172401401.jpg', '30', null);
INSERT INTO `audio` VALUES ('1575193987987', '来不及勇敢啊，我去', '1575193987987', '2019-12-01 17:53:07', '2019-12-05 21:48:49', 'resources/audio/1575193987987.mp3', '2', 'resources/image/audio/1575193987987.jpg', '30', null);

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
-- Records of audio_history
-- ----------------------------

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
-- Records of file
-- ----------------------------
INSERT INTO `file` VALUES ('1', '1', '2019-12-04 22:47:48', '2019-12-04 22:47:51', 'resource/image/resource/167', '');

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
-- Records of image
-- ----------------------------
INSERT INTO `image` VALUES ('1', '6', 'wdw', 'dwd', '1', '1', null, null);

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
-- Records of login
-- ----------------------------
INSERT INTO `login` VALUES ('1', '123456');

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

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'Lx', 'u_1', '2019-11-04 20:39:47', '2019-11-04 20:39:51', '/none/resources/images/audio/1.png');
INSERT INTO `user` VALUES ('2', '周深', 'u_2', '2019-11-04 21:01:09', '2019-11-05 21:01:13', '/none/resources/images/audio/2.png');
INSERT INTO `user` VALUES ('3', '神舞幻想', 'u_3', '2019-11-04 21:01:56', '2019-11-14 21:01:59', '/none/resources/images/audio/3.jpg');
