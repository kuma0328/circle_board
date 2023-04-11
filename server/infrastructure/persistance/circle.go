package persistance

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/kuma0328/circle_board/domain/entity"
	"github.com/kuma0328/circle_board/domain/repositry"
	"github.com/kuma0328/circle_board/infrastructure/database"
)

var _ repositry.ICircleRepositry = &CircleRepositry{}

type CircleRepositry struct {
	conn *database.Conn
}

func NewCircleRepository(conn *database.Conn) *CircleRepositry {
	return &CircleRepositry{
		conn: conn,
	}
}
// POST /circle/create
func (repo *CircleRepositry) CreateCircle(circle *entity.Circle) error {
	query := `
		INSERT INTO circle (name, image_url, genre, location, member, twitter, instagram)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := repo.conn.DB.Exec(query, circle.Name, circle.ImageURL, circle.Genre, circle.Location, circle.Member, circle.Twitter, circle.Instagram)
	if err != nil {
		return fmt.Errorf("failed to insert circle db.Exec : %v", err)
	}
	
	return nil
}
// GET /circle/all
func (repo *CircleRepositry) GetAllCircle() ([]*entity.Circle, error) {
	query := `
		SELECT * 
		FROM circle
	`

	rows, err := repo.conn.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all circle db.Query : %v", err)
	}
	defer rows.Close()

	var circles []*entity.Circle
	for rows.Next() {
		c := &entity.Circle{}
		if err := rows.Scan(&c.ID, &c.Name, &c.ImageURL, &c.Genre, &c.Location, &c.Member, &c.Twitter, &c.Instagram); err != nil {
			return nil, fmt.Errorf("failed to get all circle rows.Scan : %v", err)
		}
		circles = append(circles, c)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get all circle rows.Err : %v", err)
	}
	return circles, nil
}
// GET /circle/:id
func (repo *CircleRepositry) GetCircleById(id int64) (*entity.Circle, error) {
	query := `
		SELECT *
		FROM circle
		WHERE id = $1
	`

	c := &entity.Circle{}
	err := repo.conn.DB.QueryRow(query, id).
	Scan(&c.ID, &c.Name, &c.ImageURL, &c.Genre, &c.Location, &c.Member, &c.Twitter, &c.Instagram)
	if err != nil {
		return nil, fmt.Errorf("failed to get circle by id : %v", err)
	}
	
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to get circle ErrNoRows: %v", err)
	}

	return c, nil
}