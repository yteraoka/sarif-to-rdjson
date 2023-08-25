package rdjson

type DiagnosticResult struct {
	Diagnostics []*Diagnostic `json:"diagnostics,omitempty"`
	Source      *Source       `json:"source,omitempty"`
	Severity    string        `json:"severity,omitempty"`
}

type Diagnostic struct {
	Message        string        `json:"message,omitempty"`
	Location       *Location     `json:"location,omitempty"`
	Severity       string        `json:"severity,omitempty"`
	Source         *Source       `json:"source,omitempty"`
	Code           *Code         `json:"code,omitempty"`
	Suggestions    []*Suggestion `json:"suggestions,omitempty"`
	OriginalOutput string        `json:"original_output,omitempty"`
}

type Location struct {
	Path  string `json:"path,omitempty"`
	Range *Range `json:"range,omitempty"`
}

type Range struct {
	Start *Position `json:"start,omitempty"`
	End   *Position `json:"end,omitempty"`
}

type Position struct {
	Line   int32 `json:"line,omitempty"`
	Column int32 `json:"column,omitempty"`
}

type Source struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Code struct {
	Value string `json:"value,omitempty"`
	Url   string `json:"url,omitempty"`
}

type Suggestion struct {
	Range *Range `json:"range,omitempty"`
	Text  string `json:"text,omitempty"`
}
