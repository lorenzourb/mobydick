# mobydick

GOOS=linux CGO_ENABLED=0 go build -ldflags '-s -w' -a -installsuffix cgo -o -o main.go 
zip function.zip main  

chmod 644 $(find ./mobydick -type f)
chmod 755 $(find ./mobydick -type d)

SELECT "blockNumber" FROM public."blockNumbers" ORDER BY "blockNumber" DESC;
SELECT "timestamp", token, amount, "from", "to", "blockNumber", "txHash"
	FROM public.transfers ORDER BY timestamp DESC;


CREATE TABLE IF NOT EXISTS public.transfers
(
    "timestamp" timestamp without time zone,
    token text,
    amount bigint,
    "from" text,
    "to" text,
    "blockNumber" text,
    "txHash" text
);

ALTER TABLE public.transfers
    OWNER to {db_username};

CREATE TABLE IF NOT EXISTS public."blockNumbers"
(
   "blockNumber" bigint
);

ALTER TABLE public."blockNumbers"
   OWNER to {db_username};