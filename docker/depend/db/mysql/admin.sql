-- MySQL dump 10.13  Distrib 5.7.27, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: GoAdmin
-- ------------------------------------------------------
-- Server version	5.7.27-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `goadmin_menu`
--

DROP TABLE IF EXISTS `goadmin_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goadmin_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0',
  `type` tinyint(4) unsigned NOT NULL DEFAULT '0',
  `order` int(11) unsigned NOT NULL DEFAULT '0',
  `title` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `icon` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `uri` varchar(3000) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `header` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_menu`
--

LOCK TABLES `goadmin_menu` WRITE;
INSERT INTO `goadmin_menu` VALUES (1,0,1,14,'Admin','fa-tasks','',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_menu` VALUES (2,1,1,17,'Users','fa-users','/info/manager',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_menu` VALUES (3,1,1,15,'Roles','fa-user','/info/roles',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_menu` VALUES (4,1,1,14,'Permission','fa-ban','/info/permission',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_menu` VALUES (5,1,1,16,'Menu','fa-bars','/menu',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_menu` VALUES (6,1,1,18,'Operation log','fa-history','/info/op',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_menu` VALUES (7,0,1,1,'Dashboard','fa-bar-chart','/','','2019-09-10 00:00:00','2020-04-05 23:12:57');
INSERT INTO `goadmin_menu` VALUES (8,0,0,9,'Demo','fa-align-left','/info/demo','Demo','2020-03-26 08:23:20','2020-04-05 23:13:17');
INSERT INTO `goadmin_menu` VALUES (9,8,0,9,'grade','fa-bars','/info/demo_grade','','2020-03-26 08:24:03','2020-03-26 16:25:56');
INSERT INTO `goadmin_menu` VALUES (10,8,0,10,'class','fa-bars','/info/demo_class','','2020-03-26 08:28:45','2020-03-26 08:28:45');
INSERT INTO `goadmin_menu` VALUES (11,8,0,11,'student','fa-bars','/info/demo_student','','2020-03-26 08:31:03','2020-03-26 08:31:03');
INSERT INTO `goadmin_menu` VALUES (12,8,0,12,'student-class','fa-bars','/info/demo_student_class','','2020-03-26 08:32:58','2020-03-26 16:33:44');
INSERT INTO `goadmin_menu` VALUES (13,8,0,13,'student-score','fa-bars','/info/demo_student_score','','2020-03-26 08:33:29','2020-03-26 08:33:29');
INSERT INTO `goadmin_menu` VALUES (14,0,0,2,'配置','fa-align-left','/info/cpgame','配置','2020-03-27 13:42:27','2020-03-27 22:23:19');
INSERT INTO `goadmin_menu` VALUES (15,14,0,2,'渠道配置','fa-bars','/info/cpgame_cp?__columns=id,cp_name,cp_alias,created_at,updated_at','渠道配置','2020-03-27 13:42:49','2020-03-28 11:55:01');
INSERT INTO `goadmin_menu` VALUES (16,14,0,3,'游戏配置','fa-bars','/info/cpgame_game?__columns=id,game_name,platform,cp_id,game_sign_key,game_pay_key,created_at,updated_at','游戏配置','2020-03-27 13:43:18','2020-03-29 12:15:09');
INSERT INTO `goadmin_menu` VALUES (17,0,0,4,'用户','fa-align-left','/info/gameuser','用户','2020-03-27 14:13:52','2020-03-27 22:23:27');
INSERT INTO `goadmin_menu` VALUES (18,17,0,6,'用户信息','fa-bars','/info/gameuser_user','','2020-03-27 14:16:22','2020-03-27 14:16:22');
INSERT INTO `goadmin_menu` VALUES (19,17,0,4,'登录记录','fa-bars','/info/gameuser_record','','2020-03-27 14:16:48','2020-03-27 14:16:48');
INSERT INTO `goadmin_menu` VALUES (20,17,0,5,'登录活跃','fa-bars','/info/gameuser_token','','2020-03-29 09:32:18','2020-03-31 15:24:46');
INSERT INTO `goadmin_menu` VALUES (21,0,0,7,'订单','fa-align-left','/info/order','','2020-04-05 15:08:02','2020-04-05 15:08:02');
INSERT INTO `goadmin_menu` VALUES (22,21,0,7,'支付订单','fa-bars','/info/order_trade','','2020-04-05 15:08:48','2020-04-05 23:09:55');
INSERT INTO `goadmin_menu` VALUES (23,21,0,8,'支付记录','fa-bars','/info/order_record','','2020-04-05 15:09:15','2020-04-05 23:10:07');
UNLOCK TABLES;

--
-- Table structure for table `goadmin_permissions`
--

DROP TABLE IF EXISTS `goadmin_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goadmin_permissions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `http_method` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `http_path` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_permissions_name_unique` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_permissions`
--

LOCK TABLES `goadmin_permissions` WRITE;
INSERT INTO `goadmin_permissions` VALUES (1,'All permission','*','','*','2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_permissions` VALUES (2,'Dashboard','dashboard','GET,PUT,POST,DELETE','/','2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_permissions` VALUES (3,'demo_grade-*','demo_grade-*','GET,PUT,POST,PATCH,OPTIONS','/info/demo_grade\r\n/info/demo_grade/edit\r\n/edit/demo_grade\r\n/info/demo_grade/new\r\n/new/demo_grade\r\n/export/demo_grade\r\n/delete/demo_grade','2020-03-26 07:41:29','2020-03-31 14:45:13');
INSERT INTO `goadmin_permissions` VALUES (4,'demo_class-*','demo_class-*','','/info/demo_class\r\n/info/demo_class/edit\r\n/edit/demo_class\r\n/info/demo_class/new\r\n/new/demo_class\r\n/export/demo_class\r\n/delete/demo_class','2020-03-26 08:27:59','2020-03-31 14:45:59');
INSERT INTO `goadmin_permissions` VALUES (5,'demo_student-*','demo_student-*','','/info/demo_student\r\n/info/demo_student/edit\r\n/edit/demo_student\r\n/info/demo_student/new\r\n/new/demo_student\r\n/export/demo_student\r\n/delete/demo_student','2020-03-26 08:29:33','2020-03-31 14:46:22');
INSERT INTO `goadmin_permissions` VALUES (6,'demo_student_class-*','demo_student_class-*','','/info/demo_student_class\r\n/info/demo_student_class/edit\r\n/edit/demo_student_class\r\n/info/demo_student_class/new\r\n/new/demo_student_class\r\n/export/demo_student_class\r\n/delete/demo_student_class','2020-03-26 08:31:48','2020-03-31 14:46:35');
INSERT INTO `goadmin_permissions` VALUES (7,'demo_student_score-*','demo_student_score-*','','/info/demo_student_score\r\n/info/demo_student_score/edit\r\n/edit/demo_student_score\r\n/info/demo_student_score/new\r\n/new/demo_student_score\r\n/export/demo_student_score\r\n/delete/demo_student_score','2020-03-26 08:32:08','2020-03-31 14:46:58');
INSERT INTO `goadmin_permissions` VALUES (8,'cpgame_cp-operator','cpgame_cp-operator','','/info/cpgame_cp\r\n/info/cpgame_cp/edit\r\n/edit/cpgame_cp\r\n/info/cpgame_cp/new\r\n/new/cpgame_cp\r\n/export/cpgame_cp\r\n/info/cpgame_cp/detail','2020-03-27 13:39:25','2020-04-01 11:15:46');
INSERT INTO `goadmin_permissions` VALUES (9,'cpgame_game-operator','cpgame_game-operator','','/info/cpgame_game\r\n/info/cpgame_game/edit\r\n/edit/cpgame_game\r\n/info/cpgame_game/new\r\n/new/cpgame_game\r\n/export/cpgame_game\r\n/info/cpgame_game/detail','2020-03-27 13:40:06','2020-04-01 11:15:20');
INSERT INTO `goadmin_permissions` VALUES (10,'gameuser_user-operator','gameuser_user-operator','','/info/gameuser_user\r\n/info/gameuser_user/edit\r\n/edit/gameuser_user\r\n/info/gameuser_user/new\r\n/new/gameuser_user\r\n/export/gameuser_user\r\n/info/gameuser_user/detail','2020-03-27 14:12:02','2020-04-01 11:14:41');
INSERT INTO `goadmin_permissions` VALUES (11,'gameuser_record-operator','gameuser_record-operator','','/info/gameuser_record\r\n/info/gameuser_record/edit\r\n/edit/gameuser_record\r\n/info/gameuser_record/new\r\n/new/gameuser_record\r\n/export/gameuser_record\r\n/info/gameuser_record/detail','2020-03-27 14:12:45','2020-04-01 11:14:16');
INSERT INTO `goadmin_permissions` VALUES (12,'gameuser_token-operator','gameuser_token-operator','','/info/gameuser_token\r\n/info/gameuser_token/edit\r\n/edit/gameuser_token\r\n/info/gameuser_token/new\r\n/new/gameuser_token\r\n/export/gameuser_token\r\n/info/gameuser_token/detail','2020-03-29 09:31:42','2020-04-01 11:14:04');
INSERT INTO `goadmin_permissions` VALUES (13,'cpgame_cp-delete','cpgame_cp-delete','','/delete/cpgame_cp','2020-03-31 07:10:04','2020-03-31 15:10:04');
INSERT INTO `goadmin_permissions` VALUES (14,'cpgame_game-delete','cpgame_game-delete','','/delete/cpgame_game','2020-03-31 07:11:33','2020-03-31 15:11:33');
INSERT INTO `goadmin_permissions` VALUES (15,'gameuser_user-delete','gameuser_user-delete','','/delete/gameuser_user','2020-03-31 07:12:13','2020-03-31 15:12:13');
INSERT INTO `goadmin_permissions` VALUES (16,'gameuser_record-delete','gameuser_record-delete','','/delete/gameuser_record','2020-03-31 07:12:56','2020-03-31 15:12:56');
INSERT INTO `goadmin_permissions` VALUES (17,'gameuser_token-delete','gameuser_token-delete','','/delete/gameuser_token','2020-03-31 07:13:30','2020-03-31 15:13:30');
INSERT INTO `goadmin_permissions` VALUES (18,'order_trade-operator','order_trade-operator','','/info/order_trade\r\n/info/order_trade/edit\r\n/edit/order_trade\r\n/info/order_trade/new\r\n/new/order_trade\r\n/export/order_trade\r\n/info/order_trade/detail','2020-04-05 15:05:08','2020-04-05 23:05:08');
INSERT INTO `goadmin_permissions` VALUES (19,'order_trade-delete','order_trade-delete','','/delete/order_trade','2020-04-05 15:05:33','2020-04-05 23:05:33');
INSERT INTO `goadmin_permissions` VALUES (20,'order_record-operator','order_record-operator','','/info/order_record\r\n/info/order_record/edit\r\n/edit/order_record\r\n/info/order_record/new\r\n/new/order_record\r\n/export/order_record\r\n/info/order_record/detail','2020-04-05 15:06:04','2020-04-05 23:06:04');
INSERT INTO `goadmin_permissions` VALUES (21,'order_record-delete','order_record-delete','','/delete/order_trade','2020-04-05 15:06:41','2020-04-05 23:06:41');
UNLOCK TABLES;

--
-- Table structure for table `goadmin_role_menu`
--

DROP TABLE IF EXISTS `goadmin_role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goadmin_role_menu` (
  `role_id` int(11) unsigned NOT NULL,
  `menu_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  KEY `admin_role_menu_role_id_menu_id_index` (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_role_menu`
--

LOCK TABLES `goadmin_role_menu` WRITE;
INSERT INTO `goadmin_role_menu` VALUES (1,1,'2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_role_menu` VALUES (1,9,'2020-03-26 08:25:56','2020-03-26 08:25:56');
INSERT INTO `goadmin_role_menu` VALUES (2,9,'2020-03-26 08:25:56','2020-03-26 08:25:56');
INSERT INTO `goadmin_role_menu` VALUES (3,9,'2020-03-26 08:25:56','2020-03-26 08:25:56');
INSERT INTO `goadmin_role_menu` VALUES (1,10,'2020-03-26 08:28:45','2020-03-26 08:28:45');
INSERT INTO `goadmin_role_menu` VALUES (2,10,'2020-03-26 08:28:45','2020-03-26 08:28:45');
INSERT INTO `goadmin_role_menu` VALUES (3,10,'2020-03-26 08:28:45','2020-03-26 08:28:45');
INSERT INTO `goadmin_role_menu` VALUES (1,11,'2020-03-26 08:31:03','2020-03-26 08:31:03');
INSERT INTO `goadmin_role_menu` VALUES (2,11,'2020-03-26 08:31:03','2020-03-26 08:31:03');
INSERT INTO `goadmin_role_menu` VALUES (3,11,'2020-03-26 08:31:03','2020-03-26 08:31:03');
INSERT INTO `goadmin_role_menu` VALUES (1,13,'2020-03-26 08:33:29','2020-03-26 08:33:29');
INSERT INTO `goadmin_role_menu` VALUES (2,13,'2020-03-26 08:33:29','2020-03-26 08:33:29');
INSERT INTO `goadmin_role_menu` VALUES (3,13,'2020-03-26 08:33:29','2020-03-26 08:33:29');
INSERT INTO `goadmin_role_menu` VALUES (1,12,'2020-03-26 08:33:44','2020-03-26 08:33:44');
INSERT INTO `goadmin_role_menu` VALUES (2,12,'2020-03-26 08:33:44','2020-03-26 08:33:44');
INSERT INTO `goadmin_role_menu` VALUES (3,12,'2020-03-26 08:33:44','2020-03-26 08:33:44');
INSERT INTO `goadmin_role_menu` VALUES (1,18,'2020-03-27 14:16:22','2020-03-27 14:16:22');
INSERT INTO `goadmin_role_menu` VALUES (3,18,'2020-03-27 14:16:22','2020-03-27 14:16:22');
INSERT INTO `goadmin_role_menu` VALUES (1,19,'2020-03-27 14:16:48','2020-03-27 14:16:48');
INSERT INTO `goadmin_role_menu` VALUES (3,19,'2020-03-27 14:16:48','2020-03-27 14:16:48');
INSERT INTO `goadmin_role_menu` VALUES (1,14,'2020-03-27 14:23:19','2020-03-27 14:23:19');
INSERT INTO `goadmin_role_menu` VALUES (3,14,'2020-03-27 14:23:19','2020-03-27 14:23:19');
INSERT INTO `goadmin_role_menu` VALUES (1,17,'2020-03-27 14:23:27','2020-03-27 14:23:27');
INSERT INTO `goadmin_role_menu` VALUES (3,17,'2020-03-27 14:23:27','2020-03-27 14:23:27');
INSERT INTO `goadmin_role_menu` VALUES (1,15,'2020-03-28 03:55:01','2020-03-28 03:55:01');
INSERT INTO `goadmin_role_menu` VALUES (3,15,'2020-03-28 03:55:01','2020-03-28 03:55:01');
INSERT INTO `goadmin_role_menu` VALUES (1,16,'2020-03-29 04:15:09','2020-03-29 04:15:09');
INSERT INTO `goadmin_role_menu` VALUES (3,16,'2020-03-29 04:15:09','2020-03-29 04:15:09');
INSERT INTO `goadmin_role_menu` VALUES (1,20,'2020-03-31 07:24:46','2020-03-31 07:24:46');
INSERT INTO `goadmin_role_menu` VALUES (3,20,'2020-03-31 07:24:46','2020-03-31 07:24:46');
INSERT INTO `goadmin_role_menu` VALUES (1,21,'2020-04-05 15:08:02','2020-04-05 15:08:02');
INSERT INTO `goadmin_role_menu` VALUES (2,21,'2020-04-05 15:08:02','2020-04-05 15:08:02');
INSERT INTO `goadmin_role_menu` VALUES (3,21,'2020-04-05 15:08:02','2020-04-05 15:08:02');
INSERT INTO `goadmin_role_menu` VALUES (1,22,'2020-04-05 15:09:55','2020-04-05 15:09:55');
INSERT INTO `goadmin_role_menu` VALUES (2,22,'2020-04-05 15:09:55','2020-04-05 15:09:55');
INSERT INTO `goadmin_role_menu` VALUES (3,22,'2020-04-05 15:09:55','2020-04-05 15:09:55');
INSERT INTO `goadmin_role_menu` VALUES (1,23,'2020-04-05 15:10:07','2020-04-05 15:10:07');
INSERT INTO `goadmin_role_menu` VALUES (2,23,'2020-04-05 15:10:07','2020-04-05 15:10:07');
INSERT INTO `goadmin_role_menu` VALUES (3,23,'2020-04-05 15:10:07','2020-04-05 15:10:07');
INSERT INTO `goadmin_role_menu` VALUES (1,7,'2020-04-05 15:12:57','2020-04-05 15:12:57');
INSERT INTO `goadmin_role_menu` VALUES (1,8,'2020-04-05 15:13:17','2020-04-05 15:13:17');
UNLOCK TABLES;

--
-- Table structure for table `goadmin_role_permissions`
--

DROP TABLE IF EXISTS `goadmin_role_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goadmin_role_permissions` (
  `role_id` int(11) unsigned NOT NULL,
  `permission_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `admin_role_permissions` (`role_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_role_permissions`
--

LOCK TABLES `goadmin_role_permissions` WRITE;
INSERT INTO `goadmin_role_permissions` VALUES (1,1,'2020-04-05 15:07:02','2020-04-05 15:07:02');
INSERT INTO `goadmin_role_permissions` VALUES (1,2,'2020-04-05 15:07:02','2020-04-05 15:07:02');
INSERT INTO `goadmin_role_permissions` VALUES (1,3,'2020-04-05 15:07:02','2020-04-05 15:07:02');
INSERT INTO `goadmin_role_permissions` VALUES (1,4,'2020-04-05 15:07:02','2020-04-05 15:07:02');
INSERT INTO `goadmin_role_permissions` VALUES (1,5,'2020-04-05 15:07:02','2020-04-05 15:07:02');
INSERT INTO `goadmin_role_permissions` VALUES (1,6,'2020-04-05 15:07:02','2020-04-05 15:07:02');
INSERT INTO `goadmin_role_permissions` VALUES (1,7,'2020-04-05 15:07:02','2020-04-05 15:07:02');
INSERT INTO `goadmin_role_permissions` VALUES (1,8,'2020-04-05 15:07:02','2020-04-05 15:07:02');
INSERT INTO `goadmin_role_permissions` VALUES (1,9,'2020-04-05 15:07:02','2020-04-05 15:07:02');
INSERT INTO `goadmin_role_permissions` VALUES (1,10,'2020-04-05 15:07:03','2020-04-05 15:07:03');
INSERT INTO `goadmin_role_permissions` VALUES (1,11,'2020-04-05 15:07:02','2020-04-05 15:07:02');
INSERT INTO `goadmin_role_permissions` VALUES (1,12,'2020-04-05 15:07:02','2020-04-05 15:07:02');
INSERT INTO `goadmin_role_permissions` VALUES (1,13,'2020-04-05 15:07:02','2020-04-05 15:07:02');
INSERT INTO `goadmin_role_permissions` VALUES (1,14,'2020-04-05 15:07:02','2020-04-05 15:07:02');
INSERT INTO `goadmin_role_permissions` VALUES (1,15,'2020-04-05 15:07:03','2020-04-05 15:07:03');
INSERT INTO `goadmin_role_permissions` VALUES (1,16,'2020-04-05 15:07:02','2020-04-05 15:07:02');
INSERT INTO `goadmin_role_permissions` VALUES (1,17,'2020-04-05 15:07:02','2020-04-05 15:07:02');
INSERT INTO `goadmin_role_permissions` VALUES (1,18,'2020-04-05 15:07:03','2020-04-05 15:07:03');
INSERT INTO `goadmin_role_permissions` VALUES (1,19,'2020-04-05 15:07:03','2020-04-05 15:07:03');
INSERT INTO `goadmin_role_permissions` VALUES (1,20,'2020-04-05 15:07:03','2020-04-05 15:07:03');
INSERT INTO `goadmin_role_permissions` VALUES (1,21,'2020-04-05 15:07:03','2020-04-05 15:07:03');
INSERT INTO `goadmin_role_permissions` VALUES (2,2,'2020-03-26 08:35:05','2020-03-26 08:35:05');
INSERT INTO `goadmin_role_permissions` VALUES (2,3,'2020-03-26 08:35:05','2020-03-26 08:35:05');
INSERT INTO `goadmin_role_permissions` VALUES (2,4,'2020-03-26 08:35:05','2020-03-26 08:35:05');
INSERT INTO `goadmin_role_permissions` VALUES (2,5,'2020-03-26 08:35:05','2020-03-26 08:35:05');
INSERT INTO `goadmin_role_permissions` VALUES (2,6,'2020-03-26 08:35:05','2020-03-26 08:35:05');
INSERT INTO `goadmin_role_permissions` VALUES (2,7,'2020-03-26 08:35:05','2020-03-26 08:35:05');
INSERT INTO `goadmin_role_permissions` VALUES (3,2,'2020-04-05 15:06:55','2020-04-05 15:06:55');
INSERT INTO `goadmin_role_permissions` VALUES (3,8,'2020-04-05 15:06:55','2020-04-05 15:06:55');
INSERT INTO `goadmin_role_permissions` VALUES (3,9,'2020-04-05 15:06:55','2020-04-05 15:06:55');
INSERT INTO `goadmin_role_permissions` VALUES (3,10,'2020-04-05 15:06:55','2020-04-05 15:06:55');
INSERT INTO `goadmin_role_permissions` VALUES (3,11,'2020-04-05 15:06:55','2020-04-05 15:06:55');
INSERT INTO `goadmin_role_permissions` VALUES (3,12,'2020-04-05 15:06:55','2020-04-05 15:06:55');
INSERT INTO `goadmin_role_permissions` VALUES (3,18,'2020-04-05 15:06:55','2020-04-05 15:06:55');
INSERT INTO `goadmin_role_permissions` VALUES (3,20,'2020-04-05 15:06:55','2020-04-05 15:06:55');
UNLOCK TABLES;

--
-- Table structure for table `goadmin_role_users`
--

DROP TABLE IF EXISTS `goadmin_role_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goadmin_role_users` (
  `role_id` int(11) unsigned NOT NULL,
  `user_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `admin_user_roles` (`role_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_role_users`
--

LOCK TABLES `goadmin_role_users` WRITE;
INSERT INTO `goadmin_role_users` VALUES (1,1,'2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_role_users` VALUES (2,2,'2020-03-26 08:14:36','2020-03-26 08:14:36');
INSERT INTO `goadmin_role_users` VALUES (3,3,'2020-03-26 07:44:12','2020-03-26 07:44:12');
UNLOCK TABLES;

--
-- Table structure for table `goadmin_roles`
--

DROP TABLE IF EXISTS `goadmin_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goadmin_roles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_roles_name_unique` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_roles`
--

LOCK TABLES `goadmin_roles` WRITE;
INSERT INTO `goadmin_roles` VALUES (1,'Administrator','administrator','2019-09-10 00:00:00','2020-04-05 23:07:02');
INSERT INTO `goadmin_roles` VALUES (2,'Operator','operator','2019-09-10 00:00:00','2020-03-26 16:35:05');
INSERT INTO `goadmin_roles` VALUES (3,'Test','test','2020-03-26 07:33:50','2020-04-05 23:06:55');
UNLOCK TABLES;

--
-- Table structure for table `goadmin_session`
--

DROP TABLE IF EXISTS `goadmin_session`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goadmin_session` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `sid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `values` varchar(3000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_session`
--

LOCK TABLES `goadmin_session` WRITE;
INSERT INTO `goadmin_session` VALUES (32,'9539e355-ee24-4b71-a3d9-c94c9274e117','{\"user_id\":1}','2020-04-05 15:02:35','2020-04-05 15:02:35');
UNLOCK TABLES;

--
-- Table structure for table `goadmin_user_permissions`
--

DROP TABLE IF EXISTS `goadmin_user_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goadmin_user_permissions` (
  `user_id` int(11) unsigned NOT NULL,
  `permission_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `admin_user_permissions` (`user_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_user_permissions`
--

LOCK TABLES `goadmin_user_permissions` WRITE;
INSERT INTO `goadmin_user_permissions` VALUES (1,1,'2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_user_permissions` VALUES (2,2,'2020-03-26 08:14:36','2020-03-26 08:14:36');
UNLOCK TABLES;

--
-- Table structure for table `goadmin_users`
--

DROP TABLE IF EXISTS `goadmin_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goadmin_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_users_username_unique` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_users`
--

LOCK TABLES `goadmin_users` WRITE;
INSERT INTO `goadmin_users` VALUES (1,'admin','$2a$10$W7IuQ7/OjaePu9yy8OzoCuDjfGyq7K0NNgDrzU5l23tSQnXsw0LcG','admin','','tlNcBVK9AvfYH7WEnwB1RKvocJu8FfRy4um3DJtwdHuJy0dwFsLOgAc0xUfh','2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_users` VALUES (2,'operator','$2a$10$Trvjo8139K7KDXXORCWAl.7lMl8Q9lZ4F20GvmTbmJD6yM1cz.Tti','Operator','',NULL,'2019-09-10 00:00:00','2020-03-26 16:14:36');
INSERT INTO `goadmin_users` VALUES (3,'test','$2a$10$jgUuvG2JVfnTZGCgxFSH9OzJnxESFU2RQdMsVQbIXrMr9e/RunR3m','test','',NULL,'2020-03-26 07:34:38','2020-03-26 15:44:12');
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-04-05 15:13:44
-- MySQL dump 10.13  Distrib 5.7.27, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: GoAdmin
-- ------------------------------------------------------
-- Server version	5.7.27-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `goadmin_operation_log`
--

DROP TABLE IF EXISTS `goadmin_operation_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goadmin_operation_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL,
  `path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `method` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ip` varchar(15) COLLATE utf8mb4_unicode_ci NOT NULL,
  `input` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `admin_operation_log_user_id_index` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1637 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-04-05 15:13:44

CREATE TABLE `goadmin_site` (
                                `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                                `key` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                `value` longtext COLLATE utf8mb4_unicode_ci,
                                `description` varchar(3000) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                `state` tinyint(3) unsigned NOT NULL DEFAULT '0',
                                `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;