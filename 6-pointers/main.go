package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// SensorData struct holds temperature readings from multiple sensors
type SensorData struct {
	id          int
	temperature float64
	timestamp   time.Time
}

// Function to simulate sensor data collection
func collectSensorData(sensorID int) *SensorData {
	return &SensorData{
		id:          sensorID,
		temperature: rand.Float64()*40 + (-10), // Random temperature between -10 to 40°C
		timestamp:   time.Now(),
	}
}

// Function to process sensor data using a pointer to avoid copying
func processSensorData(data *SensorData, wg *sync.WaitGroup) {
	defer wg.Done() // Ensures goroutine exits properly

	// Check for anomalies (e.g., extreme temperature)
	if data.temperature < -5 || data.temperature > 35 {
		fmt.Printf("[ALERT] Sensor %d recorded extreme temperature: %.2f°C at %s\n",
			data.id, data.temperature, data.timestamp.Format("15:04:05"))
	} else {
		fmt.Printf("Sensor %d: Temperature = %.2f°C at %s\n",
			data.id, data.temperature, data.timestamp.Format("15:04:05"))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed random generator

	var wg sync.WaitGroup
	sensorCount := 5 // Simulating 5 sensors

	// Simulating real-time data collection from multiple sensors
	for i := 1; i <= sensorCount; i++ {
		wg.Add(1)
		data := collectSensorData(i)    // Collect sensor data (returns pointer)
		go processSensorData(data, &wg) // Process in parallel
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Sensor data processing completed.")
}
