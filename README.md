# mobydick

GOOS=linux CGO_ENABLED=0 go build -ldflags '-s -w' -a -installsuffix cgo -o -o main.go 
    zip function.zip main  

chmod 644 $(find ./mobydick -type f)
chmod 755 $(find ./mobydick -type d)

SELECT "block_number" FROM public."block_number" ORDER BY "block_number" DESC;
SELECT "timestamp", token, amount, "from", "to", "block_number", "tx_hash"
	FROM public.transfer ORDER BY timestamp DESC;

CREATE TABLE IF NOT EXISTS public.transfer
(
    "timestamp" timestamp without time zone,
    token text,
    amount bigint,
    "from" text,
    "to" text,
    "block_number" text,
    "tx_hash" text
);

ALTER TABLE public.transfer
    OWNER to {db_username};

CREATE TABLE IF NOT EXISTS public."block_number"
(
   "block_number" bigint
);

ALTER TABLE public."block_number"
   OWNER to {db_username};