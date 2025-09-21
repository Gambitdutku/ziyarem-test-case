package sensors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"ziyaremtestcase/domain"
)

type TempSensor struct {
	Endpoint string
}

func (t *TempSensor) Type() string {
	return "temperature"
}

func (t *TempSensor) Read(id string) (*domain.SensorData, error) {
	url := fmt.Sprintf("%s/%s", t.Endpoint, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("network error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("invalid response from sensor API: %d", resp.StatusCode)
	}

	var payload struct {
		Value float64 `json:"value"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, fmt.Errorf("invalid sensor data: %w", err)
	}

	return &domain.SensorData{
		ID:        id,
		DeviceID:  id,
		Value:     payload.Value,
		Timestamp: time.Now(),
	}, nil
}
