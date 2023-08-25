package sarif

type Sarif struct {
	Schema  string `json:"$schema"`
	Version string `json:"version"`
	Runs    []*Run `json:"runs"`
}

type Run struct {
	Invocations []*Invocation `json:"invocations"`
	Results     []*Result     `json:"results"`
	Tool        struct {
		Driver struct {
			Name            string  `json:"name"`
			FullName        string  `json:"fullName,omitempty"`
			InformationUri  string  `json:"informationUri,omitempty"`
			SemantecVersion string  `json:"semanticVersion"`
			Rules           []*Rule `json:"rules"`
		} `json:"driver"`
	} `json:"tool"`
}

type Invocation struct {
	ExecutionSuccessful        bool            `json:"executionSuccessful"`
	ToolExecutionNotifications []*Notification `json:"toolExecutionNotifications"`
}

type Notification struct {
	Descriptior struct {
		ID string `json:"id"`
	} `json:"descriptor"`
	Level   string `json:"level"`
	Message struct {
		Text string `json:"text"`
	} `json:"message"`
}

type Rule struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	ShortDescription struct {
		Text string `json:"text"`
	} `json:"shortDescription"`
	FullDescription struct {
		Text string `json:"text"`
	} `json:"fullDescription"`
	DefaultConfiguration struct {
		Level string `json:"level"`
	} `json:"defaultConfiguration"`
	Properties ruleProperties `json:"properties"`
	HelpURI    string         `json:"helpUri"`
	Help       struct {
		Text     string `json:"text"`
		Markdown string `json:"markdown"`
	} `json:"help"`
}

type ruleProperties struct {
	Precision        string   `json:"precision"`
	Tags             []string `json:"tags"`
	SecuritySeverity string   `json:"security-severity"`
}

type Result struct {
	RuleID    string `json:"ruleId"`
	RuleIndex int    `json:"ruleIndex"`
	Level     string `json:"level"`
	Message   struct {
		Text string `json:"text"`
	} `json:"message"`
	Locations []struct {
		PhysicalLocation struct {
			ArtifactLocation struct {
				URI       string `json:"uri"`
				URIBaseID string `json:"uriBaseId"`
			} `json:"artifactLocation"`
			Region struct {
				StartLine   int32 `json:"startLine"`
				StartColumn int32 `json:"startColumn"`
				EndLine     int32 `json:"endLine"`
				EndColumn   int32 `json:"endColumn"`
			} `json:"region"`
		} `json:"physicalLocation"`
	} `json:"locations"`
	Suppressions []struct {
		Kind   string `json:"kind"`
		Status string `json:"status,omitempty"`
		GUID   string `json:"guid,omitempty"`
	} `json:"suppressions,omitempty"`
}
