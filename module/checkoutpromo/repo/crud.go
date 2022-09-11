package repository

import (
	"checkoutpromo/graph/model"
	"checkoutpromo/internal/db"
	"strconv"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

func AddCart(customerID string, productID string, qty int) (res *model.ResponseData, err error) {
	stmt, err := db.Db.Prepare(`INSERT INTO shopping_carts (cart_id, customer_id, product_id, qty, created_at) values (?,?,?,?, now()) `)

	if err != nil {
		return nil, err
	}

	u, _ := uuid.NewV4()
	Id := u.String()

	_, err = stmt.Exec(Id, customerID, productID, qty)

	if err != nil {
		return nil, err
	}

	hasil := &model.ResponseData{
		StatusCode: 200,
		Message:    "success add product to cart",
	}

	return hasil, err
}

func UpdateCart(customerID string, productID string, qty int) (res *model.ResponseData, err error) {
	stmt, err := db.Db.Prepare(`UPDATE shopping_carts set qty = ?, updated_at = now() WHERE customer_id= ? AND product_id = ? `)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(qty, customerID, productID)

	if err != nil {
		return nil, err
	}

	hasil := &model.ResponseData{
		StatusCode: 200,
		Message:    "success update product to cart",
	}

	return hasil, err
}

func ReduceStock(productID string, qty int) (err error) {
	stmt, err := db.Db.Prepare(`UPDATE products set product_qty=product_qty-? WHERE product_id = ?`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(qty, productID)

	if err != nil {
		return err
	}

	return err
}

func AddStock(productID string, qty int) (err error) {
	stmt, err := db.Db.Prepare(`UPDATE products set product_qty=product_qty+? WHERE product_id = ?`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(qty, productID)

	if err != nil {
		return err
	}

	return err
}

func DeleteCart(customerID string, productID string) (res *model.ResponseData, err error) {
	stmt, err := db.Db.Prepare(`DELETE FROM shopping_carts WHERE customer_id = ? AND product_id = ?`)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(customerID, productID)

	if err != nil {
		return nil, err
	}

	hasil := &model.ResponseData{
		StatusCode: 200,
		Message:    "success remove product from cart",
	}

	return hasil, err
}

func Checkout(customerId string) (orderNumber string, err error) {
	stmt, err := db.Db.Prepare(`SELECT sc.customer_id , sc.product_id , p.product_price , sc.qty, p.product_price*sc.qty as total_price FROM shopping_carts sc 
	LEFT JOIN products p on p.product_id = sc.product_id where sc.customer_id = ?`)

	if err != nil {
		return orderNumber, err
	}

	rows, err := stmt.Query(customerId)

	if err != nil {
		return orderNumber, err
	}

	// generate order num
	t := time.Now()
	strYear := strconv.Itoa(t.Year())
	strMonth := strconv.Itoa(int(t.Month()))
	strDate := strconv.Itoa(t.Day())
	strHour := strconv.Itoa(t.Hour())
	strMinute := strconv.Itoa(t.Minute())
	strSecond := strconv.Itoa(t.Second())

	orderNum := `INV-` + strYear + `-` + strMonth + `-` + strDate + `-` + strHour + `-` + strMinute + `-` + strSecond

	for rows.Next() {
		var orders model.OrderDetail
		rows.Scan(&orders.CustomerID, &orders.ProductID, &orders.UnitPrice, &orders.Quantity, &orders.TotalPrice)

		//insert to table orders
		u, _ := uuid.NewV4()
		Id := u.String()

		// cek promo
		stmtp, err := db.Db.Prepare(`SELECT coalesce(price, NULL, 0), coalesce(bonus_product_id, NULL, '')  FROM promo_rules WHERE product_id = ? AND  min_qty >= ?`)

		if err != nil {
			return orderNumber, err
		}

		row, err := stmtp.Query(orders.ProductID, orders.Quantity)

		if err != nil {
			return orderNumber, err
		}

		var discountPrice float64
		var bonusProductID string

		if row.Next() {
			err = row.Scan(&discountPrice, &bonusProductID)
			if err != nil {
				return orderNumber, err
			}
		}

		if discountPrice != 0 {
			orders.UnitPrice = discountPrice
			orders.TotalPrice = discountPrice * float64(orders.Quantity)
		}

		if bonusProductID != "" {
			// check stock
			qtystock, err := CheckStock(bonusProductID)

			if err != nil {
				return orderNum, err
			}

			if qtystock > 1 {

				//reduce stock
				err := ReduceStock(bonusProductID, 1)

				if err != nil {
					return orderNum, err
				}

				ubonus, _ := uuid.NewV4()
				Idbonus := ubonus.String()

				stmtbonus, err := db.Db.Prepare(`INSERT INTO orders (order_id, order_num, customer_id, product_id, unit_price, qty, total_price, created_at) values (?,?,?,?,?,?,?,now())`)

				_, err = stmtbonus.Exec(Idbonus, orderNum, orders.CustomerID, bonusProductID, 0, 1, 0)

				if err != nil {
					return orderNumber, err
				}
			}

		}

		stmtb, err := db.Db.Prepare(`INSERT INTO orders (order_id, order_num, customer_id, product_id, unit_price, qty, total_price, created_at) values (?,?,?,?,?,?,?,now())`)

		_, err = stmtb.Exec(Id, orderNum, orders.CustomerID, orders.ProductID, orders.UnitPrice, orders.Quantity, orders.TotalPrice)

		if err != nil {
			return orderNumber, err
		}

	}

	// clean up table shopping carts
	stmtc, err := db.Db.Prepare(`DELETE FROM shopping_carts WHERE customer_id = ? `)

	if err != nil {
		return orderNumber, err
	}

	_, err = stmtc.Exec(customerId)

	if err != nil {
		return orderNumber, err
	}

	return orderNum, err
}
