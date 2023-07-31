package kusto_dashboard_client

type DashboardDetail struct {
	ID                string             `json:"id"`
	IsDashboardEditor bool               `json:"isDashboardEditor"`
	Title             string             `json:"title"`
	Tiles             []Tile             `json:"tiles"`
	DataSources       []DataSourceDetail `json:"dataSources"`
	ETag              string             `json:"eTag"`
	SchemaVersion     string             `json:"schema_version"`
	AutoRefresh       AutoRefresh        `json:"autoRefresh"`
	SharedQueries     []interface{}      `json:"sharedQueries"`
	Parameters        []Parameter        `json:"parameters"`
	Pages             []Page             `json:"pages"`
}

type Tile struct {
	ID            string        `json:"id"`
	Title         string        `json:"title"`
	VisualType    string        `json:"visualType"`
	PageID        string        `json:"pageId"`
	Layout        Layout        `json:"layout"`
	Query         Query         `json:"query"`
	VisualOptions VisualOptions `json:"visualOptions"`
}

type Layout struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Query struct {
	Kind          string        `json:"kind"`
	DataSource    DataSource    `json:"dataSource"`
	UsedVariables []interface{} `json:"usedVariables"`
	Text          string        `json:"text"`
}

type DataSource struct {
	Kind         string `json:"kind"`
	DataSourceID string `json:"dataSourceId"`
}

type VisualOptions struct {
	HideTileTitle        bool                 `json:"hideTileTitle"`
	MultiStatTextSize    string               `json:"multiStat__textSize"`
	MultiStatValueColumn MultiStatValueColumn `json:"multiStat__valueColumn"`
	ColorRulesDisabled   bool                 `json:"colorRulesDisabled"`
	ColorRules           []interface{}        `json:"colorRules"`
	PieTopNSlices        interface{}          `json:"pie__TopNSlices"`
	ColorStyle           string               `json:"colorStyle"`
}

type MultiStatValueColumn struct {
	Type string `json:"type"`
}

type DataSourceDetail struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ClusterURI string `json:"clusterUri"`
	Database   string `json:"database"`
	Kind       string `json:"kind"`
	ScopeID    string `json:"scopeId"`
}

type AutoRefresh struct {
	DefaultInterval string `json:"defaultInterval"`
	Enabled         bool   `json:"enabled"`
	MinInterval     string `json:"minInterval"`
}

type Parameter struct {
	Kind              string       `json:"kind"`
	ID                string       `json:"id"`
	DisplayName       string       `json:"displayName"`
	BeginVariableName string       `json:"beginVariableName"`
	EndVariableName   string       `json:"endVariableName"`
	DefaultValue      DefaultValue `json:"defaultValue"`
	ShowOnPages       ShowOnPages  `json:"showOnPages"`
	VariableName      string       `json:"variableName"`
	SelectionType     string       `json:"selectionType"`
	IncludeAllOption  bool         `json:"includeAllOption"`
	DataSource        DataSource   `json:"dataSource"`
}

type DefaultValue struct {
	Kind  string `json:"kind"`
	Count int    `json:"count"`
	Unit  string `json:"unit"`
}

type ShowOnPages struct {
	Kind string `json:"kind"`
}

type Page struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
