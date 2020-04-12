CREATE TABLE public.shortener_mst_url (
	id bigserial NOT NULL PRIMARY KEY,
	short_url varchar(250) NOT NULL DEFAULT '',
	long_url varchar(250) NOT NULL DEFAULT '',
	created_at timestamp NOT NULL DEFAULT now(),
	expire_at timestamp NOT NULL,
	created_by varchar(250) NULL DEFAULT 'anonymous',
	CONSTRAINT short_url_unique UNIQUE (short_url)
);
CREATE INDEX shortener_mst_url_short_url_idx ON public.shortener_mst_url (short_url);

ALTER TABLE public.shortener_mst_url OWNER TO postgres;
GRANT ALL ON TABLE public.shortener_mst_url TO postgres;
