-- MySQL dump 10.13  Distrib 8.4.6, for Linux (x86_64)
--
-- Host: localhost    Database: ziyarem
-- ------------------------------------------------------
-- Server version	8.4.6-0ubuntu0.25.04.3

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `ziyarem`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `ziyarem` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `ziyarem`;

--
-- Table structure for table `sensor_apis`
--

DROP TABLE IF EXISTS `sensor_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sensor_apis` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `device_id` varchar(64) NOT NULL,
  `endpoint` varchar(255) NOT NULL,
  `method` varchar(16) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_sensor_apis_device` (`device_id`),
  CONSTRAINT `fk_sensor_apis_device` FOREIGN KEY (`device_id`) REFERENCES `sensor_devices` (`device_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sensor_apis`
--

LOCK TABLES `sensor_apis` WRITE;
/*!40000 ALTER TABLE `sensor_apis` DISABLE KEYS */;
INSERT INTO `sensor_apis` VALUES (1,'temp-001','http://localhost:8081/temp','GET'),(2,'hum-001','http://localhost:8081/hum','GET'),(3,'air-001','http://localhost:8081/air','GET');
/*!40000 ALTER TABLE `sensor_apis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sensor_data`
--

DROP TABLE IF EXISTS `sensor_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sensor_data` (
  `id` varchar(64) NOT NULL,
  `device_id` varchar(64) NOT NULL,
  `value` double NOT NULL,
  `timestamp` datetime(3) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_sensor_data_device` (`device_id`),
  CONSTRAINT `fk_sensor_data_device` FOREIGN KEY (`device_id`) REFERENCES `sensor_devices` (`device_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sensor_data`
--

LOCK TABLES `sensor_data` WRITE;
/*!40000 ALTER TABLE `sensor_data` DISABLE KEYS */;
INSERT INTO `sensor_data` VALUES ('air-001','air-001',71.18,'2025-09-21 18:49:19.192'),('hum-001','hum-001',55.56,'2025-09-21 18:49:19.187'),('temp-001','temp-001',23.03,'2025-09-21 18:49:19.141');
/*!40000 ALTER TABLE `sensor_data` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sensor_devices`
--

DROP TABLE IF EXISTS `sensor_devices`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sensor_devices` (
  `device_id` varchar(64) NOT NULL,
  `type_id` bigint unsigned NOT NULL,
  `location` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`device_id`),
  KEY `fk_sensor_devices_type` (`type_id`),
  CONSTRAINT `fk_sensor_devices_type` FOREIGN KEY (`type_id`) REFERENCES `sensor_types` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sensor_devices`
--

LOCK TABLES `sensor_devices` WRITE;
/*!40000 ALTER TABLE `sensor_devices` DISABLE KEYS */;
INSERT INTO `sensor_devices` VALUES ('air-001',3,'Room C'),('hum-001',2,'Room B'),('temp-001',1,'Room A');
/*!40000 ALTER TABLE `sensor_devices` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sensor_types`
--

DROP TABLE IF EXISTS `sensor_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sensor_types` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_sensor_types_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sensor_types`
--

LOCK TABLES `sensor_types` WRITE;
/*!40000 ALTER TABLE `sensor_types` DISABLE KEYS */;
INSERT INTO `sensor_types` VALUES (3,'airquality'),(2,'humidity'),(1,'temperature');
/*!40000 ALTER TABLE `sensor_types` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-09-21 19:14:03
