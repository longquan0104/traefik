//go:build !(linux || freebsd || openbsd || darwin)

package server

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestNewListenConfig(t *testing.T) {
	ep := static.EntryPoint{Address: ":0"}
	listenConfig := newListenConfig(&ep)
	require.Nil(t, listenConfig.Control)
	require.Zero(t, listenConfig.KeepAlive)

	l1, err := listenConfig.Listen(context.Background(), "tcp", ep.Address)
	require.NoError(t, err)
	require.NotNil(t, l1)
	defer l1.Close()

	l2, err := listenConfig.Listen(context.Background(), "tcp", l1.Addr().String())
	require.Error(t, err)
	require.ErrorContains(t, err, "address already in use")
	require.Nil(t, l2)

	ep = static.EntryPoint{Address: ":0", ReusePort: true}
	listenConfig = newListenConfig(&ep)
	require.Nil(t, listenConfig.Control)
	require.Zero(t, listenConfig.KeepAlive)

	l3, err := listenConfig.Listen(context.Background(), "tcp", ep.Address)
	require.NoError(t, err)
	require.NotNil(t, l3)
	defer l3.Close()

	l4, err := listenConfig.Listen(context.Background(), "tcp", l3.Addr().String())
	require.Error(t, err)
	require.ErrorContains(t, err, "address already in use")
	require.Nil(t, l4)
}
