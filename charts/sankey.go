package charts

import (
	"github.com/go-echarts/go-echarts/types"
	"io"
)

// Sankey represents a sankey chart.
type Sankey struct {
	BaseConfiguration
	MultiSeries
}

// SankeyLink represents relationship between two data nodes.
type SankeyLink struct {
	// 边的源节点名称的字符串，也支持使用数字表示源节点的索引
	Source interface{} `json:"source,omitempty"`
	// 边的目标节点名称的字符串，也支持使用数字表示源节点的索引
	Target interface{} `json:"target,omitempty"`
	// 边的数值，可以在力引导布局中用于映射到边的长度
	Value float32 `json:"value,omitempty"`
}

// SankeyNode represents a data node.
type SankeyNode struct {
	// 数据项名称
	Name string `json:"name,omitempty"`
	// 数据项值
	Value string `json:"value,omitempty"`
}

func (Sankey) Type() string { return types.ChartSankey }

// NewSankey creates a new sankey chart.
func NewSankey() *Sankey {
	chart := new(Sankey)
	chart.initBaseConfiguration()
	return chart
}

// Add adds new data sets.
func (c *Sankey) Add(name string, nodes []SankeyNode, links []SankeyLink, opts ...SeriesOpts) *Sankey {
	series := SingleSeries{Name: name, Type: types.ChartSankey, Data: nodes, Links: links}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Sankey instance.
func (c *Sankey) SetGlobalOptions(opts ...GlobalOpts) *Sankey {
	c.BaseConfiguration.setBaseGlobalOptions(opts...)
	return c
}

func (c *Sankey) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *Sankey) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
