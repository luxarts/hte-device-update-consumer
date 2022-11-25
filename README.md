# go-rest-template
REST API server template with ping implemented.

# Setup
1. Run Redis on Docker
2. Set the env `REDIS_HOST` to redis host and port
3. Set the env `API_LOCATION_HOST` to Location API host and port
4. Set the env `API_STATUS_HOST` to Status API host and port

## Put data into queue
1. Connect to redis-cli
2. Send a left push to `q-device-update` with the body
  ```json
  {
    "device_id": string,
    "ts": int64,
    "coords": {
      "lat": int64,
      "lon": int64
    },
    "battery": int64
  }
  ```

Example:
```
lpush q-device-update '{"device_id":"abc","ts":1234,"coords":{"lat":1234.5678,"lon":-1234.5678},"battery":100}'
```