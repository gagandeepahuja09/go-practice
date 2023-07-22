**Entities**
* Cab (Driver) same entity as one-on-one mapping.
    * Location struct { x, y }
    * isAvailable bool
    * UpdateLocation
    * setAvailability

* Rider
    * Location
    * []Trip struct{cabId, bookingTime, endLocation}
    * Book(cab, endLocation) 
        * location (distance validation)
        * pick the first found rider
        * append to trip and update status to inprogress.
    * Status ==> inProgress, ended
    * EndTrip()
        * update trip end location
