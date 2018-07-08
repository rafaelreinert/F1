package main

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"strconv"
)

// TelemetryPack is a struct with telemetry informations
type TelemetryPack struct {
	Time                 float32
	LapTime              float32
	LapDistance          float32
	TotalDistance        float32
	X                    float32
	Y                    float32
	Z                    float32
	Speed                float32
	Xv                   float32
	Xy                   float32
	Xz                   float32
	Xr                   float32
	Yr                   float32
	Zr                   float32
	Xd                   float32
	Yd                   float32
	Zd                   float32
	SuspPos              [4]float32
	SuspVel              [4]float32
	WheelSpeed           [4]float32
	Throttle             float32
	Steer                float32
	Brake                float32
	Clutch               float32
	Gear                 float32
	GforceIat            float32
	GforceIon            float32
	Lap                  float32
	EngineRate           float32
	SliProNativeSuport   float32
	CarPosition          float32
	KersLevel            float32
	KersMaxLevel         float32
	Drs                  float32
	TractionControl      float32
	AntiLockBrakes       float32
	FuelInTank           float32
	FuelCapacity         float32
	InPits               float32
	Sector               float32
	Sector1Time          float32
	Sector2Time          float32
	BrakesTemp           [4]float32
	TyresPressure        [4]float32
	TeamInfo             float32
	TotalLaps            float32
	TrackSize            float32
	LastLapTime          float32
	MaxRPM               float32
	IdleRPM              float32
	MaxGears             float32
	SessionType          float32
	DrsAllowed           float32
	TrackNumber          float32
	VehicleFIAFlags      float32
	Era                  float32
	EngineTemperature    float32
	GforceVert           float32
	AngVelX              float32
	AngVelY              float32
	AngVelZ              float32
	TyresTemperature     [4]int8
	TyresWear            [4]int8
	TyreCompound         int8
	FrontBrakesBias      int8
	FuelMix              int8
	CurrentLapInvalid    int8
	TyresDamage          [4]int8
	FrontLeftWingDamage  int8
	FrontRightWingDamage int8
	RearWingDamage       int8
	EngineDamage         int8
	GearBoxDamage        int8
	ExhaustDamage        int8
	PitLimiterStatus     int8
	PitSpeedLimit        int8
	SessionTimeLeft      float32
	RevLightsPercent     int8
	IsSpectating         int8
	SpectatorCarIndex    int8
	NumCars              int8
	PlayerCarIndex       int8
	CarData              [20]carDataPack
	Yam                  float32
	Pitch                float32
	Roll                 float32
	XLocalVelocity       float32
	YLocalVelocity       float32
	ZLocalVelocity       float32
	SuspAcceleration     [4]float32
	AngAccX              float32
	AngAccY              float32
	AngAccZ              float32
}

type carDataPack struct {
	WordPosition      [3]float32
	LastLapTime       float32
	CurrentLapTime    float32
	BestLapTime       float32
	Sector1LapTime    float32
	Sector2LapTime    float32
	LapDistance       float32
	DriverID          int8
	TeamID            int8
	CarPosition       int8
	CurrentLapNum     int8
	TyreCompound      int8
	InPits            int8
	Sector            int8
	CurrentLapInvalid int8
	Penalties         int8
}

// ToMap convert this TelemetryPack to a Map
func (t *TelemetryPack) ToMap() map[string]interface{} {
	var m map[string]interface{}
	m = make(map[string]interface{})

	s := reflect.ValueOf(t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		switch f.Kind() {
		case reflect.Int8:
			m[typeOfT.Field(i).Name] = int(f.Interface().(int8))
		case reflect.Float32:
			m[typeOfT.Field(i).Name] = float64(f.Interface().(float32))
		case reflect.Array:
			for i2 := 0; i2 < f.Len(); i2++ {
				e := f.Index(i2)
				switch e.Kind() {
				case reflect.Int8:
					m[typeOfT.Field(i).Name+"_"+strconv.Itoa(i2)] = int(e.Interface().(int8))
				case reflect.Float32:
					m[typeOfT.Field(i).Name+"_"+strconv.Itoa(i2)] = float64(e.Interface().(float32))
				}
			}
		}
	}
	return m
}

// NewTelemetryPack return a new TelemetryPack from bytes
func NewTelemetryPack(buf []byte) (*TelemetryPack, error) {
	var pack TelemetryPack
	r := bytes.NewReader(buf)
	if err := binary.Read(r, binary.LittleEndian, &pack); err != nil {
		return nil, err
	}
	return &pack, nil
}
