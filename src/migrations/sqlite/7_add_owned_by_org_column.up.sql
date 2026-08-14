ALTER TABLE devices ADD COLUMN org_id integer NOT NULL;
ALTER TABLE interfaces ADD COLUMN org_id integer NOT NULL;
ALTER TABLE networks ADD COLUMN org_id integer NOT NULL;