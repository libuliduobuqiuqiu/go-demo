package goothers

type FlowConfig struct {
	FlowChain      any          `json:"flowChain"`
	FlowChainNodes []*ChainNode `json:"flowChainNodes"`
	StoreList      []any        `json:"storeList"`
}

type ChainNode struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ChainID   string `json:"chainId"`
	ChainType string `json:"chainType"`
	NodeTag   string `json:"nodeTag"`
	NodeType  string `json:"nodeType"`
	Props     Props  `json:"props"`
}

type Props struct {
	ChainID string `json:"chainId"`
	NodeTag string `json:"nodeTag"`

	Method   string                   `json:"method,omitempty"`
	URI      string                   `json:"uri,omitempty"`
	Props    []InternalProps          `json:"props,omitempty"`
	CodeType string                   `json:"codeType,omitempty"`
	Context  string                   `json:"context,omitempty"`
	Param    map[string]string        `json:"param,omitempty"`
	Header   map[string]string        `json:"header,omitempty"`
	Body     string                   `json:"body,omitempty"`
	Auth     httpExecuteAuthorization `json:"authorization,omitempty"`
	ResField string                   `json:"resField,omitempty"`

	ElExpression string `json:"elExpression,omitempty"`
	PrintType    string `json:"printType,omitempty"`
	PrintType0   string `json:"$printType,omitempty"`
	RespCode     string `json:"respCode,omitempty"`

	Targets []Targets `json:"targets,omitempty"`

	Type     string `json:"type,omitempty"`
	Type0    string `json:"$type,omitempty"`
	FileID   string `json:"fileId,omitempty"`
	Password string `json:"password,omitempty"`
	Port     int    `json:"port,omitempty"`
	IP       string `json:"ip,omitempty"`
	FilePath string `json:"filePath,omitempty"`
	Prop     string `json:"prop,omitempty"`
	Username string `json:"username,omitempty"`
}

type Targets struct {
	NodeCase string `json:"nodeCase"`
	NodeTag  string `json:"nodeTag"`
}

type InternalProps struct {
	Field    string `json:"field"`
	DataType string `json:"dataType"`
	BodyType string `json:"bodyType,omitempty"`
}

type httpExecuteAuthorization struct {
	AuthType string `json:"type,omitempty"`
	AuthBody struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
		Token    string `json:"token,omitempty"`
	}
}
