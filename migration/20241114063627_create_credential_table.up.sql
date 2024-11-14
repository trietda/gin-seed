BEGIN;

CREATE TABLE IF NOT EXISTS "credentials" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "user_id" uuid REFERENCES "users"("id"),
  "username" varchar(32),
  "password" varchar(255),
  "created_at" timestamp DEFAULT now(),
  "updated_at" timestamp DEFAULT now()
);

COMMIT;

