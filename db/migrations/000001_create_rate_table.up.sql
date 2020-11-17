CREATE TABLE rates
(
    id SERIAL,
    sell_currency character varying(50) COLLATE pg_catalog."default" NOT NULL,
    buy_currency character varying(50) COLLATE pg_catalog."default" NOT NULL,
    rate numeric(10,0) NOT NULL,
    created_on TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
)