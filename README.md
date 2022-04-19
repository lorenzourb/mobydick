# mobydick

GOOS=linux CGO_ENABLED=0 go build main.go 
    zip lambda.zip main  

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


SELECT "timestamp", token, amount, "from", "to", block_number, tx_hash
	FROM public.transfer WHERE "from" IN     ('0x000000000000000000000000187e3534f461d7c59a7d6899a983a5305b48f93f','0x00000000000000000000000066e092fd00c4e4eb5bd20f5392c1902d738ae7bc','0x00000000000000000000000049a2dcc237a65cc1f412ed47e0594602f6141936','0x000000000000000000000000b60c61dbb7456f024f9338c739b02be68e3f545c',
'0x0000000000000000000000005770815b0c2a09a43c9e5aecb7e2f3886075b605','0x000000000000000000000000530e0a6993ea99ffc96615af43f327225a5fe536',
'0x000000000000000000000000d628f7c481c7dd87f674870bec5d7a311fb1d9a2','0x000000000000000000000000203520f4ec42ea39b03f62b20e20cf17db5fdfa7',
'0x0000000000000000000000009b0c45d46d386cedd98873168c36efd0dcba8d46','0x000000000000000000000000d49a1525b46f9149ff665807c925bd83b5a6d7e3',
'0x0000000000000000000000000548f59fee79f8832c299e01dca5c76f034f558e','0x000000000000000000000000286af5cf60ae834199949bbc815485f07cc9c644',
'0x000000000000000000000000dc1664458d2f0b6090bea60a8793a4e66c2f1c00','0x0000000000000000000000007d812b62dc15e6f4073eba8a2ba8db19c4e40704',
'0x00000000000000000000000015abb66ba754f05cbc0165a64a11cded1543de48','0x00000000000000000000000072a53cdbbcc1b9efa39c834a540550e23463aacb','0x0000000000000000000000003d8fc1cffaa110f7a7f9f8bc237b73d54c4abf61','0x0000000000000000000000001157a2076b9bb22a85cc2c162f20fab3898f4101');

UPDATE public.whales
	SET "name"='gwen wallet'
	WHERE address='0x0000000000000000000000007d812b62dc15e6f4073eba8a2ba8db19c4e40704';