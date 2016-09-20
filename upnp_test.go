package upnp

import (
	"sync"
	"testing"
)

// TestConcurrentUPNP tests that several threads calling Discover() concurrently
// succeed.
func TestConcurrentUPNP(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	// verify that a router exists
	_, err := Discover()
	if err != nil {
		t.Skip(err)
	}

	// now try to concurrently Discover() using 20 threads
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			_, err := Discover()
			if err != nil {
				t.Fatal(err)
			}
		}()
	}
	wg.Wait()
}

func TestIGD(t *testing.T) {
	// connect to router
	d, err := Discover()
	if err != nil {
		t.Skip(err)
	}

	// discover external IP
	ip, err := d.ExternalIP()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Your external IP is:", ip)

	// forward a port
	err = d.Forward(9001, "upnp test")
	if err != nil {
		t.Fatal(err)
	}

	// un-forward a port
	err = d.Clear(9001)
	if err != nil {
		t.Fatal(err)
	}

	// record router's location
	loc := d.Location()
	if err != nil {
		t.Fatal(err)
	}

	// connect to router directly
	d, err = Load(loc)
	if err != nil {
		t.Fatal(err)
	}
}
