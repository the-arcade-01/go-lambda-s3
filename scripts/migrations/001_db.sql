CREATE TABLE bse_bhavcopy (
	code BIGINT NOT NULL,
	name VARCHAR(255) NOT NULL,
	open Decimal(10,2) NOT NULL DEFAULT 0,
	high Decimal(10,2) NOT NULL DEFAULT 0,
	low Decimal(10,2) NOT NULL DEFAULT 0,
	close Decimal(10,2) NOT NULL DEFAULT 0,
	pclose Decimal(10,2) NOT NULL DEFAULT 0,
	isin VARCHAR(60) NOT NULL,
	date DATETIME NOT NULL,
	PRIMARY KEY (code, isin)
)
