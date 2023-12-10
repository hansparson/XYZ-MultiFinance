-- MySQL dump 10.13  Distrib 8.0.34, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: xyz_multifinance
-- ------------------------------------------------------
-- Server version	5.7.44

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `bills`
--

DROP TABLE IF EXISTS `bills`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `bills` (
  `kode_bayar` varchar(255) DEFAULT NULL,
  `user_id` longtext,
  `terbayarkan` double DEFAULT NULL,
  `total_tagihan` double DEFAULT NULL,
  `sisa_tagihan_bulanan` double DEFAULT NULL,
  `tanggal_bayar` datetime(3) DEFAULT NULL,
  `sisa_limit` double DEFAULT NULL,
  `status_bayar` longtext,
  UNIQUE KEY `idx_bills_kode_bayar` (`kode_bayar`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bills`
--

LOCK TABLES `bills` WRITE;
/*!40000 ALTER TABLE `bills` DISABLE KEYS */;
/*!40000 ALTER TABLE `bills` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `monthly_billings`
--

DROP TABLE IF EXISTS `monthly_billings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `monthly_billings` (
  `billing_id` varchar(255) DEFAULT NULL,
  `user_id` longtext,
  `nama_aset` longtext,
  `tenor` double DEFAULT NULL,
  `harga_cicilan` double DEFAULT NULL,
  `tanggal_tagihan` datetime(3) DEFAULT NULL,
  UNIQUE KEY `idx_monthly_billings_billing_id` (`billing_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `monthly_billings`
--

LOCK TABLES `monthly_billings` WRITE;
/*!40000 ALTER TABLE `monthly_billings` DISABLE KEYS */;
INSERT INTO `monthly_billings` VALUES ('BILL-1702224216973318385-riOUvf','52617426','Baju Bola Messi',1,260,'2024-01-10 16:03:36.973'),('BILL-1702224216980070752-dVExyg','52617426','Baju Bola Messi',2,260,'2024-02-10 16:03:36.973'),('BILL-1702224216982667857-pHN2WX','52617426','Baju Bola Messi',3,260,'2024-03-10 16:03:36.973'),('BILL-1702224216986914771-vFciT4','52617426','Baju Bola Messi',4,260,'2024-04-10 16:03:36.973'),('BILL-1702224216989054215-ZpjndS','52617426','Baju Bola Messi',5,260,'2024-05-10 16:03:36.973'),('BILL-1702224216990997355-oPXgNr','52617426','Baju Bola Messi',6,260,'2024-06-10 16:03:36.973');
/*!40000 ALTER TABLE `monthly_billings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transactions` (
  `nomor_kontrak` varchar(255) DEFAULT NULL,
  `user_id` longtext,
  `tanggal_kontrak` datetime(3) DEFAULT NULL,
  `tanggl_update` datetime(3) DEFAULT NULL,
  `otr` double DEFAULT NULL,
  `admin_fee` double DEFAULT NULL,
  `harga_aset` double DEFAULT NULL,
  `jumlah_cicilan` double DEFAULT NULL,
  `jumlah_bunga` double DEFAULT NULL,
  `tenor` double DEFAULT NULL,
  `cicilan_bulanan` double DEFAULT NULL,
  `nama_aset` longtext,
  UNIQUE KEY `idx_transactions_nomor_kontrak` (`nomor_kontrak`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` VALUES ('1702224216961572161-vOQ6gK','52617426','2023-12-10 16:03:36.963','2023-12-10 16:03:36.963',100,100,1000,1560,5,6,260,'Baju Bola Messi');
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_limit_balances`
--

DROP TABLE IF EXISTS `user_limit_balances`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_limit_balances` (
  `user_id` varchar(255) DEFAULT NULL,
  `limit_one_month` double DEFAULT NULL,
  `limit_two_month` double DEFAULT NULL,
  `limit_three_month` double DEFAULT NULL,
  `limit_sixth` double DEFAULT NULL,
  UNIQUE KEY `idx_user_limit_balances_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_limit_balances`
--

LOCK TABLES `user_limit_balances` WRITE;
/*!40000 ALTER TABLE `user_limit_balances` DISABLE KEYS */;
INSERT INTO `user_limit_balances` VALUES ('52617426',100000,300000,500000,9998440);
/*!40000 ALTER TABLE `user_limit_balances` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `user_id` varchar(255) DEFAULT NULL,
  `user_status` longtext,
  `nik` varchar(255) DEFAULT NULL,
  `full_name` longtext,
  `legal_name` longtext,
  `tempat_lahir` longtext,
  `tanggal_lahir` longtext,
  `gaji` double DEFAULT NULL,
  `foto_ktp` longtext,
  `foto_selfie` longtext,
  UNIQUE KEY `idx_users_user_id` (`user_id`),
  UNIQUE KEY `idx_users_nik` (`nik`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('52617426','ACTIVE','212015013t35','Hans Parson','Hans Parson','Polewali','1997-03-02',12000000,'https://assets.pikiran-rakyat.com/crop/0x0:0x0/x/photo/2021/04/20/750175463.jpg','https://assets.pikiran-rakyat.com/crop/0x0:0x0/x/photo/2021/04/20/750175463.jpg');
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

-- Dump completed on 2023-12-10 23:47:52
