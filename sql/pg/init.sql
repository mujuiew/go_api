-- SET statement_timeout = 0;
-- SET lock_timeout = 0;
-- SET idle_in_transaction_session_timeout = 0;
-- SET client_encoding = 'UTF8';
-- SET standard_conforming_strings = on;
-- SELECT pg_catalog.set_config('search_path', '', false);
-- SET check_function_bodies = false;
-- SET xmloption = content;
-- SET client_min_messages = warning;
-- SET row_security = off;

-- SET default_tablespace = '';

-- SET default_table_access_method = heap;

CREATE TABLE public.Promotion
(
    "promotion_name" character varying(30) COLLATE pg_catalog."default" NOT NULL,
    "description" character varying(50) COLLATE pg_catalog."default",
    start_date date COLLATE pg_catalog."default",
    end_date date COLLATE pg_catalog."default" ,
    CONSTRAINT "Promotion_pkey" PRIMARY KEY ("promotion_name")
)

TABLESPACE pg_default;

ALTER TABLE public.Promotion
    OWNER to postgres;

INSERT INTO Promotion(promotion_name, description, start_date, end_date)
VALUES ('Promo1', 'Rate < 10', '2020-01-01', '2020-03-31');
INSERT INTO Promotion(promotion_name, description, start_date, end_date)
VALUES ('Promo2', 'Rate > 10 < 2', '22020-04-01', '2020-06-30');
INSERT INTO Promotion(promotion_name, description, start_date, end_date)
VALUES ('Promo3', 'Rate > 20 < 28', '2020-07-01', '2020-12-30');