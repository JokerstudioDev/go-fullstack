/*
Navicat MySQL Data Transfer

Source Server         : sck
Source Server Version : 50505
Source Host           : localhost:3306
Source Database       : testsck

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2018-07-09 11:58:49
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for district
-- ----------------------------
DROP TABLE IF EXISTS `district`;
CREATE TABLE `district` (
  `district_id` int(11) NOT NULL,
  `district_name` varchar(45) DEFAULT NULL,
  `province_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`district_id`),
  CONSTRAINT `district_ibfk_1` FOREIGN KEY (`district_id`) REFERENCES `province` (`province_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of district
-- ----------------------------
INSERT INTO `district` VALUES ('1', 'เมือง', '1');

-- ----------------------------
-- Table structure for province
-- ----------------------------
DROP TABLE IF EXISTS `province`;
CREATE TABLE `province` (
  `province_id` int(11) NOT NULL,
  `province_name` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`province_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of province
-- ----------------------------
INSERT INTO `province` VALUES ('1', 'กทม');

-- ----------------------------
-- Table structure for soldier_address
-- ----------------------------
DROP TABLE IF EXISTS `soldier_address`;
CREATE TABLE `soldier_address` (
  `address_id` int(11) NOT NULL,
  `home_number` varchar(45) DEFAULT NULL,
  `moo` varchar(45) DEFAULT NULL,
  `road_soi` varchar(45) DEFAULT NULL,
  `province_id` int(11) DEFAULT NULL,
  `district_id` int(11) DEFAULT NULL,
  `subdistrict_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`address_id`),
  KEY `province_id` (`province_id`),
  KEY `district_id` (`district_id`),
  KEY `subdistrict_id` (`subdistrict_id`),
  CONSTRAINT `soldier_address_ibfk_1` FOREIGN KEY (`province_id`) REFERENCES `province` (`province_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `soldier_address_ibfk_2` FOREIGN KEY (`district_id`) REFERENCES `district` (`district_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `soldier_address_ibfk_3` FOREIGN KEY (`subdistrict_id`) REFERENCES `subdictrict` (`subdistrict_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of soldier_address
-- ----------------------------
INSERT INTO `soldier_address` VALUES ('1', '2', '5', 'ซอยจ้า', '1', '1', '1');

-- ----------------------------
-- Table structure for subdictrict
-- ----------------------------
DROP TABLE IF EXISTS `subdictrict`;
CREATE TABLE `subdictrict` (
  `subdistrict_id` int(11) NOT NULL,
  `subdistrict_name` varchar(45) DEFAULT NULL,
  `province_id` int(11) DEFAULT NULL,
  `district_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`subdistrict_id`),
  KEY `province_id` (`province_id`),
  KEY `district_id` (`district_id`),
  CONSTRAINT `subdictrict_ibfk_1` FOREIGN KEY (`province_id`) REFERENCES `province` (`province_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `subdictrict_ibfk_2` FOREIGN KEY (`district_id`) REFERENCES `district` (`district_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of subdictrict
-- ----------------------------
INSERT INTO `subdictrict` VALUES ('1', 'นครสวรรค์ตก', '1', '1');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `citizen_id` varchar(13) NOT NULL,
  `firstname` varchar(45) DEFAULT NULL,
  `lastname` varchar(45) DEFAULT NULL,
  `birthyear` int(4) DEFAULT NULL,
  `firstname_father` varchar(45) DEFAULT NULL,
  `lastname_father` varchar(45) DEFAULT NULL,
  `firstname_mother` varchar(45) DEFAULT NULL,
  `lastname_mother` varchar(45) DEFAULT NULL,
  `soldier_id` int(11) DEFAULT NULL,
  `address_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  KEY `address_id` (`address_id`),
  CONSTRAINT `user_ibfk_1` FOREIGN KEY (`address_id`) REFERENCES `soldier_address` (`address_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', '1600100386841', 'อวิรุทธิ์', 'เพ็ชรเม็ดเอียม', '1990', 'วีระ', 'เพ็ชรเม็ดเอี่ยม', 'บังเอิญ', 'เพ็ชรเม็ดเอี่ยม', '50', '1');
