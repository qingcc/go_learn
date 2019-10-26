CREATE TABLE IF NOT EXISTS "coin_price1_min" (
	"id" SERIAL PRIMARY KEY NOT NULL,
	"coin_id" INT8 NOT NULL COMMENT "货币id",
	"high" NUMERIC(12, 2) DEFAULT 0 NOT NULL COMMENT "最高价",
	"low" NUMERIC(12, 2) DEFAULT 0 NOT NULL COMMENT "最低价",
	"open" NUMERIC(12, 2) DEFAULT 0 NOT NULL COMMENT "开盘价",
	"close" NUMERIC(12, 2) DEFAULT 0 NOT NULL COMMENT "收盘价",
	"count" NUMERIC(18, 8) DEFAULT 0 NOT NULL COMMENT "交易次数",
	"amount" NUMERIC(18, 8) DEFAULT 0 NOT NULL COMMENT "以基础币种计量的交易量",
	"vol" NUMERIC(12, 2) DEFAULT 0 NOT NULL COMMENT "以报价币种计量的交易量",
	"grow" NUMERIC(12, 2) DEFAULT 0 NOT NULL COMMENT "价格增长额度",
	"rise" NUMERIC(12, 2) DEFAULT 0 NOT NULL COMMENT "价格涨幅",
	"time" NUMERIC(12, 2) DEFAULT 0 NOT NULL COMMENT "插入时间",
	"updated_at" TIMESTAMP NULL
);

ALTER TABLE coin_price1_min
ALTER id TYPE INT8;

-- CREATE TABLE IF NOT EXISTS "test_1" (
-- 	"id" INT8 PRIMARY KEY NOT NULL,
-- 	"email" VARCHAR(255) NOT NULL,
-- 	"name" VARCHAR(1000) NULL,
-- 	"price" NUMERIC(12, 2) DEFAULT 0,
-- 	"new" DECIMAL(12, 2),
-- 	"is_lock" BOOL DEFAULT false NULL,
-- 	"describe" VARCHAR(255) NULL,
-- 	"num" INT8 NULL,
-- 	"created_at" TIMESTAMP NULL,
-- 	"updated_at" TIMESTAMP NULL
-- );
--
-- CREATE SEQUENCE public.test_1_seq_id
--     INCREMENT 1
--     START 1
--     MINVALUE 1
--     MAXVALUE 9223372036854775807
--     CACHE 1;
--
-- ALTER SEQUENCE public.test_1_seq_id
--     OWNER TO postgres;

CREATE TABLE IF NOT EXISTS "test_ta" (
	"id" SERIAL PRIMARY KEY NOT NULL,
	"email" VARCHAR(255) NOT NULL,
	"name" VARCHAR(1000) NULL,
	"price" NUMERIC(12, 2) DEFAULT 0,
	"new" DECIMAL(12, 2),
	"is_lock" BOOL DEFAULT false NULL,
	"describe" VARCHAR(255) NULL,
	"num" INT8 NULL,
	"created_at" TIMESTAMP NULL,
	"updated_at" TIMESTAMP NULL
);

ALTER TABLE test_ta
ALTER id TYPE INT8;