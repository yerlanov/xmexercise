package storage

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/yerlanov/xmexercise/internal/company"
	"testing"
)

func createCompany(t *testing.T) int64 {
	in := company.Company{
		Name:    "TestName",
		Code:    "TestCode",
		Country: "TestCountry",
		Website: "TestWebsite",
		Phone:   "TestPhone",
	}

	res, err := queries.Create(context.Background(), in)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.NotZero(t, res)

	return res.ID
}

func TestQueries_Create(t *testing.T) {
	createCompany(t)
}

func TestQueries_Delete(t *testing.T) {
	id := createCompany(t)

	affected, err := queries.Delete(context.Background(), id)
	require.NoError(t, err)
	require.NotEqual(t, 0, affected)
}

func TestQueries_Update(t *testing.T) {
	id := createCompany(t)

	in := company.Company{
		Name:    "TestUpdatedName",
		Code:    "TestUpdatedCode",
		Country: "TestUpdatedCountry",
		Website: "TestUpdatedWebsite",
		Phone:   "TestUpdatedPhone",
	}

	updated, err := queries.Update(context.Background(), in, id)
	if err != nil {
		return
	}

	require.NoError(t, err)
	require.NotEmpty(t, updated)

	require.Equal(t, id, updated.ID)
	require.Equal(t, in.Name, updated.Name)
	require.Equal(t, in.Code, updated.Code)
	require.Equal(t, in.Country, updated.Country)
	require.Equal(t, in.Website, updated.Website)
	require.Equal(t, in.Phone, updated.Phone)
}

func TestQueries_List(t *testing.T) {
	for i := 0; i < 2; i++ {
		createCompany(t)
	}

	list, err := queries.List(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, list)
}

func TestQueries_ListWithFilter(t *testing.T) {

	in1 := company.Company{
		Name:    "TestFilterName1",
		Code:    "TestFilterCode1",
		Country: "TestFilterCountry1",
		Website: "TestFilterWebsite1",
		Phone:   "TestFilterPhone1",
	}

	in2 := company.Company{
		Name:    "TestFilterName2",
		Code:    "TestFilterCode2",
		Country: "TestFilterCountry1",
		Website: "TestFilterWebsite2",
		Phone:   "TestFilterPhone2",
	}

	_, err := queries.Create(context.Background(), in1)
	require.NoError(t, err)

	_, err = queries.Create(context.Background(), in2)
	require.NoError(t, err)

	filter := map[string]string{"country": "TestFilterCountry1"}

	list, err := queries.ListWithFilter(context.Background(), filter)
	require.NoError(t, err)
	require.NotEmpty(t, list)

}
