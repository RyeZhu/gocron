package gocron

import (
	"fmt"
	"testing"
	"time"
)

var err = 1

func task() {
	fmt.Println("I am a running job.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b, time.Now().Format(time.RFC3339Nano))
}

func taskMilliseconds(strChan chan<- string) {
	strChan <- time.Now().Format(time.RFC3339Nano)
}

func TestJob_Millisecond(t *testing.T) {
	strChan := make(chan string)
	go func() {
		l := len("2018-01-12T07:52:44.172")
		for s := range strChan {
			fmt.Println(s[l-6:l])
		}
	}()
	//defaultScheduler.Every(10).Milliseconds().Do(task)
	//defaultScheduler.Every(2).Milliseconds().Do(taskWithParams, 2, "hello")
	defaultScheduler.Every(10).Milliseconds().Do(taskMilliseconds, strChan)
	defaultScheduler.Start()
	time.Sleep(1 * time.Second)
}

func TestMilliSecond(*testing.T) {
	strChan := make(chan string)
	go func() {
		l := len("2018-01-12T07:52:44.172")
		for s := range strChan {
			fmt.Println(s[l-6:l])
		}
	}()
	//defaultScheduler.Every(1).Millisecond().Do(task)
	defaultScheduler.Every(1).Millisecond().Do(taskMilliseconds, strChan)
	defaultScheduler.Start()
	time.Sleep(2 * time.Second)
}

func TestSecond(*testing.T) {
	defaultScheduler.Every(1).Second().Do(task)
	defaultScheduler.Every(1).Second().Do(taskWithParams, 1, "hello")
	defaultScheduler.Start()
	time.Sleep(10 * time.Second)
}

func Test_formatTime(t *testing.T) {
	tests := []struct {
		name     string
		args     string
		wantHour int
		wantMin  int
		wantErr  bool
	}{
		{
			name:     "normal",
			args:     "16:18",
			wantHour: 16,
			wantMin:  18,
			wantErr:  false,
		},
		{
			name:     "normal",
			args:     "6:18",
			wantHour: 6,
			wantMin:  18,
			wantErr:  false,
		},
		{
			name:     "notnumber",
			args:     "e:18",
			wantHour: 0,
			wantMin:  0,
			wantErr:  true,
		},
		{
			name:     "outofrange",
			args:     "25:18",
			wantHour: 25,
			wantMin:  18,
			wantErr:  true,
		},
		{
			name:     "wrongformat",
			args:     "19:18:17",
			wantHour: 0,
			wantMin:  0,
			wantErr:  true,
		},
		{
			name:     "wrongminute",
			args:     "19:1e",
			wantHour: 19,
			wantMin:  0,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHour, gotMin, err := formatTime(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("formatTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHour != tt.wantHour {
				t.Errorf("formatTime() gotHour = %v, want %v", gotHour, tt.wantHour)
			}
			if gotMin != tt.wantMin {
				t.Errorf("formatTime() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
		})
	}
}
