SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: linux_distro; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.linux_distro (
    linux_distro text NOT NULL
);


--
-- Name: linux_user; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.linux_user (
    linux_user_id bigint NOT NULL,
    linux_user_name text NOT NULL,
    linux_distro text NOT NULL
);


--
-- Name: linux_user_linux_user_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.linux_user ALTER COLUMN linux_user_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.linux_user_linux_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(255) NOT NULL
);


--
-- Name: linux_distro linux_distro_linux_distro_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.linux_distro
    ADD CONSTRAINT linux_distro_linux_distro_key UNIQUE (linux_distro);


--
-- Name: linux_user linux_user_linux_user_name_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.linux_user
    ADD CONSTRAINT linux_user_linux_user_name_key UNIQUE (linux_user_name);


--
-- Name: linux_user linux_user_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.linux_user
    ADD CONSTRAINT linux_user_pkey PRIMARY KEY (linux_user_id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: linux_user fk_distro_name; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.linux_user
    ADD CONSTRAINT fk_distro_name FOREIGN KEY (linux_distro) REFERENCES public.linux_distro(linux_distro);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20221110141229'),
    ('20221129182448');
