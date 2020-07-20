// Code generated by mockery v1.0.0. DO NOT EDIT.

package blob

import (
	context "context"

	models "github.com/goharbor/harbor/src/pkg/blob/models"
	mock "github.com/stretchr/testify/mock"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// AssociateWithArtifact provides a mock function with given fields: ctx, blobDigest, artifactDigest
func (_m *Manager) AssociateWithArtifact(ctx context.Context, blobDigest string, artifactDigest string) (int64, error) {
	ret := _m.Called(ctx, blobDigest, artifactDigest)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string) int64); ok {
		r0 = rf(ctx, blobDigest, artifactDigest)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, blobDigest, artifactDigest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssociateWithProject provides a mock function with given fields: ctx, blobID, projectID
func (_m *Manager) AssociateWithProject(ctx context.Context, blobID int64, projectID int64) (int64, error) {
	ret := _m.Called(ctx, blobID, projectID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) int64); ok {
		r0 = rf(ctx, blobID, projectID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, blobID, projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CalculateTotalSizeByProject provides a mock function with given fields: ctx, projectID, excludeForeignLayer
func (_m *Manager) CalculateTotalSizeByProject(ctx context.Context, projectID int64, excludeForeignLayer bool) (int64, error) {
	ret := _m.Called(ctx, projectID, excludeForeignLayer)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, int64, bool) int64); ok {
		r0 = rf(ctx, projectID, excludeForeignLayer)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, bool) error); ok {
		r1 = rf(ctx, projectID, excludeForeignLayer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CleanupAssociationsForArtifact provides a mock function with given fields: ctx, artifactDigest
func (_m *Manager) CleanupAssociationsForArtifact(ctx context.Context, artifactDigest string) error {
	ret := _m.Called(ctx, artifactDigest)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, artifactDigest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CleanupAssociationsForProject provides a mock function with given fields: ctx, projectID, blobs
func (_m *Manager) CleanupAssociationsForProject(ctx context.Context, projectID int64, blobs []*models.Blob) error {
	ret := _m.Called(ctx, projectID, blobs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, []*models.Blob) error); ok {
		r0 = rf(ctx, projectID, blobs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: ctx, digest, contentType, size
func (_m *Manager) Create(ctx context.Context, digest string, contentType string, size int64) (int64, error) {
	ret := _m.Called(ctx, digest, contentType, size)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) int64); ok {
		r0 = rf(ctx, digest, contentType, size)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int64) error); ok {
		r1 = rf(ctx, digest, contentType, size)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Manager) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, digest
func (_m *Manager) Get(ctx context.Context, digest string) (*models.Blob, error) {
	ret := _m.Called(ctx, digest)

	var r0 *models.Blob
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Blob); ok {
		r0 = rf(ctx, digest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Blob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, digest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, params
func (_m *Manager) List(ctx context.Context, params models.ListParams) ([]*models.Blob, error) {
	ret := _m.Called(ctx, params)

	var r0 []*models.Blob
	if rf, ok := ret.Get(0).(func(context.Context, models.ListParams) []*models.Blob); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Blob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.ListParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, _a1
func (_m *Manager) Update(ctx context.Context, _a1 *models.Blob) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Blob) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateBlobStatus provides a mock function with given fields: ctx, _a1
func (_m *Manager) UpdateBlobStatus(ctx context.Context, _a1 *models.Blob) (int64, error) {
	ret := _m.Called(ctx, _a1)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *models.Blob) int64); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Blob) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UselessBlobs provides a mock function with given fields: ctx, timeWindowHours
func (_m *Manager) UselessBlobs(ctx context.Context, timeWindowHours int64) ([]*models.Blob, error) {
	ret := _m.Called(ctx, timeWindowHours)

	var r0 []*models.Blob
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*models.Blob); ok {
		r0 = rf(ctx, timeWindowHours)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Blob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, timeWindowHours)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
