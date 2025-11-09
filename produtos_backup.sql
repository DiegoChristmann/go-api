--
-- PostgreSQL database dump
--

\restrict LCAeaskkr5w76sffsmNUOrffrsxD2XaX26gTOV0nt6s8CDsL2xgbS2xGi4MDwoR

-- Dumped from database version 17.6 (Homebrew)
-- Dumped by pg_dump version 17.6 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Data for Name: product; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.product (id, product_name, price) VALUES (1, 'Sushi', 100.00);
INSERT INTO public.product (id, product_name, price) VALUES (2, 'Pizza', 100.00);
INSERT INTO public.product (id, product_name, price) VALUES (3, 'Espetinho', 10.00);


--
-- Name: product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.product_id_seq', 3, true);


--
-- PostgreSQL database dump complete
--

\unrestrict LCAeaskkr5w76sffsmNUOrffrsxD2XaX26gTOV0nt6s8CDsL2xgbS2xGi4MDwoR

