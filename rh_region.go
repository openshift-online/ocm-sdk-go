package sdk

import (
	json2 "encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Region struct {
	URL string
	AWS []string
	GCP []string
}

func GetRhRegions(ocmServiceUrl string) (map[string]Region, error) {
	var regions map[string]Region
	url, err := DetermineRegionDiscoveryUrl(ocmServiceUrl)
	if err != nil {
		return regions, fmt.Errorf("Can't determine region discovery URL: %s\n", err)
	}
	// Adding nolint here in order to prevent linter from failing due to variable http get
	resp, err := http.Get(url) //nolint
	if err != nil {
		return regions, fmt.Errorf("Can't retrieve shards: %s", err)
	}
	err = json2.NewDecoder(resp.Body).Decode(&regions)
	if err != nil {
		return regions, fmt.Errorf("Can't decode shards: %s", err)
	}
	return regions, nil
}

func GetRhRegion(ocmServiceUrl string, regionName string) (Region, error) {
	regions, err := GetRhRegions(ocmServiceUrl)
	if err != nil {
		return Region{}, err
	}
	regionName = fmt.Sprintf("rh-%s", regionName)
	regVal, ok := regions[regionName]
	if !ok {
		return Region{}, fmt.Errorf("Can't find region %s", regionName)
	}
	return regVal, nil
}

func DetermineRegionDiscoveryUrl(ocmServiceUrl string) (string, error) {
	baseUrl, err := url.Parse(ocmServiceUrl)
	if err != nil {
		return "", err
	}
	regionDiscoveryHost := "api.openshift.com"
	//TODO: Remove the OR condition from this if statement before the MR merge
	if strings.HasSuffix(baseUrl.Hostname(), "integration.openshift.com") || true {
		regionDiscoveryHost = "api.integration.openshift.com"
	} else if strings.HasSuffix(baseUrl.Hostname(), "stage.openshift.com") {
		regionDiscoveryHost = "api.stage.openshift.com"
	}
	return fmt.Sprintf("https://%s/static/ocm-shards.json", regionDiscoveryHost), nil
}
