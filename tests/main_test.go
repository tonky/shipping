package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	// replace github.com/tonky/shipping/ClientAPI => ../ClientAPI
	"github.com/tonky/shipping/ClientAPI/portService"
	"github.com/tonky/shipping/PortDomainService/domain"
)

func TestAdd(t *testing.T) {
	dsAddr := os.Getenv("DOMAIN_SERVICE_ADDR")

	if dsAddr == "" {
		panic("Must set 'DOMAIN_SERVICE_ADDR' to domainService address:port")
	}

	ps, err := portService.New(dsAddr)

	if err != nil {
		t.Error("Service dial error: ", err)
	}

	pre := domain.Port{Key: "ABCDE", Name: "Testing", City: "Tululah_old"}

	pc1 := make(chan domain.Port, 1)
	pc1 <- pre
	close(pc1)

	ps.Send(pc1)

	want := domain.Port{Key: "ABCDE", Name: "Oracular", City: "Tululah"}

	pc := make(chan domain.Port, 1)
	pc <- want
	close(pc)

	ps.Send(pc)

	p, err := getByKey("ABCDE")

	if err != nil {
		t.Error("Get error: ", err)
	}

	if p.Key != want.Key || p.Name != want.Name || p.City != want.City {
		t.Errorf("Want: %v, got: %v", want, p)
	}
}

// 2019/12/29 13:53:03 COPY IN successful, added 10000 ports in 137ms, begin upserting data into `ports`
// 2019/12/29 13:53:03 Upserted 10000 batch in 432ms
// 2019/12/29 13:53:03 Upserted 100000 ports in 6591ms
/*
func TestUpsert100KPorts(t *testing.T) {
	dsAddr := os.Getenv("DOMAIN_SERVICE_ADDR")

	if dsAddr == "" {
		panic("Must set 'DOMAIN_SERVICE_ADDR' to domainService address:port")
	}

	ps, err := portService.New(dsAddr)

	if err != nil {
		t.Error("Service dial error: ", err)
	}

	pc := make(chan domain.Port)

	go func() {
		for i := 0; i < 100000; i++ {
			pc <- domain.Port{Key: fmt.Sprintf("TEST_KEY_%d", i), Name: fmt.Sprintf("TEST_NAME_%d", i)}
		}

		close(pc)
	}()

	ps.Send(pc)

	want := domain.Port{Key: "TEST_KEY_99999", Name: "TEST_NAME_99999"}

	p, err := getByKey(want.Key)

	if err != nil {
		t.Error("Error getting port by key: ", err)
	}

	if p.Key != want.Key || p.Name != want.Name {
		t.Errorf("want: %v, Got: %v", want, p)
	}
}
*/
func getByKey(n string) (*domain.Port, error) {
	apiAddr := os.Getenv("API_ADDR")

	if apiAddr == "" {
		panic("Must set 'API_ADDR' to clientAPI address:port")
	}

	resp, err := http.Get("http://" + apiAddr + "/?name=" + n)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var p domain.Port

	if err := json.Unmarshal(body, &p); err != nil {
		return nil, err
	}

	return &p, nil
}
