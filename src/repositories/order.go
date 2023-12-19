package repositories

import (
	"api/src/models"
	"database/sql"
	"errors"
)

type Order struct {
	db *sql.DB
}

func NewRepositoryOrder(db *sql.DB) *Order {
	return &Order{db: db}
}

func (o *Order) CreateNewOrder(order models.Order) error {
	statement, err := o.db.Prepare("insert into orders (order_id, user_id, status) values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(&order.OrderID, &order.UserID, &order.Status); err != nil {
		return err
	}
	return nil
}

func (o *Order) InsertOrderProducts(orderID string, orderProduct models.OrderItem) error {
	statement, err := o.db.Prepare("insert into order_item (product_id, order_id, amount, price) values(?, ?, ?, ?)")
	if err != nil {
		return err
	}

	if _, err = statement.Exec(&orderProduct.Product.ID, &orderID, &orderProduct.Amount, &orderProduct.Price); err != nil {
		return err
	}

	return nil
}

func (o *Order) SearchOrderByID(orderID uint64) (models.Order, error) {
	row, err := o.db.Query("select * from orders where order_id = ?", &orderID)
	if err != nil {
		return models.Order{}, err
	}
	defer row.Close()

	if row.Next() {
		var order models.Order
		if err = row.Scan(&order.OrderID, &order.UserID, &order.Date, &order.Status); err != nil {
			return models.Order{}, err
		}
		return order, nil
	}

	return models.Order{}, errors.New("no order")
}

func (o *Order) SearchOrderItens(order *models.Order) error {
	rows, err := o.db.Query("select o.*, p.*, i.filename, i.path from order_item o inner join product p on p.id = o.product_id inner join images i on i.image_id = p.product_image_id where o.order_id = ?", order.OrderID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var orderItem models.OrderItem
		if err = rows.Scan(&orderItem.Product.ID, &orderItem.OrderID, &orderItem.Amount, &orderItem.Price,
			&orderItem.Product.ID, &orderItem.Product.Name, &orderItem.Product.Description, &orderItem.Product.Price,
			&orderItem.Product.Stock, &orderItem.Product.Image.ImageID, &orderItem.Product.AddedIn,
			&orderItem.Product.Image.Filename, &orderItem.Product.Image.Path); err != nil {
			return err
		}
		order.OrderItems = append(order.OrderItems, orderItem)
	}
	return nil
}

func (o *Order) ShowOrders() ([]models.Order, error) {
	rows, err := o.db.Query("select * from orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order

		if err = rows.Scan(&order.OrderID, &order.UserID, &order.Date, &order.Status); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (o *Order) SearchOrderByUserID(userID string) ([]models.Order, error) {
	rows, err := o.db.Query("select * from orders where user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order

	for rows.Next() {
		var order models.Order

		if err = rows.Scan(&order.OrderID, &order.UserID, &order.Date, &order.Status); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (o *Order) ChangeStatus(order models.Order) error {
	statement, err := o.db.Prepare("update orders set status = ? where order_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(&order.Status, &order.OrderID); err != nil {
		return err
	}

	return nil
}
