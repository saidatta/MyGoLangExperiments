package main

//Both InfluxDB and Cassandra are distributed database management systems that can be used for storing and querying time series data. However, they have some differences that make them better suited for different use cases:
//
//InfluxDB:
//
//InfluxDB is specifically designed for storing and querying time series data. It has a native time series data model and supports efficient querying of data over a time range.
//InfluxDB is optimized for high write and read performance, making it a good choice for applications that need to handle a high volume of read and write operations.
//InfluxDB has built-in support for handling data with a high cardinality (i.e., a large number of unique values) and for performing complex aggregations on the data.
//Cassandra:
//
//Cassandra is a general-purpose database that can be used for storing and querying a wide variety of data types, including time series data.
//Cassandra is optimized for write performance and can handle a high volume of writes efficiently. However, it may not be as efficient for read operations, especially if you need to retrieve a large number of rows or perform complex aggregations on the data.
//Cassandra stores data in a denormalized, column-oriented format, which can be efficient for storing and querying data with a large number of columns. However, this may not be as efficient if you only have a few columns of data and need to retrieve a large number of rows.
//In general, InfluxDB is a better choice for applications that need to handle a high volume of time series data and need to perform complex queries on the data, such as aggregation and filtering. Cassandra, on the other hand, is a good choice for applications that need to handle a large volume of writes and are less concerned with the efficiency of read operations.
//
//Ultimately, the choice between InfluxDB and Cassandra for a Twitch chat or Facebook news feed application will depend on the specific requirements of the application, including the volume and type of data being stored, the performance and availability needs, and the complexity of the queries being performed.

//func main() {
//	// Connect to Cassandra cluster
//	cluster := gocql.NewCluster("127.0.0.1")
//	cluster.Keyspace = "twitch_chat"
//	session, err := cluster.CreateSession()
//	if err != nil {
//		panic(err)
//	}
//	defer session.Close()
//
//	// Create chat message table
//	if err := session.Query(`
//		CREATE TABLE IF NOT EXISTS chat_messages (
//			channel_id text,
//			timestamp timestamp,
//			username text,
//			message text,
//			PRIMARY KEY (channel_id, timestamp)
//		)
//	`).Exec(); err != nil {
//		panic(err)
//	}
//
//	// Send a chat message
//	channelID := "some_channel"
//	username := "some_user"
//	message := "Hello, world!"
//	if err := session.Query(`
//		INSERT INTO chat_messages (channel_id, timestamp, username, message)
//		VALUES (?, ?, ?, ?)
//	`, channelID, time.Now(), username, message).Exec(); err != nil {
//		panic(err)
//	}
//
//	// Retrieve the latest chat messages
//	var retrievedChannelID, retrievedUsername, retrievedMessage string
//	var retrievedTimestamp time.Time
//	iter := session.Query(`
//		SELECT channel_id, timestamp, username, message
//		FROM chat_messages
//		WHERE channel_id = ?
//		ORDER BY timestamp DESC
//		LIMIT 10
//	`, channelID).Iter()
//	for iter.Scan(&retrievedChannelID, &retrievedTimestamp, &retrievedUsername, &retrievedMessage) {
//		fmt.Printf("[%s] %s: %s\n", retrievedTimestamp, retrievedUsername, retrievedMessage)
//	}
//	if err := iter.Close(); err != nil {
//		panic(err)
//	}
//}
