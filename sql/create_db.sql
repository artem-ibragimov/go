-- Table: public.brand
-- DROP TABLE IF EXISTS public.brand;
CREATE TABLE
   IF NOT EXISTS PUBLIC.brand (
      id INTEGER NOT NULL DEFAULT NEXTVAL('brand_id_seq'::regclass),
      NAME CHARACTER VARYING(50) COLLATE pg_catalog."default" NOT NULL,
      CONSTRAINT brand_pkey PRIMARY KEY (id),
      CONSTRAINT brand_name_key UNIQUE (NAME)
   ) TABLESPACE pg_default;

ALTER TABLE IF EXISTS PUBLIC.brand OWNER TO artem;

-- Table: public.component_category
-- DROP TABLE IF EXISTS public.component_category;
CREATE TABLE
   IF NOT EXISTS PUBLIC.component_category (
      id INTEGER NOT NULL DEFAULT NEXTVAL('component_category_id_seq'::regclass),
      NAME CHARACTER VARYING(50) COLLATE pg_catalog."default" NOT NULL,
      CONSTRAINT component_category_pkey PRIMARY KEY (id)
   ) TABLESPACE pg_default;

ALTER TABLE
   IF EXISTS PUBLIC.component_category OWNER TO artem;

-- Table: public.countries
-- DROP TABLE IF EXISTS public.countries;
CREATE TABLE
   IF NOT EXISTS PUBLIC.countries (
      id INTEGER NOT NULL DEFAULT NEXTVAL('countries_id_seq'::regclass),
      NAME CHARACTER VARYING(50) COLLATE pg_catalog."default" NOT NULL,
      CONSTRAINT countries_pkey PRIMARY KEY (id),
      CONSTRAINT countries_name_key UNIQUE (NAME)
   ) TABLESPACE pg_default;

ALTER TABLE IF EXISTS PUBLIC.countries OWNER TO artem;

-- Table: public.defect_category
-- DROP TABLE IF EXISTS public.defect_category;
CREATE TABLE
   IF NOT EXISTS PUBLIC.defect_category (
      id INTEGER NOT NULL DEFAULT NEXTVAL('defect_category_id_seq'::regclass),
      NAME CHARACTER VARYING(150) COLLATE pg_catalog."default" NOT NULL,
      CONSTRAINT defect_category_pkey PRIMARY KEY (id)
   ) TABLESPACE pg_default;

ALTER TABLE
   IF EXISTS PUBLIC.defect_category OWNER TO artem;

-- Table: public.engines
-- DROP TABLE IF EXISTS public.engines;
CREATE TABLE
   IF NOT EXISTS PUBLIC.engines (
      id INTEGER NOT NULL DEFAULT NEXTVAL('engines_id_seq'::regclass),
      NAME CHARACTER VARYING(50) COLLATE pg_catalog."default" NOT NULL,
      displacement INTEGER,
      config CHARACTER VARYING(50) COLLATE pg_catalog."default",
      valves CHARACTER VARYING(50) COLLATE pg_catalog."default",
      aspiration CHARACTER VARYING(50) COLLATE pg_catalog."default",
      fuel_type CHARACTER VARYING(50) COLLATE pg_catalog."default",
      torque INTEGER,
      power_hp INTEGER,
      CONSTRAINT engines_pkey PRIMARY KEY (id),
      --  CONSTRAINT engines_name_key UNIQUE (name)
   ) TABLESPACE pg_default;

ALTER TABLE IF EXISTS PUBLIC.engines OWNER TO artem;

-- Table: public.transmissions
-- DROP TABLE IF EXISTS public.transmissions;
CREATE TABLE
   IF NOT EXISTS PUBLIC.transmissions (
      id INTEGER NOT NULL DEFAULT NEXTVAL('transmissions_id_seq'::regclass),
      NAME CHARACTER VARYING(50) COLLATE pg_catalog."default" NOT NULL,
      component_category_id INTEGER NOT NULL,
      CONSTRAINT transmissions_pkey PRIMARY KEY (id),
      CONSTRAINT transmissions_component_category_id_fkey FOREIGN KEY (component_category_id) REFERENCES PUBLIC.component_category (id) MATCH SIMPLE
      ON
      UPDATE NO ACTION
         ON
      DELETE CASCADE
   ) TABLESPACE pg_default;

ALTER TABLE
   IF EXISTS PUBLIC.transmissions OWNER TO artem;

-- Table: public.model
-- DROP TABLE IF EXISTS public.model;
CREATE TABLE
   IF NOT EXISTS PUBLIC.model (
      id INTEGER NOT NULL DEFAULT NEXTVAL('model_id_seq'::regclass),
      NAME CHARACTER VARYING(50) COLLATE pg_catalog."default" NOT NULL,
      VERSION CHARACTER VARYING(50) COLLATE pg_catalog."default",
      brand_id INTEGER NOT NULL,
      YEAR INTEGER,
      transmission_id INTEGER,
      engine_id INTEGER,
      CONSTRAINT model_pkey PRIMARY KEY (id),
      CONSTRAINT model_name_version_key UNIQUE (NAME, VERSION),
      CONSTRAINT model_brand_id_fkey FOREIGN KEY (brand_id) REFERENCES PUBLIC.brand (id) MATCH SIMPLE
      ON
      UPDATE NO ACTION
         ON
      DELETE
         CASCADE,
         CONSTRAINT model_engine_id_fkey FOREIGN KEY (engine_id) REFERENCES PUBLIC.engines (id) MATCH SIMPLE
         ON
      UPDATE NO ACTION
         ON
      DELETE
         CASCADE NOT VALID,
         CONSTRAINT model_transmission_id_fkey FOREIGN KEY (transmission_id) REFERENCES PUBLIC.transmissions (id) MATCH SIMPLE
         ON
      UPDATE NO ACTION
         ON
      DELETE CASCADE NOT VALID
   ) TABLESPACE pg_default;

ALTER TABLE IF EXISTS PUBLIC.model OWNER TO artem;

-- Table: public.defects
-- DROP TABLE IF EXISTS public.defects;
CREATE TABLE
   IF NOT EXISTS PUBLIC.defects (
      id INTEGER NOT NULL DEFAULT NEXTVAL('defects_id_seq'::regclass),
      brand_id INTEGER NOT NULL,
      model_id INTEGER NOT NULL,
      country_id INTEGER NOT NULL,
      defect_category_id INTEGER NOT NULL,
      production_year INTEGER NOT NULL,
      YEAR INTEGER NOT NULL,
      mileage INTEGER NOT NULL,
      description TEXT COLLATE pg_catalog."default",
      COST money,
      rating NUMERIC,
      CONSTRAINT defects_pkey PRIMARY KEY (id),
      CONSTRAINT defects_brand_id_fkey FOREIGN KEY (brand_id) REFERENCES PUBLIC.brand (id) MATCH SIMPLE
      ON
      UPDATE NO ACTION
         ON
      DELETE
         CASCADE,
         CONSTRAINT defects_country_id_fkey FOREIGN KEY (country_id) REFERENCES PUBLIC.countries (id) MATCH SIMPLE
         ON
      UPDATE NO ACTION
         ON
      DELETE
         CASCADE,
         CONSTRAINT defects_defect_category_id_fkey FOREIGN KEY (defect_category_id) REFERENCES PUBLIC.defect_category (id) MATCH SIMPLE
         ON
      UPDATE NO ACTION
         ON
      DELETE
         CASCADE,
         CONSTRAINT defects_model_id_fkey FOREIGN KEY (model_id) REFERENCES PUBLIC.model (id) MATCH SIMPLE
         ON
      UPDATE NO ACTION
         ON
      DELETE CASCADE
   ) TABLESPACE pg_default;

ALTER TABLE IF EXISTS PUBLIC.defects OWNER TO artem;