#!/bin/bash

# Configuration
CLIENT_SEED=100000
APP_TOKEN=d28721be-fd2d-4b45-869e-9f253b554e50
PROMO_ID=43e35910-c168-4634-ad4f-52fd764a843f

# Functions
make_request() {
  local url=$1
  local method=$2
  local data=$3
  local auth=$4

  if [ -z "$auth" ]; then
    curl -s -X "$method" "$url" -H "Content-Type: application/json" -d "$data"
  else
    curl -s -X "$method" "$url" -H "Authorization: Bearer $auth" -H "Content-Type: application/json" -d "$data"
  fi
}
generate_uuid() {
  cat /proc/sys/kernel/random/uuid
}
random_int() {
  local min=$1
  local max=$2
  local seed=$3
  awk -v seed="$seed" -v min="$min" -v max="$max" 'BEGIN {
    srand(seed);
    print int(min + rand() * (max - min + 1))
  }'
}

# Step 0. Generate clientId
RANDOM_INT1=$(random_int 1721926800000 1722272400000 $SEED)
RANDOM_INT2=$(random_int 100000000 999999999 $SEED)
RANDOM_INT3=$(random_int 100000000 9999999999 $SEED)
CLIENT_ID="${RANDOM_INT1}-${RANDOM_INT2}${RANDOM_INT3}"
echo "Using clientId: $CLIENT_ID"

# Step 1: Login and get clientToken
echo "Logging in to obtain clientToken..."
login_response=$(make_request "https://api.gamepromo.io/promo/login-client" "POST" "{\"appToken\": \"$APP_TOKEN\", \"clientId\": \"$CLIENT_ID\", \"clientOrigin\": \"deviceid\"}")
client_token=$(echo $login_response | jq -r '.clientToken')

# Step 2: Check if clientToken was received
if [ -z "$client_token" ] || [ "$client_token" == "null" ]; then
  echo "Failed to obtain clientToken"
  exit 1
fi
echo "Successfully obtained clientToken: $client_token"

# Step 3: Register event in a loop
echo "Starting to register events..."
while true; do
  event_id=$(generate_uuid)
  echo "Registering event with eventId: $event_id"
  register_event_response=$(make_request "https://api.gamepromo.io/promo/register-event" "POST" "{\"eventId\": \"$event_id\", \"eventOrigin\": \"undefined\", \"promoId\": \"$PROMO_ID\"}" "$client_token")
  has_code=$(echo $register_event_response | jq -r '.hasCode')

  if [ "$has_code" == "true" ]; then
    echo "Event registration successful, hasCode is true."
    break
  else
    echo "Event registration did not return hasCode. Retrying in 1 minute..."
    sleep 60
  fi
done

# Step 4: Create promo code
echo "Creating promo code..."
create_code_response=$(make_request "https://api.gamepromo.io/promo/create-code" "POST" "{\"promoId\": \"$PROMO_ID\"}" "$client_token")
promo_code=$(echo $create_code_response | jq -r '.promoCode')

# Step 5: Write out the promo code
if [ -z "$promo_code" ] || [ "$promo_code" == "null" ]; then
  echo "Failed to create promo code"
  exit 1
else
  echo "Promo Code: $promo_code"
fi
