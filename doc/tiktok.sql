-- User表
CREATE TABLE `users`
(
    `id`             bigint(0) NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `name`           varchar(32)  NOT NULL,
    `password`       varchar(200)  NULL DEFAULT NULL,
    `follow_count`   bigint(0)  NULL DEFAULT 0,
    `follower_count` bigint(0)  NULL DEFAULT 0,
    `avatar` varchar(255) DEFAULT NULL COMMENT '头像链接',
    `signature` varchar(255) DEFAULT NULL COMMENT '签名',
    `background_image` varchar(255) DEFAULT NULL COMMENT '背景图',
    PRIMARY KEY (`id`)
)
ALTER TABLE `tiktok`.`users`
    ADD UNIQUE INDEX `name&password`(`name`, `password`) USING BTREE COMMENT 'name+password的唯一组合索引';

-- 视频表
CREATE TABLE `videos`
(
    `id`             bigint(0) NOT NULL  COMMENT 'video_id',
    `play_url`       varchar(200)   DEFAULT NULL,
    `cover_url`      varchar(200)  DEFAULT NULL,
    `favorite_count` bigint(0) NULL DEFAULT 0 COMMENT '点赞量',
    `comment_count`  bigint(0) NULL DEFAULT 0 COMMENT '评论量',
    `title`          varchar(128)  DEFAULT '' COMMENT '标题',
    `user_id`        bigint(0) NOT NULL COMMENT 'FK reference user id',
    `create_time`    datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP (0),
    PRIMARY KEY (`id`)
)

-- 评论表
CREATE TABLE `comments`
(
    `id`          bigint(0) NOT NULL AUTO_INCREMENT,
    `user_id`     bigint(0) NOT NULL,
    `video_id`    bigint(0) NOT NULL,
    `content`     varchar(500) DEFAULT '',
    `create_time`    datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP (0),
    PRIMARY KEY (`id`)
)
ALTER TABLE `tiktok`.`comments`
    ADD INDEX `videoId`(`video_id`) USING BTREE COMMENT 'videoId的普通索引';

-- 点赞表
CREATE TABLE `favorites`
(
    `id`          bigint(0) NOT NULL AUTO_INCREMENT,
    `user_id`     bigint(0) NOT NULL,
    `video_id`    bigint(0) NOT NULL,
    `is_favorite` tinyint(0) NULL DEFAULT 0,
    PRIMARY KEY (`id`)
)
ALTER TABLE `tiktok`.`favorites`
    ADD UNIQUE INDEX `user_video`(`user_id`, `video_id`) USING BTREE COMMENT 'user_id+video_id的唯一索引';

-- 关注表
CREATE TABLE `relations`
(
    `id`           bigint(0) NOT NULL AUTO_INCREMENT,
    `user_id`      bigint(0) NULL DEFAULT 0,
    `following_id` bigint(0) NULL DEFAULT 0,
    PRIMARY KEY (`id`)
)
ALTER TABLE `tiktok`.`relations`
    ADD UNIQUE INDEX `user_follow`(`user_id`, `following_id`) USING BTREE COMMENT '关注着和被关注者的id构成唯一索引';