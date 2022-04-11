CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS campaign (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


CREATE TABLE IF NOT EXISTS donation (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    campaign_id uuid NOT NULL,
    amount numeric NOT NULL,
    status character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


CREATE TABLE IF NOT EXISTS wallet (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    campaign_id uuid NOT NULL,
    amount numeric NOT NULL,
    flow character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

ALTER TABLE ONLY campaign
    ADD CONSTRAINT campaign_pkey PRIMARY KEY (id);


ALTER TABLE ONLY donation
    ADD CONSTRAINT donation_pkey PRIMARY KEY (id);


ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


ALTER TABLE ONLY wallet
    ADD CONSTRAINT wallet_pkey PRIMARY KEY (id);


ALTER TABLE ONLY campaign
    ADD CONSTRAINT campaign_fk FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE;


ALTER TABLE ONLY donation
    ADD CONSTRAINT donation_fk FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY donation
    ADD CONSTRAINT campaign_fk FOREIGN KEY (campaign_id) REFERENCES public.campaign(id) ON DELETE CASCADE ON UPDATE CASCADE;


ALTER TABLE ONLY wallet
    ADD CONSTRAINT wallet_fk FOREIGN KEY (campaign_id) REFERENCES campaign(id) ON UPDATE CASCADE ON DELETE CASCADE;


CREATE INDEX campaign_user_id_idx ON campaign USING btree (user_id);

CREATE INDEX donation_user_id_idx ON donation USING btree (user_id);

CREATE INDEX donation_campaign_id_idx ON donation USING btree (campaign_id);

CREATE INDEX wallet_campaign_id_idx ON wallet USING btree (campaign_id);
