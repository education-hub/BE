// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	entities "github.com/education-hub/BE/app/entities"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"
)

// TransactionRepo is an autogenerated mock type for the TransactionRepo type
type TransactionRepo struct {
	mock.Mock
}

// CreateTranscation provides a mock function with given fields: db, data, typee
func (_m *TransactionRepo) CreateTranscation(db *gorm.DB, data entities.Transaction, typee string) error {
	ret := _m.Called(db, data, typee)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.Transaction, string) error); ok {
		r0 = rf(db, data, typee)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCart provides a mock function with given fields: db, schid, uid
func (_m *TransactionRepo) DeleteCart(db *gorm.DB, schid int, uid int) error {
	ret := _m.Called(db, schid, uid)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, int, int) error); ok {
		r0 = rf(db, schid, uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllCartByuid provides a mock function with given fields: db, uid
func (_m *TransactionRepo) GetAllCartByuid(db *gorm.DB, uid int) ([]entities.Carts, error) {
	ret := _m.Called(db, uid)

	var r0 []entities.Carts
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) ([]entities.Carts, error)); ok {
		return rf(db, uid)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) []entities.Carts); ok {
		r0 = rf(db, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Carts)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, int) error); ok {
		r1 = rf(db, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCart provides a mock function with given fields: db, schid, userid
func (_m *TransactionRepo) GetCart(db *gorm.DB, schid int, userid int) (*entities.Carts, error) {
	ret := _m.Called(db, schid, userid)

	var r0 *entities.Carts
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, int, int) (*entities.Carts, error)); ok {
		return rf(db, schid, userid)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, int, int) *entities.Carts); ok {
		r0 = rf(db, schid, userid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Carts)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, int, int) error); ok {
		r1 = rf(db, schid, userid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSchoolPayment provides a mock function with given fields: db, schid
func (_m *TransactionRepo) GetSchoolPayment(db *gorm.DB, schid int) (*entities.School, error) {
	ret := _m.Called(db, schid)

	var r0 *entities.School
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) (*entities.School, error)); ok {
		return rf(db, schid)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) *entities.School); ok {
		r0 = rf(db, schid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.School)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, int) error); ok {
		r1 = rf(db, schid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransaction provides a mock function with given fields: db, schoolid, userid
func (_m *TransactionRepo) GetTransaction(db *gorm.DB, schoolid int, userid int) (*entities.Transaction, error) {
	ret := _m.Called(db, schoolid, userid)

	var r0 *entities.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, int, int) (*entities.Transaction, error)); ok {
		return rf(db, schoolid, userid)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, int, int) *entities.Transaction); ok {
		r0 = rf(db, schoolid, userid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, int, int) error); ok {
		r1 = rf(db, schoolid, userid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByInvoice provides a mock function with given fields: db, invoice
func (_m *TransactionRepo) GetTransactionByInvoice(db *gorm.DB, invoice string) (*entities.Transaction, error) {
	ret := _m.Called(db, invoice)

	var r0 *entities.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) (*entities.Transaction, error)); ok {
		return rf(db, invoice)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) *entities.Transaction); ok {
		r0 = rf(db, invoice)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, string) error); ok {
		r1 = rf(db, invoice)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: db, invoice, status
func (_m *TransactionRepo) UpdateStatus(db *gorm.DB, invoice string, status string) error {
	ret := _m.Called(db, invoice, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, string) error); ok {
		r0 = rf(db, invoice, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTransactionRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionRepo creates a new instance of TransactionRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionRepo(t mockConstructorTestingTNewTransactionRepo) *TransactionRepo {
	mock := &TransactionRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
