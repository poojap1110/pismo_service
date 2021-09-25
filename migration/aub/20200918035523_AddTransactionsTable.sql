
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
SET @@session.sql_mode='';

CREATE TABLE IF NOT EXISTS `transactions`  (
`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
`transaction_id` varchar(32) DEFAULT NULL,
`account_id` varchar(64) DEFAULT NULL,
`operation_type_id` timestamp NOT NULL DEFAULT '1981-01-01 01:01:01',
`amount` timestamp NOT NULL DEFAULT '1981-01-01 01:01:01',
`event_date` timestamp NOT NULL DEFAULT '1981-01-01 01:01:01',
`event_updated` timestamp NOT NULL DEFAULT '1981-01-01 01:01:01',
PRIMARY KEY (id),
KEY `TRANSACTION_IDX` (`transaction_id`),
KEY `ACCOUNT_IDX` (`account_id`),
KEY `TRANSACTION_ACCOUNT_IDX` (`transaction_account_id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `transactions`;