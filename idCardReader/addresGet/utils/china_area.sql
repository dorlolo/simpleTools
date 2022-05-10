/*
 Navicat Premium Data Transfer

 Source Server         : 内网数据库90
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : 192.168.32.90:3306
 Source Schema         : dtsite_iot_test

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 26/03/2022 09:26:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for china_area
-- ----------------------------
DROP TABLE IF EXISTS `china_area`;
CREATE TABLE `china_area` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `pid` int(11) DEFAULT NULL COMMENT '父id',
  `short_name` varchar(100) DEFAULT NULL COMMENT '简称',
  `name` varchar(100) DEFAULT NULL COMMENT '名称',
  `merger_name` varchar(255) DEFAULT NULL COMMENT '全称',
  `level` tinyint(4) DEFAULT NULL COMMENT '层级 0 1 2 省市区县',
  `pinyin` varchar(100) DEFAULT NULL COMMENT '拼音',
  `phone_code` varchar(100) DEFAULT NULL COMMENT '长途区号',
  `zip_code` varchar(100) DEFAULT NULL COMMENT '邮编',
  `first` varchar(50) DEFAULT NULL COMMENT '首字母',
  `lng` varchar(100) DEFAULT NULL COMMENT '经度',
  `lat` varchar(100) DEFAULT NULL COMMENT '纬度',
  `area_code` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  KEY `ix_china_area_merge` (`merger_name`),
  KEY `ix_china_area_area_code` (`area_code`)
) ENGINE=InnoDB AUTO_INCREMENT=100000000 DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
