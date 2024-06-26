// Code generated by "core generate"; DO NOT EDIT.

package plotcore

import (
	"cogentcore.org/core/core"
	"cogentcore.org/core/tree"
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot/plotcore.PlotParams", IDName: "plot-params", Doc: "PlotParams are parameters for overall plot", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Fields: []types.Field{{Name: "Title", Doc: "optional title at top of plot"}, {Name: "Type", Doc: "type of plot to generate.  For a Bar plot, items are plotted ordinally by row and the XAxis is optional"}, {Name: "Lines", Doc: "whether to plot lines"}, {Name: "Points", Doc: "whether to plot points with symbols"}, {Name: "LineWidth", Doc: "width of lines"}, {Name: "PointSize", Doc: "size of points"}, {Name: "PointShape", Doc: "the shape used to draw points"}, {Name: "BarWidth", Doc: "width of bars for bar plot, as fraction of available space (1 = no gaps)"}, {Name: "NegativeXDraw", Doc: "if true, draw lines that connect points with a negative X-axis direction;\notherwise there is a break in the line.\ndefault is false, so that repeated series of data across the X axis\nare plotted separately."}, {Name: "Scale", Doc: "Scale multiplies the plot DPI value, to change the overall scale\nof the rendered plot.  Larger numbers produce larger scaling.\nTypically use larger numbers when generating plots for inclusion in\ndocuments or other cases where the overall plot size will be small."}, {Name: "XAxisColumn", Doc: "what column to use for the common X axis. if empty or not found,\nthe row number is used.  This optional for Bar plots, if present and\nLegendColumn is also present, then an extra space will be put between X values."}, {Name: "LegendColumn", Doc: "optional column for adding a separate colored / styled line or bar\naccording to this value, and acts just like a separate Y variable,\ncrossed with Y variables."}, {Name: "LegendPosition", Doc: "position of the Legend"}, {Name: "XAxisRotation", Doc: "rotation of the X Axis labels, in degrees"}, {Name: "XAxisLabel", Doc: "optional label to use for XAxis instead of column name"}, {Name: "YAxisLabel", Doc: "optional label to use for YAxis -- if empty, first column name is used"}, {Name: "Plot", Doc: "our plot, for update method"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot/plotcore.ColumnParams", IDName: "column-params", Doc: "ColumnParams are parameters for plotting one column of data", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Fields: []types.Field{{Name: "On", Doc: "whether to plot this column"}, {Name: "Column", Doc: "name of column we're plotting"}, {Name: "Lines", Doc: "whether to plot lines; uses the overall plot option if unset"}, {Name: "Points", Doc: "whether to plot points with symbols; uses the overall plot option if unset"}, {Name: "LineWidth", Doc: "the width of lines; uses the overall plot option if unset"}, {Name: "PointSize", Doc: "the size of points; uses the overall plot option if unset"}, {Name: "PointShape", Doc: "the shape used to draw points; uses the overall plot option if unset"}, {Name: "Range", Doc: "effective range of data to plot -- either end can be fixed"}, {Name: "FullRange", Doc: "full actual range of data -- only valid if specifically computed"}, {Name: "Color", Doc: "color to use when plotting the line / column"}, {Name: "NTicks", Doc: "desired number of ticks"}, {Name: "Lbl", Doc: "if non-empty, this is an alternative label to use in plotting"}, {Name: "TensorIndex", Doc: "if column has n-dimensional tensor cells in each row, this is the index within each cell to plot -- use -1 to plot *all* indexes as separate lines"}, {Name: "ErrColumn", Doc: "specifies a column containing error bars for this column"}, {Name: "IsString", Doc: "if true this is a string column -- plots as labels"}, {Name: "Plot", Doc: "our plot, for update method"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot/plotcore.Plot", IDName: "plot", Doc: "Plot is a widget that renders a [plot.Plot] object.\nIf it is not [states.ReadOnly], the user can pan and zoom the graph.\nSee [PlotEditor] for an interactive interface for selecting columns to view.", Methods: []types.Method{{Name: "SavePlot", Doc: "SaveSVG saves the current Plot to an SVG file", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"filename"}, Returns: []string{"error"}}, {Name: "SavePNG", Doc: "SavePNG saves the current rendered Plot image to an PNG image file.", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"filename"}, Returns: []string{"error"}}}, Embeds: []types.Field{{Name: "WidgetBase"}}, Fields: []types.Field{{Name: "Scale", Doc: "Scale multiplies the plot DPI value, to change the overall scale\nof the rendered plot.  Larger numbers produce larger scaling.\nTypically use larger numbers when generating plots for inclusion in\ndocuments or other cases where the overall plot size will be small."}, {Name: "Plot", Doc: "Plot is the Plot to display in this widget"}}})

// NewPlot returns a new [Plot] with the given optional parent:
// Plot is a widget that renders a [plot.Plot] object.
// If it is not [states.ReadOnly], the user can pan and zoom the graph.
// See [PlotEditor] for an interactive interface for selecting columns to view.
func NewPlot(parent ...tree.Node) *Plot { return tree.New[Plot](parent...) }

// SetScale sets the [Plot.Scale]:
// Scale multiplies the plot DPI value, to change the overall scale
// of the rendered plot.  Larger numbers produce larger scaling.
// Typically use larger numbers when generating plots for inclusion in
// documents or other cases where the overall plot size will be small.
func (t *Plot) SetScale(v float32) *Plot { t.Scale = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot/plotcore.PlotEditor", IDName: "plot-editor", Doc: "PlotEditor is a widget that provides an interactive 2D plot\nof selected columns of tabular data, represented by a [table.IndexView] into\na [table.Table]. Other types of tabular data can be converted into this format.\nThe user can change various options for the plot and also modify the underlying data.", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Methods: []types.Method{{Name: "SaveSVG", Doc: "SaveSVG saves the plot to an svg -- first updates to ensure that plot is current", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"fname"}}, {Name: "SavePNG", Doc: "SavePNG saves the current plot to a png, capturing current render", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"fname"}}, {Name: "SaveCSV", Doc: "SaveCSV saves the Table data to a csv (comma-separated values) file with headers (any delim)", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"fname", "delim"}}, {Name: "SaveAll", Doc: "SaveAll saves the current plot to a png, svg, and the data to a tsv -- full save\nAny extension is removed and appropriate extensions are added", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"fname"}}, {Name: "OpenCSV", Doc: "OpenCSV opens the Table data from a csv (comma-separated values) file (or any delim)", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"filename", "delim"}}, {Name: "SetColumnsByName", Doc: "SetColumnsByName turns cols On or Off if their name contains given string", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"nameContains", "on"}}}, Embeds: []types.Field{{Name: "Frame"}}, Fields: []types.Field{{Name: "Table", Doc: "the table of data being plotted"}, {Name: "Params", Doc: "the overall plot parameters"}, {Name: "Columns", Doc: "the parameters for each column of the table"}, {Name: "Plot", Doc: "Plot is the plot object."}, {Name: "ConfigPlotFunc", Doc: "ConfigPlotFunc is a function to call to configure [PlotEditor.Plot], the plot.Plot that\nactually does the plotting. It is called after [Plot] is generated, and properties\nof [Plot] can be modified in it. Properties of [Plot] should not be modified outside\nof this function, as doing so will have no effect."}, {Name: "SVGFile", Doc: "current svg file"}, {Name: "DataFile", Doc: "current csv data file"}, {Name: "InPlot", Doc: "currently doing a plot"}}})

// NewPlotEditor returns a new [PlotEditor] with the given optional parent:
// PlotEditor is a widget that provides an interactive 2D plot
// of selected columns of tabular data, represented by a [table.IndexView] into
// a [table.Table]. Other types of tabular data can be converted into this format.
// The user can change various options for the plot and also modify the underlying data.
func NewPlotEditor(parent ...tree.Node) *PlotEditor { return tree.New[PlotEditor](parent...) }

// SetParams sets the [PlotEditor.Params]:
// the overall plot parameters
func (t *PlotEditor) SetParams(v PlotParams) *PlotEditor { t.Params = v; return t }

// SetConfigPlotFunc sets the [PlotEditor.ConfigPlotFunc]:
// ConfigPlotFunc is a function to call to configure [PlotEditor.Plot], the plot.Plot that
// actually does the plotting. It is called after [Plot] is generated, and properties
// of [Plot] can be modified in it. Properties of [Plot] should not be modified outside
// of this function, as doing so will have no effect.
func (t *PlotEditor) SetConfigPlotFunc(v func()) *PlotEditor { t.ConfigPlotFunc = v; return t }

// SetSVGFile sets the [PlotEditor.SVGFile]:
// current svg file
func (t *PlotEditor) SetSVGFile(v core.Filename) *PlotEditor { t.SVGFile = v; return t }

// SetDataFile sets the [PlotEditor.DataFile]:
// current csv data file
func (t *PlotEditor) SetDataFile(v core.Filename) *PlotEditor { t.DataFile = v; return t }
