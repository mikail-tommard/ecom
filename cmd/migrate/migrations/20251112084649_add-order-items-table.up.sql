CREATE TABLE IF NOT EXISTS order_items (
  id SERIAL PRIMARY KEY,
  orderId INT NOT NULL,
  productId INT NOT NULL,
  quanity INT NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  FOREIGN KEY (orderId) REFERENCES orders(id),
  FOREIGN KEY (productId) REFERENCES products(id)
);