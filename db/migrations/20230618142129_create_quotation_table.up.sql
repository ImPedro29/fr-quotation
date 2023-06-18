CREATE TABLE quotations
(
    id         serial
        PRIMARY KEY
        UNIQUE,
    carrier    TEXT    NOT NULL,
    price      NUMERIC NOT NULL,
    days       NUMERIC NOT NULL,
    service    TEXT    NOT NULl,
    created_at TIME DEFAULT now()
);

CREATE INDEX quotations_carrier_index
    ON quotations (carrier);

CREATE INDEX quotations_price_index
    ON quotations (price);

CREATE INDEX quotations_created_at_index
    ON quotations (created_at);

