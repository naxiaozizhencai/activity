CREATE TABLE `anwser` (
                          `anwser_id` int(11) NOT NULL COMMENT '问题id',
                          `anwser_name` varchar(50) DEFAULT NULL COMMENT '关卡名称',
                          `last_answer_id` int(11) DEFAULT '0' COMMENT '上一个关卡id',
                          `start_time` datetime NOT NULL COMMENT '开启时间',
                          `end_time` datetime NOT NULL COMMENT '结束时间',
                          `item_id` varchar(50) DEFAULT NULL COMMENT '道具id',
                          `item_num` int(11) DEFAULT NULL COMMENT '道具数量',
                          `result` varchar(10) NOT NULL COMMENT '结果',
                          `add_time` datetime NOT NULL COMMENT '添加时间',
                          PRIMARY KEY (`anwser_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='答题列表';
INSERT INTO `anwser` VALUES ('1', '关卡1', '0', '2023-01-04 00:00:00', '2023-01-31 00:00:00', 'l_coin', '800', 'C', '2022-12-16 17:18:36');
INSERT INTO `anwser` VALUES ('2', '关卡2', '1', '2023-01-05 00:00:00', '2023-01-31 00:00:00', 'rush_common_60m', '5', 'B', '2022-12-16 17:18:59');
INSERT INTO `anwser` VALUES ('3', '关卡3', '2', '2023-01-06 00:00:00', '2023-01-31 00:00:00', 'top_key', '5', 'C', '2022-12-16 17:19:13');
INSERT INTO `anwser` VALUES ('4', '关卡4', '3', '2023-01-07 00:00:00', '2023-01-31 00:00:00', 'camp_key', '3', 'D', '2022-12-16 17:19:38');
INSERT INTO `anwser` VALUES ('5', '关卡5', '4', '2023-01-08 00:00:00', '2023-01-31 00:00:00', 'decoration_ice_tank', '1', 'X', '2022-12-16 17:20:03');
CREATE TABLE `reward` (
                          `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
                          `game_uid` varchar(100) NOT NULL COMMENT '游戏uid',
                          `answer_id` int(11) NOT NULL COMMENT '题目id',
                          `item_id` varchar(50) NOT NULL COMMENT '道具id',
                          `nums` int(11) NOT NULL DEFAULT '0' COMMENT '数量',
                          `status` tinyint(4) unsigned NOT NULL COMMENT '0失败1成功',
                          `add_time` datetime DEFAULT NULL COMMENT '添加时间',
                          PRIMARY KEY (`id`),
                          KEY `game_uid` (`game_uid`,`answer_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=79 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `anwser_result` (
                                 `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
                                 `game_uid` varchar(100) NOT NULL COMMENT '游戏uid',
                                 `answer_id` int(11) NOT NULL COMMENT '题目id',
                                 `status` tinyint(4) unsigned NOT NULL COMMENT '0失败1成功',
                                 `add_time` datetime DEFAULT NULL COMMENT '添加时间',
                                 PRIMARY KEY (`id`),
                                 UNIQUE KEY `game_uid` (`game_uid`,`answer_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8mb4 COMMENT='答题结果';