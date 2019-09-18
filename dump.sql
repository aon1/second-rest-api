--
-- PostgreSQL database dump
--

-- Dumped from database version 11.5
-- Dumped by pg_dump version 11.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: author; Type: TABLE; Schema: public; Owner: arlen
--

CREATE TABLE public.author (
    id integer NOT NULL,
    name text
);


ALTER TABLE public.author OWNER TO arlen;

--
-- Name: author_id_seq; Type: SEQUENCE; Schema: public; Owner: arlen
--

CREATE SEQUENCE public.author_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.author_id_seq OWNER TO arlen;

--
-- Name: author_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: arlen
--

ALTER SEQUENCE public.author_id_seq OWNED BY public.author.id;


--
-- Name: books; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.books (
    isbn text NOT NULL,
    title text NOT NULL,
    price double precision NOT NULL,
    author_id integer
);


ALTER TABLE public.books OWNER TO postgres;

--
-- Name: author id; Type: DEFAULT; Schema: public; Owner: arlen
--

ALTER TABLE ONLY public.author ALTER COLUMN id SET DEFAULT nextval('public.author_id_seq'::regclass);


--
-- Data for Name: author; Type: TABLE DATA; Schema: public; Owner: arlen
--

COPY public.author (id, name) FROM stdin;
1	Clarice Lispector
\.


--
-- Data for Name: books; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.books (isbn, title, price, author_id) FROM stdin;
123	Crepusculo	12	1
\.


--
-- Name: author_id_seq; Type: SEQUENCE SET; Schema: public; Owner: arlen
--

SELECT pg_catalog.setval('public.author_id_seq', 1, true);


--
-- Name: author author_pkey; Type: CONSTRAINT; Schema: public; Owner: arlen
--

ALTER TABLE ONLY public.author
    ADD CONSTRAINT author_pkey PRIMARY KEY (id);


--
-- Name: books books_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (isbn);


--
-- Name: books author_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT author_id_fk FOREIGN KEY (author_id) REFERENCES public.author(id);


--
-- PostgreSQL database dump complete
--

