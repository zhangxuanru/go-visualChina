# Host: 127.0.0.1  (Version: 5.6.11)
# Date: 2018-10-14 17:31:01
# Generator: MySQL-Front 5.3  (Build 4.214)

/*!40101 SET NAMES utf8 */;

#
# Structure for table "visual_page_nav"
#

DROP TABLE IF EXISTS `visual_page_nav`;
CREATE TABLE `visual_page_nav` (
  `Id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `type` tinyint(3) unsigned DEFAULT NULL COMMENT '0:编辑图片, 1:创意壁纸 2:创意图片 3:设计素材',
  `nav_title` varchar(50) DEFAULT NULL,
  `nav_url` varchar(255) DEFAULT NULL,
  `nav_crc32` int(11) unsigned DEFAULT NULL,
  `add_date` int(11) DEFAULT NULL,
  `last_crawl_time` int(11) unsigned DEFAULT NULL COMMENT '最后抓取时间',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='所有栏目页-导航表';
