/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50728
 Source Host           : localhost:3306
 Source Schema         : logadmin

 Target Server Type    : MySQL
 Target Server Version : 50728
 File Encoding         : 65001

 Date: 06/03/2020 23:32:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tbl_app_info
-- ----------------------------
DROP TABLE IF EXISTS `tbl_app_info`;
CREATE TABLE `tbl_app_info` (
  `app_id` int(11) NOT NULL AUTO_INCREMENT,
  `app_name` varchar(1024) NOT NULL,
  `app_type` varchar(64) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `develop_path` varchar(256) NOT NULL,
  PRIMARY KEY (`app_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for tbl_app_ip
-- ----------------------------
DROP TABLE IF EXISTS `tbl_app_ip`;
CREATE TABLE `tbl_app_ip` (
  `app_id` int(11) DEFAULT NULL,
  `ip` varchar(64) DEFAULT NULL,
  KEY `app_id_ip_index` (`app_id`,`ip`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
