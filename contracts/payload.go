package contracts

import (
	"encoding/json"
	"github.com/jacadenac/liftit/logging"
)

type Payload struct{
	Body 	[]byte	`json:"body"`
	Params 	[]byte	`json:"params"`
}
func (payload* Payload)ToJson()(json_payload []byte) {
	json_payload, err := json.Marshal(payload)
	logging.FailOnError(err, "Error encoding Payload struct")
	return
}
func (payload* Payload)Set(json_payload []byte){
	err := json.Unmarshal(json_payload, payload)
	logging.FailOnError(err, "Failed to convert json to payload")
}