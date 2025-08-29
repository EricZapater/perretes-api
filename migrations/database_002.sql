-- Table: public.action_logs

-- DROP TABLE IF EXISTS public.action_logs;

CREATE TABLE IF NOT EXISTS public.action_logs
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    user_id uuid,
    action_type character varying(100) COLLATE pg_catalog."default" NOT NULL,
    metadata jsonb,
    timezone character varying(100) COLLATE pg_catalog."default",
    performed_at timestamp with time zone DEFAULT now(),
    CONSTRAINT action_logs_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.action_logs
    OWNER to ubuntu;
-- Index: idx_action_type

-- DROP INDEX IF EXISTS public.idx_action_type;

CREATE INDEX IF NOT EXISTS idx_action_type
    ON public.action_logs USING btree
    (action_type COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;
-- Index: idx_action_user

-- DROP INDEX IF EXISTS public.idx_action_user;

CREATE INDEX IF NOT EXISTS idx_action_user
    ON public.action_logs USING btree
    (user_id ASC NULLS LAST)
    TABLESPACE pg_default;