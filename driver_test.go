package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetDriver_Postgres(t *testing.T) {
	drv, err := GetDriver("postgres")
	require.Nil(t, err)
	_, ok := drv.(PostgresDriver)
	require.Equal(t, true, ok)
}

func TestGetDriver_MySQL(t *testing.T) {
	drv, err := GetDriver("mysql")
	require.Nil(t, err)
	_, ok := drv.(MySQLDriver)
	require.Equal(t, true, ok)
}

func TestGetDriver_Error(t *testing.T) {
	drv, err := GetDriver("foo")
	require.Equal(t, "Unknown driver: foo", err.Error())
	require.Nil(t, drv)
}
