package migrator

type RuleData struct {
	RefID             string `json:"refId"`
	RelativeTimeRange struct {
		From int `json:"from"`
		To   int `json:"to"`
	} `json:"relativeTimeRange"`
	DatasourceUID string `json:"datasourceUid"`
	Model         struct {
		EditorMode    string `json:"editorMode"`
		Expr          string `json:"expr"`
		Instant       bool   `json:"instant"`
		IntervalMs    int    `json:"intervalMs"`
		LegendFormat  string `json:"legendFormat"`
		MaxDataPoints int    `json:"maxDataPoints"`
		Range         bool   `json:"range"`
		RefID         string `json:"refId"`
	} `json:"model"`
	Model10 struct {
		Conditions []struct {
			Evaluator struct {
				Params []int  `json:"params"`
				Type   string `json:"type"`
			} `json:"evaluator"`
			Operator struct {
				Type string `json:"type"`
			} `json:"operator"`
			Query struct {
				Params []string `json:"params"`
			} `json:"query"`
			Reducer struct {
				Params []interface{} `json:"params"`
				Type   string        `json:"type"`
			} `json:"reducer"`
			Type string `json:"type"`
		} `json:"conditions"`
		Datasource struct {
			Type string `json:"type"`
			UID  string `json:"uid"`
		} `json:"datasource"`
		Expression    string `json:"expression"`
		IntervalMs    int    `json:"intervalMs"`
		MaxDataPoints int    `json:"maxDataPoints"`
		RefID         string `json:"refId"`
		Type          string `json:"type"`
	} `json:"model10"`
}

type Rule struct {
	UID          string     `json:"uid"`
	Title        string     `json:"title"`
	Condition    string     `json:"condition"`
	Data         []RuleData `json:"data"`
	NoDataState  string     `json:"noDataState"`
	ExecErrState string     `json:"execErrState"`
	For          string     `json:"for"`
	Annotations  struct {
		Description string `json:"description"`
		RunbookURL  string `json:"runbook_url"`
		Summary     string `json:"summary"`
	} `json:"annotations"`
	Labels   struct{} `json:"labels"`
	IsPaused bool     `json:"isPaused"`
}

type Group struct {
	OrgID    int    `json:"orgId"`
	Name     string `json:"name"`
	Folder   string `json:"folder"`
	Interval string `json:"interval"`
	Rules    []Rule `json:"rules"`
}

type Response struct {
	APIVersion int     `json:"apiVersion"`
	Groups     []Group `json:"groups"`
}
