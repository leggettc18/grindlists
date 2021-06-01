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
	ListUsers(ctx context.Context) ([]User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)

	// list queries
	CreateList(ctx context.Context, arg CreateListParams) (List, error)
	DeleteList(ctx context.Context, id int64) (List, error)
	GetList(ctx context.Context, id int64) (List, error)
	ListLists(ctx context.Context) ([]List, error)
	UpdateLists(ctx context.Context, arg UpdateListParams) (List, error)

	// item queries
	CreateItems(ctx context.Context, arg CreateItemParams, listIDs []int64) (*Item, error)
	DeleteItems(ctx context.Context, id int64) (Item, error)
	GetItems(ctx context.Context, id int64) (Item, error)
	ListItems(ctx context.Context) ([]Item, error)
	UpdateItems(ctx context.Context, arg UpdateItemParams, listIDs []int64) (*Item, error)
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
func (r *repoSvc) CreateItem(ctx context.Context, itemArg CreateItemParams, listIDs []int64) (*Item, error) {
	item := new(Item)
	err := r.withTx(ctx, func(q *Queries) error {
		res, err := q.CreateItem(ctx, itemArg)
		if err != nil {
			return err
		}
		for _, listID := range listIDs {
			if err := q.SetListItem(ctx, SetListItemParams{
				ListID: listID,
				ItemID: res.ID,
			}); err != nil {
				return err
			}
		}
		item = &res
		return nil
	})
	return item, err
}