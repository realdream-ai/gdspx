package engine

import (
	"grow.graphics/gd"
	_ "grow.graphics/gd/gdextension"
)

type EngineNode struct {
	gd.Class[EngineNode, gd.Node2D] `gd:"EngineNode"`
}

func (pself *EngineNode) Ready() {
	initEngine(pself)
}

func (pself *EngineNode) Process(delta gd.Float) {
	tickEngine(delta)
}
