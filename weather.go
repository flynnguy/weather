package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Weather struct {
	Response struct {
		Version        string `json:"version"`
		TermsofService string `json:"termsofService"`
		Features       struct {
			Alerts           int `json:"alerts"`
			Almanac          int `json:"almanac"`
			Astronomy        int `json:"astronomy"`
			Conditions       int `json:"conditions"`
			Currenthurricane int `json:"currenthurricane"`
			Forecast         int `json:"forecast"`
			Forecast10day    int `json:"forecast10day"`
			Hourly           int `json:"hourly"`
			Hourly10day      int `json:"hourly10day"`
			Tide             int `json:"tide"`
			Yesterday        int `json:"yesterday"`
		} `json:"features"`
	} `json:"response"`
	CurrentObservation struct {
		DisplayLocation struct {
			Full           string  `json:"full"`
			City           string  `json:"city"`
			State          string  `json:"state"`
			StateName      string  `json:"state_name"`
			Country        string  `json:"country"`
			CountryIso3166 string  `json:"country_iso3166"`
			Zip            string  `json:"zip"`
			Magic          string  `json:"magic"`
			Wmo            string  `json:"wmo"`
			Latitude       float32 `json:"latitude,string"`
			Longitude      float32 `json:"longitude,string"`
			Elevation      string  `json:"elevation"`
		} `json:"display_location"`
		ObservationLocation struct {
			Full           string  `json:"full"`
			City           string  `json:"city"`
			State          string  `json:"state"`
			Country        string  `json:"country"`
			CountryIso3166 string  `json:"country_iso3166"`
			Latitude       float32 `json:"latitude,string"`
			Longitude      float32 `json:"longitude,string"`
			Elevation      string  `json:"elevation"`
		} `json:"observation_location"`
		Estimated struct {
		} `json:"estimated"`
		StationId             string  `json:"station_id"`
		ObservationTime       string  `json:"observation_time"`
		ObservationTimeRfc822 string  `json:"observation_time_rfc822"`
		ObservationEpoch      string  `json:"observation_epoch"`
		LocalTimeRfc822       string  `json:"local_time_rfc822"`
		LocalEpoch            string  `json:"local_epoch"`
		LocalTzShort          string  `json:"local_tz_short"`
		LocalTzLong           string  `json:"local_tz_long"`
		LocalTzOffset         string  `json:"local_tz_offset"`
		Weather               string  `json:"weather"`
		TemperatureString     string  `json:"temperature_string"`
		TempF                 float32 `json:"temp_f"`
		TempC                 float32 `json:"temp_c"`
		RelativeHumidity      string  `json:"relative_humidity"`
		WindString            string  `json:"wind_string"`
		WindDir               string  `json:"wind_dir"`
		WindDegrees           int     `json:"wind_degrees"`
		WindMph               float32 `json:"wind_mph"`
		WindGustMph           float32 `json:"wind_gust_mph"`
		WindKph               float32 `json:"wind_kph"`
		WindGustKph           float32 `json:"wind_gust_kph"`
		PressureMb            string  `json:"pressure_mb"`
		PressureIn            string  `json:"pressure_in"`
		PressureTrend         string  `json:"pressure_trend"`
		DewpointString        string  `json:"dewpoint_string"`
		DewpointF             int     `json:"dewpoint_f"`
		DewpointC             int     `json:"dewpoint_c"`
		HeatIndexString       string  `json:"heat_index_string"`
		HeatIndexF            string  `json:"heat_index_f"`
		HeatIndexC            string  `json:"heat_index_c"`
		WindchillString       string  `json:"windchill_string"`
		WindchillF            string  `json:"windchill_f"`
		WindchillC            string  `json:"windchill_c"`
		FeelslikeString       string  `json:"feelslike_string"`
		FeelslikeF            string  `json:"feelslike_f"`
		FeelslikeC            string  `json:"feelslike_c"`
		VisibilityMi          string  `json:"visibility_mi"`
		VisibilityKm          string  `json:"visibility_km"`
		Solarradiation        string  `json:"solarradiation"`
		UV                    string  `json:"UV"`
		Precip1hrString       string  `json:"precip_1hr_string"`
		Precip1hrIn           string  `json:"precip_1hr_in"`
		Precip1hrMetric       string  `json:"precip_1hr_metric"`
		PrecipTodayString     string  `json:"precip_today_string"`
		PrecipTodayIn         string  `json:"precip_today_in"`
		PrecipTodayMetric     string  `json:"precip_today_metric"`
		Icon                  string  `json:"icon"`
		IconUrl               string  `json:"icon_url"`
		ForecastUrl           string  `json:"forecast_url"`
		HistoryUrl            string  `json:"history_url"`
		ObUrl                 string  `json:"ob_url"`
		Nowcast               string  `json:"nowcast"`
	} `json:"current_observation"`
	Forecast struct {
		TxtForecast struct {
			Date        string `json:"date"`
			Forecastday []struct {
				Period        int    `json:"period"`
				Icon          string `json:"icon"`
				IconUrl       string `json:"icon_url"`
				Title         string `json:"title"`
				Fcttext       string `json:"fcttext"`
				FcttextMetric string `json:"fcttext_metric"`
				Pop           string `json:"pop"`
			} `json:"forecastday"`
		} `json:"txt_forecast"`
		Simpleforecast struct {
			Forecastday []struct {
				Date struct {
					Epoch          string `json:"epoch"`
					Pretty         string `json:"pretty"`
					Day            int    `json:"day"`
					Month          int    `json:"month"`
					Year           int    `json:"year"`
					Yday           int    `json:"yday"`
					Hour           int    `json:"hour"`
					Min            string `json:"min"`
					Sec            int    `json:"sec"`
					Isdst          string `json:"isdst"`
					Monthname      string `json:"monthname"`
					MonthnameShort string `json:"monthname_short"`
					WeekdayShort   string `json:"weekday_short"`
					Weekday        string `json:"weekday"`
					Ampm           string `json:"ampm"`
					TzShort        string `json:"tz_short"`
					TzLong         string `json:"tz_long"`
				} `json:"date"`
				Period int `json:"period"`
				High   struct {
					Fahrenheit int `json:"fahrenheit,string"`
					Celsius    int `json:"celsius,string"`
				} `json:"high"`
				Low struct {
					Fahrenheit int `json:"fahrenheit,string"`
					Celsius    int `json:"celsius,string"`
				} `json:"low"`
				Conditions string `json:"conditions"`
				Icon       string `json:"icon"`
				IconUrl    string `json:"icon_url"`
				Skyicon    string `json:"skyicon"`
				Pop        int    `json:"pop"`
				QpfAllday  struct {
					In float32 `json:"in"`
					Mm int     `json:"mm"`
				} `json:"qpf_allday"`
				QpfDay struct {
					In float32 `json:"in"`
					Mm int     `json:"mm"`
				} `json:"qpf_day"`
				QpfNight struct {
					In float32 `json:"in"`
					Mm int     `json:"mm"`
				} `json:"qpf_night"`
				SnowAllday struct {
					In float32 `json:"in"`
					Cm int     `json:"cm"`
				} `json:"snow_allday"`
				SnowDay struct {
					In float32 `json:"in"`
					Cm int     `json:"cm"`
				} `json:"snow_day"`
				SnowNight struct {
					In float32 `json:"in"`
					Cm int     `json:"cm"`
				} `json:"snow_night"`
				Maxwind struct {
					Mph     int    `json:"mph"`
					Kph     int    `json:"kph"`
					Dir     string `json:"dir"`
					Degrees int    `json:"degrees"`
				} `json:"maxwind"`
				Avewind struct {
					Mph     int    `json:"mph"`
					Kph     int    `json:"kph"`
					Dir     string `json:"dir"`
					Degrees int    `json:"degrees"`
				} `json:"avewind"`
				Avehumidity int `json:"avehumidity"`
				Maxhumidity int `json:"maxhumidity"`
				Minhumidity int `json:"minhumidity"`
			} `json:"forecastday"`
		} `json:"simpleforecast"`
	} `json:"forecast"`
	HourlyForecast []struct {
		FCTTIME struct {
			Hour                   string `json:"hour"`
			HourPadded             string `json:"hour_padded"`
			Min                    string `json:"min"`
			MinUnpadded            string `json:"min_unpadded"`
			Sec                    string `json:"sec"`
			Year                   string `json:"year"`
			Mon                    string `json:"mon"`
			MonPadded              string `json:"mon_padded"`
			MonAbbrev              string `json:"mon_abbrev"`
			Mday                   string `json:"mday"`
			MdayPadded             string `json:"mday_padded"`
			Yday                   string `json:"yday"`
			Isdst                  string `json:"isdst"`
			Epoch                  string `json:"epoch"`
			Pretty                 string `json:"pretty"`
			Civil                  string `json:"civil"`
			MonthName              string `json:"month_name"`
			MonthNameAbbrev        string `json:"month_name_abbrev"`
			WeekdayName            string `json:"weekday_name"`
			WeekdayNameNight       string `json:"weekday_name_night"`
			WeekdayNameAbbrev      string `json:"weekday_name_abbrev"`
			WeekdayNameUnlang      string `json:"weekday_name_unlang"`
			WeekdayNameNightUnlang string `json:"weekday_name_night_unlang"`
			Ampm                   string `json:"ampm"`
			Tz                     string `json:"tz"`
			Age                    string `json:"age"`
			UTCDATE                string `json:"UTCDATE"`
		} `json:"FCTTIME"`
		Temp struct {
			English string `json:"english"`
			Metric  string `json:"metric"`
		} `json:"temp"`
		Dewpoint struct {
			English string `json:"english"`
			Metric  string `json:"metric"`
		} `json:"dewpoint"`
		Condition string `json:"condition"`
		Icon      string `json:"icon"`
		IconUrl   string `json:"icon_url"`
		Fctcode   string `json:"fctcode"`
		Sky       string `json:"sky"`
		Wspd      struct {
			English string `json:"english"`
			Metric  string `json:"metric"`
		} `json:"wspd"`
		Wdir struct {
			Dir     string `json:"dir"`
			Degrees string `json:"degrees"`
		} `json:"wdir"`
		Wx        string `json:"wx"`
		Uvi       string `json:"uvi"`
		Humidity  string `json:"humidity"`
		Windchill struct {
			English string `json:"english"`
			Metric  string `json:"metric"`
		} `json:"windchill"`
		Heatindex struct {
			English string `json:"english"`
			Metric  string `json:"metric"`
		} `json:"heatindex"`
		Feelslike struct {
			English string `json:"english"`
			Metric  string `json:"metric"`
		} `json:"feelslike"`
		Qpf struct {
			English string `json:"english"`
			Metric  string `json:"metric"`
		} `json:"qpf"`
		Snow struct {
			English string `json:"english"`
			Metric  string `json:"metric"`
		} `json:"snow"`
		Pop  string `json:"pop"`
		Mslp struct {
			English string `json:"english"`
			Metric  string `json:"metric"`
		} `json:"mslp"`
	} `json:"hourly_forecast"`
	MoonPhase struct {
		PercentIlluminated string `json:"percentIlluminated"`
		AgeOfMoon          string `json:"ageOfMoon"`
		PhaseofMoon        string `json:"phaseofMoon"`
		Hemisphere         string `json:"hemisphere"`
		CurrentTime        struct {
			Hour   string `json:"hour"`
			Minute string `json:"minute"`
		} `json:"current_time"`
		Sunrise struct {
			Hour   string `json:"hour"`
			Minute string `json:"minute"`
		} `json:"sunrise"`
		Sunset struct {
			Hour   string `json:"hour"`
			Minute string `json:"minute"`
		} `json:"sunset"`
	} `json:"moon_phase"`
	SunPhase struct {
		Sunrise struct {
			Hour   string `json:"hour"`
			Minute string `json:"minute"`
		} `json:"sunrise"`
		Sunset struct {
			Hour   string `json:"hour"`
			Minute string `json:"minute"`
		} `json:"sunset"`
	} `json:"sun_phase"`
	QueryZone string        `json:"query_zone"`
	Alerts    []interface{} `json:"alerts"`
	History   struct {
		Date struct {
			Pretty string `json:"pretty"`
			Year   string `json:"year"`
			Mon    string `json:"mon"`
			Mday   string `json:"mday"`
			Hour   string `json:"hour"`
			Min    string `json:"min"`
			Tzname string `json:"tzname"`
		} `json:"date"`
		Utcdate struct {
			Pretty string `json:"pretty"`
			Year   string `json:"year"`
			Mon    string `json:"mon"`
			Mday   string `json:"mday"`
			Hour   string `json:"hour"`
			Min    string `json:"min"`
			Tzname string `json:"tzname"`
		} `json:"utcdate"`
		Observations []struct {
			Date struct {
				Pretty string `json:"pretty"`
				Year   string `json:"year"`
				Mon    string `json:"mon"`
				Mday   string `json:"mday"`
				Hour   string `json:"hour"`
				Min    string `json:"min"`
				Tzname string `json:"tzname"`
			} `json:"date"`
			Utcdate struct {
				Pretty string `json:"pretty"`
				Year   string `json:"year"`
				Mon    string `json:"mon"`
				Mday   string `json:"mday"`
				Hour   string `json:"hour"`
				Min    string `json:"min"`
				Tzname string `json:"tzname"`
			} `json:"utcdate"`
			Tempm      string `json:"tempm"`
			Tempi      string `json:"tempi"`
			Dewptm     string `json:"dewptm"`
			Dewpti     string `json:"dewpti"`
			Hum        string `json:"hum"`
			Wspdm      string `json:"wspdm"`
			Wspdi      string `json:"wspdi"`
			Wgustm     string `json:"wgustm"`
			Wgusti     string `json:"wgusti"`
			Wdird      string `json:"wdird"`
			Wdire      string `json:"wdire"`
			Vism       string `json:"vism"`
			Visi       string `json:"visi"`
			Pressurem  string `json:"pressurem"`
			Pressurei  string `json:"pressurei"`
			Windchillm string `json:"windchillm"`
			Windchilli string `json:"windchilli"`
			Heatindexm string `json:"heatindexm"`
			Heatindexi string `json:"heatindexi"`
			Precipm    string `json:"precipm"`
			Precipi    string `json:"precipi"`
			Conds      string `json:"conds"`
			Icon       string `json:"icon"`
			Fog        string `json:"fog"`
			Rain       string `json:"rain"`
			Snow       string `json:"snow"`
			Hail       string `json:"hail"`
			Thunder    string `json:"thunder"`
			Tornado    string `json:"tornado"`
			Metar      string `json:"metar"`
		} `json:"observations"`
		Dailysummary []struct {
			Date struct {
				Pretty string `json:"pretty"`
				Year   string `json:"year"`
				Mon    string `json:"mon"`
				Mday   string `json:"mday"`
				Hour   string `json:"hour"`
				Min    string `json:"min"`
				Tzname string `json:"tzname"`
			} `json:"date"`
			Fog                                string `json:"fog"`
			Rain                               string `json:"rain"`
			Snow                               string `json:"snow"`
			Snowfallm                          string `json:"snowfallm"`
			Snowfalli                          string `json:"snowfalli"`
			Monthtodatesnowfallm               string `json:"monthtodatesnowfallm"`
			Monthtodatesnowfalli               string `json:"monthtodatesnowfalli"`
			Since1julsnowfallm                 string `json:"since1julsnowfallm"`
			Since1julsnowfalli                 string `json:"since1julsnowfalli"`
			Snowdepthm                         string `json:"snowdepthm"`
			Snowdepthi                         string `json:"snowdepthi"`
			Hail                               string `json:"hail"`
			Thunder                            string `json:"thunder"`
			Tornado                            string `json:"tornado"`
			Meantempm                          string `json:"meantempm"`
			Meantempi                          string `json:"meantempi"`
			Meandewptm                         string `json:"meandewptm"`
			Meandewpti                         string `json:"meandewpti"`
			Meanpressurem                      string `json:"meanpressurem"`
			Meanpressurei                      string `json:"meanpressurei"`
			Meanwindspdm                       string `json:"meanwindspdm"`
			Meanwindspdi                       string `json:"meanwindspdi"`
			Meanwdire                          string `json:"meanwdire"`
			Meanwdird                          string `json:"meanwdird"`
			Meanvism                           string `json:"meanvism"`
			Meanvisi                           string `json:"meanvisi"`
			Humidity                           string `json:"humidity"`
			Maxtempm                           string `json:"maxtempm"`
			Maxtempi                           string `json:"maxtempi"`
			Mintempm                           string `json:"mintempm"`
			Mintempi                           string `json:"mintempi"`
			Maxhumidity                        string `json:"maxhumidity"`
			Minhumidity                        string `json:"minhumidity"`
			Maxdewptm                          string `json:"maxdewptm"`
			Maxdewpti                          string `json:"maxdewpti"`
			Mindewptm                          string `json:"mindewptm"`
			Mindewpti                          string `json:"mindewpti"`
			Maxpressurem                       string `json:"maxpressurem"`
			Maxpressurei                       string `json:"maxpressurei"`
			Minpressurem                       string `json:"minpressurem"`
			Minpressurei                       string `json:"minpressurei"`
			Maxwspdm                           string `json:"maxwspdm"`
			Maxwspdi                           string `json:"maxwspdi"`
			Minwspdm                           string `json:"minwspdm"`
			Minwspdi                           string `json:"minwspdi"`
			Maxvism                            string `json:"maxvism"`
			Maxvisi                            string `json:"maxvisi"`
			Minvism                            string `json:"minvism"`
			Minvisi                            string `json:"minvisi"`
			Gdegreedays                        string `json:"gdegreedays"`
			Heatingdegreedays                  string `json:"heatingdegreedays"`
			Coolingdegreedays                  string `json:"coolingdegreedays"`
			Precipm                            string `json:"precipm"`
			Precipi                            string `json:"precipi"`
			Precipsource                       string `json:"precipsource"`
			Heatingdegreedaysnormal            string `json:"heatingdegreedaysnormal"`
			Monthtodateheatingdegreedays       string `json:"monthtodateheatingdegreedays"`
			Monthtodateheatingdegreedaysnormal string `json:"monthtodateheatingdegreedaysnormal"`
			Since1sepheatingdegreedays         string `json:"since1sepheatingdegreedays"`
			Since1sepheatingdegreedaysnormal   string `json:"since1sepheatingdegreedaysnormal"`
			Since1julheatingdegreedays         string `json:"since1julheatingdegreedays"`
			Since1julheatingdegreedaysnormal   string `json:"since1julheatingdegreedaysnormal"`
			Coolingdegreedaysnormal            string `json:"coolingdegreedaysnormal"`
			Monthtodatecoolingdegreedays       string `json:"monthtodatecoolingdegreedays"`
			Monthtodatecoolingdegreedaysnormal string `json:"monthtodatecoolingdegreedaysnormal"`
			Since1sepcoolingdegreedays         string `json:"since1sepcoolingdegreedays"`
			Since1sepcoolingdegreedaysnormal   string `json:"since1sepcoolingdegreedaysnormal"`
			Since1jancoolingdegreedays         string `json:"since1jancoolingdegreedays"`
			Since1jancoolingdegreedaysnormal   string `json:"since1jancoolingdegreedaysnormal"`
		} `json:"dailysummary"`
	} `json:"history"`
	Almanac struct {
		AirportCode string `json:"airport_code"`
		TempHigh    struct {
			Normal struct {
				F string `json:"F"`
				C string `json:"C"`
			} `json:"normal"`
			Record struct {
				F string `json:"F"`
				C string `json:"C"`
			} `json:"record"`
			Recordyear string `json:"recordyear"`
		} `json:"temp_high"`
		TempLow struct {
			Normal struct {
				F string `json:"F"`
				C string `json:"C"`
			} `json:"normal"`
			Record struct {
				F string `json:"F"`
				C string `json:"C"`
			} `json:"record"`
			Recordyear string `json:"recordyear"`
		} `json:"temp_low"`
	} `json:"almanac"`
	Tide struct {
		TideInfo []struct {
			TideSite string `json:"tideSite"`
			Lat      string `json:"lat"`
			Lon      string `json:"lon"`
			Units    string `json:"units"`
			Type     string `json:"type"`
			Tzname   string `json:"tzname"`
		} `json:"tideInfo"`
		TideSummary []struct {
			Date struct {
				Pretty string `json:"pretty"`
				Year   string `json:"year"`
				Mon    string `json:"mon"`
				Mday   string `json:"mday"`
				Hour   string `json:"hour"`
				Min    string `json:"min"`
				Tzname string `json:"tzname"`
				Epoch  string `json:"epoch"`
			} `json:"date"`
			Utcdate struct {
				Pretty string `json:"pretty"`
				Year   string `json:"year"`
				Mon    string `json:"mon"`
				Mday   string `json:"mday"`
				Hour   string `json:"hour"`
				Min    string `json:"min"`
				Tzname string `json:"tzname"`
				Epoch  string `json:"epoch"`
			} `json:"utcdate"`
			Data struct {
				Height string `json:"height"`
				Type   string `json:"type"`
			} `json:"data"`
		} `json:"tideSummary"`
		TideSummaryStats []struct {
			Maxheight float32 `json:"maxheight"`
			Minheight float32 `json:"minheight"`
		} `json:"tideSummaryStats"`
	} `json:"tide"`
	Currenthurricane []struct {
		StormInfo struct {
			StormName     string `json:"stormName"`
			StormNameNice string `json:"stormName_Nice"`
			StormNumber   string `json:"stormNumber"`
			Requesturl    string `json:"requesturl"`
			Wuiurl        string `json:"wuiurl"`
		} `json:"stormInfo"`
		Current struct {
			Lat                   float32 `json:"lat"`
			Lon                   float32 `json:"lon"`
			SaffirSimpsonCategory int     `json:"SaffirSimpsonCategory"`
			Category              string  `json:"Category"`
			Time                  struct {
				Hour                   string `json:"hour"`
				HourPadded             string `json:"hour_padded"`
				Year                   string `json:"year"`
				Mon                    string `json:"mon"`
				MonPadded              string `json:"mon_padded"`
				MonAbbrev              string `json:"mon_abbrev"`
				Mday                   string `json:"mday"`
				MdayPadded             string `json:"mday_padded"`
				Yday                   string `json:"yday"`
				Epoch                  string `json:"epoch"`
				Pretty                 string `json:"pretty"`
				Civil                  string `json:"civil"`
				MonthName              string `json:"month_name"`
				MonthNameAbbrev        string `json:"month_name_abbrev"`
				WeekdayName            string `json:"weekday_name"`
				WeekdayNameNight       string `json:"weekday_name_night"`
				WeekdayNameAbbrev      string `json:"weekday_name_abbrev"`
				WeekdayNameUnlang      string `json:"weekday_name_unlang"`
				WeekdayNameNightUnlang string `json:"weekday_name_night_unlang"`
				Ampm                   string `json:"ampm"`
			} `json:"Time"`
			TimeGMT struct {
				Hour                   string `json:"hour"`
				HourPadded             string `json:"hour_padded"`
				Year                   string `json:"year"`
				Mon                    string `json:"mon"`
				MonPadded              string `json:"mon_padded"`
				MonAbbrev              string `json:"mon_abbrev"`
				Mday                   string `json:"mday"`
				MdayPadded             string `json:"mday_padded"`
				Yday                   string `json:"yday"`
				Epoch                  string `json:"epoch"`
				Pretty                 string `json:"pretty"`
				Civil                  string `json:"civil"`
				MonthName              string `json:"month_name"`
				MonthNameAbbrev        string `json:"month_name_abbrev"`
				WeekdayName            string `json:"weekday_name"`
				WeekdayNameNight       string `json:"weekday_name_night"`
				WeekdayNameAbbrev      string `json:"weekday_name_abbrev"`
				WeekdayNameUnlang      string `json:"weekday_name_unlang"`
				WeekdayNameNightUnlang string `json:"weekday_name_night_unlang"`
				Ampm                   string `json:"ampm"`
			} `json:"TimeGMT"`
			WindSpeed struct {
				Kts int `json:"Kts"`
				Mph int `json:"Mph"`
				Kph int `json:"Kph"`
			} `json:"WindSpeed"`
			WindGust struct {
				Kts int `json:"Kts"`
				Mph int `json:"Mph"`
				Kph int `json:"Kph"`
			} `json:"WindGust"`
			Fspeed struct {
				Kts int `json:"Kts"`
				Mph int `json:"Mph"`
				Kph int `json:"Kph"`
			} `json:"Fspeed"`
			Movement struct {
				Degrees string `json:"Degrees"`
				Text    string `json:"Text"`
			} `json:"Movement"`
			Pressure struct {
				Mb     interface{} `json:"mb"`
				Inches interface{} `json:"inches"`
			} `json:"Pressure"`
			WindQuadrants struct {
				Comment string `json:"comment"`
				Quad1   string `json:"quad_1"`
				Quad2   string `json:"quad_2"`
				Quad3   string `json:"quad_3"`
				Quad4   string `json:"quad_4"`
			} `json:"WindQuadrants"`
			WindRadius struct {
				Number34 struct {
					NE int `json:"NE"`
					SE int `json:"SE"`
					SW int `json:"SW"`
					NW int `json:"NW"`
				} `json:"34"`
				Number50 struct {
					NE int `json:"NE"`
					SE int `json:"SE"`
					SW int `json:"SW"`
					NW int `json:"NW"`
				} `json:"50"`
				Number64 struct {
					NE int `json:"NE"`
					SE int `json:"SE"`
					SW int `json:"SW"`
					NW int `json:"NW"`
				} `json:"64"`
			} `json:"WindRadius"`
			SeaQuadrants struct {
				Comment string `json:"comment"`
			} `json:"SeaQuadrants"`
			SeaRadius struct {
				Number12 struct {
				} `json:"12"`
			} `json:"SeaRadius"`
		} `json:"Current"`
		Forecast []struct {
			ForecastHour          string  `json:"ForecastHour"`
			SaffirSimpsonCategory int     `json:"SaffirSimpsonCategory"`
			Lat                   float32 `json:"lat"`
			Lon                   int     `json:"lon"`
			Category              string  `json:"Category"`
			Time                  struct {
				Hour                   string `json:"hour"`
				HourPadded             string `json:"hour_padded"`
				Year                   string `json:"year"`
				Mon                    string `json:"mon"`
				MonPadded              string `json:"mon_padded"`
				MonAbbrev              string `json:"mon_abbrev"`
				Mday                   string `json:"mday"`
				MdayPadded             string `json:"mday_padded"`
				Yday                   string `json:"yday"`
				Epoch                  string `json:"epoch"`
				Pretty                 string `json:"pretty"`
				Civil                  string `json:"civil"`
				MonthName              string `json:"month_name"`
				MonthNameAbbrev        string `json:"month_name_abbrev"`
				WeekdayName            string `json:"weekday_name"`
				WeekdayNameNight       string `json:"weekday_name_night"`
				WeekdayNameAbbrev      string `json:"weekday_name_abbrev"`
				WeekdayNameUnlang      string `json:"weekday_name_unlang"`
				WeekdayNameNightUnlang string `json:"weekday_name_night_unlang"`
				Ampm                   string `json:"ampm"`
			} `json:"Time"`
			TimeGMT struct {
				Year                   string `json:"year"`
				Mon                    string `json:"mon"`
				MonPadded              string `json:"mon_padded"`
				MonAbbrev              string `json:"mon_abbrev"`
				Mday                   string `json:"mday"`
				MdayPadded             string `json:"mday_padded"`
				Yday                   string `json:"yday"`
				Epoch                  string `json:"epoch"`
				Pretty                 string `json:"pretty"`
				Civil                  string `json:"civil"`
				MonthName              string `json:"month_name"`
				MonthNameAbbrev        string `json:"month_name_abbrev"`
				WeekdayName            string `json:"weekday_name"`
				WeekdayNameNight       string `json:"weekday_name_night"`
				WeekdayNameAbbrev      string `json:"weekday_name_abbrev"`
				WeekdayNameUnlang      string `json:"weekday_name_unlang"`
				WeekdayNameNightUnlang string `json:"weekday_name_night_unlang"`
				Ampm                   string `json:"ampm"`
			} `json:"TimeGMT"`
			WindSpeed struct {
				Kts int `json:"Kts"`
				Mph int `json:"Mph"`
				Kph int `json:"Kph"`
			} `json:"WindSpeed"`
			WindGust struct {
				Kts int `json:"Kts"`
				Mph int `json:"Mph"`
				Kph int `json:"Kph"`
			} `json:"WindGust"`
			WindQuadrants struct {
				Comment string `json:"comment"`
				Quad1   string `json:"quad_1"`
				Quad2   string `json:"quad_2"`
				Quad3   string `json:"quad_3"`
				Quad4   string `json:"quad_4"`
			} `json:"WindQuadrants"`
			WindRadius struct {
				Number34 struct {
					NE int `json:"NE"`
					SE int `json:"SE"`
					SW int `json:"SW"`
					NW int `json:"NW"`
				} `json:"34"`
				Number50 struct {
					NE int `json:"NE"`
					SE int `json:"SE"`
					SW int `json:"SW"`
					NW int `json:"NW"`
				} `json:"50"`
				Number64 struct {
					NE int `json:"NE"`
					SE int `json:"SE"`
					SW int `json:"SW"`
					NW int `json:"NW"`
				} `json:"64"`
			} `json:"WindRadius"`
			ErrorRadius string `json:"ErrorRadius"`
		} `json:"forecast"`
		ExtendedForecast []interface{} `json:"ExtendedForecast"`
		Track            []struct {
			SaffirSimpsonCategory int     `json:"SaffirSimpsonCategory"`
			Category              string  `json:"Category"`
			Lat                   float32 `json:"lat"`
			Lon                   int     `json:"lon"`
			Time                  struct {
				Hour                   string `json:"hour"`
				HourPadded             string `json:"hour_padded"`
				Year                   string `json:"year"`
				Mon                    string `json:"mon"`
				MonPadded              string `json:"mon_padded"`
				MonAbbrev              string `json:"mon_abbrev"`
				Mday                   string `json:"mday"`
				MdayPadded             string `json:"mday_padded"`
				Yday                   string `json:"yday"`
				Epoch                  string `json:"epoch"`
				Pretty                 string `json:"pretty"`
				Civil                  string `json:"civil"`
				MonthName              string `json:"month_name"`
				MonthNameAbbrev        string `json:"month_name_abbrev"`
				WeekdayName            string `json:"weekday_name"`
				WeekdayNameNight       string `json:"weekday_name_night"`
				WeekdayNameAbbrev      string `json:"weekday_name_abbrev"`
				WeekdayNameUnlang      string `json:"weekday_name_unlang"`
				WeekdayNameNightUnlang string `json:"weekday_name_night_unlang"`
				Ampm                   string `json:"ampm"`
			} `json:"Time"`
			TimeGMT struct {
				Year                   string `json:"year"`
				Mon                    string `json:"mon"`
				MonPadded              string `json:"mon_padded"`
				MonAbbrev              string `json:"mon_abbrev"`
				Mday                   string `json:"mday"`
				MdayPadded             string `json:"mday_padded"`
				Yday                   string `json:"yday"`
				Epoch                  string `json:"epoch"`
				Pretty                 string `json:"pretty"`
				Civil                  string `json:"civil"`
				MonthName              string `json:"month_name"`
				MonthNameAbbrev        string `json:"month_name_abbrev"`
				WeekdayName            string `json:"weekday_name"`
				WeekdayNameNight       string `json:"weekday_name_night"`
				WeekdayNameAbbrev      string `json:"weekday_name_abbrev"`
				WeekdayNameUnlang      string `json:"weekday_name_unlang"`
				WeekdayNameNightUnlang string `json:"weekday_name_night_unlang"`
				Ampm                   string `json:"ampm"`
			} `json:"TimeGMT"`
			WindSpeed struct {
				Kts int `json:"Kts"`
				Mph int `json:"Mph"`
				Kph int `json:"Kph"`
			} `json:"WindSpeed"`
			Pressure struct {
				Mb     interface{} `json:"mb"`
				Inches interface{} `json:"inches"`
			} `json:"Pressure"`
			WindQuadrants struct {
				Comment string `json:"comment"`
				Quad1   string `json:"quad_1"`
				Quad2   string `json:"quad_2"`
				Quad3   string `json:"quad_3"`
				Quad4   string `json:"quad_4"`
			} `json:"WindQuadrants"`
			WindRadius struct {
				Number34 struct {
					NE interface{} `json:"NE"`
					SE interface{} `json:"SE"`
					SW interface{} `json:"SW"`
					NW interface{} `json:"NW"`
				} `json:"34"`
				Number50 struct {
					NE interface{} `json:"NE"`
					SE interface{} `json:"SE"`
					SW interface{} `json:"SW"`
					NW interface{} `json:"NW"`
				} `json:"50"`
				Number64 struct {
					NE interface{} `json:"NE"`
					SE interface{} `json:"SE"`
					SW interface{} `json:"SW"`
					NW interface{} `json:"NW"`
				} `json:"64"`
			} `json:"WindRadius"`
			SeaQuadrants struct {
				Comment string `json:"comment"`
			} `json:"SeaQuadrants"`
			SeaRadius struct {
				Number12 struct {
				} `json:"12"`
			} `json:"SeaRadius"`
		} `json:"track"`
	} `json:"currenthurricane"`
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// This function goes out and saves the json data to a file so we don't have to make so many json calls while testing
func readUrlToFile(url string) (err error) {
	// Examples of json endpoints
	// fmt.Sprintf("http://api.wunderground.com/api/%s/geolookup/q/autoip.json", API_KEY) // auto ip
	// fmt.Sprintf("http://api.wunderground.com/api/%s/geolookup/q/07005.json" , API_KEY) // Lookup zip
	// fmt.Sprintf("http://api.wunderground.com/api/%s/geolookup/q/EWR.json" , API_KEY)   // Lookup airport code
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Download new data? [y/n] ")
	text, _ := reader.ReadString('\n')
	switch text[:1] {
	case "Y", "y":
		fmt.Printf("%s\n", url)
		resp, err := http.Get(url)
		check(err)

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		check(err)

		var ModePerm os.FileMode = 0777
		err = ioutil.WriteFile("weather_data.json", body, ModePerm)
		check(err)
	}
	return
}

func main() {
	api_key_byte, err := ioutil.ReadFile("API_KEY.txt")
	check(err)
	API_KEY := strings.Trim(string(api_key_byte), "\n")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Zip Code: ")
	zip_code, _ := reader.ReadString('\n')
	zip_code = strings.Trim(zip_code, "\n")

	url := fmt.Sprintf("http://api.wunderground.com/api/%s/alerts/almanac/astronomy/conditions/currenthurricane/forecast/forecast10day/hourly/hourly10day/tide/yesterday/q/geolookup/%s.json", API_KEY, zip_code)
	err = readUrlToFile(url)
	check(err)

	jsonStr, err := ioutil.ReadFile("weather_data.json")
	check(err)

	var w Weather
	err = json.Unmarshal(jsonStr, &w) // There still seems to be some errors marshalling the data, but so far nothing we're using
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The weather for %s:\n", w.CurrentObservation.DisplayLocation.Full)
	fmt.Printf("%s; %.2f F (%.2f C)\n", w.CurrentObservation.Weather, w.CurrentObservation.TempF, w.CurrentObservation.TempC)
}
