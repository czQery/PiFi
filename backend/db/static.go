package db

const dbCreate = `
	CREATE TABLE IF NOT EXISTS ap (
		bssid TEXT NOT NULL PRIMARY KEY,
		ssid TEXT NOT NULL,
		mode TEXT NOT NULL,
		discovered TEXT NOT NULL,
		channel INTEGER NOT NULL,
		frequency INTEGER NOT NULL,
		rssi REAL NOT NULL,
		latitude REAL NOT NULL,
		longitude REAL NOT NULL,
		altitude INTEGER NOT NULL,
		accuracy INTEGER NOT NULL,
		device TEXT NOT NULL,
		wigle BOOLEAN DEFAULT FALSE,
		beacondb BOOLEAN DEFAULT FALSE,
		dwpa boolean DEFAULT FALSE
	);
`
