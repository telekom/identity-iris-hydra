// SPDX-FileCopyrightText: 2025 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package iris

import (
	"context"
	"errors"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"

	"github.com/ory/x/networkx"
	"github.com/ory/x/sqlcon"
)

type NetworkManager struct {
	c   *pop.Connection
	nid uuid.UUID
}

func NewNetworkManager(
	c *pop.Connection,
	nid uuid.UUID,
) *NetworkManager {
	return &NetworkManager{
		c:   c,
		nid: nid,
	}
}

func (m *NetworkManager) getNetwork(ctx context.Context) (*networkx.Network, error) {
	var p networkx.Network
	c := m.c.WithContext(ctx)
	if err := sqlcon.HandleError(c.Q().Find(&p, m.nid)); err != nil {
		if errors.Is(err, sqlcon.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (m *NetworkManager) addNetwork(ctx context.Context) error {
	if m.nid == uuid.Nil {
		return errors.New("network manager cannot be used without a valid network ID")
	}
	c := m.c.WithContext(ctx)
	network := networkx.Network{ID: m.nid, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	return sqlcon.HandleError(c.Create(&network))
}

func (m *NetworkManager) Determine(ctx context.Context) (*networkx.Network, error) {
	network, err := m.getNetwork(ctx)
	if err != nil {
		return nil, err
	}
	if network == nil {
		err = m.addNetwork(ctx)
		if err != nil {
			return nil, err
		}
		return m.getNetwork(ctx)
	}
	return network, nil
}
