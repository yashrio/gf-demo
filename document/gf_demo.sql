/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 50733
 Source Host           : localhost:3306
 Source Schema         : gf_demo

 Target Server Type    : MySQL
 Target Server Version : 50733
 File Encoding         : 65001

 Date: 19/02/2021 17:49:55
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint(20) NOT NULL COMMENT '用户ID',
  `passport` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户账号',
  `password` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户密码',
  `nickname` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户昵称',
  `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'ctx2020', '123456', 'ctx2020', '2021-02-16 15:45:28', '2021-02-16 15:45:28');
INSERT INTO `user` VALUES (2, 'ctx2020', '123456', 'ctx2020', '2021-02-16 15:45:43', '2021-02-16 15:45:43');
INSERT INTO `user` VALUES (10, 'ctx2020', '123456', 'ctx2020', '2021-02-16 17:54:02', '2021-02-16 20:58:40');
INSERT INTO `user` VALUES (1361662496477286400, 'ctx2020', '123456', 'ctx2020', '2021-02-16 21:03:32', '2021-02-16 21:03:32');
INSERT INTO `user` VALUES (1361663008891211776, 'ctx2020', '123456', 'ctx2020', '2021-02-16 21:05:34', '2021-02-16 21:05:34');
INSERT INTO `user` VALUES (1361663025345466368, 'ctx2020', '123456', 'ctx2020', '2021-02-16 21:05:38', '2021-02-16 21:05:38');
INSERT INTO `user` VALUES (1361663029627850752, 'ctx2020', '123456', 'ctx2020', '2021-02-16 21:05:39', '2021-02-16 21:05:39');
INSERT INTO `user` VALUES (1361663032236707840, 'ctx2020', '123456', 'ctx2020', '2021-02-16 21:05:40', '2021-02-16 21:05:40');
INSERT INTO `user` VALUES (1361663066881658880, 'ctx2020', '123456', 'ctx2020', '2021-02-16 21:05:48', '2021-02-16 21:05:48');
INSERT INTO `user` VALUES (1361663069381464064, 'ctx2020', '123456', 'ctx2020', '2021-02-16 21:05:48', '2021-02-16 21:05:48');
INSERT INTO `user` VALUES (1361663071734468608, 'ctx2020', '123456', 'ctx2020', '2021-02-16 21:05:49', '2021-02-16 21:05:49');
INSERT INTO `user` VALUES (1361663073965838336, 'ctx2020', '123456', 'ctx2020', '2021-02-16 21:05:49', '2021-02-16 21:05:49');
INSERT INTO `user` VALUES (1361663249732341760, 'ctx2020', '123456', 'ctx2020', '2021-02-16 21:06:31', '2021-02-16 21:06:31');
INSERT INTO `user` VALUES (1361663332498542592, 'ctx2020', '123456', 'ctx2020', '2021-02-16 21:06:51', '2021-02-16 21:06:51');

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info`  (
  `id` bigint(20) NOT NULL COMMENT 'id',
  `nick_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '昵称',
  `avatar_url` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '头像',
  `gender` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '性别',
  `birthday` date NULL DEFAULT NULL COMMENT '生日',
  `zodiac` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '星座',
  `country` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '中国' COMMENT '国家',
  `province` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '省份',
  `city` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '城市',
  `area` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '地区',
  `profile` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '个人简介',
  `id_number` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '平台ID',
  `tags` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '个性标签',
  `create_at` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `update_at` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '用户信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_info
-- ----------------------------
INSERT INTO `user_info` VALUES (0, '', '', '0', NULL, '', '', '', '', '', '', '123213213', '', '2021-02-18 21:23:40', '2021-02-18 21:23:40');
INSERT INTO `user_info` VALUES (1, '', '', '0', NULL, NULL, '中国', '', '', '', NULL, NULL, NULL, NULL, NULL);
INSERT INTO `user_info` VALUES (22, 'ctx', '', '0', NULL, '', '', '', '', '', '', '123213213', '', '2021-02-18 21:35:52', '2021-02-18 21:35:52');

-- ----------------------------
-- Table structure for user_login
-- ----------------------------
DROP TABLE IF EXISTS `user_login`;
CREATE TABLE `user_login`  (
  `id` bigint(20) NOT NULL COMMENT 'id',
  `user_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '登录账号',
  `password` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '登录密码',
  `user_status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '用户状态',
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `create_at` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `update_at` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '用户平台账号' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_login
-- ----------------------------
INSERT INTO `user_login` VALUES (0, '13718878727', '', '', '', '2021-02-17 19:29:42', '2021-02-17 19:29:42');
INSERT INTO `user_login` VALUES (1, '13623331789', '', '0', NULL, '2021-02-18 00:00:00', '2021-02-18 00:00:00');

-- ----------------------------
-- Table structure for user_social
-- ----------------------------
DROP TABLE IF EXISTS `user_social`;
CREATE TABLE `user_social`  (
  `id` bigint(20) NOT NULL COMMENT 'id',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `social_type` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '账户类型',
  `open_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'OpenId',
  `union_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'UnionId',
  `session_key` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT 'SessionKey',
  `create_at` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `update_at` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '用户社交账号' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_social
-- ----------------------------
INSERT INTO `user_social` VALUES (0, 22, 'wechat', 'orLQB5UdIZTKmSwlyupMLJ6WsyXI', '', 'etnf4861n5QplKFTswspBg==', '2021-02-18 21:37:10', '2021-02-18 21:37:10');

SET FOREIGN_KEY_CHECKS = 1;
