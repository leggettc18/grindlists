package pg

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Repository is the applications data layer functionality
type Repository interface {
	// user queries
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteUser(ctx context.Context, id int64) (User, error)
	GetUser(ctx context.Context, id int64) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	ListUsers(ctx context.Context) ([]User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)

	// list queries
	CreateList(ctx context.Context, arg CreateListParams) (List, error)
	DeleteList(ctx context.Context, id int64) (List, error)
	GetList(ctx context.Context, id int64) (List, error)
	ListLists(ctx context.Context) ([]List, error)
	UpdateList(ctx context.Context, arg UpdateListParams) (List, error)
	GetUserLists(ctx context.Context, user_id int64) ([]List, error)

	// item queries
	DeleteItem(ctx context.Context, id int64) (Item, error)
	GetItem(ctx context.Context, id int64) (Item, error)
	ListItems(ctx context.Context) ([]Item, error)
	UpdateItem(ctx context.Context, arg UpdateItemParams) (Item, error)

	// listItem queries
	CreateListItem(ctx context.Context, itemArg CreateItemParams, listItemArg SetListItemParams) (*Item, error)
	SetListItem(ctx context.Context, arg SetListItemParams) (error)
	UpdateListItem(ctx context.Context, arg UpdateListItemParams) (error)
	UnsetListItem(ctx context.Context, id int64) (error)
	GetListListItems(ctx context.Context, list_id int64) ([]ListItem, error)

	// listHeart queries
	GetListHearts(ctx context.Context, list_id int64) ([]ListHeart, error)
	CountListHearts(ctx context.Context, list_id int64) (int64, error)
	GetListHeartsByUser(ctx context.Context, user_id int64) ([]ListHeart, error)
	CountListHeartsByUser(ctx context.Context, user_id int64) (int64, error)
	SetListHeart(ctx context.Context, arg SetListHeartParams) (error)
	UnsetListHeart(ctx context.Context, id int64) (error)
}

type repoSvc struct {
	*Queries
	db *sql.DB
}

func (r *repoSvc) withTx(ctx context.Context, txFn func(*Queries) error) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = txFn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			err = fmt.Errorf("tx failed: %v, unable to rollback: %v", err, rbErr)
		}
	} else {
		err = tx.Commit()
	}
	return err
}

// CreateItem creates an item and links it to a list right away
func (r *repoSvc) CreateListItem(ctx context.Context, itemArg CreateItemParams, listItemArg SetListItemParams) (*Item, error) {
	item := new(Item)
	err := r.withTx(ctx, func(q *Queries) error {
		res, err := q.CreateItem(ctx, itemArg)
		if err != nil {
			return err
		}
		if err := q.SetListItem(ctx, SetListItemParams{
			Quantity: listItemArg.Quantity,
			Collected: listItemArg.Collected,
			ListID: listItemArg.ListID,
			ItemID: res.ID,
		}); err != nil {
			return err
		}
		item = &res
		return nil
	})
	return item, err
}

// NewRepository returns an implementation of the Repository interface.
func NewRepository(db *sql.DB) Repository {
	return &repoSvc {
		Queries: New(db),
		db: db,
	}
}

// Open opens a database specified by the data source name.
// Format: "host=foo port=5432 user=bar password=baz dbname=qux sslmode=disable"
func Open(dataSourceName string) (*sql.DB, error) {
	return sql.Open("postgres", dataSourceName)
}

// StringPtrToNullString converts *string to sql.NullString.
func StringPtrToNullString(s *string) sql.NullString {
	if s != nil {
		return sql.NullString{String: *s, Valid: true}
	}
	return sql.NullString{}
}

// IntPtrToNullInt64 converts *int to sql.NullInt64.
func IntPtrToNullInt64(i *int) sql.NullInt64 {
	if i != nil {
		return sql.NullInt64{Int64: int64(*i), Valid: true}
	}
	return sql.NullInt64{}
}