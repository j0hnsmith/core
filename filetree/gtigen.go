// Code generated by "core generate"; DO NOT EDIT.

package filetree

import (
	"cogentcore.org/core/events"
	"cogentcore.org/core/giv"
	"cogentcore.org/core/gti"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/ki"
	"cogentcore.org/core/mat32"
	"cogentcore.org/core/units"
	"cogentcore.org/core/vci"
)

// NodeType is the [gti.Type] for [Node]
var NodeType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/filetree.Node", IDName: "node", Doc: "Node represents a file in the file system, as a TreeView node.\nThe name of the node is the name of the file.\nFolders have children containing further nodes.", Directives: []gti.Directive{{Tool: "core", Directive: "embedder"}}, Methods: []gti.Method{{Name: "OpenFilesDefault", Doc: "OpenFilesDefault opens selected files with default app for that file type (os defined).\nruns open on Mac, xdg-open on Linux, and start on Windows", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}}, {Name: "DuplicateFiles", Doc: "DuplicateFiles makes a copy of selected files", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}}, {Name: "DeleteFiles", Doc: "deletes any selected files or directories. If any directory is selected,\nall files and subdirectories in that directory are also deleted.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}}, {Name: "RenameFiles", Doc: "renames any selected files", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}}, {Name: "RenameFile", Doc: "RenameFile renames file to new name", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"newpath"}, Returns: []string{"error"}}, {Name: "NewFiles", Doc: "NewFiles makes a new file in selected directory", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"filename", "addToVCS"}}, {Name: "NewFile", Doc: "NewFile makes a new file in this directory node", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"filename", "addToVCS"}}, {Name: "NewFolders", Doc: "makes a new folder in the given selected directory", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"foldername"}}, {Name: "NewFolder", Doc: "NewFolder makes a new folder (directory) in this directory node", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"foldername"}}, {Name: "ShowFileInfo", Doc: "Shows file information about selected file(s)", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}}, {Name: "SortBys", Doc: "SortBys determines how to sort the selected files in the directory.\nDefault is alpha by name, optionally can be sorted by modification time.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"modTime"}}, {Name: "OpenAll", Doc: "OpenAll opens all directories under this one", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}}, {Name: "CloseAll", Doc: "CloseAll closes all directories under this one, this included", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}}, {Name: "RemoveFromExterns", Doc: "RemoveFromExterns removes file from list of external files", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}}, {Name: "AddToVCSSel", Doc: "AddToVCSSel adds selected files to version control system", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}}, {Name: "DeleteFromVCSSel", Doc: "DeleteFromVCSSel removes selected files from version control system", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}}, {Name: "CommitToVCSSel", Doc: "CommitToVCSSel commits to version control system based on last selected file", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}}, {Name: "RevertVCSSel", Doc: "RevertVCSSel removes selected files from version control system", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}}, {Name: "DiffVCSSel", Doc: "DiffVCSSel shows the diffs between two versions of selected files, given by the\nrevision specifiers -- if empty, defaults to A = current HEAD, B = current WC file.\n-1, -2 etc also work as universal ways of specifying prior revisions.\nDiffs are shown in a DiffViewDialog.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"rev_a", "rev_b"}}, {Name: "LogVCSSel", Doc: "LogVCSSel shows the VCS log of commits for selected files, optionally with a\nsince date qualifier: If since is non-empty, it should be\na date-like expression that the VCS will understand, such as\n1/1/2020, yesterday, last year, etc.  SVN only understands a\nnumber as a maximum number of items to return.\nIf allFiles is true, then the log will show revisions for all files, not just\nthis one.\nReturns the Log and also shows it in a VCSLogView which supports further actions.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"allFiles", "since"}}, {Name: "BlameVCSSel", Doc: "BlameVCSSel shows the VCS blame report for this file, reporting for each line\nthe revision and author of the last change.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}}}, Embeds: []gti.Field{{Name: "TreeView"}}, Fields: []gti.Field{{Name: "FPath", Doc: "full path to this file"}, {Name: "Info", Doc: "full standard file info about this file"}, {Name: "Buf", Doc: "file buffer for editing this file"}, {Name: "FRoot", Doc: "root of the tree -- has global state"}, {Name: "DirRepo", Doc: "version control system repository for this directory,\nonly non-nil if this is the highest-level directory in the tree under vcs control"}, {Name: "RepoFiles", Doc: "version control system repository file status -- only valid during ReadDir"}}, Instance: &Node{}})

// NewNode adds a new [Node] with the given name to the given parent:
// Node represents a file in the file system, as a TreeView node.
// The name of the node is the name of the file.
// Folders have children containing further nodes.
func NewNode(par ki.Ki, name ...string) *Node {
	return par.NewChild(NodeType, name...).(*Node)
}

// KiType returns the [*gti.Type] of [Node]
func (t *Node) KiType() *gti.Type { return NodeType }

// New returns a new [*Node] value
func (t *Node) New() ki.Ki { return &Node{} }

// NodeEmbedder is an interface that all types that embed Node satisfy
type NodeEmbedder interface {
	AsNode() *Node
}

// AsNode returns the given value as a value of type Node if the type
// of the given value embeds Node, or nil otherwise
func AsNode(k ki.Ki) *Node {
	if k == nil || k.This() == nil {
		return nil
	}
	if t, ok := k.(NodeEmbedder); ok {
		return t.AsNode()
	}
	return nil
}

// AsNode satisfies the [NodeEmbedder] interface
func (t *Node) AsNode() *Node { return t }

// SetTooltip sets the [Node.Tooltip]
func (t *Node) SetTooltip(v string) *Node { t.Tooltip = v; return t }

// SetText sets the [Node.Text]
func (t *Node) SetText(v string) *Node { t.Text = v; return t }

// SetIcon sets the [Node.Icon]
func (t *Node) SetIcon(v icons.Icon) *Node { t.Icon = v; return t }

// SetIconOpen sets the [Node.IconOpen]
func (t *Node) SetIconOpen(v icons.Icon) *Node { t.IconOpen = v; return t }

// SetIconClosed sets the [Node.IconClosed]
func (t *Node) SetIconClosed(v icons.Icon) *Node { t.IconClosed = v; return t }

// SetIconLeaf sets the [Node.IconLeaf]
func (t *Node) SetIconLeaf(v icons.Icon) *Node { t.IconLeaf = v; return t }

// SetIndent sets the [Node.Indent]
func (t *Node) SetIndent(v units.Value) *Node { t.Indent = v; return t }

// SetOpenDepth sets the [Node.OpenDepth]
func (t *Node) SetOpenDepth(v int) *Node { t.OpenDepth = v; return t }

// SetViewIndex sets the [Node.ViewIndex]
func (t *Node) SetViewIndex(v int) *Node { t.ViewIndex = v; return t }

// SetWidgetSize sets the [Node.WidgetSize]
func (t *Node) SetWidgetSize(v mat32.Vec2) *Node { t.WidgetSize = v; return t }

// SetRootView sets the [Node.RootView]
func (t *Node) SetRootView(v *giv.TreeView) *Node { t.RootView = v; return t }

// SetSelectedNodes sets the [Node.SelectedNodes]
func (t *Node) SetSelectedNodes(v ...giv.TreeViewer) *Node { t.SelectedNodes = v; return t }

// TreeType is the [gti.Type] for [Tree]
var TreeType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/filetree.Tree", IDName: "tree", Doc: "Tree is the root of a tree representing files in a given directory\n(and subdirectories thereof), and has some overall management state for how to\nview things.  The Tree can be viewed by a TreeView to provide a GUI\ninterface into it.", Embeds: []gti.Field{{Name: "Node"}}, Fields: []gti.Field{{Name: "ExtFiles", Doc: "external files outside the root path of the tree -- abs paths are stored -- these are shown in the first sub-node if present -- use AddExtFile to add and update"}, {Name: "Dirs", Doc: "records state of directories within the tree (encoded using paths relative to root),\ne.g., open (have been opened by the user) -- can persist this to restore prior view of a tree"}, {Name: "DirsOnTop", Doc: "if true, then all directories are placed at the top of the tree view\notherwise everything is mixed"}, {Name: "NodeType", Doc: "type of node to create -- defaults to filetree.Node but can use custom node types"}, {Name: "DoubleClickFun", Doc: "DoubleClickFun is a function to call when a node receives a DoubleClick event.\nif not set, defaults to OpenEmptyDir() (for folders)"}, {Name: "InOpenAll", Doc: "if true, we are in midst of an OpenAll call -- nodes should open all dirs"}, {Name: "Watcher", Doc: "change notify for all dirs"}, {Name: "DoneWatcher", Doc: "channel to close watcher watcher"}, {Name: "WatchedPaths", Doc: "map of paths that have been added to watcher -- only active if bool = true"}, {Name: "LastWatchUpdt", Doc: "last path updated by watcher"}, {Name: "LastWatchTime", Doc: "timestamp of last update"}, {Name: "UpdtMu", Doc: "Update mutex"}}, Instance: &Tree{}})

// NewTree adds a new [Tree] with the given name to the given parent:
// Tree is the root of a tree representing files in a given directory
// (and subdirectories thereof), and has some overall management state for how to
// view things.  The Tree can be viewed by a TreeView to provide a GUI
// interface into it.
func NewTree(par ki.Ki, name ...string) *Tree {
	return par.NewChild(TreeType, name...).(*Tree)
}

// KiType returns the [*gti.Type] of [Tree]
func (t *Tree) KiType() *gti.Type { return TreeType }

// New returns a new [*Tree] value
func (t *Tree) New() ki.Ki { return &Tree{} }

// SetDirsOnTop sets the [Tree.DirsOnTop]:
// if true, then all directories are placed at the top of the tree view
// otherwise everything is mixed
func (t *Tree) SetDirsOnTop(v bool) *Tree { t.DirsOnTop = v; return t }

// SetNodeType sets the [Tree.NodeType]:
// type of node to create -- defaults to filetree.Node but can use custom node types
func (t *Tree) SetNodeType(v *gti.Type) *Tree { t.NodeType = v; return t }

// SetDoubleClickFun sets the [Tree.DoubleClickFun]:
// DoubleClickFun is a function to call when a node receives a DoubleClick event.
// if not set, defaults to OpenEmptyDir() (for folders)
func (t *Tree) SetDoubleClickFun(v func(e events.Event)) *Tree { t.DoubleClickFun = v; return t }

// SetTooltip sets the [Tree.Tooltip]
func (t *Tree) SetTooltip(v string) *Tree { t.Tooltip = v; return t }

// SetText sets the [Tree.Text]
func (t *Tree) SetText(v string) *Tree { t.Text = v; return t }

// SetIcon sets the [Tree.Icon]
func (t *Tree) SetIcon(v icons.Icon) *Tree { t.Icon = v; return t }

// SetIconOpen sets the [Tree.IconOpen]
func (t *Tree) SetIconOpen(v icons.Icon) *Tree { t.IconOpen = v; return t }

// SetIconClosed sets the [Tree.IconClosed]
func (t *Tree) SetIconClosed(v icons.Icon) *Tree { t.IconClosed = v; return t }

// SetIconLeaf sets the [Tree.IconLeaf]
func (t *Tree) SetIconLeaf(v icons.Icon) *Tree { t.IconLeaf = v; return t }

// SetIndent sets the [Tree.Indent]
func (t *Tree) SetIndent(v units.Value) *Tree { t.Indent = v; return t }

// SetOpenDepth sets the [Tree.OpenDepth]
func (t *Tree) SetOpenDepth(v int) *Tree { t.OpenDepth = v; return t }

// SetViewIndex sets the [Tree.ViewIndex]
func (t *Tree) SetViewIndex(v int) *Tree { t.ViewIndex = v; return t }

// SetWidgetSize sets the [Tree.WidgetSize]
func (t *Tree) SetWidgetSize(v mat32.Vec2) *Tree { t.WidgetSize = v; return t }

// SetRootView sets the [Tree.RootView]
func (t *Tree) SetRootView(v *giv.TreeView) *Tree { t.RootView = v; return t }

// SetSelectedNodes sets the [Tree.SelectedNodes]
func (t *Tree) SetSelectedNodes(v ...giv.TreeViewer) *Tree { t.SelectedNodes = v; return t }

// VCSLogViewType is the [gti.Type] for [VCSLogView]
var VCSLogViewType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/filetree.VCSLogView", IDName: "vcs-log-view", Doc: "VCSLogView is a view of the VCS log data", Embeds: []gti.Field{{Name: "Layout"}}, Fields: []gti.Field{{Name: "Log", Doc: "current log"}, {Name: "File", Doc: "file that this is a log of -- if blank then it is entire repository"}, {Name: "Since", Doc: "date expression for how long ago to include log entries from"}, {Name: "Repo", Doc: "version control system repository"}, {Name: "RevA", Doc: "revision A -- defaults to HEAD"}, {Name: "RevB", Doc: "revision B -- blank means current working copy"}, {Name: "SetA", Doc: "double-click will set the A revision -- else B"}}, Instance: &VCSLogView{}})

// NewVCSLogView adds a new [VCSLogView] with the given name to the given parent:
// VCSLogView is a view of the VCS log data
func NewVCSLogView(par ki.Ki, name ...string) *VCSLogView {
	return par.NewChild(VCSLogViewType, name...).(*VCSLogView)
}

// KiType returns the [*gti.Type] of [VCSLogView]
func (t *VCSLogView) KiType() *gti.Type { return VCSLogViewType }

// New returns a new [*VCSLogView] value
func (t *VCSLogView) New() ki.Ki { return &VCSLogView{} }

// SetLog sets the [VCSLogView.Log]:
// current log
func (t *VCSLogView) SetLog(v vci.Log) *VCSLogView { t.Log = v; return t }

// SetFile sets the [VCSLogView.File]:
// file that this is a log of -- if blank then it is entire repository
func (t *VCSLogView) SetFile(v string) *VCSLogView { t.File = v; return t }

// SetSince sets the [VCSLogView.Since]:
// date expression for how long ago to include log entries from
func (t *VCSLogView) SetSince(v string) *VCSLogView { t.Since = v; return t }

// SetRepo sets the [VCSLogView.Repo]:
// version control system repository
func (t *VCSLogView) SetRepo(v vci.Repo) *VCSLogView { t.Repo = v; return t }

// SetSetA sets the [VCSLogView.SetA]:
// double-click will set the A revision -- else B
func (t *VCSLogView) SetSetA(v bool) *VCSLogView { t.SetA = v; return t }

// SetTooltip sets the [VCSLogView.Tooltip]
func (t *VCSLogView) SetTooltip(v string) *VCSLogView { t.Tooltip = v; return t }
