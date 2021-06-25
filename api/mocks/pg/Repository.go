// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	pg "github.com/leggettc18/grindlists/api/pg"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CountListHearts provides a mock function with given fields: ctx, list_id
func (_m *Repository) CountListHearts(ctx context.Context, list_id int64) (int64, error) {
	ret := _m.Called(ctx, list_id)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, int64) int64); ok {
		r0 = rf(ctx, list_id)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, list_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountListHeartsByUser provides a mock function with given fields: ctx, user_id
func (_m *Repository) CountListHeartsByUser(ctx context.Context, user_id int64) (int64, error) {
	ret := _m.Called(ctx, user_id)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, int64) int64); ok {
		r0 = rf(ctx, user_id)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, user_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateList provides a mock function with given fields: ctx, arg
func (_m *Repository) CreateList(ctx context.Context, arg pg.CreateListParams) (pg.List, error) {
	ret := _m.Called(ctx, arg)

	var r0 pg.List
	if rf, ok := ret.Get(0).(func(context.Context, pg.CreateListParams) pg.List); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(pg.List)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pg.CreateListParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateListItem provides a mock function with given fields: ctx, itemArg, listItemArg
func (_m *Repository) CreateListItem(ctx context.Context, itemArg pg.CreateItemParams, listItemArg pg.SetListItemParams) (*pg.Item, error) {
	ret := _m.Called(ctx, itemArg, listItemArg)

	var r0 *pg.Item
	if rf, ok := ret.Get(0).(func(context.Context, pg.CreateItemParams, pg.SetListItemParams) *pg.Item); ok {
		r0 = rf(ctx, itemArg, listItemArg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pg.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pg.CreateItemParams, pg.SetListItemParams) error); ok {
		r1 = rf(ctx, itemArg, listItemArg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: ctx, arg
func (_m *Repository) CreateUser(ctx context.Context, arg pg.CreateUserParams) (pg.User, error) {
	ret := _m.Called(ctx, arg)

	var r0 pg.User
	if rf, ok := ret.Get(0).(func(context.Context, pg.CreateUserParams) pg.User); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(pg.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pg.CreateUserParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteItem provides a mock function with given fields: ctx, id
func (_m *Repository) DeleteItem(ctx context.Context, id int64) (pg.Item, error) {
	ret := _m.Called(ctx, id)

	var r0 pg.Item
	if rf, ok := ret.Get(0).(func(context.Context, int64) pg.Item); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(pg.Item)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteList provides a mock function with given fields: ctx, id
func (_m *Repository) DeleteList(ctx context.Context, id int64) (pg.List, error) {
	ret := _m.Called(ctx, id)

	var r0 pg.List
	if rf, ok := ret.Get(0).(func(context.Context, int64) pg.List); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(pg.List)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *Repository) DeleteUser(ctx context.Context, id int64) (pg.User, error) {
	ret := _m.Called(ctx, id)

	var r0 pg.User
	if rf, ok := ret.Get(0).(func(context.Context, int64) pg.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(pg.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetItem provides a mock function with given fields: ctx, id
func (_m *Repository) GetItem(ctx context.Context, id int64) (pg.Item, error) {
	ret := _m.Called(ctx, id)

	var r0 pg.Item
	if rf, ok := ret.Get(0).(func(context.Context, int64) pg.Item); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(pg.Item)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetList provides a mock function with given fields: ctx, id
func (_m *Repository) GetList(ctx context.Context, id int64) (pg.List, error) {
	ret := _m.Called(ctx, id)

	var r0 pg.List
	if rf, ok := ret.Get(0).(func(context.Context, int64) pg.List); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(pg.List)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListHearts provides a mock function with given fields: ctx, list_id
func (_m *Repository) GetListHearts(ctx context.Context, list_id int64) ([]pg.ListHeart, error) {
	ret := _m.Called(ctx, list_id)

	var r0 []pg.ListHeart
	if rf, ok := ret.Get(0).(func(context.Context, int64) []pg.ListHeart); ok {
		r0 = rf(ctx, list_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pg.ListHeart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, list_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListHeartsByUser provides a mock function with given fields: ctx, user_id
func (_m *Repository) GetListHeartsByUser(ctx context.Context, user_id int64) ([]pg.ListHeart, error) {
	ret := _m.Called(ctx, user_id)

	var r0 []pg.ListHeart
	if rf, ok := ret.Get(0).(func(context.Context, int64) []pg.ListHeart); ok {
		r0 = rf(ctx, user_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pg.ListHeart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, user_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListListItems provides a mock function with given fields: ctx, list_id
func (_m *Repository) GetListListItems(ctx context.Context, list_id int64) ([]pg.ListItem, error) {
	ret := _m.Called(ctx, list_id)

	var r0 []pg.ListItem
	if rf, ok := ret.Get(0).(func(context.Context, int64) []pg.ListItem); ok {
		r0 = rf(ctx, list_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pg.ListItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, list_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: ctx, id
func (_m *Repository) GetUser(ctx context.Context, id int64) (pg.User, error) {
	ret := _m.Called(ctx, id)

	var r0 pg.User
	if rf, ok := ret.Get(0).(func(context.Context, int64) pg.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(pg.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *Repository) GetUserByEmail(ctx context.Context, email string) (pg.User, error) {
	ret := _m.Called(ctx, email)

	var r0 pg.User
	if rf, ok := ret.Get(0).(func(context.Context, string) pg.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(pg.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserLists provides a mock function with given fields: ctx, user_id
func (_m *Repository) GetUserLists(ctx context.Context, user_id int64) ([]pg.List, error) {
	ret := _m.Called(ctx, user_id)

	var r0 []pg.List
	if rf, ok := ret.Get(0).(func(context.Context, int64) []pg.List); ok {
		r0 = rf(ctx, user_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pg.List)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, user_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListItems provides a mock function with given fields: ctx
func (_m *Repository) ListItems(ctx context.Context) ([]pg.Item, error) {
	ret := _m.Called(ctx)

	var r0 []pg.Item
	if rf, ok := ret.Get(0).(func(context.Context) []pg.Item); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pg.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListLists provides a mock function with given fields: ctx
func (_m *Repository) ListLists(ctx context.Context) ([]pg.List, error) {
	ret := _m.Called(ctx)

	var r0 []pg.List
	if rf, ok := ret.Get(0).(func(context.Context) []pg.List); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pg.List)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUsers provides a mock function with given fields: ctx
func (_m *Repository) ListUsers(ctx context.Context) ([]pg.User, error) {
	ret := _m.Called(ctx)

	var r0 []pg.User
	if rf, ok := ret.Get(0).(func(context.Context) []pg.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pg.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetListHeart provides a mock function with given fields: ctx, arg
func (_m *Repository) SetListHeart(ctx context.Context, arg pg.SetListHeartParams) error {
	ret := _m.Called(ctx, arg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, pg.SetListHeartParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetListItem provides a mock function with given fields: ctx, arg
func (_m *Repository) SetListItem(ctx context.Context, arg pg.SetListItemParams) error {
	ret := _m.Called(ctx, arg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, pg.SetListItemParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnsetListHeart provides a mock function with given fields: ctx, id
func (_m *Repository) UnsetListHeart(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnsetListItem provides a mock function with given fields: ctx, id
func (_m *Repository) UnsetListItem(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateItem provides a mock function with given fields: ctx, arg
func (_m *Repository) UpdateItem(ctx context.Context, arg pg.UpdateItemParams) (pg.Item, error) {
	ret := _m.Called(ctx, arg)

	var r0 pg.Item
	if rf, ok := ret.Get(0).(func(context.Context, pg.UpdateItemParams) pg.Item); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(pg.Item)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pg.UpdateItemParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateList provides a mock function with given fields: ctx, arg
func (_m *Repository) UpdateList(ctx context.Context, arg pg.UpdateListParams) (pg.List, error) {
	ret := _m.Called(ctx, arg)

	var r0 pg.List
	if rf, ok := ret.Get(0).(func(context.Context, pg.UpdateListParams) pg.List); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(pg.List)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pg.UpdateListParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateListItem provides a mock function with given fields: ctx, arg
func (_m *Repository) UpdateListItem(ctx context.Context, arg pg.UpdateListItemParams) error {
	ret := _m.Called(ctx, arg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, pg.UpdateListItemParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: ctx, arg
func (_m *Repository) UpdateUser(ctx context.Context, arg pg.UpdateUserParams) (pg.User, error) {
	ret := _m.Called(ctx, arg)

	var r0 pg.User
	if rf, ok := ret.Get(0).(func(context.Context, pg.UpdateUserParams) pg.User); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(pg.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pg.UpdateUserParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
