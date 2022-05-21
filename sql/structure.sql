CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS purchase (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    item character varying NOT NULL,
    price numeric NOT NULL,
    status character varying NOT NULL,
    process_key numeric NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS approval(
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    purchase_id uuid NOT NULL,
    username character varying NOT NULL,
    action character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

ALTER TABLE ONLY purchase
    ADD CONSTRAINT purchase_pkey PRIMARY KEY (id);

ALTER TABLE ONLY approval
    ADD CONSTRAINT approval_pkey PRIMARY KEY (id);

ALTER TABLE ONLY approval
    ADD CONSTRAINT purchase_fk FOREIGN KEY (purchase_id) REFERENCES purchase(id) ON UPDATE CASCADE ON DELETE CASCADE;

CREATE INDEX purchase_process_key_idx ON purchase USING btree (process_key);

CREATE INDEX approval_purchase_id_idx ON approval USING btree (purchase_id);