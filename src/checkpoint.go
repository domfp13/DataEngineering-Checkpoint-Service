package src

import "encoding/json"

type CheckpointObject struct {
	Time string `json:"Time"`
}

// MarshalBinary Encodes a CheckpointObject struct using Marshal Encode package.
// Inputs:
//     *CheckpointObject Reference
// Output:
// 		[]byte encoding of the Checkpoint struct
// 		error Returns an error if something goes wrong with the function.
func (t *CheckpointObject) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

// UnmarshalBinary Decodes a CheckpointObject struct using Marshal Encode package.
// Inputs:
//     *CheckpointObject Reference
// 	   []byte encoding of the Checkpoint struct
// Output:
// 		error Returns an error if something goes wrong with the function.
func (t *CheckpointObject) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	return nil
}
