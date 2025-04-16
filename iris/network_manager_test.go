// Copyright 2025 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package iris

import (
	"context"
	"testing"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/x/dbal"
	"github.com/ory/x/networkx"
)

func initConnection(t *testing.T) *pop.Connection {
	c, err := pop.NewConnection(&pop.ConnectionDetails{URL: dbal.NewSQLiteInMemoryDatabase("test")})
	require.NoError(t, err)
	require.NoError(t, c.Open())

	err = c.RawQuery("create table networks (id uuid primary_key, created_at datetime, updated_at datetime)").Exec()
	require.NoError(t, err)

	return c
}

func TestDetermineNetwork(t *testing.T) {
	ctx := context.Background()
	nid := uuid.FromStringOrNil("4b26c191-e56f-4513-ba16-55c1f32d689b")
	c := initConnection(t)
	defer c.Close()
	m := NewNetworkManager(c, nid)

	net, err := m.Determine(ctx)
	require.NoError(t, err)
	assert.Equal(t, net.ID, nid)
}

func TestAddNetwork(t *testing.T) {
	ctx := context.Background()
	nid := uuid.FromStringOrNil("4b26c191-e56f-4513-ba16-55c1f32d689b")
	c := initConnection(t)
	defer c.Close()
	m := NewNetworkManager(c, nid)

	err := m.addNetwork(ctx)
	require.NoError(t, err)

	var network networkx.Network
	err = c.Find(&network, nid)
	require.NoError(t, err)
	require.Equal(t, nid, network.ID, nid)
}

func TestAddNetworkNilID(t *testing.T) {
	ctx := context.Background()
	nid := uuid.Nil
	c := initConnection(t)
	defer c.Close()
	m := NewNetworkManager(c, nid)

	_, err := m.Determine(ctx)
	require.Error(t, err)
}

func TestGetNetworkNotExists(t *testing.T) {
	ctx := context.Background()
	nid := uuid.FromStringOrNil("4b26c191-e56f-4513-ba16-55c1f32d689b")
	c := initConnection(t)
	defer c.Close()
	m := NewNetworkManager(c, nid)

	network, err := m.getNetwork(ctx)
	require.NoError(t, err)
	require.Nil(t, network)
}

func TestGetNetworkExists(t *testing.T) {
	ctx := context.Background()
	nid := uuid.FromStringOrNil("4b26c191-e56f-4513-ba16-55c1f32d689b")
	c := initConnection(t)
	defer c.Close()
	m := NewNetworkManager(c, nid)

	err := m.addNetwork(ctx)
	require.NoError(t, err)

	network, err := m.getNetwork(ctx)
	require.NoError(t, err)
	require.Equal(t, network.ID, nid)
}
