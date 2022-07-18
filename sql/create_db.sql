--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2
-- Dumped by pg_dump version 14.2

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
-- Name: brand; Type: TABLE; Schema: public; Owner: artem
--

CREATE TABLE public.brand (
    id integer NOT NULL,
    name character varying(50) NOT NULL
);


ALTER TABLE public.brand OWNER TO artem;

--
-- Name: brand_id_seq; Type: SEQUENCE; Schema: public; Owner: artem
--

CREATE SEQUENCE public.brand_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.brand_id_seq OWNER TO artem;

--
-- Name: brand_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: artem
--

ALTER SEQUENCE public.brand_id_seq OWNED BY public.brand.id;


--
-- Name: country; Type: TABLE; Schema: public; Owner: artem
--

CREATE TABLE public.country (
    id integer NOT NULL,
    name character varying(50) NOT NULL
);


ALTER TABLE public.country OWNER TO artem;

--
-- Name: countries_id_seq; Type: SEQUENCE; Schema: public; Owner: artem
--

CREATE SEQUENCE public.countries_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.countries_id_seq OWNER TO artem;

--
-- Name: countries_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: artem
--

ALTER SEQUENCE public.countries_id_seq OWNED BY public.country.id;


--
-- Name: defect; Type: TABLE; Schema: public; Owner: artem
--

CREATE TABLE public.defect (
    id integer NOT NULL,
    brand_id integer NOT NULL,
    model_id integer NOT NULL,
    country_id integer NOT NULL,
    mileage integer,
    description text,
    cost numeric,
    rating numeric,
    major_category_id integer NOT NULL,
    minor_category_id integer NOT NULL,
    category_id integer NOT NULL,
    freq integer,
    year integer NOT NULL
);


ALTER TABLE public.defect OWNER TO artem;

--
-- Name: defect_category; Type: TABLE; Schema: public; Owner: artem
--

CREATE TABLE public.defect_category (
    id integer NOT NULL,
    name character varying(150) NOT NULL
);


ALTER TABLE public.defect_category OWNER TO artem;

--
-- Name: defect_category_id_seq; Type: SEQUENCE; Schema: public; Owner: artem
--

CREATE SEQUENCE public.defect_category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.defect_category_id_seq OWNER TO artem;

--
-- Name: defect_category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: artem
--

ALTER SEQUENCE public.defect_category_id_seq OWNED BY public.defect_category.id;


--
-- Name: defects_id_seq; Type: SEQUENCE; Schema: public; Owner: artem
--

CREATE SEQUENCE public.defects_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.defects_id_seq OWNER TO artem;

--
-- Name: defects_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: artem
--

ALTER SEQUENCE public.defects_id_seq OWNED BY public.defect.id;


--
-- Name: engine; Type: TABLE; Schema: public; Owner: artem
--

CREATE TABLE public.engine (
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    displacement smallint,
    fuel_type character varying(50),
    torque smallint,
    power_hp smallint,
    valves smallint,
    cylinders smallint,
    img text
);


ALTER TABLE public.engine OWNER TO artem;

--
-- Name: engines_id_seq; Type: SEQUENCE; Schema: public; Owner: artem
--

CREATE SEQUENCE public.engines_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.engines_id_seq OWNER TO artem;

--
-- Name: engines_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: artem
--

ALTER SEQUENCE public.engines_id_seq OWNED BY public.engine.id;


--
-- Name: generation; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.generation (
    id integer NOT NULL,
    name character varying NOT NULL,
    model_id integer NOT NULL,
    img text,
    start smallint,
    finish smallint
);


ALTER TABLE public.generation OWNER TO postgres;

--
-- Name: generation_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.generation_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.generation_id_seq OWNER TO postgres;

--
-- Name: generation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.generation_id_seq OWNED BY public.generation.id;


--
-- Name: model; Type: TABLE; Schema: public; Owner: artem
--

CREATE TABLE public.model (
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    brand_id integer NOT NULL
);


ALTER TABLE public.model OWNER TO artem;

--
-- Name: model_id_seq; Type: SEQUENCE; Schema: public; Owner: artem
--

CREATE SEQUENCE public.model_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.model_id_seq OWNER TO artem;

--
-- Name: model_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: artem
--

ALTER SEQUENCE public.model_id_seq OWNED BY public.model.id;


--
-- Name: transmission; Type: TABLE; Schema: public; Owner: artem
--

CREATE TABLE public.transmission (
    id integer NOT NULL,
    consumption numeric(5,2),
    acceleration numeric(5,2),
    name character varying,
    brand_id integer NOT NULL,
    gears smallint
);


ALTER TABLE public.transmission OWNER TO artem;

--
-- Name: transmissions_id_seq; Type: SEQUENCE; Schema: public; Owner: artem
--

CREATE SEQUENCE public.transmissions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transmissions_id_seq OWNER TO artem;

--
-- Name: transmissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: artem
--

ALTER SEQUENCE public.transmissions_id_seq OWNED BY public.transmission.id;


--
-- Name: version; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.version (
    id integer NOT NULL,
    name character varying NOT NULL,
    generation_id integer NOT NULL,
    transmission_id integer,
    engine_id integer
);


ALTER TABLE public.version OWNER TO postgres;

--
-- Name: version_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.version_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.version_id_seq OWNER TO postgres;

--
-- Name: version_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.version_id_seq OWNED BY public.version.id;


--
-- Name: brand id; Type: DEFAULT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.brand ALTER COLUMN id SET DEFAULT nextval('public.brand_id_seq'::regclass);


--
-- Name: country id; Type: DEFAULT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.country ALTER COLUMN id SET DEFAULT nextval('public.countries_id_seq'::regclass);


--
-- Name: defect id; Type: DEFAULT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.defect ALTER COLUMN id SET DEFAULT nextval('public.defects_id_seq'::regclass);


--
-- Name: defect_category id; Type: DEFAULT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.defect_category ALTER COLUMN id SET DEFAULT nextval('public.defect_category_id_seq'::regclass);


--
-- Name: engine id; Type: DEFAULT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.engine ALTER COLUMN id SET DEFAULT nextval('public.engines_id_seq'::regclass);


--
-- Name: generation id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.generation ALTER COLUMN id SET DEFAULT nextval('public.generation_id_seq'::regclass);


--
-- Name: model id; Type: DEFAULT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.model ALTER COLUMN id SET DEFAULT nextval('public.model_id_seq'::regclass);


--
-- Name: transmission id; Type: DEFAULT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.transmission ALTER COLUMN id SET DEFAULT nextval('public.transmissions_id_seq'::regclass);


--
-- Name: version id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.version ALTER COLUMN id SET DEFAULT nextval('public.version_id_seq'::regclass);


--
-- Name: brand brand_name_key; Type: CONSTRAINT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.brand
    ADD CONSTRAINT brand_name_key UNIQUE (name);


--
-- Name: brand brand_pkey; Type: CONSTRAINT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.brand
    ADD CONSTRAINT brand_pkey PRIMARY KEY (id);


--
-- Name: country countries_name_key; Type: CONSTRAINT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.country
    ADD CONSTRAINT countries_name_key UNIQUE (name);


--
-- Name: country countries_pkey; Type: CONSTRAINT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.country
    ADD CONSTRAINT countries_pkey PRIMARY KEY (id);


--
-- Name: defect defect_brand_id_description_model_id_key; Type: CONSTRAINT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.defect
    ADD CONSTRAINT defect_brand_id_description_model_id_key UNIQUE (brand_id, description, model_id);


--
-- Name: defect_category defect_category_name_key; Type: CONSTRAINT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.defect_category
    ADD CONSTRAINT defect_category_name_key UNIQUE (name);


--
-- Name: defect_category defect_category_pkey; Type: CONSTRAINT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.defect_category
    ADD CONSTRAINT defect_category_pkey PRIMARY KEY (id);


--
-- Name: defect defects_pkey; Type: CONSTRAINT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.defect
    ADD CONSTRAINT defects_pkey PRIMARY KEY (id);


--
-- Name: engine engine_name_key; Type: CONSTRAINT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.engine
    ADD CONSTRAINT engine_name_key UNIQUE (name);


--
-- Name: engine engines_pkey; Type: CONSTRAINT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.engine
    ADD CONSTRAINT engines_pkey PRIMARY KEY (id);


--
-- Name: generation generation_name_model_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.generation
    ADD CONSTRAINT generation_name_model_id_key UNIQUE (name, model_id);


--
-- Name: generation generation_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.generation
    ADD CONSTRAINT generation_pkey PRIMARY KEY (id);


--
-- Name: model model_name_brand_id_key; Type: CONSTRAINT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.model
    ADD CONSTRAINT model_name_brand_id_key UNIQUE (name, brand_id);


--
-- Name: model model_pkey; Type: CONSTRAINT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.model
    ADD CONSTRAINT model_pkey PRIMARY KEY (id);


--
-- Name: transmission transmissions_pkey; Type: CONSTRAINT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.transmission
    ADD CONSTRAINT transmissions_pkey PRIMARY KEY (id);


--
-- Name: version version_name_genaration_id_transmission_id_engine_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.version
    ADD CONSTRAINT version_name_genaration_id_transmission_id_engine_id_key UNIQUE (name, generation_id, transmission_id, engine_id);


--
-- Name: version version_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.version
    ADD CONSTRAINT version_pkey PRIMARY KEY (id);


--
-- Name: defect defect_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.defect
    ADD CONSTRAINT defect_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.defect_category(id) NOT VALID;


--
-- Name: defect defect_major_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.defect
    ADD CONSTRAINT defect_major_category_id_fkey FOREIGN KEY (major_category_id) REFERENCES public.defect_category(id) NOT VALID;


--
-- Name: defect defect_minor_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.defect
    ADD CONSTRAINT defect_minor_category_id_fkey FOREIGN KEY (minor_category_id) REFERENCES public.defect_category(id) NOT VALID;


--
-- Name: defect defects_brand_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.defect
    ADD CONSTRAINT defects_brand_id_fkey FOREIGN KEY (brand_id) REFERENCES public.brand(id) ON DELETE CASCADE;


--
-- Name: defect defects_country_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.defect
    ADD CONSTRAINT defects_country_id_fkey FOREIGN KEY (country_id) REFERENCES public.country(id) ON DELETE CASCADE;


--
-- Name: defect defects_model_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.defect
    ADD CONSTRAINT defects_model_id_fkey FOREIGN KEY (model_id) REFERENCES public.model(id) ON DELETE CASCADE;


--
-- Name: generation generation_model_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.generation
    ADD CONSTRAINT generation_model_id_fkey FOREIGN KEY (model_id) REFERENCES public.model(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: model model_brand_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.model
    ADD CONSTRAINT model_brand_id_fkey FOREIGN KEY (brand_id) REFERENCES public.brand(id) ON DELETE CASCADE;


--
-- Name: transmission transmission_brand_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: artem
--

ALTER TABLE ONLY public.transmission
    ADD CONSTRAINT transmission_brand_id_fkey FOREIGN KEY (brand_id) REFERENCES public.brand(id) NOT VALID;


--
-- Name: version version_engine_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.version
    ADD CONSTRAINT version_engine_id_fkey FOREIGN KEY (engine_id) REFERENCES public.engine(id) ON DELETE CASCADE NOT VALID;


--
-- Name: version version_generation_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.version
    ADD CONSTRAINT version_generation_id_fkey FOREIGN KEY (generation_id) REFERENCES public.generation(id) ON DELETE CASCADE NOT VALID;


--
-- Name: version version_transmission_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.version
    ADD CONSTRAINT version_transmission_id_fkey FOREIGN KEY (transmission_id) REFERENCES public.transmission(id) ON DELETE CASCADE NOT VALID;


--
-- PostgreSQL database dump complete
--

