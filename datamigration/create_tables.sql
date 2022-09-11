-- checkoutpromo.customers definition
CREATE TABLE `customers` (
  `customer_id` varchar(36) NOT NULL,
  `customer_name` varchar(50) NOT NULL,
  `customer_email` varchar(255) NOT NULL,
  PRIMARY KEY (`customer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- checkoutpromo.orders definition
CREATE TABLE `orders` (
  `order_id` varchar(36) NOT NULL,
  `order_num` varchar(50) NOT NULL,
  `customer_id` varchar(36) NOT NULL,
  `product_id` varchar(36) NOT NULL,
  `unit_price` float NOT NULL,
  `qty` int(11) DEFAULT '0',
  `total_price` float NOT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- checkoutpromo.orders definition
CREATE TABLE `orders` (
  `order_id` varchar(36) NOT NULL,
  `order_num` varchar(50) NOT NULL,
  `customer_id` varchar(36) NOT NULL,
  `product_id` varchar(36) NOT NULL,
  `unit_price` float NOT NULL,
  `qty` int(11) DEFAULT '0',
  `total_price` float NOT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- checkoutpromo.promo_rules definition
CREATE TABLE `promo_rules` (
  `promo_id` varchar(36) NOT NULL,
  `promo_name` varchar(255) NOT NULL,
  `product_id` varchar(36) NOT NULL,
  `min_qty` int(11) NOT NULL,
  `price` float DEFAULT '0',
  `bonus_product_id` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`promo_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- checkoutpromo.shopping_carts definition
CREATE TABLE `shopping_carts` (
  `cart_id` varchar(36) NOT NULL,
  `customer_id` varchar(36) NOT NULL,
  `product_id` varchar(36) NOT NULL,
  `qty` int(11) DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`cart_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;