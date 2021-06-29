package composer

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Manifest of composer.json based on https://getcomposer.org/schema.json
type Manifest struct {
	Name                string                     `json:"name,omitempty"`
	Type                string                     `json:"type,omitempty"`
	TargetDir           string                     `json:"target-dir,omitempty"`
	Description         string                     `json:"description,omitempty"`
	Keywords            []string                   `json:"keywords,omitempty"`
	Homepage            string                     `json:"homepage,omitempty"`
	Readme              string                     `json:"readme,omitempty"`
	Version             string                     `json:"version,omitempty"`
	Time                []string                   `json:"time,omitempty"`
	License             StringOrStrings            `json:"license,omitempty"`
	Authors             []Author                   `json:"authors,omitempty"`
	Require             map[string]string          `json:"require,omitempty"`
	Replace             map[string]string          `json:"replace,omitempty"`
	Conflict            map[string]string          `json:"conflict,omitempty"`
	Provide             map[string]string          `json:"provide,omitempty"`
	RequireDev          map[string]string          `json:"require-dev,omitempty"`
	Suggest             map[string]string          `json:"suggest,omitempty"`
	Config              Config                     `json:"config,omitempty"`
	Autoload            Autoload                   `json:"autoload,omitempty"`
	AutoloadDev         Autoload                   `json:"autoload-dev,omitempty"`
	Archive             map[string]ValueOrMap      `json:"archive,omitempty"`
	Repositories        Repositories               `json:"repositories,omitempty"`
	MinimumStability    string                     `json:"minimum-stability,omitempty"`
	PreferStable        Bool                       `json:"prefer-stable,omitempty"`
	Bin                 StringOrStrings            `json:"bin,omitempty"`
	IncludePath         []string                   `json:"include-path,omitempty"`
	Scripts             map[string]StringOrStrings `json:"scripts,omitempty"`
	ScriptsDescriptions map[string]string          `json:"scripts-descriptions,omitempty"`
	Support             []Support                  `json:"support,omitempty"`
	Funding             Funding                    `json:"funding,omitempty"`
	NonFeatureBranches  []string                   `json:"non-feature-branches,omitempty"`
	DefaultBranch       Bool                       `json:"default-branch,omitempty"`
	Abandoned           BoolOrString               `json:"abandoned,omitempty"`
	Comment             StringOrStrings            `json:"_comment,omitempty"`
	//Extra               TODO     `json:"extra,omitempty"` // Todo implement both map[string]interface{} and []interface{}
}

// StringOrStrings convert "string" or array of "strings" into []string
type StringOrStrings []string

// MarshalJSON convert into an array of strings
func (l StringOrStrings) MarshalJSON() ([]byte, error) {
	return json.Marshal([]string(l))
}

// UnmarshalJSON convert string or array of strings into []string
func (l *StringOrStrings) UnmarshalJSON(bytes []byte) error {
	var str string
	if err := json.Unmarshal(bytes, &str); err == nil {
		*l = []string{str}
		return nil
	}

	var strs []string
	if err := json.Unmarshal(bytes, &strs); err == nil {
		*l = strs
		return nil
	}

	return errors.New(fmt.Sprintf("cannot unmarshal %s", bytes))
}

type Author struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Role     string `json:"role,omitempty"`
}

type Support struct {
	Email    string `json:"email,omitempty"`
	Homepage string `json:"issues,omitempty"`
	Wiki     string `json:"wiki,omitempty"`
	Irc      string `json:"irc,omitempty"`
	Source   string `json:"source,omitempty"`
	Doce     string `json:"docs,omitempty"`
	Rss      string `json:"rss,omitempty"`
	Chat     string `json:"chat,omitempty"`
}

type Config struct {
	ProcessTimeout        int               `json:"process-timeout,omitempty"`
	UseIncludePath        Bool              `json:"use-include-path,omitempty"`
	PreferredInstall      ValueOrMap        `json:"preferred-install,omitempty"`
	NotifyOnInstall       Bool              `json:"notify-on-install,omitempty"`
	GithubProtocols       []string          `json:"github-protocols,omitempty"`
	GithubOauth           map[string]string `json:"github-oauth,omitempty"`
	GitlabOauth           map[string]string `json:"gitlab-oauth,omitempty"`
	GitlabToken           map[string]string `json:"gitlab-token,omitempty"`
	Bearer                map[string]string `json:"bearer,omitempty"`
	DisableTls            Bool              `json:"disable-tls,omitempty"`
	SecureHttp            Bool              `json:"secure-http,omitempty"`
	Cafile                string            `json:"cafile,omitempty"`
	Capath                string            `json:"capath,omitempty"`
	HttpBasic             HttpBasic         `json:"http-basic,omitempty"`
	StoreAuths            BoolOrString      `json:"store-auths,omitempty"`
	Platform              map[string]string `json:"platform,omitempty"`
	VendorDir             string            `json:"vendor-dir,omitempty"`
	BinDir                string            `json:"bin-dir,omitempty"`
	DataDir               string            `json:"data-dir,omitempty"`
	CacheDir              string            `json:"cache-dir,omitempty"`
	CacheFilesDir         string            `json:"cache-files-dir,omitempty"`
	CacheRepoDir          string            `json:"cache-repo-dir,omitempty"`
	CacheVcsDir           string            `json:"cache-vcs-dir,omitempty"`
	CacheTtl              int               `json:"cache-ttl,omitempty"`
	CacheFilesTtl         int               `json:"cache-files-ttl,omitempty"`
	CacheFilesMaxsize     IntString         `json:"cache-files-maxsize,omitempty"`
	CacheReadOnly         Bool              `json:"cache-read-only,omitempty"`
	BinCompat             string            `json:"bin-compat,omitempty"`
	DiscardChanges        BoolOrString      `json:"discard-changes,omitempty"`
	AutoloaderSuffix      string            `json:"autoloader-suffix,omitempty"`
	OptimizeAutoloader    Bool              `json:"optimize-autoloader,omitempty"`
	PrependAutoloader     Bool              `json:"prepend-autoloader,omitempty"`
	ClassmapAuthoritative Bool              `json:"classmap-authoritative,omitempty"`
	ApcuAutoloader        Bool              `json:"apcu-autoloader,omitempty"`
	GithubDomains         []string          `json:"github-domains,omitempty"`
	GithubExposeHostname  Bool              `json:"github-expose-hostname,omitempty"`
	GitlabDomains         []string          `json:"gitlab-domains,omitempty"`
	UseGithubApi          Bool              `json:"use-github-api,omitempty"`
	ArchiveFormat         string            `json:"archive-format,omitempty"`
	ArchiveDir            string            `json:"archive-dir,omitempty"`
	HtaccessProtect       Bool              `json:"htaccess-protect,omitempty"`
	SortPackages          Bool              `json:"sort-packages,omitempty"`
	Lock                  Bool              `json:"lock,omitempty"`
	PlatformCheck         BoolOrString      `json:"platform-check,omitempty"`
}

// Bool convert string, integer or bool variations into a boolean
//
// Examples:
//  { "value": 1 }
//  { "value": "1" }
//  { "value": true }
//  { "value": "true" }
//  { "value": "True" }
//  { "value": True }
// and respectively for all false values
//  { "value": 0 }
//  { "value": "0" }
//  { "value": false }
//  { "value": "false" }
//  { "value": "False" }
//  { "value": False }
// into Go boolean type
type Bool bool

// MarshalJSON json representation of a boolean
func (b Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(b))
}

// UnmarshalJSON converts any string, integer or boolean into a boolean
func (b *Bool) UnmarshalJSON(bytes []byte) error {
	switch string(bytes) {
	case "true", "True", "\"true\"", "\"True\"", "1", "\"1\"":
		*b = true
	case "false", "False", "\"false\"", "\"False\"", "0", "\"0\"":
		*b = false
	default:
		return errors.New("cannot unmarshal bool " + string(bytes))
	}
	return nil
}

// ValueOrMap convert a string or a map of strings into struct
// Example
// "preferred-install": {
//                    "type": ["string", "object"],
//                    "description": "The install method Composer will prefer to use, defaults to auto and can be any of source, dist, auto, or a hash of {\"pattern\": \"preference\"}."
//                }
type ValueOrMap struct {
	Value string
	Map   map[string]string
}

// MarshalJSON convert into either a string or a string map
func (pi ValueOrMap) MarshalJSON() ([]byte, error) {
	if pi.Value != "" {
		return json.Marshal(pi.Value)
	}
	return json.Marshal(pi.Map)
}

// UnmarshalJSON unmarshal string or map of strings into struct
func (pi *ValueOrMap) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &pi.Value); err == nil {
		return nil
	}

	if err := json.Unmarshal(bytes, &pi.Map); err == nil {
		return nil
	}

	return errors.New(fmt.Sprintf("cannot unmarshal %s", bytes))
}

type HttpBasic map[string]struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// BoolOrString convert Bool or String into structure
// Example
// "discard-changes": {
//                    "type": ["string", "boolean"],
//                    "description": "The default style of handling dirty updates, defaults to false and can be any of true, false or \"stash\"."
//                }
type BoolOrString struct {
	Bool   Bool
	String string
}

// MarshalJSON marshal JSON into struct
func (bs BoolOrString) MarshalJSON() ([]byte, error) {
	if bs.String != "" {
		return json.Marshal(bs.String)
	}
	return bs.Bool.MarshalJSON()
}

// UnmarshalJSON convert string or boolean into a custom structure to hold values separately
// If string value is empty, then boolean value is used
func (bs *BoolOrString) UnmarshalJSON(bytes []byte) error {
	if err := bs.Bool.UnmarshalJSON(bytes); err == nil {
		return nil
	}
	if err := json.Unmarshal(bytes, &bs.String); err == nil {
		return nil
	}
	return errors.New(fmt.Sprintf("cannot unmarshal %s", bytes))
}

// IntString convert integer or string into string
// Example
// "cache-files-maxsize": {
//                    "type": ["string", "integer"],
//                    "description": "The cache max size for the files cache, defaults to \"300MiB\"."
//                }
type IntString string

// MarshalJSON marshal JSON into string
func (c IntString) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(c))
}

// UnmarshalJSON convert integer or string into string
// Example JSON values: 300, "300", "300MiB"
func (c *IntString) UnmarshalJSON(bytes []byte) error {
	r := make([]byte, 0)
	for _, b := range bytes {
		switch string(b) {
		case "\"":
		default:
			r = append(r, b)
		}
	}
	*c = (IntString)(r)
	return nil
}

type Autoload struct {
	Psr0                Psr      `json:"psr-0,omitempty"`
	Psr4                Psr      `json:"psr-4,omitempty"`
	Classmap            []string `json:"classmap,omitempty"`
	Files               []string `json:"files,omitempty"`
	ExcludeFromClassmap []string `json:"exclude-from-classmap,omitempty"`
}

// Psr convert a map or a map of arrays into a map of arrays
// Example values
// "psr-0": {
//            "key1": "/string/value/",
//            "key2": [
//                "/rnd/folder1/",
//                "/rnd/folder2/"
//            ]
//        }
type Psr map[string][]string

// MarshalJSON marshal JSON into a map of arrays
func (p Psr) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string][]string(p))
}

// UnmarshalJSON convert a map or a map of arrays into a map of arrays
// Example values
// "psr-0": {
//            "key1": "/string/value/",
//            "key2": [
//                "/rnd/folder1/",
//                "/rnd/folder2/"
//            ]
//        }
func (p *Psr) UnmarshalJSON(bytes []byte) error {
	var i map[string]interface{}
	if err := json.Unmarshal(bytes, &i); err != nil {
		return err
	}

	m := map[string][]string{}
	for k, v := range i {

		var arrStr []string
		switch v.(type) {
		case []interface{}:
			arrValues := v.([]interface{})
			arrStr = make([]string, len(arrValues))
			for a, s := range arrValues {
				arrStr[a] = s.(string)
			}
		case interface{}:
			if arr, ok := v.(string); ok {
				arrStr = []string{arr}
			} else {
				return errors.New("cannot unmarshal psr " + string(bytes))
			}
		}
		m[k] = arrStr
	}
	*p = m
	return nil
}

type Archive struct {
	Name    string   `json:"name,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}

// Repositories convert a map or an array into an array
// Example values
// {
//     "composer": {
//         "type": "composer",
//         "url": "https://composer.example.com/"
//     }
// }
// or
// [{
//     "type": "composer",
//     "url": "https://composer.example.com/"
//  }]
type Repositories []Repository

// MarshalJSON marshal JSON into an array of structs
func (p Repositories) MarshalJSON() ([]byte, error) {
	return json.Marshal([]Repository(p))
}

// UnmarshalJSON convert a map or an array into an array
// Example values
// {
//     "composer": {
//         "type": "composer",
//         "url": "https://composer.example.com/"
//     }
// }
// or
// [{
//     "type": "composer",
//     "url": "https://composer.example.com/"
//  }]
func (p *Repositories) UnmarshalJSON(bytes []byte) error {
	var arr []Repository
	if err := json.Unmarshal(bytes, &arr); err == nil {
		*p = arr
		return nil
	}
	var m map[string]Repository
	if err := json.Unmarshal(bytes, &m); err == nil {
		arr := make([]Repository, 0)
		for _, v := range m {
			arr = append(arr, v)
		}
		*p = arr
		return nil
	}
	return errors.New("cannot unmarshal " + string(bytes))
}

type Repository struct {
	Type                     string                 `json:"type,omitempty"`
	Url                      string                 `json:"url,omitempty"`
	Canonical                Bool                   `json:"canonical,omitempty"`
	Only                     []string               `json:"only,omitempty"`
	Exclude                  []string               `json:"exclude,omitempty"`
	Options                  map[string]interface{} `json:"options,omitempty"`
	AllowSslDowngrade        Bool                   `json:"allow_ssl_downgrade,omitempty"`
	ForceLazyProviders       Bool                   `json:"force-lazy-providers,omitempty"`
	NoApi                    Bool                   `json:"no-api,omitempty"`
	SecureHttp               Bool                   `json:"secure-http,omitempty"`
	SvnCacheCredentials      Bool                   `json:"svn-cache-credentials,omitempty"`
	TrunkPath                *BoolOrString          `json:"trunk-path,omitempty"`
	BranchesPath             *BoolOrString          `json:"branches-path,omitempty"`
	TagsPath                 *BoolOrString          `json:"tags-path,omitempty"`
	PackagePath              string                 `json:"package-path,omitempty"`
	Depot                    string                 `json:"depot,omitempty"`
	Branch                   string                 `json:"branch,omitempty"`
	UniquePerforceClientName string                 `json:"unique_perforce_client_name,omitempty"`
	P4user                   string                 `json:"p4user,omitempty"`
	P4password               string                 `json:"p4password,omitempty"`
	VendorAlias              string                 `json:"vendor-alias,omitempty"`
	//Package   TODO  `json:"package"` //Todo implement package structure
}

type Funding struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}
