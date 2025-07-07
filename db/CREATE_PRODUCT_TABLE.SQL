CREATE TABLE IF NOT EXISTS product (
  id         VARCHAR(128) UNIQUE NOT NULL,
  name       VARCHAR(128) NOT NULL,
  quantity   INT NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO product (id, name, quantity) VALUES ('prod-001', 'Gaming Mouse', 100);
INSERT INTO product (id, name, quantity) VALUES ('prod-002', 'Mechanical Keyboard', 50);