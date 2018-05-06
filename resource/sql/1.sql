CREATE TABLE "fronius"."current_io_state" (
  "id"         SERIAL8,
  "pin_number" INTEGER      NOT NULL,
  "function"   VARCHAR(255) NOT NULL,
  "type"       VARCHAR(255) NOT NULL,
  "direction"  VARCHAR(255) NOT NULL,
  "set"        BOOL         NOT NULL DEFAULT FALSE,
  "created_at" timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT "pk_current_io_state" PRIMARY KEY ("id")
);
