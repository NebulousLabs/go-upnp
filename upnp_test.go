package upnp

import (
	"testing"
)

// TestConcurrentUPNP tests that several threads calling Discover() concurrently
// succeed.
func TestConcurrentUPNP(t *testing.T) {
	// verify that a router exists
	_, err := Discover()
	if err != nil {
		t.Fatal(err)
	}

	// now try to concurrently Discover() using 10 threads
	for i := 0; i < 10; i++ {
		go func() {
			_, err := Discover()
			if err != nil {
				t.Fatal(err)
			}
		}()
	}
}

func TestIGD(t *testing.T) {
	// connect to router
	d, err := Discover()
	if err != nil {
		t.Fatal(err)
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
