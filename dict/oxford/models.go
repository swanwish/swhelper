package oxford

type Lemmatron struct {
	Metadata interface{}         `json:"metadata"` // Additional Information provided by OUP ,
	Results  []HeadwordLemmatron `json:"results"`  // A list of inflections matching a given word
}

type HeadwordLemmatron struct {
	Id             string                  `json:"id"`             // The identifier of a word ,
	Language       string                  `json:"language"`       // IANA language code ,
	LexicalEntries []LemmatronLexicalEntry `json:"lexicalEntries"` // A grouping of various senses in a specific language, and a lexical category that relates to a word ,
	Type           string                  `json:"type"`           // The json object type. Could be 'headword', 'inflection' or 'phrase' ,
	Word           string                  `json:"word"`           // A given written or spoken realisation of a an entry, lowercased.
}

type LemmatronLexicalEntry struct {
	GrammaticalFeatures GrammaticalFeaturesList `json:"grammaticalFeatures"`
	InflectionOf        InflectionsList         `json:"inflectionOf"`    // The canonical form of words for which the entry is an inflection ,
	Language            string                  `json:"language"`        // IANA language code ,
	LexicalCategory     string                  `json:"lexicalCategory"` // A linguistic category of words (or more precisely lexical items), generally defined by the syntactic or morphological behaviour of the lexical item in question, such as noun or verb ,
	Text                string                  `json:"text"`
}

type GrammaticalFeaturesList []GrammaticalFeature

type GrammaticalFeature struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type InflectionsList []Inflection

type Inflection struct {
	Id   string `json:"id"` // The identifier of the word ,
	Text string `json:"text"`
}

type RetrieveEntry struct {
	Metadata interface{}     `json:"metadata"` // metadata (object, optional): Additional Information provided by OUP ,
	Results  []HeadwordEntry `json:"results"`  // results (Array[HeadwordEntry], optional): A list of entries and all the data related to them
}

type HeadwordEntry struct {
	Id             string             `json:"id"`             // id (string): The identifier of a word ,
	Language       string             `json:"language"`       // language (string): IANA language code ,
	LexicalEntries []LexicalEntry     `json:"lexicalEntries"` // exicalEntries (Array[lexicalEntry]): A grouping of various senses in a specific language, and a lexical category that relates to a word ,
	Pronunciations PronunciationsList `json:"pronunciations"` // pronunciations (PronunciationsList, optional),
	Type           string             `json:"type"`           // type (string, optional): The json object type. Could be 'headword', 'inflection' or 'phrase' ,
	Word           string             `json:"word"`           //word (string): A given written or spoken realisation of a an entry, lowercased.
}

type LexicalEntry struct {
	DerivativeOf        ArrayOfRelatedEntries   `json:"derivativeOf"`            // derivativeOf (ArrayOfRelatedEntries, optional): Other words from which this one derives ,
	Derivatives         ArrayOfRelatedEntries   `json:"derivatives"`             // derivatives (ArrayOfRelatedEntries, optional): Other words from which their Sense derives ,
	Entries             []Entry                 `json:"entries"`                 // entries (Array[Entry], optional),
	GrammaticalFeatures GrammaticalFeaturesList `json:"GrammaticalFeaturesList"` // grammaticalFeatures (GrammaticalFeaturesList, optional),
	Language            string                  `json:"language"`                // language (string): IANA language code ,
	LexicalCategory     string                  `json:"lexicalCategory"`         // lexicalCategory (string): A linguistic category of words (or more precisely lexical items), generally defined by the syntactic or morphological behaviour of the lexical item in question, such as noun or verb ,
	Notes               CategorizedTextList     `json:"notes"`                   // notes (CategorizedTextList, optional),
	Pronunciations      PronunciationsList      `json:"pronunciations"`          // pronunciations (PronunciationsList, optional),
	Text                string                  `json:"text"`                    // text (string): A given written or spoken realisation of a an entry. ,
	VariantForms        VariantFormsList        `json:"variantForms"`            // variantForms (VariantFormsList, optional): Various words that are used interchangeably depending on the context, e.g 'a' and 'an'
}

type PronunciationsList []Pronunciation

type ArrayOfRelatedEntries []RelatedEntry

type Entry struct {
	Etymologies         []string                `json:"etymologies"`         // etymologies (arrayofstrings, optional): The origin of the word and the way in which its meaning has changed throughout history ,
	GrammaticalFeatures GrammaticalFeaturesList `json:"grammaticalFeatures"` // grammaticalFeatures (GrammaticalFeaturesList, optional),
	HomographNumber     string                  `json:"homographNumber"`     // homographNumber (string, optional): Identifies the homograph grouping. The last two digits identify different entries of the same homograph. The first one/two digits identify the homograph number. ,
	Notes               CategorizedTextList     `json:"notes"`               // notes (CategorizedTextList, optional),
	Pronunciations      PronunciationsList      `json:"pronunciations"`      // pronunciations (PronunciationsList, optional),
	Senses              []Sense                 `json:"senses"`              // senses (Array[Sense], optional): Complete list of senses ,
	VariantForms        VariantFormsList        `json:"variantForms"`        // variantForms (VariantFormsList, optional): Various words that are used interchangeably depending on the context, e.g 'a' and 'an'
}

type CategorizedTextList []CategorizedText

type VariantFormsList []VariantForm

type Pronunciation struct {
	AudioFile        string   `json:"audioFile"`        // audioFile (string, optional): The URL of the sound file ,
	Dialects         []string `json:"dialects"`         // dialects (arrayofstrings, optional): A local or regional variation where the pronunciation occurs, e.g. 'British English' ,
	PhoneticNotation string   `json:"phoneticNotation"` // phoneticNotation (string, optional): The alphabetic system used to display the phonetic spelling ,
	PhoneticSpelling string   `json:"phoneticSpelling"` // phoneticSpelling (string, optional): Phonetic spelling is the representation of vocal sounds which express pronunciations of words. It is a system of spelling in which each letter represents invariably the same spoken sound ,
	Regions          []string `json:"regions"`          // regions (arrayofstrings, optional): A particular area in which the pronunciation occurs, e.g. 'Great Britain'
}

type RelatedEntry struct {
	Domains   []string `json:"domains"`   // domains (arrayofstrings, optional): A subject, discipline, or branch of knowledge particular to the Sense ,
	Id        string   `json:"id"`        // id (string): The identifier of the word ,
	Language  string   `json:"language"`  // language (string, optional): IANA language code specifying the language of the word ,
	Regions   []string `json:"regions"`   // regions (arrayofstrings, optional): A particular area in which the pronunciation occurs, e.g. 'Great Britain' ,
	Registers []string `json:"registers"` // registers (arrayofstrings, optional): A level of language usage, typically with respect to formality. e.g. 'offensive', 'informal' ,
	Text      string   `json:"text"`      // text (string)
}

type Sense struct {
	CrossReferenceMarkers []string            `json:"crossReferenceMarkers"` // crossReferenceMarkers (arrayofstrings, optional): A grouping of crossreference notes. ,
	CrossReferences       CrossReferencesList `json:"crossReferences"`       // crossReferences (CrossReferencesList, optional),
	Definitions           []string            `json:"definitions"`           // definitions (arrayofstrings, optional): A list of statements of the exact meaning of a word ,
	Domains               []string            `json:"domains"`               // domains (arrayofstrings, optional): A subject, discipline, or branch of knowledge particular to the Sense ,
	Examples              ExamplesList        `json:"examples"`              // examples (ExamplesList, optional),
	Id                    string              `json:"id"`                    // id (string, optional): The id of the sense that is required for the delete procedure ,
	Notes                 CategorizedTextList `json:"notes"`                 // notes (CategorizedTextList, optional),
	Pronunciations        PronunciationsList  `json:"pronunciations"`        // pronunciations (PronunciationsList, optional),
	Regions               []string            `json:"regions"`               // regions (arrayofstrings, optional): A particular area in which the Sense occurs, e.g. 'Great Britain' ,
	Registers             []string            `json:"registers"`             // registers (arrayofstrings, optional): A level of language usage, typically with respect to formality. e.g. 'offensive', 'informal' ,
	Subsenses             []Sense             `json:"subsenses"`             // subsenses (Array[Sense], optional): Ordered list of subsenses of a sense ,
	Translations          TranslationsList    `json:"translations"`          // translations (TranslationsList, optional),
	VariantForms          VariantFormsList    `json:"variantForms"`          // variantForms (VariantFormsList, optional): Various words that are used interchangeably depending on the context, e.g 'duck' and 'duck boat'
}

type CategorizedText struct {
	Id   string `json:"id"`   // id (string, optional): The identifier of the word ,
	Text string `json:"text"` // text (string): A note text ,
	Type string `json:"type"` // type (string): The descriptive category of the text
}

type VariantForm struct {
	Regions []string `json:"regions"` // regions (arrayofstrings, optional): A particular area in which the variant form occurs, e.g. 'Great Britain' ,
	Text    string   `json:"text"`    // text (string)
}

type CrossReferencesList []CrossReference

//CrossReferencesList [
//Inline Model 6
//]

type ExamplesList []Example

//ExamplesList [
//Inline Model 7
//]

type TranslationsList []Translation

//TranslationsList [
//Inline Model 8
//]

type CrossReference struct {
	Id   string `json:"id"`   // id (string): The word id of cooccurrence ,
	Text string `json:"text"` // text (string): The word of cooccurrence ,
	Type string `json:"type"` // type (string): The type of relation between the two words. Possible values are 'close match', 'related', 'see also', 'variant spelling', and 'abbreviation' in case of crossreferences, or 'pre', 'post' in case of collocates.
}

type Example struct {
	Definitions  []string            `json:"definitions"`  // definitions (arrayofstrings, optional): A list of statements of the exact meaning of a word ,
	Domains      []string            `json:"domains"`      // domains (arrayofstrings, optional): A subject, discipline, or branch of knowledge particular to the Sense ,
	Notes        CategorizedTextList `json:"notes"`        // notes (CategorizedTextList, optional),
	Regions      []string            `json:"regions"`      // regions (arrayofstrings, optional): A particular area in which the pronunciation occurs, e.g. 'Great Britain' ,
	Registers    []string            `json:"registers"`    // registers (arrayofstrings, optional): A level of language usage, typically with respect to formality. e.g. 'offensive', 'informal' ,
	SenseIds     []string            `json:"senseIds"`     // senseIds (arrayofstrings, optional): The list of sense identifiers related to the example. Provided in the sentences endpoint only. ,
	Text         string              `json:"text"`         // text (string),
	Translations TranslationsList    `json:"translations"` // translations (TranslationsList, optional)
}

type Translation struct {
	Domains             []string                `json:"domains"`             // domains (arrayofstrings, optional): A subject, discipline, or branch of knowledge particular to the translation ,
	GrammaticalFeatures GrammaticalFeaturesList `json:"grammaticalFeatures"` // grammaticalFeatures (GrammaticalFeaturesList, optional),
	Language            string                  `json:"language"`            // language (string): IANA language code specifying the language of the translation ,
	Notes               CategorizedTextList     `json:"notes"`               // notes (CategorizedTextList, optional),
	Regions             []string                `json:"regions"`             // regions (arrayofstrings, optional): A particular area in which the translation occurs, e.g. 'Great Britain' ,
	Registers           []string                `json:"registers"`           // registers (arrayofstrings, optional): A level of language usage, typically with respect to formality. e.g. 'offensive', 'informal' ,
	Text                string                  `json:"text"`                // text (string)
}
