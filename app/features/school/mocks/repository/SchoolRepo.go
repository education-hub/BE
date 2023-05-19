// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	entities "github.com/education-hub/BE/app/entities"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"
)

// SchoolRepo is an autogenerated mock type for the SchoolRepo type
type SchoolRepo struct {
	mock.Mock
}

// AddAchievement provides a mock function with given fields: db, achv
func (_m *SchoolRepo) AddAchievement(db *gorm.DB, achv entities.Achievement) (int, error) {
	ret := _m.Called(db, achv)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.Achievement) (int, error)); ok {
		return rf(db, achv)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.Achievement) int); ok {
		r0 = rf(db, achv)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, entities.Achievement) error); ok {
		r1 = rf(db, achv)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddExtracurricular provides a mock function with given fields: db, achv
func (_m *SchoolRepo) AddExtracurricular(db *gorm.DB, achv entities.Extracurricular) (int, error) {
	ret := _m.Called(db, achv)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.Extracurricular) (int, error)); ok {
		return rf(db, achv)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.Extracurricular) int); ok {
		r0 = rf(db, achv)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, entities.Extracurricular) error); ok {
		r1 = rf(db, achv)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddFaq provides a mock function with given fields: db, faq
func (_m *SchoolRepo) AddFaq(db *gorm.DB, faq entities.Faq) (int, error) {
	ret := _m.Called(db, faq)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.Faq) (int, error)); ok {
		return rf(db, faq)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.Faq) int); ok {
		r0 = rf(db, faq)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, entities.Faq) error); ok {
		r1 = rf(db, faq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddPayment provides a mock function with given fields: db, paym
func (_m *SchoolRepo) AddPayment(db *gorm.DB, paym entities.Payment) (int, error) {
	ret := _m.Called(db, paym)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.Payment) (int, error)); ok {
		return rf(db, paym)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.Payment) int); ok {
		r0 = rf(db, paym)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, entities.Payment) error); ok {
		r1 = rf(db, paym)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddReview provides a mock function with given fields: db, data
func (_m *SchoolRepo) AddReview(db *gorm.DB, data entities.Reviews) (int, error) {
	ret := _m.Called(db, data)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.Reviews) (int, error)); ok {
		return rf(db, data)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.Reviews) int); ok {
		r0 = rf(db, data)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, entities.Reviews) error); ok {
		r1 = rf(db, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: db, school
func (_m *SchoolRepo) Create(db *gorm.DB, school entities.School) (int, error) {
	ret := _m.Called(db, school)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.School) (int, error)); ok {
		return rf(db, school)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.School) int); ok {
		r0 = rf(db, school)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, entities.School) error); ok {
		r1 = rf(db, school)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSubmission provides a mock function with given fields: db, subm
func (_m *SchoolRepo) CreateSubmission(db *gorm.DB, subm entities.Submission) (int, error) {
	ret := _m.Called(db, subm)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.Submission) (int, error)); ok {
		return rf(db, subm)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.Submission) int); ok {
		r0 = rf(db, subm)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, entities.Submission) error); ok {
		r1 = rf(db, subm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: db, id, uid
func (_m *SchoolRepo) Delete(db *gorm.DB, id int, uid int) error {
	ret := _m.Called(db, id, uid)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, int, int) error); ok {
		r0 = rf(db, id, uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAchievement provides a mock function with given fields: db, id
func (_m *SchoolRepo) DeleteAchievement(db *gorm.DB, id int) error {
	ret := _m.Called(db, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) error); ok {
		r0 = rf(db, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteExtracurricular provides a mock function with given fields: db, id
func (_m *SchoolRepo) DeleteExtracurricular(db *gorm.DB, id int) error {
	ret := _m.Called(db, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) error); ok {
		r0 = rf(db, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteFaq provides a mock function with given fields: db, id
func (_m *SchoolRepo) DeleteFaq(db *gorm.DB, id int) error {
	ret := _m.Called(db, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) error); ok {
		r0 = rf(db, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePayment provides a mock function with given fields: db, id
func (_m *SchoolRepo) DeletePayment(db *gorm.DB, id int) error {
	ret := _m.Called(db, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) error); ok {
		r0 = rf(db, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByNPSN provides a mock function with given fields: db, npsn
func (_m *SchoolRepo) FindByNPSN(db *gorm.DB, npsn string) error {
	ret := _m.Called(db, npsn)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) error); ok {
		r0 = rf(db, npsn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: db, limit, offset, search
func (_m *SchoolRepo) GetAll(db *gorm.DB, limit int, offset int, search string) ([]entities.School, int, error) {
	ret := _m.Called(db, limit, offset, search)

	var r0 []entities.School
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, int, int, string) ([]entities.School, int, error)); ok {
		return rf(db, limit, offset, search)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, int, int, string) []entities.School); ok {
		r0 = rf(db, limit, offset, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.School)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, int, int, string) int); ok {
		r1 = rf(db, limit, offset, search)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(*gorm.DB, int, int, string) error); ok {
		r2 = rf(db, limit, offset, search)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAllProgressAndSubmissionByuid provides a mock function with given fields: db, uid
func (_m *SchoolRepo) GetAllProgressAndSubmissionByuid(db *gorm.DB, uid int) (*entities.School, error) {
	ret := _m.Called(db, uid)

	var r0 *entities.School
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) (*entities.School, error)); ok {
		return rf(db, uid)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) *entities.School); ok {
		r0 = rf(db, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.School)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, int) error); ok {
		r1 = rf(db, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllProgressByuid provides a mock function with given fields: db, uid
func (_m *SchoolRepo) GetAllProgressByuid(db *gorm.DB, uid int) ([]entities.Progress, error) {
	ret := _m.Called(db, uid)

	var r0 []entities.Progress
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) ([]entities.Progress, error)); ok {
		return rf(db, uid)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) []entities.Progress); ok {
		r0 = rf(db, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Progress)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, int) error); ok {
		r1 = rf(db, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: db, id
func (_m *SchoolRepo) GetById(db *gorm.DB, id int) (*entities.School, error) {
	ret := _m.Called(db, id)

	var r0 *entities.School
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) (*entities.School, error)); ok {
		return rf(db, id)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) *entities.School); ok {
		r0 = rf(db, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.School)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, int) error); ok {
		r1 = rf(db, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUid provides a mock function with given fields: db, uid
func (_m *SchoolRepo) GetByUid(db *gorm.DB, uid int) (*entities.School, error) {
	ret := _m.Called(db, uid)

	var r0 *entities.School
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) (*entities.School, error)); ok {
		return rf(db, uid)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) *entities.School); ok {
		r0 = rf(db, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.School)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, int) error); ok {
		r1 = rf(db, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProgressByid provides a mock function with given fields: db, id
func (_m *SchoolRepo) GetProgressByid(db *gorm.DB, id int) (*entities.Progress, error) {
	ret := _m.Called(db, id)

	var r0 *entities.Progress
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) (*entities.Progress, error)); ok {
		return rf(db, id)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) *entities.Progress); ok {
		r0 = rf(db, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Progress)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, int) error); ok {
		r1 = rf(db, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubmissionByid provides a mock function with given fields: db, id
func (_m *SchoolRepo) GetSubmissionByid(db *gorm.DB, id int) (*entities.Submission, error) {
	ret := _m.Called(db, id)

	var r0 *entities.Submission
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) (*entities.Submission, error)); ok {
		return rf(db, id)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) *entities.Submission); ok {
		r0 = rf(db, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Submission)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, int) error); ok {
		r1 = rf(db, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: db, school
func (_m *SchoolRepo) Update(db *gorm.DB, school entities.School) (*entities.School, error) {
	ret := _m.Called(db, school)

	var r0 *entities.School
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.School) (*entities.School, error)); ok {
		return rf(db, school)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.School) *entities.School); ok {
		r0 = rf(db, school)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.School)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, entities.School) error); ok {
		r1 = rf(db, school)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAchievement provides a mock function with given fields: db, achv
func (_m *SchoolRepo) UpdateAchievement(db *gorm.DB, achv entities.Achievement) (*entities.Achievement, error) {
	ret := _m.Called(db, achv)

	var r0 *entities.Achievement
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.Achievement) (*entities.Achievement, error)); ok {
		return rf(db, achv)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.Achievement) *entities.Achievement); ok {
		r0 = rf(db, achv)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Achievement)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, entities.Achievement) error); ok {
		r1 = rf(db, achv)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateExtracurricular provides a mock function with given fields: db, achv
func (_m *SchoolRepo) UpdateExtracurricular(db *gorm.DB, achv entities.Extracurricular) (*entities.Extracurricular, error) {
	ret := _m.Called(db, achv)

	var r0 *entities.Extracurricular
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.Extracurricular) (*entities.Extracurricular, error)); ok {
		return rf(db, achv)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.Extracurricular) *entities.Extracurricular); ok {
		r0 = rf(db, achv)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Extracurricular)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, entities.Extracurricular) error); ok {
		r1 = rf(db, achv)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFaq provides a mock function with given fields: db, extrac
func (_m *SchoolRepo) UpdateFaq(db *gorm.DB, extrac entities.Faq) (*entities.Faq, error) {
	ret := _m.Called(db, extrac)

	var r0 *entities.Faq
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.Faq) (*entities.Faq, error)); ok {
		return rf(db, extrac)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.Faq) *entities.Faq); ok {
		r0 = rf(db, extrac)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Faq)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, entities.Faq) error); ok {
		r1 = rf(db, extrac)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePayment provides a mock function with given fields: db, paym
func (_m *SchoolRepo) UpdatePayment(db *gorm.DB, paym entities.Payment) (*entities.Payment, error) {
	ret := _m.Called(db, paym)

	var r0 *entities.Payment
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.Payment) (*entities.Payment, error)); ok {
		return rf(db, paym)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, entities.Payment) *entities.Payment); ok {
		r0 = rf(db, paym)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Payment)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, entities.Payment) error); ok {
		r1 = rf(db, paym)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProgress provides a mock function with given fields: db, id, status
func (_m *SchoolRepo) UpdateProgress(db *gorm.DB, id int, status string) (*entities.Progress, error) {
	ret := _m.Called(db, id, status)

	var r0 *entities.Progress
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, int, string) (*entities.Progress, error)); ok {
		return rf(db, id, status)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, int, string) *entities.Progress); ok {
		r0 = rf(db, id, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Progress)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, int, string) error); ok {
		r1 = rf(db, id, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProgressByUid provides a mock function with given fields: db, uid, schid, status
func (_m *SchoolRepo) UpdateProgressByUid(db *gorm.DB, uid int, schid int, status string) (int, error) {
	ret := _m.Called(db, uid, schid, status)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, int, int, string) (int, error)); ok {
		return rf(db, uid, schid, status)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, int, int, string) int); ok {
		r0 = rf(db, uid, schid, status)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, int, int, string) error); ok {
		r1 = rf(db, uid, schid, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewSchoolRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewSchoolRepo creates a new instance of SchoolRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSchoolRepo(t mockConstructorTestingTNewSchoolRepo) *SchoolRepo {
	mock := &SchoolRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
