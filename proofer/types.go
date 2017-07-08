package proofer

const (
	URL     string = `https://api.meaningcloud.com/stilus-1.2?`
	OF_JSON string = `json`
	OF_XML  string = `xml`
)

type Proofer struct {
	MeaningcloudKey string
}

type StatusData struct {
	//	* 0: OK
	//	* 100: Operation denied
	//	* 101: License expired
	//	* 102: Credits per subscription exceeded
	//	* 103: Request too large
	//	* 104: Request rate limit exceeded
	//	* 200: Missing required parameter(s) - [name of the parameter]
	//	* 202: Engine internal error
	//	* 203: Cannot connect to service
	//	* 205: Language not supported
	//	* 212: No content to analyze
	//	* 215: Timeout exceeded for service response
	Code             string `json:"code"`
	Msg              string `json:"msg"`
	Credits          string `json:"credits"`
	RemainingCredits string `json:"remaining_credits"`
}

type Result struct {
	Status      StatusData     `json:"status"`
	ResultLists []ResultList   `json:"result_list"`
	Statistics  StatisticsData `json:"statistics"`
}

type ResultList struct {
	// text: text in which the error or warning is located.
	Text string `json:"text"`
	//	type: type of issue. There are five types or errors defined:
	//		+ S: spelling error or warning
	//		+ G: grammar error or warning
	//		+ Y: style error or warning
	//		+ T: typography error or warning
	//		+ M: semantic-related error or warning
	Type string `json:"type"`
	// inip: initial position of the result, starting on 0
	Inip string `json:"inip"`
	// end position of the result
	Endp string `json:"endp"`
	// level: CEFR level of the error.
	Level string `json:"level"`
	// rule:  ID of the rule associated to the issue
	Rule string `json:"rule"`
	// msg: message that explains the issue
	Msg         string    `json:"msg"`
	Suggestions []SugList `json:"sug_list"`
}

type SugList struct {
	Form       string `json:"form"`
	Confidence string `json:"confidence"`
}

type StatisticsData struct {
	Words                   string `json:"words"`
	NonRepeatedWords        string `json:"non-repeated_words"`
	CharactersWithoutSpaces string `json:"characters_without_spaces"`
	TotalCharacters         string `json:"total_characters"`
	Sentences               string `json:"sentences"`
	Paragraphs              string `json:"paragraphs"`
}
