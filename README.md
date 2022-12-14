## Description
Read the data from the device update queue and process it.

## Setup
1. Run Redis on Docker
2. Set the env `REDIS_HOST` to redis host and port
3. Set the env `API_LOCATION_HOST` to Location API host and port
4. Set the env `API_STATUS_HOST` to Status API host and port
5. Run WireMock on Docker
  ```
  docker run \
    --name wiremock-studio \
    -p 9000:9000 \
    -p 8000-8100:8000-8100 \
    up9inc/wiremock-studio:2.32.0-18
  ```

## Put data into queue
1. Connect to redis-cli
2. Send a left push to `q-device-update` with the body
  ```json
  {
    "device_id": string,
    "ts": int64,
    "coords": {
      "lat": float64,
      "lon": float64
    },
    "battery": int64
  }
  ```

Example:
```
lpush q-device-update '{"device_id":"abc","ts":1234,"coords":{"lat":1234.5678,"lon":-1234.5678},"battery":100}'
```