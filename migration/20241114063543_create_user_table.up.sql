BEGIN;

CREATE TABLE IF NOT EXISTS "users" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "created_at" timestamp DEFAULT now(),
  "updated_at" timestamp DEFAULT now()
);

COMMIT;
