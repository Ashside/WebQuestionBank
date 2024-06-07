-- MySQL dump 10.13  Distrib 8.0.37, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: seproject
-- ------------------------------------------------------
-- Server version	8.0.37

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
-- Table structure for table `choice_question_keywords`
--

DROP TABLE IF EXISTS `choice_question_keywords`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `choice_question_keywords` (
  `question_id` int NOT NULL,
  `keyword_id` int NOT NULL,
  PRIMARY KEY (`question_id`,`keyword_id`),
  KEY `keyword_id` (`keyword_id`),
  CONSTRAINT `choice_question_keywords_ibfk_1` FOREIGN KEY (`question_id`) REFERENCES `choicequestions` (`id`),
  CONSTRAINT `choice_question_keywords_ibfk_2` FOREIGN KEY (`keyword_id`) REFERENCES `keywords` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `choice_question_keywords`
--

LOCK TABLES `choice_question_keywords` WRITE;
/*!40000 ALTER TABLE `choice_question_keywords` DISABLE KEYS */;
INSERT INTO `choice_question_keywords` VALUES (13,61),(13,62),(13,63),(13,64),(13,65);
/*!40000 ALTER TABLE `choice_question_keywords` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `choicequestions`
--

DROP TABLE IF EXISTS `choicequestions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `choicequestions` (
  `id` int NOT NULL,
  `subject` varchar(45) DEFAULT NULL,
  `content` longtext,
  `options` longtext,
  `answer` longtext,
  `difficulty` varchar(45) DEFAULT NULL,
  `author` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `choicequestions`
--

LOCK TABLES `choicequestions` WRITE;
/*!40000 ALTER TABLE `choicequestions` DISABLE KEYS */;
INSERT INTO `choicequestions` VALUES (6,'math','aaaaaaaa','{\"option1\":\"1\",\"option2\":\"2\",\"option3\":\"3\",\"option4\":\"4\"}','option2','1','admin@hit.com'),(13,'history','从下面三个题目中任选一题，按要求作答。不超过150字。不透露所在区、学校及个人信息。','{\"option1\":\"1）微信朋友圈有“点赞”功能。有人关注“点赞”数量，有人热衷于给人“点赞”……对“点赞”现象，你有什么看法？请说明你的观点和理由。要求：观点明确，言之有理。\",\"option2\":\"（2）年级准备开展“走进名人故乡”主题研学活动，计划在目的地研学两天，现征询同学们对目的地的建议。你建议去哪里？请说说理由。要求：明确写出名人及其故乡，重点陈述理由，理由合理。\",\"option3\":\"3）请以“月的独白”为题目，用月亮的口吻，写一首小诗或一段抒情文字。要求：感情真挚，语言生动，有感染力。\",\"option4\":\"略\"}','option1, option2, option3','2','admin@hit.com');
/*!40000 ALTER TABLE `choicequestions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `keywords`
--

DROP TABLE IF EXISTS `keywords`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `keywords` (
  `id` int NOT NULL,
  `keyword` varchar(45) NOT NULL,
  `score` float NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `keywords`
--

LOCK TABLES `keywords` WRITE;
/*!40000 ALTER TABLE `keywords` DISABLE KEYS */;
INSERT INTO `keywords` VALUES (1,'问题',0.143854),(2,'人工智能',0.0529375),(3,'个人信息',0.0528843),(4,'应用',0.0528545),(5,'互联网',0.052852),(6,'材料',0.0527445),(7,'联想',0.0527068),(8,'越来越少',0.0527038),(9,'答案',0.0522334),(10,'篇文章',0.050825),(11,'普及',0.0503329),(12,'角度',0.0500716),(13,'标题',0.0489665),(14,'思考',0.0458662),(15,'文体',0.0431271),(16,'立意',0.0348434),(17,'问题',0.194547),(18,'人工智能',0.0715916),(19,'个人信息',0.0715192),(20,'应用',0.0714798),(21,'互联网',0.0714768),(22,'越来越少',0.0712789),(23,'答案',0.0706733),(24,'普及',0.0680628),(25,'角度',0.0678662),(26,'标题',0.0657987),(27,'文体',0.058041),(28,'立意',0.0478526),(29,'地球',0.095962),(30,'个人信息',0.0372959),(31,'月球',0.0370993),(32,'联想',0.0370345),(33,'任务',0.0367034),(34,'目光',0.0365708),(35,'航天人',0.0364374),(36,'篇文章',0.0358776),(37,'地球',0.135814),(38,'联想',0.0525811),(39,'月球',0.0525129),(40,'任务',0.0517376),(41,'目光',0.0517204),(42,'航天人',0.0516082),(43,'人类',0.0508552),(44,'篇文章',0.0507899),(45,'思考',0.0497734),(46,'现代文',0.0477002),(47,'面纱',0.0457808),(48,'未知',0.0454514),(49,'天问一号',0.0447524),(50,'探月',0.0399047),(51,'嫦娥',0.039384),(52,'月背',0.0385259),(53,'试卷',0.0303027),(54,'个人信息',0.0807262),(55,'想法',0.151806),(56,'联想',0.151364),(57,'篇文章',0.146039),(58,'冲突',0.144281),(59,'思考',0.131299),(60,'相遇',0.12804),(61,'学校',0.233135),(62,'个人信息',0.232795),(63,'题目',0.229617),(64,'一题',0.164321),(65,'要求',0.140132);
/*!40000 ALTER TABLE `keywords` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `subjective_question_keywords`
--

DROP TABLE IF EXISTS `subjective_question_keywords`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `subjective_question_keywords` (
  `question_id` int NOT NULL,
  `keyword_id` int NOT NULL,
  PRIMARY KEY (`question_id`,`keyword_id`),
  KEY `keyword_id` (`keyword_id`),
  CONSTRAINT `subjective_question_keywords_ibfk_1` FOREIGN KEY (`question_id`) REFERENCES `subjectivequestions` (`id`),
  CONSTRAINT `subjective_question_keywords_ibfk_2` FOREIGN KEY (`keyword_id`) REFERENCES `keywords` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `subjective_question_keywords`
--

LOCK TABLES `subjective_question_keywords` WRITE;
/*!40000 ALTER TABLE `subjective_question_keywords` DISABLE KEYS */;
INSERT INTO `subjective_question_keywords` VALUES (12,55),(12,56),(12,57),(12,58),(12,59),(12,60);
/*!40000 ALTER TABLE `subjective_question_keywords` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `subjectivequestions`
--

DROP TABLE IF EXISTS `subjectivequestions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `subjectivequestions` (
  `id` int NOT NULL,
  `subject` varchar(45) DEFAULT NULL,
  `content` longtext,
  `answer` longtext,
  `difficulty` varchar(45) DEFAULT NULL,
  `author` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `subjectivequestions`
--

LOCK TABLES `subjectivequestions` WRITE;
/*!40000 ALTER TABLE `subjectivequestions` DISABLE KEYS */;
INSERT INTO `subjectivequestions` VALUES (1,'history','123','123','1','admin@hit.com'),(2,'history','测试','option1, option4','1','admin@hit.com'),(3,'history','1','option3','1','admin@hit.com'),(4,'math','123123123','option1','1','admin@hit.com'),(5,'math','aaa','option3','1','admin@hit.com'),(7,'history','随着互联网的普及、人工智能的应用，越来越多的问题能很快得到答案。那么，我们的问题是否会越来越少？\n\n	以上材料引发了你怎样的联想和思考？请写一篇文章。\n\n	要求：选准角度，确定立意，明确文体，自拟标题；不要套作，不得抄袭；不得泄露个人信息；不少于800字。','略','2','admin@hit.com'),(8,'history','随着互联网的普及、人工智能的应用，越来越多的问题能很快得到答案。那么，我们的问题是否会越来越少？\n\n	要求：选准角度，确定立意，明确文体，自拟标题；不要套作，不得抄袭；不得泄露个人信息；不少于800字。','略','2','admin@hit.com'),(9,'history','本试卷现代文阅读I提到，长久以来，人们只能看到月球固定朝向地球的一面，“嫦娥四号”探月任务揭开了月背的神秘面纱；随着“天问一号”飞离地球，航天人的目光又投向遥远的深空……\n\n	正如人类的太空之旅，我们每个人也都在不断抵达未知之境。\n\n	这引发了你怎样的联想与思考？请写一篇文章。\n\n	要求：选准角度，确定立意，明确文体，自拟标题；不要套作，不得抄袭；不得泄露个人信息；不少于800字。','略','2','admin@hit.com'),(10,'history','本试卷现代文阅读I提到，长久以来，人们只能看到月球固定朝向地球的一面，“嫦娥四号”探月任务揭开了月背的神秘面纱；随着“天问一号”飞离地球，航天人的目光又投向遥远的深空……\n\n	正如人类的太空之旅，我们每个人也都在不断抵达未知之境。\n\n	这引发了你怎样的联想与思考？请写一篇文章。\n\n	','略','2','admin@hit.com'),(11,'history','每个人都要学习与他人相处。有时，我们为避免冲突而不愿表达自己的想法。其实，坦诚交流才有可能迎来真正的相遇。\n\n	这引发了你怎样的联想和思考？请写一篇文章。\n\n	要求：选准角度，确定立意，明确文体，自拟标题；不要套作，不得抄袭；不得泄露个人信息；不少于800字。','略','1','admin@hit.com'),(12,'math','每个人都要学习与他人相处。有时，我们为避免冲突而不愿表达自己的想法。其实，坦诚交流才有可能迎来真正的相遇。\n\n	这引发了你怎样的联想和思考？请写一篇文章。\n\n','略','2','admin@hit.com');
/*!40000 ALTER TABLE `subjectivequestions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `type` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('admin@hit.com','admin','ADMIN'),('test@hit.com','test','user');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-06-07 20:03:26
