CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `app_key` varchar(20) DEFAULT '' COMMENT 'Key',
  `app_secret` varchar(50) DEFAULT '' COMMENT 'Secret',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'cjsj',
  `created_by` varchar(100) DEFAULT '' COMMENT 'cjr',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'xgsj',
  `modified_by` varchar(100) DEFAULT '' COMMENT 'xgr',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'scsj',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '0:no 1:yes',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='认证管理';