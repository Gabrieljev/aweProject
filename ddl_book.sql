-- book_store.book definition

CREATE TABLE `book` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(6500) NOT NULL,
  `author` varchar(255) DEFAULT NULL,
  `pub_id` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `is_deleted` int(1) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `book_pub_id_is_deleted` (`pub_id`,`is_deleted`) USING BTREE,
  CONSTRAINT `book_FK` FOREIGN KEY (`pub_id`) REFERENCES `publisher` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
