package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tonky/shipping/PortDomainService/domain"
)

type mockPortService struct {
	ports map[string]domain.Port
}

func (ms mockPortService) Get(name string) (*domain.Port, error) {
	p := ms.ports[name]

	return &p, nil
}

func (ms mockPortService) Send(ps chan domain.Port) error {
	for p := range ps {
		ms.ports[p.Key] = p
	}

	return nil
}

func TestLoadFileAndSendToMockService(t *testing.T) {
	pjs := `{"AEAJM":{
  "name": "Ajman",
  "city": "Ajman",
  "country": "United Arab Emirates",
  "alias": [],
  "regions": [],
  "coordinates": [
    55.5136433,
    25.4052165
  ],
  "province": "Ajman",
  "timezone": "Asia/Dubai",
  "unlocs": [
    "AEAJM"
  ],
  "code": "52000"
},
 "UAIEV": {
    "name": "Kiev",
    "city": "Kiev",
    "country": "Ukraine",
    "alias": [],
    "regions": [],
    "coordinates": [
      30.5234,
      50.4501
    ],
    "province": "Kyiv city",
    "timezone": "Europe/Kiev",
    "unlocs": [
      "UAIEV"
    ]
  }}`

	ports := map[string]domain.Port{
		"MockPort": domain.Port{Name: "MockPortOld"},
		"UAIEV":    domain.Port{Name: "Kiev", Code: "1"},
	}

	mps := mockPortService{ports}

	pc, _ := streamDecodePorts(strings.NewReader(pjs))

	mps.Send(pc)

	if mps.ports["AEAJM"].Code != "52000" {
		t.Fatal("Code wasn't added", mps.ports)
	}

	if ports["UAIEV"].Code != "" {
		t.Fatal("Code wasn't updated")
	}

	if ports["UAIEV"].Latitude != 30.5234 {
		t.Fatal("Missing latitude")
	}
}

func TestGetPort(t *testing.T) {
	ps := map[string]domain.Port{
		"MockPort": domain.Port{Name: "MockPort"},
		"UAIEV":    domain.Port{Name: "Kiev"},
	}

	ah := api{mockPortService{ps}}

	ts := httptest.NewServer(ah)

	defer ts.Close()

	res, err := http.Get(ts.URL + "?name=UAIEV")

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Fatal("Failed to read all body: ", err)
	}

	res.Body.Close()

	var port domain.Port

	if err := json.Unmarshal(body, &port); err != nil {
		t.Fatalf("Can't decode '%s' as json: %s", string(body), err)
	}

	if port.Name != "Kiev" {
		t.Fatalf("Expected fixture, got: %s", port.String())
	}
}
