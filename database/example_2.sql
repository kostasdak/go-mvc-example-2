-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.4.11-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             11.3.0.6295
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- Dumping structure for table golang.colors
CREATE TABLE IF NOT EXISTS `colors` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product_id` int(11) NOT NULL,
  `color` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table golang.colors: ~4 rows (approximately)
DELETE FROM `colors`;
/*!40000 ALTER TABLE `colors` DISABLE KEYS */;
INSERT INTO `colors` (`id`, `product_id`, `color`) VALUES
	(11, 1, 'Blue'),
	(12, 1, 'Red'),
	(13, 2, 'Green'),
	(15, 2, 'Black');
/*!40000 ALTER TABLE `colors` ENABLE KEYS */;

-- Dumping structure for table golang.pages
CREATE TABLE IF NOT EXISTS `pages` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `content` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table golang.pages: ~0 rows (approximately)
DELETE FROM `pages`;
/*!40000 ALTER TABLE `pages` DISABLE KEYS */;
INSERT INTO `pages` (`id`, `name`, `content`) VALUES
	(1, 'home', '<p>\r\n<strong>page content from database table [pages]</strong></br> \r\npage content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] -  \r\n</p>\r\n<p>\r\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur justo dolor, sagittis vitae elit vel, vehicula dignissim purus. Etiam ut erat interdum arcu ultrices egestas eu fringilla sem. Donec pulvinar fringilla quam, vitae tristique ipsum venenatis eu. Curabitur egestas porttitor dolor, quis suscipit massa euismod eget. Nullam velit ligula, rutrum at pellentesque sit amet, maximus quis leo. Ut laoreet ex sem, vel posuere nisl viverra et. Nullam suscipit volutpat metus. Sed mattis aliquet commodo. Pellentesque cursus elit eget est blandit, ac posuere risus posuere. Nunc lobortis tortor ante, eu sagittis sapien molestie non. Pellentesque ultricies orci at enim pulvinar, et rutrum turpis egestas.\r\n</p>\r\n<p>\r\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur justo dolor, sagittis vitae elit vel, vehicula dignissim purus. Etiam ut erat interdum arcu ultrices egestas eu fringilla sem. Donec pulvinar fringilla quam, vitae tristique ipsum venenatis eu. Curabitur egestas porttitor dolor, quis suscipit massa euismod eget. Nullam velit ligula, rutrum at pellentesque sit amet, maximus quis leo. Ut laoreet ex sem, vel posuere nisl viverra et. Nullam suscipit volutpat metus. Sed mattis aliquet commodo. Pellentesque cursus elit eget est blandit, ac posuere risus posuere. Nunc lobortis tortor ante, eu sagittis sapien molestie non. Pellentesque ultricies orci at enim pulvinar, et rutrum turpis egestas.\r\n</p>\r\n<p>\r\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur justo dolor, sagittis vitae elit vel, vehicula dignissim purus. Etiam ut erat interdum arcu ultrices egestas eu fringilla sem. Donec pulvinar fringilla quam, vitae tristique ipsum venenatis eu. Curabitur egestas porttitor dolor, quis suscipit massa euismod eget. Nullam velit ligula, rutrum at pellentesque sit amet, maximus quis leo. Ut laoreet ex sem, vel posuere nisl viverra et. Nullam suscipit volutpat metus. Sed mattis aliquet commodo. Pellentesque cursus elit eget est blandit, ac posuere risus posuere. Nunc lobortis tortor ante, eu sagittis sapien molestie non. Pellentesque ultricies orci at enim pulvinar, et rutrum turpis egestas.\r\n</p> ');
/*!40000 ALTER TABLE `pages` ENABLE KEYS */;

-- Dumping structure for table golang.products
CREATE TABLE IF NOT EXISTS `products` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `code` varchar(50) NOT NULL,
  `type` varchar(50) NOT NULL,
  `name` varchar(250) NOT NULL,
  `description` text NOT NULL,
  `price` float DEFAULT NULL,
  `images` text DEFAULT NULL,
  `status` text DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table golang.products: ~3 rows (approximately)
DELETE FROM `products`;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` (`id`, `code`, `type`, `name`, `description`, `price`, `images`, `status`) VALUES
	(1, 'F-150', 'Truck', 'Ford F-1500', 'The redesigned 2021 Ford F-150 claims one of the top spots in our full-size pickup truck rankings because of its tremendous capability and spacious, comfortable cabin', 29990, NULL, 'avail'),
	(2, 'F-250', 'Truck', 'Ford F-2500', 'The all new 2022 Ford Super Duty F-250 Crew Cab comes fully ready with interior comforts, advanced technology, and exterior conveniences to make every trip you decide to go on in this powerful 4-door Pickup so enjoyable! With an available 7.3L gas V8 engine, the Super Duty F-250 Crew Cab can easily cut through any terrain and get even the toughest of jobs done! Discover how much its massive available engine can do for you. ', 35990, NULL, NULL);
/*!40000 ALTER TABLE `products` ENABLE KEYS */;

-- Dumping structure for table golang.users
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `firstname` varchar(50) NOT NULL DEFAULT '0',
  `lastname` varchar(50) NOT NULL DEFAULT '0',
  `username` text NOT NULL,
  `password` text NOT NULL,
  `auth` text NOT NULL,
  `expire` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Dumping data for table golang.users: ~0 rows (approximately)
DELETE FROM `users`;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
