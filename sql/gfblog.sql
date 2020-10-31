/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50727
 Source Host           : 127.0.0.1:33060
 Source Schema         : gfblog

 Target Server Type    : MySQL
 Target Server Version : 50727
 File Encoding         : 65001

 Date: 31/10/2020 22:41:45
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for articles
-- ----------------------------
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL COMMENT '文章标题',
  `summary` varchar(255) DEFAULT NULL COMMENT '文章简介绍',
  `category_id` int(11) DEFAULT '0' COMMENT '文章栏目ID',
  `tags` varchar(255) DEFAULT NULL COMMENT '文章标签集合',
  `cover` varchar(255) DEFAULT NULL COMMENT '文章封面',
  `content` text COMMENT '文章内容',
  `md_content` text COMMENT '文章markdown内容',
  `from` tinyint(1) DEFAULT '0' COMMENT '文章来源:0-原创/1-转载/2-其他',
  `status` tinyint(1) DEFAULT '0' COMMENT '文章状态:1-发布/2-草稿/3-隐藏',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(50) DEFAULT NULL COMMENT '栏目名称',
  `sort` tinyint(1) DEFAULT '0' COMMENT '栏目排序',
  `status` int(1) DEFAULT '0' COMMENT '栏目状态:0-正常/1-隐藏',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
