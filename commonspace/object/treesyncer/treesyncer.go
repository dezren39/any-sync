//go:generate mockgen -destination mock_treesyncer/mock_treesyncer.go github.com/anyproto/any-sync/commonspace/object/treesyncer TreeSyncer
package treesyncer

import (
	"context"

	"github.com/anyproto/any-sync/app"
)

const CName = "common.object.treesyncer"

type TreeSyncer interface {
	app.ComponentRunnable
	StartSync()
	SyncAll(ctx context.Context, peerId string, existing, missing []string) error
}
