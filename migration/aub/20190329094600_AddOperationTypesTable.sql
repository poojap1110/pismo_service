
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
SET @@session.sql_mode='';

CREATE TABLE IF NOT EXISTS `operation_types`  (
`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
`operation_id` varchar(32) DEFAULT NULL,
`description` varchar(64) DEFAULT NULL,
`date_added` timestamp NOT NULL DEFAULT '1981-01-01 01:01:01',
`date_updated` timestamp NOT NULL DEFAULT '1981-01-01 01:01:01',
PRIMARY KEY (id),
KEY `OPERATION_IDX` (`operation_id`),
KEY `DESCRIPTION_IDX` (`description`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `operation_types`;