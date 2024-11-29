CREATE TABLE payments (
  id uuid PRIMARY KEY,
  id_order uuid NOT NULL unique,
  amount NUMERIC(10, 2) NOT NULL,
  payment_status VARCHAR(20) NOT NULL DEFAULT 'PENDING'
);
