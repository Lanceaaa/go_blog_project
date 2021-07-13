/*
 Navicat MySQL Data Transfer

 Source Server         : 本地mysql
 Source Server Type    : MySQL
 Source Server Version : 50728
 Source Host           : localhost:3306
 Source Schema         : go

 Target Server Type    : MySQL
 Target Server Version : 50728
 File Encoding         : 65001

 Date: 12/07/2021 23:03:57
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_article_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_article_tag`;
CREATE TABLE `blog_article_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NOT NULL COMMENT 'id',
  `tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'id',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'cjsj',
  `created_by` varchar(100) DEFAULT '' COMMENT 'cjr',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'xgsj',
  `modified_by` varchar(100) DEFAULT '' COMMENT 'xgr',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'scsj',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '0:no 1:yes',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
