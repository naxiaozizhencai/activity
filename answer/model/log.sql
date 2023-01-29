CREATE TABLE `answer_log` (
                              `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                              `uid` varchar(50) NOT NULL,
                              `languge` varchar(50) NOT NULL,
                              `answer_id` int(11) NOT NULL,
                              `user_result` varchar(20) NOT NULL,
                              `answer_result` varchar(20) NOT NULL,
                              `log_time` datetime NOT NULL,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `page_log` (
                            `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                            `uid` varchar(50) NOT NULL,
                            `page_url` varchar(200) NOT NULL,
                            `ip` varchar(50) NOT NULL,
                            `log_time` datetime NOT NULL,
                            PRIMARY KEY (`id`),
                            KEY `uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `login_log` (
                             `id` int(20) NOT NULL AUTO_INCREMENT,
                             `language` varchar(50) NOT NULL,
                             `uid` varchar(50) NOT NULL,
                             `login_time` datetime NOT NULL,
                             `log_time` datetime NOT NULL,
                             PRIMARY KEY (`id`),
                             KEY `uid` (`uid`),
                             KEY `login_time` (`login_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;


CREATE TABLE `reward_log` (
                              `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                              `anwser_id` int(11) NOT NULL,
                              `item_id` varchar(50) NOT NULL,
                              `item_num` int(11) NOT NULL,
                              `uid` varchar(50) NOT NULL,
                              `log_time` datetime NOT NULL,
                              PRIMARY KEY (`id`),
                              KEY `log_time` (`log_time`),
                              KEY `uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;