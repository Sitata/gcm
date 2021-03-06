package ccs

import (
	"os"
	"testing"
)

// GCM environment variables
var host = os.Getenv("GCM_CCS_HOST")
var senderID = os.Getenv("GCM_SENDER_ID")
var apiKey = os.Getenv("GOOGLE_API_KEY")

// optional registration ID from an Android device, used for testing outgoing messages
var regID = os.Getenv("GCM_REG_ID")

func TestConnect(t *testing.T) {
	c := getConn(t)
	c.Close()
}

func TestClosedConnection(t *testing.T) {
	// test guard against nil/closed connections
}

func TestConnectError(t *testing.T) {
}

func TestSend(t *testing.T) {
	if regID == "" {
		t.Skip("Skipping integration test due to missing GCM registration ID (GCM_REG_ID) environment variable")
	}

	c := getConn(t)

	outmsg := OutMsg{To: regID, Data: map[string]string{"test_message": "GCM CCS client testing message"}}
	t.Logf("Testing out message: %+v to device with registration ID: %+v", outmsg, regID)
	send(t, c, &outmsg)

	inmsg := receive(t, c)
	if inmsg != nil {
		t.Fatalf("Received a message for some reason even though delivery request was not requested. Received message: %+v", inmsg)
	}

	c.Close()
}

func TestSendError(t *testing.T) {
	if regID == "" {
		t.Skip("Skipping integration test due to missing GCM registration ID (GCM_REG_ID) environment variable")
	}

	// // JSON error
	// c := getConn(t)
	//
	// outmsg := OutMsg{}
	// t.Logf("Testing out message: %+v to device with registration ID: %+v", outmsg, regID)
	// send(t, c, &outmsg)
	//
	// inmsg := receive(t, c)
	//
	// // stanza error
	//
	// c.Close()

}

func TestAck(t *testing.T) {
	// messages should be removed from the queue once they are ACKed (or delivery receipt arrives if requested)
}

func TestNack(t *testing.T) {
}

func TestReceipt(t *testing.T) {
}

func TestMessageId(t *testing.T) {
	id, err := getMsgID()
	if err != nil {
		t.Fatal(err)
	}
	if len(id) != 26 {
		t.Fatalf("Failed to generate unique message ID of lenght 26 chars")
	}
}

// Test to see if we can handle all known GCM message types properly.
func TestGCMMessages(t *testing.T) {
	c := getConn(t)
	c.Close()
}

// Test to see if our message structure's fields match the incoming message fields exactly.
func TestMessageFields(t *testing.T) {
	c := getConn(t)
	c.Close()
}

func getConn(t *testing.T) *Conn {
	if testing.Short() {
		t.Skip("Skipping integration test in short testing mode")
	} else if host == "" || senderID == "" || apiKey == "" {
		t.Skip("Skipping integration test due to missing GCM environment variables")
	}

	c, err := Connect(host, senderID, apiKey, true)
	if err != nil {
		t.Fatalf("CCS error while connecting to server: %v", err)
	}
	return c
}

func receive(t *testing.T, c *Conn) *InMsg {
	m, err := c.Receive()
	if err != nil {
		t.Fatalf("CCS error while receiving message: %v", err)
	}
	return m
}

func send(t *testing.T, c *Conn, m *OutMsg) (n int) {
	n, err := c.Send(m)
	if err != nil {
		t.Fatalf("CCS error while sending message: %v", err)
	}
	if n == 0 {
		t.Fatal("CCS error while sending message: 0 bytes were written to the underlying socket connection")
	}
	return
}
