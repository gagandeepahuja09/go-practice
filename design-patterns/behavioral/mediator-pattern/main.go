package main

func main() {
	stationManager := newStationManager()
	passengerTrain := &passengerTrain{
		mediator: stationManager,
	}
	goodsTrain := &goodsTrain{
		mediator: stationManager,
	}
	intercityTrain := &intercityTrain{
		mediator: stationManager,
	}
	passengerTrain.requestArrival()
	goodsTrain.requestArrival()
	intercityTrain.requestArrival()
	passengerTrain.departure()
	goodsTrain.departure()
}
