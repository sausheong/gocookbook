package errors

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"testing"
	"time"
)

func TestUnmarshaling(t *testing.T) {
	person := unmarshal()
	if person.Name != "Luke Skywalker" {
		t.Error("Wrong name")
	}
}

func TestAnotherUnmarshaling(t *testing.T) {
	person := anotherUnmarshal()
	if person.Name != "Luke Skywalker" {
		t.Error("Wrong name")
	}
}

func TestHelperUnmarshaling(t *testing.T) {
	person := helperUnmarshal()
	if person.Name != "Luke Skywalker" {
		t.Error("Wrong name")
	}
}

func TestUnmarshaller(t *testing.T) {
	u := &unmarshaller{}
	u.get("https://swapi.dev/api/people/1")
	u.read()
	u.unmarshal()
	if u.err != nil {
		t.Error("Cannot unmarshal:", u.err)
	}
	if u.person.Name != "Luke Skywalker" {
		t.Error("Wrong name")
	}

}

func TestMust(t *testing.T) {
	ret := must(strconv.ParseInt("10", 10, 64)).(int64)
	if ret != 10 {
		t.Error("helper not working")
	}
}

func TestMustUnmarshal(t *testing.T) {
	person := mustUnmarshal()
	if person.Name != "Luke Skywalker" {
		t.Error("Wrong name")
	}
}

func TestConversion(t *testing.T) {
	var err error
	convert := func(num string) (i int64) {
		if err != nil {
			return
		}
		i, err = strconv.ParseInt(num, 10, 64)
		return
	}

	i1 := convert("10")
	i2 := convert("aa")
	i3 := convert("30")
	if err != nil {
		t.Error(err)
	}
	t.Log(i1, i2, i3)
}

func TestCommsError(t *testing.T) {
	err := send([]byte("data"))
	if err != nil {
		t.Error("Error is:", err)
	}
}

func TestSyntaxError(t *testing.T) {
	err := run()
	if err != nil {
		err, ok := err.(*SyntaxError)
		if ok {
			t.Log("Error is:", err)
		}
	}
}

func TestWrappedError(t *testing.T) {
	err1 := errors.New("Oops something happened.")
	err2 := fmt.Errorf("An error was encountered - %w", err1)
	e := errors.Unwrap(err2)
	if !errors.Is(e, err1) {
		t.Error("Errors are not the same", e, "is not the same as ", err1)
	}
}

func TestWrappedStructError(t *testing.T) {
	err := connect()
	e := errors.Unwrap(err)
	if !errors.Is(e, err) {
		t.Error("Errors are not the same", e, "is not the same as ", err)
	} else {
		t.Log(e)
	}

}

func TestErrorsAs(t *testing.T) {
	err := connect()

	var connErr *ConnectionError
	if !errors.As(err, &connErr) {
		t.Error("Error is not a Connection Error")
	} else {
		t.Log("Host:", connErr.Host, "Port:", connErr.Port)
	}
}

func TestErrorsIs(t *testing.T) {
	err := connectAPI() // connectAPI returns ApiErr
	if !errors.Is(err, ApiErr) {
		t.Error("The errors are NOT the same")
	}
}

func TestErrorsIsWrap(t *testing.T) {
	err := connect() // connect returns an error that wraps ApiErr
	if !errors.Is(err, ApiErr) {
		t.Error("The errors are NOT the same")
	}
}

func TestInterrupt(t *testing.T) {
	var exited = false
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)

	go func(*bool) {
		<-ch
		// clean up before graceful exit
		t.Log("Cleaning up and exiting gracefully")
		exited = true
	}(&exited)

	time.Sleep(1 * time.Second)
	ch <- syscall.SIGINT
	time.Sleep(1 * time.Second)
	if !exited {
		t.Error("Did not handle interrupt")
	}

}

func TestPanic(t *testing.T) {
	A()
}
