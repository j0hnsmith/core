// Code generated by "core generate"; DO NOT EDIT.

package tree

import (
	"cogentcore.org/core/types"
)

// NodeBaseType is the [types.Type] for [NodeBase]
var NodeBaseType = types.AddType(&types.Type{Name: "cogentcore.org/core/tree.NodeBase", IDName: "node-base", Doc: "NodeBase implements the [Node] interface and provides the core functionality\nfor the Cogent Core tree system. You must use NodeBase as an embedded struct\nin all higher-level tree types.\n\nAll nodes must be properly initialized by using one of [New], [NodeBase.NewChild],\n[NodeBase.AddChild], [NodeBase.InsertChild], [NodeBase.InsertNewChild],\n[NodeBase.Clone], [Update], or [cogentcore.org/core/core.Plan]. This ensures\nthat the [Node.This] field is set correctly and that the [Node.Init] method is\ncalled.\n\nAll nodes support JSON marshalling and unmarshalling through the standard [encoding/json]\ninterfaces, so you can use the standard functions for loading and saving trees. However,\nif you want to load a root node of the correct type from JSON, you need to use the\n[UnmarshalRootJSON] function.\n\nAll node types must be added to the Cogent Core type registry via typegen,\nso you must add a go:generate line that runs `core generate` to any packages\nyou write that have new node types defined.", Fields: []types.Field{{Name: "Name", Doc: "Name is the name of this node, which is typically unique relative to other children of\nthe same parent. It can be used for finding and serializing nodes. If not otherwise set,\nit defaults to the ID (kebab-case) name of the node type combined with the total number\nof children that have ever been added to the node's parent."}, {Name: "This", Doc: "This is the value of this Node as its true underlying type. This allows methods\ndefined on base types to call methods defined on higher-level types, which\nis necessary for various parts of tree and widget functionality. This is set\nto nil when the node is deleted."}, {Name: "Parent", Doc: "Parent is the parent of this node, which is set automatically when this node is\nadded as a child of a parent. To change the parent of a node, use [MoveToParent];\nyou should typically not set this field directly. Nodes can only have one parent\nat a time."}, {Name: "Children", Doc: "Children is the list of children of this node. All of them are set to have this node\nas their parent. You can directly modify this list, but you should typically use the\nvarious NodeBase child helper functions when applicable so that everything is updated\nproperly, such as when deleting children."}, {Name: "Properties", Doc: "Properties is a property map for arbitrary key-value properties.\nWhen possible, use typed fields on a new type embedding NodeBase instead of this.\nYou should typically use the [NodeBase.SetProperty], [NodeBase.Property], and\n[NodeBase.DeleteProperty] methods for modifying and accessing properties."}, {Name: "numLifetimeChildren", Doc: "numLifetimeChildren is the number of children that have ever been added to this\nnode, which is used for automatic unique naming."}, {Name: "index", Doc: "index is the last value of our index, which is used as a starting point for\nfinding us in our parent next time. It is not guaranteed to be accurate;\nuse the [NodeBase.IndexInParent] method."}, {Name: "depth", Doc: "depth is the depth of the node while using [NodeBase.WalkDownBreadth]."}}, Instance: &NodeBase{}})

// NewNodeBase returns a new [NodeBase] with the given optional parent:
// NodeBase implements the [Node] interface and provides the core functionality
// for the Cogent Core tree system. You must use NodeBase as an embedded struct
// in all higher-level tree types.
//
// All nodes must be properly initialized by using one of [New], [NodeBase.NewChild],
// [NodeBase.AddChild], [NodeBase.InsertChild], [NodeBase.InsertNewChild],
// [NodeBase.Clone], [Update], or [cogentcore.org/core/core.Plan]. This ensures
// that the [Node.This] field is set correctly and that the [Node.Init] method is
// called.
//
// All nodes support JSON marshalling and unmarshalling through the standard [encoding/json]
// interfaces, so you can use the standard functions for loading and saving trees. However,
// if you want to load a root node of the correct type from JSON, you need to use the
// [UnmarshalRootJSON] function.
//
// All node types must be added to the Cogent Core type registry via typegen,
// so you must add a go:generate line that runs `core generate` to any packages
// you write that have new node types defined.
func NewNodeBase(parent ...Node) *NodeBase { return New[*NodeBase](parent...) }

// NodeType returns the [*types.Type] of [NodeBase]
func (t *NodeBase) NodeType() *types.Type { return NodeBaseType }

// New returns a new [*NodeBase] value
func (t *NodeBase) New() Node { return &NodeBase{} }

// SetName sets the [NodeBase.Name]:
// Name is the name of this node, which is typically unique relative to other children of
// the same parent. It can be used for finding and serializing nodes. If not otherwise set,
// it defaults to the ID (kebab-case) name of the node type combined with the total number
// of children that have ever been added to the node's parent.
func (t *NodeBase) SetName(v string) *NodeBase { t.Name = v; return t }
