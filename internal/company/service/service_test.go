package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/yerlanov/xmexercise/internal/company"
	"github.com/yerlanov/xmexercise/internal/company/service/mock"
	"testing"
)

func TestService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := company.Company{
		ID:      1,
		Name:    "Name",
		Code:    "Code",
		Country: "Country",
		Website: "Website",
		Phone:   "Phone",
	}

	mockService := mock.NewMockService(ctrl)

	mockService.EXPECT().Create(context.Background(), expected).Return(expected, nil).Times(1)

	res, err := mockService.Create(context.Background(), expected)
	require.NoError(t, err)
	require.Equal(t, expected.ID, res.ID)
	require.Equal(t, expected.Name, res.Name)
	require.Equal(t, expected.Code, res.Code)
	require.Equal(t, expected.Country, res.Country)
	require.Equal(t, expected.Website, res.Website)
	require.Equal(t, expected.Phone, res.Phone)
}

func TestService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedId := int64(1)
	expectedAffectedCount := int64(1)

	mockService := mock.NewMockService(ctrl)

	mockService.EXPECT().Delete(context.Background(), expectedId).Return(expectedAffectedCount, nil).Times(1)

	affected, err := mockService.Delete(context.Background(), expectedId)
	require.NoError(t, err)
	require.NotEqual(t, 0, affected)
}

func TestService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedId := int64(1)

	expectedUpdate := company.Company{
		ID:      1,
		Name:    "NameUpdated",
		Code:    "CodeUpdated",
		Country: "CountryUpdated",
		Website: "WebsiteUpdated",
		Phone:   "PhoneUpdated",
	}

	mockService := mock.NewMockService(ctrl)
	mockService.EXPECT().Update(context.Background(), expectedUpdate, expectedId).Return(expectedUpdate, nil).Times(1)

	updated, err := mockService.Update(context.Background(), expectedUpdate, expectedId)
	require.NoError(t, err)
	require.Equal(t, expectedUpdate.ID, updated.ID)
	require.Equal(t, expectedUpdate.Name, updated.Name)
	require.Equal(t, expectedUpdate.Code, updated.Code)
	require.Equal(t, expectedUpdate.Country, updated.Country)
	require.Equal(t, expectedUpdate.Website, updated.Website)
	require.Equal(t, expectedUpdate.Phone, updated.Phone)
}

func TestService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedList := []company.Company{{
		ID:      1,
		Name:    "Name1",
		Code:    "Code1",
		Country: "Country1",
		Website: "Website1",
		Phone:   "Phone1",
	},
		{
			ID:      2,
			Name:    "Name2",
			Code:    "Code2",
			Country: "Country2",
			Website: "Website2",
			Phone:   "Phone2",
		},
	}

	mockService := mock.NewMockService(ctrl)
	mockService.EXPECT().List(context.Background()).Return(expectedList, nil).Times(1)

	list, err := mockService.List(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, list)
}
