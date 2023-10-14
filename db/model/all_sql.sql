CREATE TABLE `node` (
                        `id` int NOT NULL AUTO_INCREMENT,
                        `node_id` int DEFAULT NULL,
                        `version_id` int DEFAULT NULL,
                        `name` varchar(255) DEFAULT NULL,
                        `content` varchar(255) DEFAULT NULL,
                        `comment` varchar(255) DEFAULT NULL,
                        `create_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `node_id_version_id` (`version_id`,`node_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `comment` (
                           `id` int NOT NULL AUTO_INCREMENT,
                           `node_id` int DEFAULT NULL,
                           `content` varchar(255) DEFAULT NULL,
                           `avatar_url` varchar(255) DEFAULT NULL,
                           `name` varchar(255) DEFAULT NULL,
                           `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
                           `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `edge` (
                        `id` int NOT NULL AUTO_INCREMENT,
                        `src_id` int NOT NULL,
                        `tar_id` int NOT NULL,
                        `description` varchar(255) DEFAULT NULL,
                        `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
                        `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;