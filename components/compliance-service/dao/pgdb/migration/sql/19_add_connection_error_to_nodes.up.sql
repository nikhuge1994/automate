ALTER TABLE IF EXISTS nodes
  ADD COLUMN IF NOT EXISTS connection_error text NOT NULL DEFAULT '';