package main

// var hostname string
// var port string

// // DBTimeout is the maximum response time from DB
// const DBTimeout = 5000

// func setMongoParameters() {
// 	if os.Getenv("MONGO_HOSTNAME") != "" {
// 		hostname = os.Getenv("MONGO_HOSTNAME")
// 	} else {
// 		customWarn("USING LOCAL DATABASE")
// 		hostname = "localhost"
// 	}
// 	if os.Getenv("MONGO_PORT") != "" {
// 		port = os.Getenv("MONGO_PORT")
// 	} else {
// 		customWarn("USING DEFAULT DATABASE")
// 		port = "27017"
// 	}
// 	customLog("DB: {name: mongo, hostname:" + hostname + ", port:" + port + "}")
// }

// func getClient(c *gin.Context) (*mongo.Client, error) {

// 	client, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://"+hostname+":27017"))
// 	if err != nil {
// 		return nil, errors.New("Failed to generate Mongo Client")
// 	}

// 	customLog("pinging database with this FQDN: " + hostname)

// 	// Short timeout to test mongo connection
// 	shortCtx, cancelFunc := context.WithTimeout(c, DBTimeout*time.Millisecond)
// 	defer cancelFunc()
// 	err = client.Ping(shortCtx, readpref.Primary())
// 	if err != nil {
// 		return nil, errors.New("Unable to reach database within " + strconv.Itoa(DBTimeout) + "ms")
// 	}
// 	customLog("Acces granted !")
// 	return client, nil
// }

// func getDatabase(c *mongo.Client) *mongo.Database {
// 	name := "smartsearch"
// 	database := c.Database(name)
// 	return database
// }
