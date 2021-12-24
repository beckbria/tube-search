package tube_search

// Schema: https://api-portal.tfl.gov.uk/api-details#api=Line&operation=Line_RouteByModeByPathModesQueryServiceTypes

type Tfl3 struct {
	TimeSlice string `json:"timeSlice"`
	Value     int32  `json:"value"`
}

type Tfl4 struct {
	Line                string `json:"line"`
	LineDirection       string `json:"lineDirection"`
	PlatformDestination string `json:"platformDestination"`
	Direction           string `json:"direction"`
	NaptanTo            string `json:"naptanTo"`
	TimeSlice           string `json:"timeSlice"`
	Value               int32  `json:"value"`
}

type Tfl5 struct {
	PassengerFlows []Tfl3 `json:"passengerFlows"`
	TrainLoadings  []Tfl4 `json:"trainLoadings"`
}

type Tfl6 struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Uri       string `json:"uri"`
	FullName  string `json:"fullName"`
	Type      string `json:"type"`
	Crowding  Tfl5   `json:"crowding"`
	RouteType string `json:"routeType"`
	Status    string `json:"status"`
}

type Tfl7 struct {
	NaptanIdReference string   `json:"naptanIdReference"`
	StationAtcoCode   string   `json:"stationAtcoCode"`
	LineIdentifier    []string `json:"lineIdentifier"`
}

type Tfl8 struct {
	ModeName       string   `json:"modeName"`
	LineIdentifier []string `json:"lineIdentifier"`
}

type Tfl9 struct {
	Category        string `json:"category"`
	Key             string `json:"key"`
	SourceSystemKey string `json:"sourceSystemKey"`
	Value           string `json:"value"`
	Modified        string `json:"modified"`
}

type Tfl10 struct {
	Id                   string   `json:"id"`
	Url                  string   `json:"url"`
	CommonName           string   `json:"commonName"`
	Distance             float64  `json:"distance"`
	PlaceType            string   `json:"placeType"`
	AdditionalProperties []Tfl9   `json:"additionalProperties"`
	Children             []Tfl10  `json:"children"`
	ChildrenUrls         []string `json:"childrenUrls"`
	Lat                  float64  `json:"lat"`
	Lon                  float64  `json:"lon"`
}

type Tfl11 struct {
	NaptanId             string   `json:"naptanId"`
	PlatformName         string   `json:"platformName"`
	Indicator            string   `json:"indicator"`
	StopLetter           string   `json:"stopLetter"`
	Modes                []string `json:"modes"`
	IcsCode              string   `json:"icsCode"`
	SmsCode              string   `json:"smsCode"`
	StopType             string   `json:"stopType"`
	StationNaptan        string   `json:"stationNaptan"`
	AccessibilitySumary  string   `json:"accessibilitySumary"`
	HubNaptanCode        string   `json:"hubNaptanCode"`
	Lines                []Tfl6   `json:"lines"`
	LineGroup            []Tfl7   `json:"lineGroup"`
	LineModeGroups       []Tfl8   `json:"lineModeGroups"`
	FullName             string   `json:"fullName"`
	NaptanMode           string   `json:"naptanMode"`
	Status               bool     `json:"status"`
	Id                   string   `json:"id"`
	Url                  string   `json:"url"`
	CommonName           string   `json:"commonName"`
	Distance             float64  `json:"distance"`
	PlaceType            string   `json:"placeType"`
	AdditionalProperties []Tfl9   `json:"additionalProperties"`
	Children             []Tfl10  `json:"children"`
	ChildrenUrls         []string `json:"childrenUrls"`
	Lat                  float64  `json:"lat"`
	Lon                  float64  `json:"lon"`
}

type Tfl12 struct {
	Ordinal   int32   `json:"ordinal"`
	StopPoint []Tfl11 `json:"stopPoint"`
}

type Tfl13 struct {
	Id                              string  `json:"id"`
	LineId                          string  `json:"lineId"`
	RouteCode                       string  `json:"routeCode"`
	Name                            string  `json:"name"`
	LineString                      string  `json:"lineString"`
	Direction                       string  `json:"direction"`
	PriginationName                 string  `json:"originationName"`
	DestinationName                 string  `json:"destinationName"`
	ValidTo                         string  `json:"validTo"`
	ValidFrom                       string  `json:"validFrom"`
	RouteSectionNaptanEntrySequence []Tfl12 `json:"routeSectionNaptanEntrySequence"`
}

type Tfl14 struct {
	Category            string  `json:"category"`
	Type                string  `json:"type"`
	CategoryDescription string  `json:"categoryDescription"`
	Description         string  `json:"description"`
	Summary             string  `json:"summary"`
	AdditionalInfo      string  `json:"additionalInfo"`
	Created             string  `json:"created"`
	LastUpdate          string  `json:"lastUpdate"`
	AffectedRoutes      []Tfl13 `json:"affectedRoutes"`
	AffectedStops       []Tfl11 `json:"affectedStops"`
}

type Tfl15 struct {
	FromDate string `json:"fromDate"`
	ToDate   string `json:"toDate"`
	IsNow    bool   `json:"isNow"`
}

type Tfl16 struct {
	Id                        int32   `json:"id"`
	LineId                    string  `json:"lineId"`
	StatusSeverity            int32   `json:"statusSeverity"`
	StatusSeverityDescription string  `json:"statusSeverityDescription"`
	Reason                    string  `json:"reason"`
	Created                   string  `json:"created"`
	Modified                  string  `json:"modified"`
	ValidityPeriods           []Tfl15 `json:"validityPeriods"`
	Disruption                Tfl14   `json:"disruption"`
}

type Tfl17 struct {
	RouteCode       string `json:"routeCode"`
	Name            string `json:"name"`
	Direction       string `json:"direction"`
	OriginationName string `json:"originationName"`
	DestinationName string `json:"destinationName"`
	Originator      string `json:"originator"`
	Destination     string `json:"destination"`
	ServiceType     string `json:"serviceType"`
	ValidTo         string `json:"validTo"`
	ValidFrom       string `json:"validFrom"`
}

type Tfl18 struct {
	Name string `json:"name"`
	Uri  string `json:"uri"`
}

type Tfl19 struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	ModeName      string  `json:"modeName"`
	Fisruptions   []Tfl14 `json:"disruptions"`
	Created       string  `json:"created"`
	Modified      string  `json:"modified"`
	LineStatuses  []Tfl16 `json:"lineStatuses"`
	RouteSections []Tfl17 `json:"routeSections"`
	ServiceTypes  []Tfl18 `json:"serviceTypes"`
	Crowding      Tfl5    `json:"crowding"`
}
