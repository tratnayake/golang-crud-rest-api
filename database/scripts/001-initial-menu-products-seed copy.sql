# ************************************************************
# Sequel Ace SQL dump
# Version 20046
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# Host: 127.0.0.1 (MySQL 8.0.32)
# Database: crud_demo
# Generation Time: 2023-02-04 18:10:57 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table menu_products
# ------------------------------------------------------------

DROP TABLE IF EXISTS `menu_products`;

CREATE TABLE `menu_products` (
  `menu_id` bigint unsigned NOT NULL,
  `product_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`menu_id`,`product_id`),
  KEY `fk_menu_products_product` (`product_id`),
  CONSTRAINT `fk_menu_products_menu` FOREIGN KEY (`menu_id`) REFERENCES `menus` (`id`),
  CONSTRAINT `fk_menu_products_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `menu_products` WRITE;
/*!40000 ALTER TABLE `menu_products` DISABLE KEYS */;

INSERT INTO `menu_products` (`menu_id`, `product_id`)
VALUES
	(1,1),
  (1,2),
  (1,3),
  (2,4),
  (2,5),
  (2,6);

/*!40000 ALTER TABLE `menu_products` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table menus
# ------------------------------------------------------------

DROP TABLE IF EXISTS `menus`;

CREATE TABLE `menus` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `menus` WRITE;
/*!40000 ALTER TABLE `menus` DISABLE KEYS */;

INSERT INTO `menus` (`id`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,'2023-01-04 18:09:57.054','2023-01-04 18:09:57.056',NULL),
  (2,'2023-01-11 18:09:57.054','2023-01-11 18:09:57.056',NULL);

/*!40000 ALTER TABLE `menus` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table products
# ------------------------------------------------------------

DROP TABLE IF EXISTS `products`;

CREATE TABLE `products` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `price` double DEFAULT NULL,
  `description` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_products_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;

INSERT INTO `products` (`id`, `name`, `price`, `description`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,'Greek Chicken, Roasted Nugget Potatoes & Market Vegetables',15,'Fresh herbs, lemon marinated chicken breast topped with tzatziki','2023-02-04 18:09:57.054','2023-02-04 18:09:57.054',NULL),
  (2,'Miso Glazed Halibut with Rice Pilaf & Mixed Vegetables',15,'Oven-baked halibut glazed with miso, topped with green onions and sesame seeds','2023-02-04 18:09:57.054','2023-02-04 18:09:57.054',NULL),
  (3,'Canneloni & Zucchini Ribbons',15,'Spinach stuffed with canneloni with marinara and bechamel, topped with feta cheese and sun-dried tomato pesto','2023-02-04 18:09:57.054','2023-02-04 18:09:57.054',NULL),
  (4,'Alsation Chicken with Mashed Potatoes and Mixed Vegetables',15,'chicken breast with cream sauce topped with sautéed onions, sliced mushrooms, and crispy bacon bits','2023-02-04 18:09:57.054','2023-02-04 18:09:57.054',NULL),
  (5,'Seafood Spaghettini',15,'spaghettini pasta tossed in basil-plum tomato sauce with sautéed scallops, prawns, and smoked sockeye salmon with grated regiano parmesan cheese','2023-02-04 18:09:57.054','2023-02-04 18:09:57.054',NULL),
  (6,'Coconut Curry Pork Shank with Steamed Rice & Market Vegetables ',15,'braised pork shank in coconut cream curry topped with chopped cilantro and thai chili','2023-02-04 18:09:57.054','2023-02-04 18:09:57.054',NULL);


/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
