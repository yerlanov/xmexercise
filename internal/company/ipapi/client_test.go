package ipapi

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_GetIPInformation(t *testing.T) {
	cl := NewClient("https://ipapi.co")

	res, err := cl.GetIPInformation(context.Background(), "91.246.103.86")
	fmt.Println(res)
	require.NoError(t, err)
	require.NotZero(t, res.IP)
	require.NotZero(t, res.CountryCode)

}
