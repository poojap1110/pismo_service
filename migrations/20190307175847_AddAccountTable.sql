
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
SET @@session.sql_mode='';

CREATE TABLE IF NOT EXISTS `accounts`  (
`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
`account_hash` varchar(32) DEFAULT NULL,
`document_numer` varchar(64) DEFAULT NULL,
`date_added` timestamp NOT NULL DEFAULT '1981-01-01 01:01:01',
`date_updated` timestamp NOT NULL DEFAULT '1981-01-01 01:01:01',
PRIMARY KEY (id),
KEY `ACCOUNT_DOCUMENT_IDX` (`account_id`, `document_numer`),
KEY `ACCOUNT_IDX` (`account_id`),
KEY `DOCUMENT_IDX` (`document_numer`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `accounts`;