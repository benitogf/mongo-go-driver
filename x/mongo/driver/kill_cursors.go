// Copyright (C) MongoDB, Inc. 2017-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package driver

import (
	"context"

	"github.com/benitogf/mongo-go-driver/x/mongo/driver/topology"
	"github.com/benitogf/mongo-go-driver/x/network/command"
	"github.com/benitogf/mongo-go-driver/x/network/description"
	"github.com/benitogf/mongo-go-driver/x/network/result"
)

// KillCursors handles the full cycle dispatch and execution of an aggregate command against the provided
// topology.
func KillCursors(
	ctx context.Context,
	cmd command.KillCursors,
	topo *topology.Topology,
	selector description.ServerSelector,
) (result.KillCursors, error) {
	ss, err := topo.SelectServer(ctx, selector)
	if err != nil {
		return result.KillCursors{}, err
	}
	desc := ss.Description()
	conn, err := ss.Connection(ctx)
	if err != nil {
		return result.KillCursors{}, err
	}
	defer conn.Close()
	return cmd.RoundTrip(ctx, desc, conn)
}
