DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
                         `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
                         `user_id` bigint NOT NULL,
                         `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                         `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                         `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                         `gender` tinyint NULL DEFAULT 0,
                         `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`) USING BTREE,
                         UNIQUE INDEX `idx_username`(`username`) USING BTREE,
                         UNIQUE INDEX `idx_user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;